package elements

import (
	"database/sql"
	"myCity/database"
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
	// query
	err := database.DBCon.QueryRow("SELECT name,owner, id, progression FROM cities WHERE name = $1 AND owner = $2 ", name, owner).Scan(&city.Name, &city.Owner, &city.Id, &city.Progression)
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
