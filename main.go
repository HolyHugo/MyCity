package main

import (
	"encoding/json"
	"html/template"
	"myCity/elements"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/annuaire", recherchePage)
	http.HandleFunc("/recherche/ville", montreVille)
	http.HandleFunc("/savemap/", saveMap)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("views/assets"))))
	err := http.ListenAndServe(":"+port, nil)
	checkErr(err)
}

func recherchePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/recherche.gtpl")
	checkErr(err)
	err = t.Execute(w, nil)
	checkErr(err)
}

func montreVille(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/montreVille.gtpl")
	checkErr(err)
	villeRecherche, err := elements.GetCity(r.FormValue("name"), r.FormValue("owner"))
	checkErr(err)
	err = t.Execute(w, villeRecherche)
	checkErr(err)
}

func saveMap(w http.ResponseWriter, r *http.Request) {
	strId := strings.TrimPrefix(r.URL.Path, "/savemap/")
	r.ParseForm()
	jsonData := r.PostFormValue("data")
	mappedJson := map[string]string{}
	json.Unmarshal([]byte(jsonData), &mappedJson)
	id, err := strconv.Atoi(strId)
	checkErr(err)
	elements.SaveCity(id, mappedJson)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
