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
	"github.com/google/generative-ai-go/genai"
)


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


	client, err := ai.NewGeminiClient(ctx,apiKey)
	if err != nil {
		log.Println("error creating gemini client",err)
		return err
	}
	defer client.Close()

	// Upload to Gemini the Equipments PDF
	files, err := ai.UploadManyFileGemini(ctx,client, pathList)
	if err != nil {
		log.Println("Error uploading file to gemini")
		return err
	}
	defer ai.CloseManyFilesGemini(ctx,client,files)

	// For each equipment ask the questions
	model := client.GenerativeModel("gemini-1.5-pro")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(ai.Intruction)},
	}

	for i, f := range files{
		for bi, j := range jsonBlocks{
			resp, err := model.GenerateContent(ctx,
				genai.Text(j),
				genai.FileData{URI: f.URI})
			if err != nil {
				log.Println("Failed to generate content",err)
				time.Sleep(time.Second*60)
				return err
			}
			endTime := time.Now()
			database.CreateResult(&models.Result{RFPId: rfp.Id,EquipmentId: rfp.Equipments[i].Id,Text : printResponse(resp),EndDate: &endTime})
			log.Printf("[%d] PROCESSED RESULT %s %d/%d\n",rfp.Id, f.DisplayName,bi+1,len(jsonBlocks))
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