package game

import "io"

type Game struct {
	Title       string
	Publisher   string
	Description string
}

func Parse(r io.Reader) ([]Game, error) {
	return nil, nil
}
