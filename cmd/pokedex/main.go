package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

var all []pokemon

func init() {
	data, err := ioutil.ReadFile("data/pokemon.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &all)
	if err != nil {
		panic(err)
	}
}

func main() {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/dex", http.StatusFound)
	})
	http.HandleFunc("/dex", func(w http.ResponseWriter, r *http.Request) {
		filter := r.URL.Query().Get("filter")
		if filter != "" {
			var ps []pokemon
			for _, p := range all {
				if strings.Contains(strings.ToLower(p.Name.English), strings.ToLower(filter)) {
					ps = append(ps, p)
				}
			}
			tmpl.Execute(w, ps)
		} else {
			tmpl.Execute(w, all)
		}
	})
	// http.HandleFunc("/action", func(w http.ResponseWriter, r *http.Request) {
	// 	bs, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	filter := string(bs)
	// 	http.Redirect(w, r, fmt.Sprintf("/dex?filter=%s", filter), http.StatusFound)
	// })
	fmt.Println("Serving http://localhost:8080/dex")
	http.ListenAndServe(":8080", nil)
}
