package main

import (
	"app/db"
	"app/hazards"
	"app/structures"
	"fmt"
	"log"
)

var ssdExample *structures.StructureSimpleDeterministic = new(structures.StructureSimpleDeterministic)
var occType *structures.OccupancyTypeDeterministic
var depth *hazards.DepthEvent = new(hazards.DepthEvent)

func main() {

	ddfCurve := structures.CustomCurve()

	dbConn := db.Init()
	buildingData := db.GetBuildingAttributes(dbConn)

	ssdExample.OccType = ddfCurve.CentralTendency()
	ssdExample.DamCat = "RES"

	var building db.Building

	for i := 0; i < len(buildingData); i++ {

		building = buildingData[i]

		ssdExample.StructVal = building.BldValue
		ssdExample.FoundHt = 4

		depth.SetDepth(4)

		sDamage, cDamage, err := ssdExample.Compute(depth)
		if err != nil {
			log.Print("ERROR: ", err)
		}

		fmt.Println(building.Foundation, "Structure Damage: ", sDamage, "Content Damage: ", cDamage)
	}

}
