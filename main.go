package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"myCity/database"
	"myCity/elements"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DATABASE_URL")
	var err error
	database.DBCon, err = sql.Open("postgres", dbUrl)
	checkErr(err)
	http.HandleFunc("/annuaire", recherchePage)
	http.HandleFunc("/recherche/ville", montreVille)
	http.HandleFunc("/savemap/", saveMap)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("views/assets"))))
	err = http.ListenAndServe(":"+port, nil)
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
