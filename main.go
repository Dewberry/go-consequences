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

	// Read in Damage Curves from JSON
	ddfs, err := structures.LoadCurves("structures/coastal-ddfs.json")
	if err != nil {
		log.Fatal(err)
	}

	// Read in Structure Data from Postgres
	dbConn := db.Init()
	buildingData := db.GetBuildingAttributes(dbConn)

	// Iterate over the buildings and compute damages
	for i := 0; i < len(buildingData); i++ {
		var ssd structures.StructureSimpleDeterministic

		ssd = buildingData[i]

		// Set Manually until fields are in place in the db matching up with curves
		ssd.DamageCategory = "CPFRA_1901"
		ssd.FoundationHeight = 2

		damageCurve, err := structures.GetDeterministicCurve(ddfs, ssd.DamageCategory)
		if err != nil {
			log.Fatal(err)
		}

		ssd.OccupancyType = damageCurve

		depth.SetDepth(10)

		structures.ComputeConsequences2(depth, &ssd)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(ssd.FID, "| ", ssd.StructureValue, ssd.StructureDamageValue, ssd.ContentDamageValue)
	}

}
