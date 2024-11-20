package worker

import (
	"log"
	"time"

	"github.com/VitorBonella/mindworks-rfp-completion-go/database"
	"github.com/VitorBonella/mindworks-rfp-completion-go/models"
)

const WaitTime = 30*time.Second
func worker(id int, taskQueue <-chan *models.RFP, rfpsInQueue map[uint]struct{}) {
	for task := range taskQueue {
		log.Printf("Worker %d processing task name: %s - to user %d\n", id, task.Name, task.UserId)
		err := ProcessRFP(task)
		if err != nil {
			// Change Status To Finished With Error
			database.SetRFPStatus(task, models.RFPStatusFinishedWithError)
		}
		
		// Remove the task from the map after processing
		delete(rfpsInQueue, task.Id)

		log.Printf("Worker %d finished task ID: %s\n", id, task.Name)
	}
}


func RunQueue() {
	// Define the number of workers
	const numWorkers = 1

	// Create the task queue
	taskQueue := make(chan *models.RFP, 100)

	// A map to track RFPs that are already in the queue
	rfpsInQueue := make(map[uint]struct{})

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go worker(i, taskQueue, rfpsInQueue)
	}

	// Task producer
	go func() {
		for {
			task, err := database.GetNewestCreatedRFP()
			if err != nil {
				//log.Println("No new task", err)
				time.Sleep(WaitTime)
				continue
			}

			// Check if the task is already in the queue
			if _, exists := rfpsInQueue[task.Id]; exists {
				// Skip adding the task if it's already in the queue
				log.Printf("Task %s already in the queue, skipping.\n", task.Name)
			} else {
				// Add the task to the queue and the map
				taskQueue <- task
				rfpsInQueue[task.Id] = struct{}{}
				log.Printf("Added RFP name: %s to the queue\n", task.Name)
			}
			time.Sleep(WaitTime) // Simulate task generation interval
		}
	}()

	// Check if there are processing RFPs that need to be reverted to created
	go func() {
		for {
			// Check if taskQueue is empty
			if len(taskQueue) == 0 {
				// Get all RFPs in processing status
				processingRFPs, err := database.ListProcessingRFP()
				if err != nil {
					log.Println("Error listing processing RFPs:", err)
				} else {
					// Change the status of the processing RFPs to created
					for _, rfp := range processingRFPs {
						// Delete preview results
						err = database.DeleteResults(rfp.Id)
						if err != nil {
							log.Println("Error Deleting previous results:", err)
						}


						err := database.SetRFPStatus(rfp, models.RFPStatusCreated)
						if err != nil {
							log.Println("Error changing RFP status to created:", err)
						} else {
							log.Printf("RFP %s status changed to created\n", rfp.Name)
						}
					}
				}
			}

			// Sleep for a while before checking again
			time.Sleep(time.Second*90)
		}
	}()
}