package main

import (
	"app/db"
	"app/hazards"
	"app/structures"
	"fmt"
	"log"
	"time"
)

var depth *hazards.DepthEvent = new(hazards.DepthEvent)

func main() {

	var run bool = true
	st := time.Now()

	if run {

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

		fmt.Println("Total Depths:", len(buildingData))
		ssdBatch := make([]structures.StructureSimpleDeterministicResult, 0)
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

				ssdResult := structures.StructureSimpleDeterministicResult{
					FID:                    ssd.FID,
					Epoch:                  ssd.Epoch,
					Event:                  ssd.Event,
					StructureDamagePercent: ssd.StructureDamagePercent,
					ContentDamagePercent:   ssd.ContentDamagePercent,
					StructureDamageValue:   ssd.StructureDamageValue,
					ContentDamageValue:     ssd.ContentDamageValue,
				}

				ssdBatch = append(ssdBatch, ssdResult)

				// fmt.Println(ssd.ContentDamagePercent)
				// Upsert Loss to database

				if i%3000 == 0 {
					err = db.UpsertBuildingBatchLoss(ssdBatch, dbConn)
					if err != nil {
						log.Fatal(err)
					}
					ssdBatch = make([]structures.StructureSimpleDeterministicResult, 0)

					pctComplete := float64(i) * 100.0 / float64(len(buildingData))
					message := fmt.Sprintf("%.2f Percent Complete in %v", pctComplete, time.Since(st))
					fmt.Println(message)
				}

			}

		}

	}
	fmt.Println("All Done!")
}
