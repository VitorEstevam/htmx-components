package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Restaurant struct {
	Restaurant string `json:"restaurant"`
	Menu       []Menu `json:"menu"`
}
type Menu struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}

func main2() {
	var restaurant Restaurant
	err := json.Unmarshal([]byte(mockedData), &restaurant)
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/listMenu", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("components/listMenu.html"))
		tmpl.Execute(w, restaurant)
	})

	http.HandleFunc("/newDish", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			tmpl := template.Must(template.ParseFiles("components/newDish.html"))
			tmpl.Execute(w, restaurant)
		case http.MethodPost:
			var m map[string]string

			err := json.NewDecoder(r.Body).Decode(&m)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(m)

			tmpl := template.Must(template.ParseFiles("components/newDish.html"))
			tmpl.Execute(w, restaurant)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}
}
