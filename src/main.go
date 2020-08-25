package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type empolyee struct {
	firstName, lastName string
	salary              int
}

// Post structure
type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Post{})
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(1000000))
	posts = append(posts, post)
	json.NewEncoder(w).Encode(&post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)

			var post Post
			dec := json.NewDecoder(r.Body)
			dec.DisallowUnknownFields()
			_ = dec.Decode(&post)
			// Not allow specify id in body, it has been specified in URL
			// Now assign it back
			post.ID = params["id"]
			posts = append(posts, post)
			json.NewEncoder(w).Encode(&post)
			return
		}
	}
	json.NewEncoder(w).Encode(posts)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(posts)
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

	posts = append(posts, Post{ID: "1", Title: "My first post", Body: "This is the content of my first post"})
	posts = append(posts, Post{ID: "2", Title: "My second post", Body: "This is the content of my second post"})

	m.HandleFunc("/posts", getPosts).Methods("GET")
	m.HandleFunc("/posts", createPost).Methods("POST")
	m.HandleFunc("/posts/{id}", getPost).Methods("GET")
	m.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	m.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	m.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":81", m)
}
