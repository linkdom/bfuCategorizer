package game

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

type Game struct {
	Title       string
	Publisher   string
	Description string
}

func Parse(r io.Reader) ([]Game, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc, "")
	return nil, nil
}

func dfs(node *html.Node, padding string) {
	fmt.Println(padding, node.Data)
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
