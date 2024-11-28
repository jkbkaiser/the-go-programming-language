package main

import (
	"exercise4-10/issue"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := issue.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	categories := map[string][]*issue.Issue{
		"last month":       {},
		"last year":        {},
		"more than a year": {},
	}

	n := time.Now()
	for _, item := range result.Items {
		if item.CreatedAt.After(n.Add(time.Hour * 24 * 7 * 3)) {
			categories["last month"] = append(categories["last month"], item)
		} else if item.CreatedAt.After(n.Add(time.Hour * 24 * 52)) {
			categories["last month"] = append(categories["last month"], item)
		} else {
			categories["last month"] = append(categories["last month"], item)
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for key, items := range categories {

		fmt.Println(key)
		for _, item := range items {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}
