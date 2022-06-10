package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template.html")
	t.Execute(w, nil)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template2.html")
	t.Execute(w, nil)
}

func main() {
	fmt.Println("サーバースタート([Ctrl] + [C]で終了)")
	http.HandleFunc("/", handler)
	http.HandleFunc("/temp", handler2)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
