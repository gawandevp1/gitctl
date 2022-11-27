package main

import (
	"fmt"
	"log"

	"github.com/gawandevp1/gitctl/controller"
	"github.com/gawandevp1/gitctl/models"
	"github.com/gawandevp1/gitctl/utils"
)

const (
	filePath = "input.json"
)

func main() {
	fmt.Println("Get PR data from github repo.....")
	if err := GatherData(); err != nil {
		log.Panic(err.Error())
	}

}

func GatherData() (err error) {
	//get data from Input file
	var input models.Input
	input, err = utils.GetConfigValues(filePath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	githandler := controller.GetNewGR(input)

	summary, err := githandler.FetchPRHistory()
	if err != nil {
		log.Println(err.Error())
		return
	}
	//print data to send mail to server.
	//This only print the data to send in mail.
	return githandler.EmailNotification(summary)
}
