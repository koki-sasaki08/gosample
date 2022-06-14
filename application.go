package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

//parsefiles 中身を解読
// func index(w http.ResponseWriter, r *http.Request) {
// 	t, _ := template.ParseFiles("template.html")
// 	list := []string{"Apple", "Banana", "Crocodile"}
// 	//データの受け渡し
// 	t.Execute(w, list)
// }

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template.html")
	//データの受け渡し
	t.Execute(w, nil)
}

func result(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	msg := r.FormValue("msg")
	t, _ := template.ParseFiles("template2.html")
	//データの受け渡し
	t.Execute(w, msg)
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/form", form)
	http.HandleFunc("/result", result)
	http.HandleFunc("/", dbtest)
	server.ListenAndServe()
}

type User struct {
	ID   string
	Name string
}

func dbtest(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgers", "user = postgres password = posflam dbname = dbtest host = localhost")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT * FROM test")
	if err != nil {
		panic(err.Error())
	}

	//rowsのNextが入っているか
	var u User
	for rows.Next() {
		rows.Scan(&u.ID, &u.Name)
		fmt.Fprintln(w, u.ID, u.Name)
	}

	defer rows.Close()
}
