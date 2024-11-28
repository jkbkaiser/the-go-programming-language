package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func createCmd() *cobra.Command {
	var owner string
	var repo string

	// Open editor for these if they are not passed
	var title string
	var description string

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Creates a new github issue",
		Long:  "Creates a new github issue",
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
			req, _ := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(j))
			req.Header.Set("Authorization", "Bearer "+API_KEY)

			res, err := (&http.Client{}).Do(req)
			if err != nil {
				log.Fatal(err)
			}

			log.Println(res.Status)
		},
	}

	createCmd.Flags().StringVarP(&owner, "owner", "o", "", "owner name (required)")
	createCmd.Flags().StringVarP(&repo, "repo", "r", "", "repository name (required)")
	createCmd.MarkFlagRequired("owner")
	createCmd.MarkFlagRequired("repo")

	createCmd.Flags().StringVarP(&title, "title", "t", "", "owner name")
	createCmd.Flags().StringVarP(&description, "description", "d", "", "repository name")

	return createCmd
}
