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
	fmt.Println("git Runner started")
	if err := Runner(); err != nil {
		log.Panic(err.Error())
	}

}

func Runner() (err error) {
	//read data from config
	var input models.Input
	input, err = utils.GetConfigValues(filePath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	githandler := controller.GetNewGR(input)

	summary, err := githandler.FetchGitPRSummary()
	if err != nil {
		log.Println(err.Error())
		return
	}
	//send mail to server
	return githandler.MailToAdmin(summary)
}
