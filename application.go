package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"

	//"text/template"

	_ "github.com/lib/pq"
)

// func result(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	msg := r.FormValue("msg")
// 	t, _ := template.ParseFiles("template2.html")
// 	//データの受け渡し
// 	t.Execute(w, msg)
// }

func main() {
	http.HandleFunc("/create", create)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/", dbtest)
	port := os.Getenv("PORT")
	if port == "" {
		port = "5432"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}

}

type User struct {
	ID   int
	Name string
}

type UserList []User

func dbtest(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "user = lbfjmvpuhlkgqt password = 57d0c849762cceda81669b7a700fd1fd42ed0e9038d07c612efe60cba4c76d41  dbname = d8ro88e9kahr66 host = ec2-34-225-159-178.compute-1.amazonaws.com sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	var ul UserList

	rows, err := db.Query("SELECT * FROM testtable")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	//rowsのNextが入っているか

	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name)
		ul = append(ul, u)
	}

	t, _ := template.ParseFiles("template3.html")
	t.Execute(w, ul)

}

func create(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "user = lbfjmvpuhlkgqt password = 57d0c849762cceda81669b7a700fd1fd42ed0e9038d07c612efe60cba4c76d41  dbname = d8ro88e9kahr66 host = ec2-34-225-159-178.compute-1.amazonaws.com sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS testtable ( id integer, name varchar(32) )")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO testtable (id, name) VALUES (1, 'flamingo')")
	if err != nil {
		log.Fatal(err)
	}

	// rows, err := db.Query("SELECT * FROM testtable")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// //rowsのNextが入っているか
	// var u User
	// for rows.Next() {
	// 	rows.Scan(&u.ID, &u.Name)
	// 	fmt.Fprintln(w, u.ID, u.Name)
	// }

	// defer rows.Close()
	t, _ := template.ParseFiles("template.html")
	t.Execute(w, nil)
}

func delete(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("postgres", "user = lbfjmvpuhlkgqt password = 57d0c849762cceda81669b7a700fd1fd42ed0e9038d07c612efe60cba4c76d41  dbname = d8ro88e9kahr66 host = ec2-34-225-159-178.compute-1.amazonaws.com sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	_, err = db.Exec("DROP TABLE testtable")
	if err != nil {
		log.Fatal(err)
	}
	t, _ := template.ParseFiles("template2.html")
	t.Execute(w, nil)

	// rows, err := db.Query("SELECT * FROM testtable")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// var u User
	// for rows.Next() {
	// 	rows.Scan(&u.ID, &u.Name)
	// 	fmt.Fprintln(w, u.ID, u.Name)
	// }

	// defer rows.Close()
	//"user = postgres password = posflam  dbname = test host = localhost sslmode=disable"
}
