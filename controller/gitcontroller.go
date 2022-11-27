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

type gitCtl struct {
	Input models.Input
}

func GetNewGR(input models.Input) (gr *gitCtl) {
	return &gitCtl{
		Input: input,
	}
}

// get PR history like, closed merged open total PR.
// FetchPRHistory ....
func (gr *gitCtl) FetchPRHistory() (summary map[string]int, err error) {
	next := true
	page := 1
	summary = make(map[string]int)
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

// EmailNotification ....
func (pr *gitCtl) EmailNotification(prData map[string]int) (err error) {
	// printing the content of email.
	repoName := strings.Split(pr.Input.Url, "/")[4:]
	fmt.Println("To: " + pr.Input.RecieverID)
	fmt.Println("From: " + pr.Input.SenderID)
	fmt.Println("<<<<<<------- Here is the PR Data fro gitctl------->>>>>>")
	fmt.Println("Subject: [DoNotReply] PR Report of last weeks github PRs for repo ", strings.Join(repoName, "/"))
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("<<<<<   State of PR    ::      Count      >>>>>")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	for key, val := range prData {
		fmt.Println("<<<<" + key + "    ->       " + strconv.Itoa(val) + ">>>>")
	}
	return nil
}
