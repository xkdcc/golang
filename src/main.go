package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type empolyee struct {
	firstName, lastName string
	salary              int
}

func main() {
	var r = "☞"
	fmt.Printf("Print the lower casee base 16 by %%x : %x \n", r)
	fmt.Printf("Print the default format by %%v      : %v \n", r)
	fmt.Printf("Print type value by %%T              : %T \n", r)

	fmt.Println()

	var e empolyee
	fmt.Println("Print Employee struct :\n", e)

	var m = mux.NewRouter()

	m.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":81", m)
}
