package structures

import (
	"app/hazards"
	"errors"
)

// StructureSimpleDeterministic is a paired down version of  struct StructureDeterministic
type StructureSimpleDeterministic struct {
	BaseStructure
	OccType                     OccupancyTypeDeterministic
	StructVal, ContVal, FoundHt float64
}

// StructureSimpleDeterministic computes loss at a single structure
func (ssd StructureSimpleDeterministic) Compute(d hazards.HazardEvent) (float64, error) {
	return computeSingleConsequence(d, ssd)
}

func computeSingleConsequence(e hazards.HazardEvent, ssd StructureSimpleDeterministic) (float64, error) {
	if e.Has(hazards.Depth) {
		depthAboveFFE := e.Depth() - ssd.FoundHt
		damageValue := ssd.OccType.GetStructureDamageFunctionForHazard(e).SampleValue(depthAboveFFE) / 100
		return damageValue, nil
	}
	return 0, errors.New("Verify depth provided")
}
