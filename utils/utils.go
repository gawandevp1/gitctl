package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gawandevp1/gitctl/models"
)

const (
	contentType = "application/json"
)

func GetConfigValues(filepath string) (input models.Input, err error) {
	var file *os.File
	file, err = os.Open(filepath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = json.NewDecoder(file).Decode(&input)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

func MakeRequest(method, url string) (response *http.Response, err error) {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	request.Header.Add("Accept", contentType)
	request.Header.Add("Content-Type", contentType)
	return client.Do(request)
}
