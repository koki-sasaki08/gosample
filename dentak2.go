package main

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
)

type Num struct {
	Num1, Num2, Answer                   float64
	Calc, Error1, Error2, Error3, Error4 string
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template.html")
	t.Execute(w, nil)
}

func result(w http.ResponseWriter, r *http.Request) {
	n := new(Num)
	r.ParseForm()
	t, _ := template.ParseFiles("template2.html")

	// if err != nil {
	// 	n.Error1 = "※未入力があります"
	// }
	// err.Error()

	if len(r.Form.Get("Num1")) == 0 {
		n.Error1 = "※左項目が入力されていません。"
	}

	if len(r.Form.Get("Num2")) == 0 {
		n.Error2 = "※右項目が入力されていません。"
	}

	if len(r.Form.Get("Calc")) == 0 {
		n.Error3 = "※演算子が選択されていません。"
	}

	n.Num1, _ = strconv.ParseFloat(r.FormValue("Num1"), 64)
	n.Num2, _ = strconv.ParseFloat(r.FormValue("Num2"), 64)
	n.Calc = r.FormValue("Calc")

	if n.Calc == "add" {
		n.Answer = n.Num1 + n.Num2
	} else if n.Calc == "sub" {
		n.Answer = n.Num1 - n.Num2
	} else if n.Calc == "multi" {
		n.Answer = n.Num1 * n.Num2
	} else if n.Calc == "div" {
		if n.Num2 != 0 {
			n.Answer = math.Round(n.Num1/n.Num2*1000) / 1000
		} else {
			n.Error4 = "undefind"
		}
	}

	//データの受け渡し
	t.Execute(w, n)

}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/result", result)
	http.HandleFunc("/", form)
	// http.HandleFunc("/", form)
	server.ListenAndServe()
}
