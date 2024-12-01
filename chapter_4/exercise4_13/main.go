package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type SearchResult struct {
	Title  string
	Poster string
	Type   string
	Year   string
	IMDBID string `json:"imdbId"`
}

type SearchResponse struct {
	Response string
	Search   []SearchResult
}

func main() {
	godotenv.Load(".env")
	API_KEY := os.Getenv("API_KEY")
	baseUrl, _ := url.Parse("https://omdbapi.com/")

	// Defines the movieFlag flag
	var movieFlag string
	flag.StringVar(&movieFlag, "movie", "", "The movie for which to download the poster image (required)")
	flag.Usage = func() {
		fmt.Println("poster is a tool that downloads poster images for movies using the Open Movie Datatbase")
		fmt.Println()
		flag.PrintDefaults()
	}
	flag.Parse()

	if movieFlag == "" {
		fmt.Println("please set the movie flag")
		return
	}

	params := url.Values{}
	params.Add("s", movieFlag)
	params.Add("apikey", API_KEY)
	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	searchResp := SearchResponse{}
	err = json.NewDecoder(resp.Body).Decode(&searchResp)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if len(searchResp.Search) == 0 {
		fmt.Println("No results")
		os.Exit(0)
	}

	resp, err = http.Get(searchResp.Search[0].Poster)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = os.WriteFile("output.jpg", data, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("Written the poster image to output.jpg")
}
