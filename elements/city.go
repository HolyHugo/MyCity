package elements

import (
	"database/sql"
)

type City struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Owner       string `json:"owner"`
	Progression int    `json:"progression"`
	Ressources  []Ressource
}

func GetCity(name string, owner string) (*City, error) {
	var city City
	db, err := sql.Open("sqlite3", "./MyCity.db")
	checkErr(err)
	// query
	err = db.QueryRow("SELECT name,owner, id, progression FROM cities WHERE name = $name AND owner = $owner ", name, owner).Scan(&city.Name, &city.Owner, &city.Id, &city.Progression)
	if err == sql.ErrNoRows {
		return &City{}, nil
	} else {
		RessourcesList, err := GetRessources(city.Id)
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
