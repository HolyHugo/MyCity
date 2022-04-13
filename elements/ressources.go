package elements

import (
	"database/sql"
)

type Ressource struct {
	Type       string `json:"type"`
	IndexBoard string `json:"indexBoard"`
}

func GetRessources(city_reference int) ([]Ressource, error) {
	var RessourcesList []Ressource
	db, err := sql.Open("sqlite3", "./MyCity.db")
	checkErr(err)
	// query
	rows, err := db.Query("SELECT indexBoard, ressourceType FROM ressourceNodes WHERE ref_city_id = $city_reference ORDER BY indexBoard", city_reference)
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		var ressource Ressource
		err = rows.Scan(&ressource.IndexBoard, &ressource.Type)
		if err != nil {
			// handle this error
			panic(err)
		}
		RessourcesList = append(RessourcesList, ressource)
	}
	// get any error encountered during iteration
	err = rows.Err()
	checkErr(err)
	return RessourcesList, nil
}
