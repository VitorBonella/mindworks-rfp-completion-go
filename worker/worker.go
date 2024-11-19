package worker

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/VitorBonella/mindworks-rfp-completion-go/database"
	"github.com/VitorBonella/mindworks-rfp-completion-go/models"
	"github.com/VitorBonella/mindworks-rfp-completion-go/worker/ai"
	gemini "github.com/VitorBonella/mindworks-rfp-completion-go/worker/ai"
	"github.com/google/generative-ai-go/genai"
)

const WaitTime = 30*time.Second

func worker(id int, taskQueue <-chan *models.RFP) {
	for task := range taskQueue {
		log.Printf("Worker %d processing task name: %s - to user %d\n", id, task.Name, task.UserId)
		err := ProcessRFP(task)
		if err != nil{
			//Change Status To Finished With Error
			database.SetRFPStatus(task,models.RFPStatusFinishedWithError)
		}
		time.Sleep(WaitTime)
		log.Printf("Worker %d finished task ID: %s\n", id, task.Name)
	}
}


func ProcessRFP(rfp *models.RFP) error{
	
	ctx := context.Background()
	
	// Change status to processing
	err := database.SetRFPStatus(rfp,models.RFPtatusProcessing)
	if err != nil{
		log.Println("error setting status to processing",err)
		return err
	}

	// Download Equipment Manuals
	pathList, _ := models.DownloadManyFile(rfp.Equipments)
	if len(pathList) == 0 {
		log.Println("Error downloading all files")
		return errors.New("error downloading all files")
	}
	defer models.DeleteManyFiles(pathList)

	//generate the json of RFP in blocks of 50.
	jsonBlocks, err := models.GenerateQuestionJSON(rfp.Requirements)
	if err != nil {
		log.Println("Error generation json blocks:", err)
		return err
	}

	//user 
	user, err := database.GetUserById(strconv.Itoa(int(rfp.UserId)))
	if err != nil {
		log.Println("Error getting user:", err)
		return err
	}

	// Get Gemini Client
	apiKey,err := user.GetAPIKey()
	if err != nil {
		log.Println("Error getting user api key:", err)
		return err
	}


	client, err := gemini.NewGeminiClient(ctx,apiKey)
	if err != nil {
		log.Println("error creating gemini client",err)
		return err
	}
	defer client.Close()

	// Upload to Gemini the Equipments PDF
	files, err := gemini.UploadManyFileGemini(ctx,client, pathList)
	if err != nil {
		log.Println("Error uploading file to gemini")
		return err
	}
	defer gemini.CloseManyFilesGemini(ctx,client,files)

	// For each equipment ask the questions
	model := client.GenerativeModel("gemini-1.5-pro")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(ai.Intruction)},
	}

	for i, f := range files{
		for _, j := range jsonBlocks{
			resp, err := model.GenerateContent(ctx,
				genai.Text(j),
				genai.FileData{URI: f.URI})
			if err != nil {
				log.Println("Failed to generate content",err)
				time.Sleep(time.Second*60)
				continue
			}
			endTime := time.Now()
			database.CreateResult(&models.Result{RFPId: rfp.Id,EquipmentId: rfp.Equipments[i].Id,Text : printResponse(resp),EndDate: &endTime})
			log.Println("PROCESSED RESULT")
			time.Sleep(time.Second*60)
		}

	}

	rfpEndTime := time.Now()
	//Set as completed with no errors
	rfp.EndDate = &rfpEndTime
	err = database.SetRFPStatus(rfp,models.RFPStatusFinished)
	if err != nil{
		log.Println("error setting status to finished",err)
		return err
	}

	return nil
}

func printResponse(resp *genai.GenerateContentResponse) string{
	text := ""
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				text += fmt.Sprint(part)
			}
		}
	}

	return strings.ReplaceAll(text, "json", "")
}

func RunQueue() {

	// Define the number of workers
	const numWorkers = 1

	// Create the task queue
	taskQueue := make(chan *models.RFP,100)

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go worker(i, taskQueue)
	}

	// Task producer
	go func() {
		for {
			task, err := database.ListNewestCreatedRFP()
			if err != nil{
				//log.Println("No new task",err)
				time.Sleep(WaitTime)
				continue
			}
			taskQueue <- task
			log.Printf("Added RFP name: %s to the queue\n", task.Name)
			time.Sleep(WaitTime) // Simulate task generation interval
		}
	}()

}