package main

import (
	"os"
	"path"
)

type WikiPage struct {
	Title string
	Body  []byte
}

func LoadWikiPage(title string) (*WikiPage, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(path.Join("media", filename))
	if err != nil {
		return nil, err
	}
	return &WikiPage{Title: title, Body: body}, nil
}

func (w *WikiPage) Save() error {
	filename := w.Title + ".txt"
	err := os.WriteFile(path.Join("media", filename), w.Body, 0600)
	return err
}
