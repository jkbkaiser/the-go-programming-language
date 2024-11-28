package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func getCmd() *cobra.Command {
	var owner string
	var repo string
	var issue string

	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Retrieves a github issue",
		Long:  "Retrieves a github issue",
		Run: func(cmd *cobra.Command, args []string) {
			apiUrl := fmt.Sprintf("%s/repos/%s/%s/issues/%s", GITHUB_URL, owner, repo, issue)
			req, _ := http.NewRequest(http.MethodGet, apiUrl, nil)
			req.Header.Set("Authorization", "Bearer "+API_KEY)

			resp, err := (&http.Client{}).Do(req)

			if err != nil {
				log.Fatal(err)
			}

			if resp.StatusCode != 200 {
				log.Fatal(resp.StatusCode)
			}

			var issue = Issue{}
			json.NewDecoder(resp.Body).Decode(&issue)
			fmt.Println(issue)
		},
	}

	getCmd.Flags().StringVarP(&owner, "owner", "o", "", "owner name (required)")
	getCmd.Flags().StringVarP(&repo, "repo", "r", "", "repository name (required)")
	getCmd.Flags().StringVarP(&issue, "issue", "i", "", "issue number (required)")
	getCmd.MarkFlagRequired("owner")
	getCmd.MarkFlagRequired("repo")
	getCmd.MarkFlagRequired("issue")

	return getCmd
}
