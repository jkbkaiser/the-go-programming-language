package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func updateCmd() *cobra.Command {
	var owner string
	var repo string

	// Open editor for these if they are not passed
	var title string
	var description string

	var updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Updates an existing github issue",
		Long:  "Updates an existing github issue",
		Run: func(cmd *cobra.Command, args []string) {
			if title == "" {
				title = GetUserInput("Enter your issue title below:\n")
			}

			if description == "" {
				description = GetUserInput("Enter your issue description below:\n")
			}

			j, _ := json.Marshal(
				&Issue{
					Description: description,
					Title:       title,
				},
			)

			apiUrl := fmt.Sprintf("%s/repos/%s/%s/issues", GITHUB_URL, owner, repo)
			req, _ := http.NewRequest(http.MethodPatch, apiUrl, bytes.NewBuffer(j))
			req.Header.Set("Authorization", "Bearer "+API_KEY)

			res, err := (&http.Client{}).Do(req)
			if err != nil {
				log.Fatal(err)
			}

			log.Println(res.Status)
		},
	}

	updateCmd.Flags().StringVarP(&owner, "owner", "o", "", "owner name (required)")
	updateCmd.Flags().StringVarP(&repo, "repo", "r", "", "repository name (required)")
	updateCmd.MarkFlagRequired("owner")
	updateCmd.MarkFlagRequired("repo")

	updateCmd.Flags().StringVarP(&title, "title", "t", "", "owner name")
	updateCmd.Flags().StringVarP(&description, "description", "d", "", "repository name")

	return updateCmd
}
