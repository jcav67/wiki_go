package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jcav67/wiki_go/models"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]

	p, _ := models.LoadPage(title)

	fmt.Fprintf(w, "<h1>%s</h1> <div>%s</div>", p.Title, p.Body)
}

func main() {
	// page1 := models.Page{Title: "TestPage", Body: []byte("Esta es una prueba")}
	// err := page1.Save()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// p2, _ := models.LoadPage(page1.Title)
	// fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", viewHandler)
	fmt.Println("App started!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
