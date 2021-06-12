package main

import (
	"app/db"
	"app/hazards"
	"app/structures"
	"fmt"
	"log"
)

var depth *hazards.DepthEvent = new(hazards.DepthEvent)

func main() {

	var run bool = false

	// Read in Damage Curves from JSON
	ddfs, err := structures.LoadCurves("structures/coastal-ddfs.json")
	if err != nil {
		log.Fatal(err)
	}

	// Read in Structure Data from Postgres
	dbConn := db.Init()
	buildingData, err := db.QueryBuildingAttributes(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	if run {

		// Iterate over the buildings and compute damages
		for i := 0; i < len(buildingData); i++ {

			// Read in Structure data from database
			var ssd structures.StructureSimpleDeterministic
			ssd = buildingData[i]

			// Pair DDF curve from curves database (JSON)
			err := structures.GetDeterministicCurve(ddfs, &ssd)
			if err != nil {
				fmt.Println(err)
				// log.Fatal(err) // Need to turn this off when all curves in place
			} else {

				// Get the recorded event depth
				depth.SetDepth(ssd.HazardDepth)

				// Calculate Loss
				structures.ComputeConsequences2(depth, &ssd)
				if err != nil {
					log.Fatal(err)
				}

				// // Upsert Loss to database
				// err = db.UpsertBuildingLoss(ssd, dbConn)
				// if err != nil {
				// 	log.Fatal(err)
				// }
			}

		}

	}
	fmt.Println("All Done!")
}
