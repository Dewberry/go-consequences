package structures

import (
	"app/hazards"
	"app/paireddata"
	"errors"
)

// StructureSimpleDeterministic is a paired down version of struct StructureDeterministic
type StructureSimpleDeterministic struct {
	Name                        string
	DamCat                      string
	OccType                     OccupancyTypeDeterministic
	StructVal, ContVal, FoundHt float64
}

// StructureSimpleDeterministic computes loss at a single structure
func (ssd StructureSimpleDeterministic) Compute(d hazards.HazardEvent) (float64, float64, error) {
	return computeSingleConsequence(d, ssd)
}

func computeSingleConsequence(e hazards.HazardEvent, ssd StructureSimpleDeterministic) (float64, float64, error) {
	if e.Has(hazards.Depth) {
		depthAboveFFE := e.Depth() - ssd.FoundHt
		damagePercent := ssd.OccType.GetStructureDamageFunctionForHazard(e).SampleValue(depthAboveFFE) / 100
		damageValue := damagePercent * ssd.StructVal

		cdamagePercent := ssd.OccType.GetContentDamageFunctionForHazard(e).SampleValue(depthAboveFFE) / 100
		cdamageValue := cdamagePercent * ssd.StructVal

		return damageValue, cdamageValue, nil
	}
	return 0, 0, errors.New("Verify depth is provided")
}

func CustomCurve() OccupancyTypeStochastic {

	structurexs := []float64{-4.0, -3.0, -2.0, -1.0, 0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0}
	structureys := []float64{0, 0, 0, 0, 11, 29, 38, 44, 51, 56, 63, 66, 71, 75, 77, 79, 81, 84, 86, 88, 89}
	var structuredamagefunction = paireddata.PairedData{Xvals: structurexs, Yvals: structureys}

	contentxs := []float64{-4.0, -3.0, -2.0, -1.0, 0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0}
	contentys := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var contentdamagefunction = paireddata.PairedData{Xvals: contentxs, Yvals: contentys}

	sm := make(map[hazards.Parameter]interface{})
	var sdf = DamageFunctionFamilyStochastic{DamageFunctions: sm}

	cm := make(map[hazards.Parameter]interface{})
	var cdf = DamageFunctionFamilyStochastic{DamageFunctions: cm}

	//Default hazard.
	sdf.DamageFunctions[hazards.Default] = structuredamagefunction
	cdf.DamageFunctions[hazards.Default] = contentdamagefunction

	//Depth, salinity hazard.
	sdf.DamageFunctions[hazards.Depth|hazards.Salinity] = structuredamagefunction
	cdf.DamageFunctions[hazards.Depth|hazards.Salinity] = contentdamagefunction

	return OccupancyTypeStochastic{Name: "VACurve", StructureDFF: sdf, ContentDFF: cdf}
}
