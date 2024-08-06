package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

const GithubURL = "https://api.github.com"

var apiKey string

type IssueResponse struct {
	State       string `json:"state"`
	Title       string `json:"title"`
	Description string `json:"body"`
}

func (issueResp IssueResponse) String() string {
	return fmt.Sprintf("%v\n%v\n%v", issueResp.Title)
}

type Issue struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (issue Issue) open(owner, repo string) error {
	u, err := url.JoinPath(GithubURL, "repos", owner, repo, "issues")
	if err != nil {
		return err
	}

	body, err := json.Marshal(issue)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func getIssue(owner, repo string, issueNumber int) (IssueResponse, error) {
	u, err := url.JoinPath(GithubURL, "repos", owner, repo, "issues", fmt.Sprintf("%d", issueNumber))
	if err != nil {
		return IssueResponse{}, err
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return IssueResponse{}, err
	}
	defer res.Body.Close()
	content, err := io.ReadAll(res.Body)

	if err != nil {
		return IssueResponse{}, err
	}

	resp := IssueResponse{}
	if err = json.Unmarshal(content, &resp); err != nil {
		return IssueResponse{}, err
	}

	return resp, nil
}

var owner = flag.String("owner", "", "repository owner")
var repo = flag.String("repo", "", "repository name")
var create = flag.Bool("create", false, "create issue")
var read = flag.Bool("read", false, "read issue")
var title = flag.String("title", "ISSUE title", "issue title")
var description = flag.String("description", "ISSUE description", "issue description")
var issueNumber = flag.Int("issue-num", 1, "issue number")

func main() {
	godotenv.Load(".env")
	apiKey = os.Getenv("API_KEY")

	flag.Parse()

	var err error = nil
	if *create {
		fmt.Println("Creating issue")
		issue := Issue{*title, *description}
		err = issue.open(*owner, *repo)
		if err != nil {
			log.Fatal(err)
		}
	} else if *read {
		fmt.Println("Retrieving issue")
		issue, err := getIssue(*owner, *repo, *issueNumber)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(issue)
	}
}
