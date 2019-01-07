package posts

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var templateCache = template.Must(template.ParseFiles(
	"tmpl/home.tmpl",
		"tmpl/post.tmpl",
		"tmpl/post-list.tmpl"))

type Post struct {
	ID string
	Title string
	Date string
	Body template.HTML
}

type PostList []*Post

func (post *Post) save() error {
	filename := "posts/" + post.ID + ".post"
	contents := []byte(
		post.Title + "\n" +
			post.Date + "\n" +
			string(post.Body))

	return ioutil.WriteFile(filename, contents, 0600)
}

func (post *Post) RenderAs(tmplName string, w http.ResponseWriter) {
	filename := tmplName + ".tmpl"
	err := templateCache.ExecuteTemplate(w, filename, post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (posts PostList) RenderAs(tmplName string, w http.ResponseWriter) {
	filename := tmplName + ".tmpl"
	err := templateCache.ExecuteTemplate(w, filename, posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAllPosts() PostList {
	var ids []string

	err := filepath.Walk("posts",
		func (path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".post" {
				ids = append(ids, path[6:len(path)-5])
			}
			return nil
		})

	if err != nil {
		panic(err)
	}

	var posts PostList

	for _, id := range ids {
		post, err := LoadPost(id)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	return posts
}

func LoadPost(id string) (*Post, error) {
	filename := "posts/" + id + ".post"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	splitContents := strings.Split(string(contents), "\n")

	title := splitContents[0]
	date := splitContents[1]
	body := template.HTML(strings.Join(splitContents[2:], "\n"))

	return &Post{id, title, date, body}, nil
}
