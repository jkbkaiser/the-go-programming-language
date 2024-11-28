package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

const GITHUB_URL = "https://api.github.com"
const INPUT_FILE = "/tmp/gh_input.txt"

var API_KEY string
var EDITOR string
var rootCmd = &cobra.Command{
	Use:   "gh",
	Short: "Github issue CLI tool",
	Long:  "This is a simple CLI tool for reading, creating, updating and deleting github issues",
}

func Execute() {
	// Load environment variables and API key
	godotenv.Load(".env")
	API_KEY = os.Getenv("GITHUB_API_KEY")
	EDITOR = os.Getenv("EDITOR")

	// Hide completion command
	rootCmd.AddCommand(&cobra.Command{Use: "completion", Hidden: true})

	// Add commands
	rootCmd.AddCommand(createCmd())
	rootCmd.AddCommand(getCmd())
	rootCmd.AddCommand(updateCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
