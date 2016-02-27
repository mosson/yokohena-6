package main

import (
	"fmt"
	"lesson2/tree"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	t.templ.Execute(w, nil)
}

func main() {
	tree := tree.New(3)

	http.Handle("/", &templateHandler{filename: "index.html"})
	http.Handle("/api", tree)

	fmt.Println("サーバーを開始します: ポート :8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
