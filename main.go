package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("サーバースタート([Ctrl] + [C]で終了)")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html>\n<body>\n")
	fmt.Fprintf(w, "<h1>httpサーバーへようこそ</h1>\n")
	fmt.Fprintf(w, "<p>サーバーは:%q</p>\n", r.Host)
	fmt.Fprintf(w, "<p>リモードアドレスは:%q</p>\n", r.RemoteAddr)
	fmt.Fprintf(w, "</body>\n</html>\n")
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
}
