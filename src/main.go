package main

import (
	"fmt"

	"github.com/jcav67/wiki_go/models"
)

func main() {
	page1 := models.Page{Title: "TestPage", Body: []byte("Esta es una prueba")}
	err := page1.Save()
	if err != nil {
		fmt.Println(err)
		return
	}

	p2, _ := models.LoadPage(page1.Title)
	fmt.Println(string(p2.Body))
}
