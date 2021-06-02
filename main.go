package main

import (
	"app/hazards"
	"app/structures"
	"fmt"
	"log"
)

func swapOcctypeMap(m map[string]structures.OccupancyTypeStochastic) map[string]structures.OccupancyTypeDeterministic {
	m2 := make(map[string]structures.OccupancyTypeDeterministic)
	for name, ot := range m {
		m2[name] = ot.CentralTendency()
	}
	return m2
}

var ssdExample *structures.StructureSimpleDeterministic = new(structures.StructureSimpleDeterministic)
var occType *structures.OccupancyTypeDeterministic
var depth *hazards.DepthEvent = new(hazards.DepthEvent)

// var depth *hazards.CoastalEvent = new(hazards.CoastalEvent)

func main() {

	m := structures.OccupancyTypeMap()
	occTypes := swapOcctypeMap(m)

	ssdExample.OccType = occTypes["RES1-1SNB"]
	ssdExample.DamCat = "RES"
	ssdExample.StructVal = 200000
	ssdExample.FoundHt = 4

	depth.SetDepth(4)
	// depth.SetSalinity(true)

	damage, err := ssdExample.Compute(depth)
	if err != nil {
		log.Print("ERROR: ", err)
	}

	fmt.Println("Calculated Damage: ", damage)
}
