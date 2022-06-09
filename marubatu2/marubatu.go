package main

import (
	"net/http"
	"text/template"
)

var Judgment string = "×"
var WinOrLose string
var data []string = []string{"", "", "", "", "", "", "", "", ""}

type Val struct {
	Val, Val1, Val2, Val3, Val4, Val5, Val6, Val7, Val8, Val9, Disp string
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template.html")
	v := new(Val)
	v.Disp = "○の手番"
	for i := 0; i < 9; i++ {
		data[i] = ""
	}
	Judgment = "×"
	t.Execute(w, v)
}

func TicTacToe(w http.ResponseWriter, r *http.Request) {

	v := new(Val)
	t, _ := template.ParseFiles("template.html")
	v.Val = r.FormValue("Val")

	if Judgment == "×" {
		Judgment = "○"
	} else if Judgment == "○" {
		Judgment = "×"
	}

	if v.Val == "num1" {
		data[0] = Judgment
	} else if r.FormValue("id") == "2" {
		data[1] = Judgment
	} else if r.FormValue("id") == "3" {
		data[2] = Judgment
	} else if r.FormValue("id") == "4" {
		data[3] = Judgment
	} else if r.FormValue("id") == "5" {
		data[4] = Judgment
	} else if r.FormValue("id") == "6" {
		data[5] = Judgment
	} else if r.FormValue("id") == "7" {
		data[6] = Judgment
	} else if r.FormValue("id") == "8" {
		data[7] = Judgment
	} else if r.FormValue("id") == "9" {
		data[8] = Judgment
	}

	v.Val1 = data[0]
	v.Val2 = data[1]
	v.Val3 = data[2]
	v.Val4 = data[3]
	v.Val5 = data[4]
	v.Val6 = data[5]
	v.Val7 = data[6]
	v.Val8 = data[7]
	v.Val9 = data[8]

	if data[0] == data[1] && data[0] == data[2] {
		WinOrLose = data[0]
	} else if data[3] == data[4] && data[3] == data[5] {
		WinOrLose = data[3]
	} else if data[6] == data[7] && data[6] == data[8] {
		WinOrLose = data[8]
	} else if data[0] == data[3] && data[0] == data[6] {
		WinOrLose = data[0]
	} else if data[1] == data[4] && data[1] == data[7] {
		WinOrLose = data[1]
	} else if data[2] == data[5] && data[2] == data[8] {
		WinOrLose = data[2]
	} else if data[0] == data[4] && data[0] == data[8] {
		WinOrLose = data[0]
	} else if data[2] == data[4] && data[2] == data[6] {
		WinOrLose = data[2]
	} else {
		count := 0
		for i := range data {
			if data[i] == "○" {
				count++
			}
		}
		if count == 5 {
			WinOrLose = "引き分け"
		}
	}

	if len(WinOrLose) == 0 {
		if Judgment == "×" {
			Judgment = "○"
		} else if Judgment == "○" {
			Judgment = "×"
		}

		v.Disp = Judgment + "の手番"

		if Judgment == "×" {
			Judgment = "○"
		} else if Judgment == "○" {
			Judgment = "×"
		}
	} else if WinOrLose == "○" || WinOrLose == "×" {
		v.Disp = WinOrLose + "の勝ち"
	} else {
		v.Disp = WinOrLose
	}

	t.Execute(w, v)
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/Tic-tac-toe", TicTacToe)
	server.ListenAndServe()
}
