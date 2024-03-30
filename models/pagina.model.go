package models

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	filename := fmt.Sprintf("pages/%v.txt", p.Title)

	return os.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := fmt.Sprintf("pages/%v.txt", title)

	if body, erro := os.ReadFile(filename); erro != nil {
		return nil, erro
	} else {
		return &Page{Title: title, Body: body}, nil
	}

}
