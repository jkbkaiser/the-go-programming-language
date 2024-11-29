package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const API = "https://xkcd.com/%d/info.0.json"
const COMICS_URL = "https://xkcd.com/%d"
const INDEX_FILE = "./index.json"
const COMICS_FILE = "./comics.json"

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func getComics() []Comic {
	comicId := 1
	comics := []Comic{}

	for {
		if comicId == 404 {
			comicId += 1
			continue
		}

		comic := Comic{}
		resp, err := http.Get(fmt.Sprintf(API, comicId))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			break
		}

		err = json.NewDecoder(resp.Body).Decode(&comic)
		if err != nil {
			log.Fatal(err)
		}

		comics = append(comics, comic)

		log.Println(resp.StatusCode, comicId)

		comicId += 1
	}

	return comics
}

type Index = map[string]map[int]bool

func buildIndex(comics []Comic) Index {
	index := map[string]map[int]bool{}

	for _, comic := range comics {
		words := strings.Split(comic.Title, " ")
		words = append(words, strings.Split(comic.Transcript, " ")...)
		wordsLowerCase := []string{}

		for _, word := range words {
			wordsLowerCase = append(wordsLowerCase, strings.ToLower(word))
		}

		for _, word := range wordsLowerCase {
			if ids, ok := index[word]; ok {
				ids[comic.Num] = true
				index[word] = ids
			} else {
				index[word] = map[int]bool{comic.Num: true}
			}
		}
	}

	return index
}

func storeComics(comics []Comic) {
	f, err := os.Create(COMICS_FILE)

	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(f).Encode(comics)

	if err != nil {
		log.Fatal(err)
	}
}

func storeIndex(index Index) {
	f, err := os.Create(INDEX_FILE)

	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(f).Encode(index)

	if err != nil {
		log.Fatal(err)
	}
}

func loadComics() ([]Comic, error) {
	f, err := os.Open(COMICS_FILE)

	if err != nil {
		return nil, err
	}

	comics := []Comic{}
	err = json.NewDecoder(f).Decode(&comics)

	if err != nil {
		return nil, err
	}

	return comics, nil
}

func loadIndex() (Index, error) {
	f, err := os.Open(INDEX_FILE)

	if err != nil {
		return nil, err
	}

	index := Index{}
	err = json.NewDecoder(f).Decode(&index)

	if err != nil {
		return nil, err
	}

	return index, nil
}

func searchIndex(comics []Comic, index Index, term string) {
	res := index[term]

	for k := range res {
		comic := comics[k]
		fmt.Println("URL:")
		fmt.Println(fmt.Sprintf(COMICS_URL, k))
		fmt.Println()
		fmt.Println("Transcript")
		fmt.Println(comic.Transcript)
	}
}

func main() {
	var buildIndexFlag bool
	var searchTerm string
	flag.BoolVar(&buildIndexFlag, "build-index", false, "Builds an index")
	flag.StringVar(&searchTerm, "term", "", "Term to search for")
	flag.Parse()

	var index Index
	var comics []Comic
	var err error

	if buildIndexFlag {
		comics = getComics()
		index = buildIndex(comics)
		storeIndex(index)
		storeComics(comics)
	}

	if searchTerm == "" {
		return
	}

	if !buildIndexFlag {
		index, err = loadIndex()

		if err != nil {
			fmt.Println("Please first build an index using the --build-index flag")
			return
		}

		comics, err = loadComics()

		if err != nil {
			fmt.Println("Please first build an index using the --build-index flag")
			return
		}
	}

	searchIndex(comics, index, searchTerm)
}
