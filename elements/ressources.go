package elements

import (
	"database/sql"
)

type Ressource struct {
	Type      string `json:"type"`
	Visual    string `json:"visual"`
	Buildable bool   `json:"buildable"`
	CoordX    string `json:"coordX"`
	CoordY    string `json:"coordY"`
}

func GetRessources(map_reference int) ([]Ressource, error) {
	var RessourcesList []Ressource
	db, err := sql.Open("sqlite3", "./MyCity.db")
	checkErr(err)
	// query
	rows, err := db.Query("SELECT coordX,coordY, ressourceType, isBuildable,visual FROM ressourceNodes WHERE ref_city_id = $map_reference ORDER BY coordX, coordY", map_reference)
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		var ressource Ressource
		err = rows.Scan(&ressource.CoordX, &ressource.CoordY, &ressource.Type, &ressource.Buildable, &ressource.Visual)
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
