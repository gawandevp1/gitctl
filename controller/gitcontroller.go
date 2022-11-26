package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gawandevp1/gitctl/models"
	"github.com/gawandevp1/gitctl/utils"
)

// GitFunctions
// type GitFunctions interface {
// 	FetchGitPRSummary() (summary map[string]int, err error)
// 	MailToAdmin(summary map[string]int) (err error)
// }

type gitRunner struct {
	Input models.Input
}

type summaryData struct {
	OpenPR   int
	ClosedPR int
	TotalPR  int
}

func GetNewGR(input models.Input) (gr *gitRunner) {
	return &gitRunner{
		Input: input,
	}
}

// FetchGitData() to get data from git repo
func (gr *gitRunner) FetchGitPRSummary() (summary map[string]int, err error) {
	//code to get data from github
	next := true
	page := 1
	summary = make(map[string]int)
	// find what was the day, 1 week back so we can get the PRs only after that day
	//need to get data for all PRs updated within last week to check if they are still open or closed or merged
	previousDays := time.Now().AddDate(0, 0, -1*gr.Input.PrevDays)
	for next {
		gitUrl := gr.Input.Url + "/pulls?state=all&&sort=updated&&direction=desc&&page=" + strconv.Itoa(page)
		response, er := utils.MakeRequest(http.MethodGet, gitUrl)
		if er != nil {
			log.Println(er.Error())
			err = er
			return
		}
		prData := []models.GitResponseStruct{}
		err = json.NewDecoder(response.Body).Decode(&prData)
		if err != nil {
			log.Println(err.Error())
			return
		}
		for _, pr := range prData {
			if pr.UpdatedAt.After(previousDays) {
				summary[pr.State] += 1
				summary["total"] += 1
				if pr.MergedAt != nil {
					summary["merged"] += 1
				}
			} else {
				next = false
				break
			}
		}
		//go to next page
		page++
	}
	return
}

// MailToAdmin() to send mail to admin
func (gr *gitRunner) MailToAdmin(summaryData map[string]int) (err error) {
	//code to send data to admin
	//as stated in assgnment this is the content of mail to be sent.
	repoName := strings.Split(gr.Input.Url, "/")[4:]
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("To: " + gr.Input.RecieverID)
	fmt.Println("From: " + gr.Input.SenderID)
	fmt.Println("Subject: Summary Report of last weeks github PRs for repo ", strings.Join(repoName, "/"))
	fmt.Println(" The summary table is as follows   ")
	fmt.Println("--------------------------------------")
	fmt.Println("|   State of PR    |       Count      |")
	fmt.Println("--------------------------------------")
	for key, val := range summaryData {
		fmt.Println("| " + key + "    |       " + strconv.Itoa(val) + "         |")
	}
	fmt.Println("------------------------------------------------------------------------")
	return nil
}
