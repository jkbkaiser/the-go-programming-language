package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func GetUserInput(prompt string) string {
	os.Remove(INPUT_FILE)
	file, err := os.Create(INPUT_FILE)
	if err != nil {
		log.Fatal(err)
	}
	file.Write([]byte(prompt))

	cmd := exec.Command(EDITOR, INPUT_FILE)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Start()
	cmd.Wait()

	contents, err := os.ReadFile(INPUT_FILE)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(fmt.Sprintf("%s(.*)", prompt))
	userInput := re.FindStringSubmatch(string(contents))

	return userInput[len(userInput)-1]
}
