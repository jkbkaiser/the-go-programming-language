package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const issueListTempl = `
  <h1>Issues</h1>
  <table>
  <tr style='text-align: left'>
    <th>#</th>
    <th>State</th>
    <th>User</th>
    <th>Title</th>
  </tr>
  {{range .Issues}}
    <tr>
    <td><a href='/issues/{{.Number}}'>{{.Number}}</td>
    <td>{{.State}}</td>
    <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
    <td><a href='/issues/{{.Number}}'>{{.Title}}</a></td>
    </tr>
  {{end}}
  </table>
`

const issueTempl = `
  <h1>#{{.Number}} {{.Title}}</h1>
  <p>State: {{.State}}</p>
  <p>Opened by: {{.User.Login}}</p>
  <p>Descrption: {{.Body}}</p>
`

type Server struct {
	IssueList IssueList
}

func (s Server) handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		templ := template.Must(template.New("issueList").Parse(issueListTempl))
		templ.Execute(w, s.IssueList)
		return
	}
	parts := strings.Split(r.URL.Path, "/")
	number, err := strconv.ParseInt(parts[len(parts)-1], 10, 0)

	if err != nil {
		log.Println(err)
		return
	}

	if int(number) > len(s.IssueList.Issues) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	number -= 1

	templ := template.Must(template.New("issue").Parse(issueTempl))
	templ.Execute(w, s.IssueList.Issues[number])
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: githubserver OWNER REPO")
		os.Exit(1)
	}
	fmt.Println("listening on port 8080")

	owner := os.Args[1]
	repo := os.Args[2]

	client := Client{
		Owner: owner,
		Repo:  repo,
	}

	issueList := client.ListIssues()

	server := Server{
		IssueList: issueList,
	}

	http.HandleFunc("/", server.handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
