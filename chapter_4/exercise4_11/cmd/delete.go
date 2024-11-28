package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func deleteCmd() *cobra.Command {
	var owner string
	var repo string
	var issue string

	var deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Deletes a github issue",
		Long:  "Deletes a github issue",
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

			apiUrl = "https://api.github.com/graphql"
			body := []byte(fmt.Sprintf(
				`query:
          mutation {
            deleteIssue(input: {
              clientMutationId: "foobar",
              issueId: %s
            }) {
              clientMutationId
            }
				}`,
				issue.Id,
			))

			req, _ = http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(body))
			req.Header.Set("Authorization", "Bearer "+API_KEY)
			req.Header.Set("Content-Type", "application/json")
			resp, err = (&http.Client{}).Do(req)

			if err != nil {
				log.Fatal(err)
			}

			// if resp.StatusCode != 200 {
			// 	log.Fatal(resp.StatusCode)
			// }

			var t = map[string]interface{}{}
			json.NewDecoder(resp.Body).Decode(&t)

			log.Println(resp.StatusCode)
			log.Println(t)
		},
	}

	deleteCmd.Flags().StringVarP(&owner, "owner", "o", "", "owner name (required)")
	deleteCmd.Flags().StringVarP(&repo, "repo", "r", "", "repository name (required)")
	deleteCmd.Flags().StringVarP(&issue, "issue", "i", "", "issue number (required)")
	deleteCmd.MarkFlagRequired("owner")
	deleteCmd.MarkFlagRequired("repo")
	deleteCmd.MarkFlagRequired("issue")

	return deleteCmd
}
