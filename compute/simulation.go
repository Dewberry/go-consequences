package compute

import (
	"fmt"

	"github.com/USACE/go-consequences/consequences"
	"github.com/USACE/go-consequences/hazards"
	"github.com/USACE/go-consequences/nsi"
)

type RequestArgs struct {
	Args interface{}
}
type FipsCodeCompute struct {
	ID         string
	FIPS       string
	HazardArgs interface{}
}
type BboxCompute struct {
	ID         string
	BBOX       string
	HazardArgs interface{}
}
type StatusReportRequest struct {
	ID string
}
type ResultsRequest struct {
	ID string
}
type StructureSimulation struct {
	Structures []consequences.StructureStochastic
	Status     string
	Result     consequences.ConsequenceDamageResult
}
type NSIStructureSimulation struct {
	RequestArgs
	StructureSimulation
}
type Computable interface {
	Compute(args RequestArgs)
}

type ProgressReportable interface {
	ReportProgress() string
}
type SimulationSummary struct {
	RowHeader       string
	StructureCount  int64
	StructureDamage float64
	ContentDamage   float64
}

func (s NSIStructureSimulation) ReportProgress() string {
	return s.Status
}
func (s StructureSimulation) ReportProgress() string {
	return s.Status
}

func (s NSIStructureSimulation) Compute(args RequestArgs) {
	var depthevent = hazards.DepthEvent{Depth: 5.32}
	okd := false
	fips, okfips := args.Args.(FipsCodeCompute)
	if okfips {
		s.Status = "Downloading NSI by fips " + fips.FIPS
		s.Structures = nsi.GetByFips(fips.FIPS)
		depthevent, okd = fips.HazardArgs.(hazards.DepthEvent)
	} else {
		bbox, okbbox := args.Args.(BboxCompute)
		if okbbox {
			s.Status = "Downloading NSI by bbox " + bbox.BBOX
			s.Structures = nsi.GetByBbox(bbox.BBOX)
			depthevent, okd = bbox.HazardArgs.(hazards.DepthEvent)
		}
	}
	s.Status = "Computing Depths"
	//depths
	var d = hazards.DepthEvent{Depth: 5.32}
	if okd {
		d = depthevent
	}

	//ideally get from some sort of source.
	rmap := make(map[string]SimulationSummary)
	s.Status = fmt.Sprintf("Computing Damages %d of %d", 0, len(s.Structures))
	for i, str := range s.Structures {
		r := str.ComputeConsequences(d)
		if val, ok := rmap[str.DamCat]; ok {
			val.StructureCount += 1
			val.StructureDamage += r.Results[0].(float64) //based on convention - super risky
			val.ContentDamage += r.Results[1].(float64)   //based on convention - super risky
		} else {
			rmap[str.DamCat] = SimulationSummary{RowHeader: str.DamCat, StructureCount: 1, StructureDamage: r.Results[0].(float64), ContentDamage: r.Results[1].(float64)}
		}
		s.Status = fmt.Sprintf("Computing Damages %d of %d", i, len(s.Structures))
	}
	header := []string{"Damage Category", "Structure Count", "Total Structure Damage", "Total Content Damage"}
	rows := make([]SimulationSummary, len(rmap))
	results := []interface{}{rows}
	var ret = consequences.ConsequenceDamageResult{Headers: header, Results: results}
	idx := 0
	for _, val := range rmap {
		ret.Results[idx] = val
		idx++
	}
	s.Status = "Complete"
	s.Result = ret
}
