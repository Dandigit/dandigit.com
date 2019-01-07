package main

import (
	"log"
	"net/http"
	"./posts"
)

func handlePost(w http.ResponseWriter, request *http.Request) {
	id := request.URL.Path[7:]
	post, err := posts.LoadPost(id)
	if err != nil {
		http.NotFound(w, request)
		return
	}

	post.RenderAs("post", w)
}

func handlePostList(w http.ResponseWriter, request *http.Request) {
	posts.GetAllPosts().RenderAs("post-list", w)
}

func handleHomepage(w http.ResponseWriter, request *http.Request) {
	posts.GetAllPosts().RenderAs("home", w)
}

func main() {
	// Handle all asset requests with a file server
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/posts/all", handlePostList)
	http.HandleFunc("/posts/", handlePost)
	http.HandleFunc("/", handleHomepage)

	log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
}
