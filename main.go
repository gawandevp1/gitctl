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
	if err := gatherData(); err != nil {
		log.Panic(err.Error())
	}

}

// gatherData.. 
func gatherData() (err error) {
	// get data from Input file
	var input models.Input
	input, err = utils.GetConfigValues(filePath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	gitHandler := controller.GetNewGR(input)

	summary, err := gitHandler.FetchPRHistory()
	if err != nil {
		log.Println(err.Error())
		return
	}

	// Reason to add this comment: Sending email notification, we can make this async.
	// go gitHandler.EmailNotification(summary)

	// print content of email.
	// This only print the data to send in mail.
	return gitHandler.EmailNotification(summary)
}
