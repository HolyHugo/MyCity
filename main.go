package main

import (
	"html/template"
	"myCity/elements"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/annuaire", recherchePage)
	http.HandleFunc("/recherche/ville", montreVille)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("views/assets"))))
	err := http.ListenAndServe(":9090", nil)
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

func checkErr(err error) {
	if err != nil  {
		panic(err)
	}
}
