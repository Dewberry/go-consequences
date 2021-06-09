package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var buildingAttributes string = `SELECT foundation, bldg_val FROM geo_context.buildings LIMIT 25;`

type Building struct {
	Foundation string  `json:"foundation" db:"foundation"`
	BldValue   float64 `json:"bldg_val" db:"bldg_val"`
}

func GetBuildingAttributes(db *sqlx.DB) []Building {

	rows, err := db.Queryx(buildingAttributes)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()
	result := make([]Building, 0)
	for rows.Next() {
		var b Building
		err = rows.StructScan(&b)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, b)
	}
	return result
}
