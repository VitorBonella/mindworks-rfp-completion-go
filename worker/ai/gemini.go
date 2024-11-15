package ai

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)


func NewGeminiClient(ctx context.Context,apiKey string) (*genai.Client, error){

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Println("Error creating genai client")
		return nil, err
	}

	return client,nil
}

func UploadFileToGemini(ctx context.Context,client *genai.Client, filePath string) (*genai.File, error){

	file, err := client.UploadFileFromPath(ctx, filePath, nil)
	if err != nil {
		log.Println("Error uploading file")
		return nil, err
	}

	return file, nil
}

func UploadManyFileGemini(ctx context.Context,client *genai.Client, filePath []string) ([]*genai.File, error){

	var files []*genai.File

	for _, f := range  filePath{
		file, err := client.UploadFileFromPath(ctx, f, nil)
		if err != nil {
			log.Println("Error uploading file",err)
		}
		files = append(files,file)
	}
	return files,nil
}

func CloseManyFilesGemini(ctx context.Context,client *genai.Client,files []*genai.File) (){

	for _, f := range  files{
		client.DeleteFile(ctx, f.Name)
	}

}