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
	DamageCategory         string                     `json:"damage_category"`
	Hazard                 hazards.HazardEvent        `json:"hazard"`
	OccupancyType          OccupancyTypeDeterministic `json:"occupancy_type"`
	StructureValue         float64                    `json:"structure_value" db:"bldg_val"`
	ContentsValue          float64                    `json:"contents_value"`
	Foundation             string                     `json:"foundation" db:"foundation"`
	FoundationHeight       float64                    `json:"foundation_height"`
	StructureDamagePercent float64                    `json:"structure_damage_percent"`
	ContentDamagePercent   float64                    `json:"content_damage_percent"`
	StructureDamageValue   float64                    `json:"structure_damage_value"`
	ContentDamageValue     float64                    `json:"content_damage_value"`
}

// // StructureSimpleDeterministic computes loss at a single structure
// func (ssd StructureSimpleDeterministic) Location() geography.Location {
// 	return geography.Location{X: 12011000.000, Y: 3870500.000, SRID: "2284"}
// }

func ComputeConsequences2(e hazards.HazardEvent, s *StructureSimpleDeterministic) {

	if e.Has(hazards.Depth) {
		depthAboveFFE := e.Depth() - s.FoundationHeight
		structureDamagePercent := s.OccupancyType.GetStructureDamageFunctionForHazard(e).SampleValue(depthAboveFFE) / 100
		contentDamagePercent := s.OccupancyType.GetContentDamageFunctionForHazard(e).SampleValue(depthAboveFFE) / 100

		s.StructureDamagePercent = structureDamagePercent * s.StructureValue
		s.ContentDamagePercent = contentDamagePercent * s.ContentsValue
		s.StructureDamageValue = structureDamagePercent * s.StructureValue
		s.ContentDamageValue = contentDamagePercent * s.ContentsValue
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

func GetDeterministicCurve(ddfs DDFS, name string) (OccupancyTypeDeterministic, error) {
	var ddf DDF
	sm := make(map[hazards.Parameter]interface{})
	sdf := DamageFunctionFamilyStochastic{DamageFunctions: sm}

	for i := 0; i < len(ddfs.Curves); i++ {

		useDDF := ddfs.Curves[i]

		switch useDDF.Name {
		case name:
			ddf = useDDF
		default:
			continue
		}
	}

	if ddf.Name == "" {
		errorMessage := fmt.Sprintf("Curve not found --> %s", name)
		return OccupancyTypeDeterministic{}, errors.New(errorMessage)
	}

	cm := make(map[hazards.Parameter]interface{})
	cdf := DamageFunctionFamilyStochastic{DamageFunctions: cm}

	sdf.DamageFunctions[hazards.Depth] = paireddata.PairedData{Xvals: ddf.Depth, Yvals: ddf.StructureDamage}
	cdf.DamageFunctions[hazards.Depth] = paireddata.PairedData{Xvals: ddf.Depth, Yvals: ddf.ContentDamage}

	stochasticCurves := OccupancyTypeStochastic{Name: ddf.Name, StructureDFF: sdf, ContentDFF: cdf}
	return stochasticCurves.CentralTendency(), nil
}

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
