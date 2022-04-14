package elements

import (
	"myCity/database"
)

type Ressource struct {
	Type       string `json:"type"`
	IndexBoard string `json:"indexBoard"`
}

func GetRessources(city_reference int) ([]Ressource, error) {
	var RessourcesList []Ressource
	// query
	rows, err := database.DBCon.Query("SELECT \"indexBoard\", \"ressourceType\" FROM \"ressourceNodes\" WHERE \"ref_city_id\" = $1 ORDER BY \"indexBoard\"", city_reference)
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

func SaveCity(cityId int, nodes map[string]string) (err error) {
	checkErr(err)
	for i, t := range nodes {
		// query
		_, err := database.DBCon.Exec("INSERT INTO \"ressourceNodes\" (\"indexBoard\", \"ressourceType\", \"ref_city_id\") VALUES ($1,$2,$3)", i, t, cityId)
		checkErr(err)
	}
	return err
}
