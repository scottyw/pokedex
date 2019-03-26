package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

type pokemon struct {
	ID   int `json:"id"`
	Name struct {
		English  string `json:"english"`
		Japanese string `json:"japanese"`
		Chinese  string `json:"chinese"`
	} `json:"name"`
	Type []string `json:"type"`
	Base struct {
		HP        int `json:"HP"`
		Attack    int `json:"Attack"`
		Defense   int `json:"Defense"`
		SpAttack  int `json:"Sp. Attack"`
		SpDefense int `json:"Sp. Defense"`
		Speed     int `json:"Speed"`
	} `json:"base"`
}

var ps []pokemon

func init() {
	data, err := ioutil.ReadFile("data/pokemon.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &ps)
	if err != nil {
		panic(err)
	}
}

func main() {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, ps)
	})
	http.HandleFunc("/action", func(w http.ResponseWriter, r *http.Request) {
		ps = ps[1:]
		http.Redirect(w, r, "/", http.StatusFound)
	})
	fmt.Println("Serving http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
