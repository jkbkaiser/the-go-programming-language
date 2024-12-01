package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

const GITHUB_URL = "https://api.github.com"

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number int
	Title  string
	State  string
	User   User
	Body   string
}

type IssueList struct {
	Issues []Issue
}

type Client struct {
	Owner string
	Repo  string
}

func (c Client) ListIssues() IssueList {
	apiUrl := fmt.Sprintf("%s/repos/%s/%s/issues", GITHUB_URL, c.Owner, c.Repo)
	req, _ := http.NewRequest(http.MethodGet, apiUrl, nil)

	resp, err := (&http.Client{}).Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal(resp.StatusCode)
	}

	var issues = []Issue{}
	err = json.NewDecoder(resp.Body).Decode(&issues)

	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(issues, func(i, j int) bool {
		return issues[i].Number < issues[j].Number
	})

	return IssueList{Issues: issues}
}
