package db

import (
	"app/structures"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var buildingAttributes string = `SELECT uid, ST_X(ST_Centroid(geom)) as x, ST_Y(ST_Centroid(geom)) as y, foundation, bldg_val FROM geo_context.buildings LIMIT 10;`

// type Building struct {
// 	UID        string  `json:"uid" db:"uid"`
// 	X          float64 `json:"x" db:"x"`
// 	Y          float64 `json:"y" db:"y"`
// 	Foundation string  `json:"foundation" db:"foundation"`
// 	BldValue   float64 `json:"bldg_val" db:"bldg_val"`
// }

// func GetBuildingAttributes(db *sqlx.DB) []Building {

// 	rows, err := db.Queryx(buildingAttributes)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer rows.Close()
// 	result := make([]Building, 0)
// 	for rows.Next() {
// 		var b Building
// 		err = rows.StructScan(&b)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		result = append(result, b)
// 	}
// 	return result
// }

func GetBuildingAttributes(db *sqlx.DB) []structures.StructureSimpleDeterministic {

	rows, err := db.Queryx(buildingAttributes)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()
	result := make([]structures.StructureSimpleDeterministic, 0)
	for rows.Next() {
		var s structures.StructureSimpleDeterministic
		err = rows.StructScan(&s)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, s)
	}
	return result
}
