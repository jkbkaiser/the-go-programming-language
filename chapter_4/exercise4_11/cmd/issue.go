package cmd

import "fmt"

type Issue struct {
	Id          string `json:"node_id"`
	Title       string `json:"title"`
	Description string `json:"body"`
}

func (i Issue) String() string {
	return fmt.Sprintf("Title:\n%s\n\nDescription:\n%s\n", i.Title, i.Description)
}
