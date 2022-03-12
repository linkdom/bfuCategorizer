package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type Games struct {
	Games []Game `json:"games"`
}

type Game struct {
	Name        string `json:"name"`
	Publisher   string `json:"publisher"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Genre       string
}

func main() {

	jsonFile, err := os.Open("games.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var games Games

	json.Unmarshal(byteValue, &games)

	for i := 0; i < len(games.Games); i++ {

		regex := regexp.MustCompile(`\[([^\[\]]*)\]`)
		match := regex.FindAllStringSubmatch(games.Games[i].Name, -1)

		if len(match) != 0 {

			r := strings.Replace(games.Games[i].Name, "["+match[0][1]+"]", "", 10)

			s := strings.Replace(r, " ", "-", 10)
			s = strings.Replace(s, ".", "-dot-", 10)
			s = strings.TrimPrefix(s, "-")
			s = strings.TrimSuffix(s, "-")
			fmt.Println("Name: " + strings.ToLower(s))

		} else {
			s := strings.Replace(games.Games[i].Name, " ", "-", 10)
			s = strings.Replace(s, ".", "-dot-", 10)
			fmt.Println("Name: " + strings.ToLower(s))
		}

		fmt.Println("Publisher: " + games.Games[i].Publisher)
		fmt.Println("Link: " + games.Games[i].Link)
		fmt.Println("Description: " + games.Games[i].Description + "\n")
	}

	/*
		resp, err := http.Get("https://www.igdb.com/games/"+s)

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

	*/
}
