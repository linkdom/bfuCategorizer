package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Games struct {
	Games []Game `json:"games"`
}

type Game struct {
	Name        string `json:"name"`
	Publisher   string `json:"publisher"`
	Link        string `json:"link"`
	Description string `json:"description"`
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
		fmt.Println("Name: " + games.Games[i].Name)
		fmt.Println("Publisher: " + games.Games[i].Publisher)
		fmt.Println("Link: " + games.Games[i].Link)
		fmt.Println("Description: " + games.Games[i].Description)
	}
}
