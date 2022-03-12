package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//.index_game_cell_widget game_cell 	-- the parent div
//.label 								-- the game name
//.user_row 							-- the publisher name
//.short_text 							-- the description

func main() {

	resp, err := http.Get("https://itch.io/b/1316/bundle-for-ukraine")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

}
