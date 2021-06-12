package structures

import (
	"app/hazards"
	"app/paireddata"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// StructureSimpleDeterministic is a paired down version of  struct StructureDeterministic
type StructureSimpleDeterministic struct {
	FID                    string                     `json:"fid" nsi:"fd_id" db:"uid"`
	X                      float64                    `json:"x" db:"x"`
	Y                      float64                    `json:"y" db:"y"`
	Epoch                  string                     `json:"epoch" db:"epoch"`      // e.g. 2020 or 2040
	Event                  string                     `json:"event" db:"event_type"` // e.g. mlw or mhw or 500yr
	DamageCategory         string                     `json:"damage_category" db:"ddf"`
	HazardDepth            float64                    `json:"hazard_depth" db:"depth"`
	Hazard                 hazards.HazardEvent        `json:"hazard"`
	OccupancyType          OccupancyTypeDeterministic `json:"occupancy_type"`
	StructureValue         float64                    `json:"structure_value" db:"structure_value"`
	ContentsValue          float64                    `json:"contents_value" db:"content_value"`
	Foundation             string                     `json:"foundation" db:"foundation"`
	FoundationHeight       float64                    `json:"foundation_height" db:"ffh"`
	StructureDamagePercent float64                    `json:"structure_damage_percent" db:"structure_damage_percent"`
	ContentDamagePercent   float64                    `json:"content_damage_percent" db:"content_damage_percent"`
	StructureDamageValue   float64                    `json:"structure_damage_value" db:"structure_damage_value"`
	ContentDamageValue     float64                    `json:"content_damage_value" db:"content_damage_value"`
}

func ComputeConsequences2(e hazards.HazardEvent, ssd *StructureSimpleDeterministic) {

	if e.Has(hazards.Depth) {
		depthAboveFFE := e.Depth() - ssd.FoundationHeight
		structureDamagePercent := ssd.OccupancyType.GetStructureDamageFunctionForHazard(e).SampleValue(depthAboveFFE) / 100
		contentDamagePercent := ssd.OccupancyType.GetContentDamageFunctionForHazard(e).SampleValue(depthAboveFFE) / 100

		ssd.StructureDamagePercent = structureDamagePercent
		ssd.ContentDamagePercent = contentDamagePercent
		ssd.StructureDamageValue = structureDamagePercent * ssd.StructureValue
		ssd.ContentDamageValue = contentDamagePercent * ssd.ContentsValue
	}
}

// -------------------Read in Damage Curves
type DDF struct {
	Name            string    `json:"name"`
	Depth           []float64 `json:"depth"`
	StructureDamage []float64 `json:"structure_damage"`
	ContentDamage   []float64 `json:"content_damage"`
}

type DDFS struct {
	Data   string
	Curves []DDF
}

func GetDeterministicCurve(ddfs DDFS, ssd *StructureSimpleDeterministic) error {
	var ddf DDF
	sm := make(map[hazards.Parameter]interface{})
	sdf := DamageFunctionFamilyStochastic{DamageFunctions: sm}

	for i := 0; i < len(ddfs.Curves); i++ {

		useDDF := ddfs.Curves[i]

		switch useDDF.Name {
		case ssd.DamageCategory:
			ddf = useDDF
		default:
			continue
		}
	}

	if ddf.Name == "" {
		errorMessage := fmt.Sprintf("Curve not found --> %s", ssd.DamageCategory)
		return errors.New(errorMessage)
	}

	cm := make(map[hazards.Parameter]interface{})
	cdf := DamageFunctionFamilyStochastic{DamageFunctions: cm}

	sdf.DamageFunctions[hazards.Depth] = paireddata.PairedData{Xvals: ddf.Depth, Yvals: ddf.StructureDamage}
	cdf.DamageFunctions[hazards.Depth] = paireddata.PairedData{Xvals: ddf.Depth, Yvals: ddf.ContentDamage}

	stochasticCurves := OccupancyTypeStochastic{Name: ddf.Name, StructureDFF: sdf, ContentDFF: cdf}
	ssd.OccupancyType = stochasticCurves.CentralTendency()
	return nil
}

// func GetDeterministicCurve(ddfs DDFS, name string) (OccupancyTypeDeterministic, error) {
// 	var ddf DDF
// 	sm := make(map[hazards.Parameter]interface{})
// 	sdf := DamageFunctionFamilyStochastic{DamageFunctions: sm}

// 	for i := 0; i < len(ddfs.Curves); i++ {

// 		useDDF := ddfs.Curves[i]

// 		switch useDDF.Name {
// 		case name:
// 			ddf = useDDF
// 		default:
// 			continue
// 		}
// 	}

// 	if ddf.Name == "" {
// 		errorMessage := fmt.Sprintf("Curve not found --> %s", name)
// 		return OccupancyTypeDeterministic{}, errors.New(errorMessage)
// 	}

// 	cm := make(map[hazards.Parameter]interface{})
// 	cdf := DamageFunctionFamilyStochastic{DamageFunctions: cm}

// 	sdf.DamageFunctions[hazards.Depth] = paireddata.PairedData{Xvals: ddf.Depth, Yvals: ddf.StructureDamage}
// 	cdf.DamageFunctions[hazards.Depth] = paireddata.PairedData{Xvals: ddf.Depth, Yvals: ddf.ContentDamage}

// 	stochasticCurves := OccupancyTypeStochastic{Name: ddf.Name, StructureDFF: sdf, ContentDFF: cdf}
// 	return stochasticCurves.CentralTendency(), nil
// }

func LoadCurves(dataFile string) (DDFS, error) {

	var ddfs DDFS
	jsonFile, err := os.Open(dataFile)
	if err != nil {
		return ddfs, nil
	}

	defer jsonFile.Close()
	jsonData, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(jsonData, &ddfs)

	return ddfs, nil

}
