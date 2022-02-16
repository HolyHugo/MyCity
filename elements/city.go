package elements

import (
	"database/sql"
)

type City struct {
	Name        string `json:"name"`
	Owner       string `json:"owner"`
	Progression int    `json:"progression"`
	Ressources  []Ressource
}

func GetCity(name string, owner string) (*City, error) {
	var city City
	var map_reference int
	db, err := sql.Open("sqlite3", "./MyCity.db")
	checkErr(err)
	// query
	err = db.QueryRow("SELECT name,owner, id, progression FROM cities WHERE name = $name AND owner = $owner ", name, owner).Scan(&city.Name, &city.Owner, &map_reference, &city.Progression)
	if err != nil && err == sql.ErrNoRows {
		return &City{}, err
	} else {
		RessourcesList, err := GetRessources(map_reference)
		if err == nil && RessourcesList != nil {
			city.Ressources = RessourcesList
		}
		return &city, nil
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
