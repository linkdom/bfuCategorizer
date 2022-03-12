package main

import (
	"fmt"
	"github.com/linko1994/bfuCategorizer/game"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

	games, err := game.Parse(strings.NewReader(string(body)))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%+v\n", games)

}
