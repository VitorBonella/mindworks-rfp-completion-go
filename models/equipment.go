package models

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

type Equipment struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" gorm:"unique" validate:"required"`
	DownloadLink     string `json:"download_link" gorm:"unique" validate:"required"`
	UserId uint `json:"user_id"`
}

func NewEquipment(name string, downloadLink string) (*Equipment, error) {

	if name == ""{
		return nil, errors.New("invalid name")
	}

	equip := Equipment{
		Name: name,
		DownloadLink: downloadLink,
	}

	ok, err := equip.TestDownloadLink()
	if err != nil{
		log.Println("Error on Test Download Link",err)
		return nil, errors.New("invalid link")
	}
	if !ok{
		return nil, errors.New("invalid link")
	}

	return &equip, err

}

func (e *Equipment) TestDownloadLink() (bool,error){

	err := e.DownLoadFile(true)
	if err != nil{
		return false,err
	}

	return true,nil
}

func (e *Equipment) DownLoadFile(test bool) error{


	// Create an HTTP request
	req, err := http.NewRequest("GET", e.DownloadLink, nil)
	if err != nil {
		log.Println("Error creating request")
		return err
	}

	// Set custom headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Referer", e.DownloadLink)

	// Create an HTTP client and execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error doing request")
		return err
	}
	defer resp.Body.Close()

	// Check if the response status is OK (200)
	if resp.StatusCode != http.StatusOK {
		log.Println("Failed to download file: ", resp.Status)
		return errors.New("failed to download file")
	}

	var fileName string
	if test{
		fileName = "tmp.pdf"
	} else{
		fileName = e.Name+".pdf"
	}

	// Create the file to save the downloaded content
	out, err := os.Create(fileName)
	if err != nil {
		log.Println("Error creating file")
		return err
	}
	defer out.Close()

	// Copy the response body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println("Error copying file body")
		return err
	}

	//log.Println("PDF downloaded successfully with headers!")
	return nil
}