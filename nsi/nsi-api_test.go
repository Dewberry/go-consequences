package nsi

import (
	"fmt"
	"math"
	"sync"
	"testing"

	"github.com/USACE/go-consequences/census"
)

func TestNsiByFips(t *testing.T) {
	var fips string = "15005" //Kalawao county (smallest county in the us by population)
	structures := GetByFips(fips)
	if len(structures.Features) != 101 {
		t.Errorf("GetByFips(%s) yeilded %d structures; expected 101", fips, len(structures.Features))
	}
}
func TestNsiStatsByFips(t *testing.T) {
	var fips string = "15005" //Kalawao county (smallest county in the us by population)stats?bbox=-81.58418,30.25165,-81.58161,30.26939,-81.55898,30.26939,-81.55281,30.24998,-81.58418,30.25165
	stats := GetStatsByFips(fips)
	fmt.Println(stats)
	if stats.SumStructVal != 101.111 {
		t.Errorf("GetByFips(%s) yeilded %f structures; expected 101", fips, stats.SumStructVal)
	}
}
func TestNsiByFipsStream(t *testing.T) {
	var fips string = "15005" //Kalawao county (smallest county in the us by population)
	index := 0
	GetByFipsStream(fips, func(str NsiFeature) {
		index++
	})
	if index != 101 {
		t.Errorf("GetByFipsStream(%s) yeilded %d structures; expected 101", fips, index)
	}
}
func TestNsiByFipsStream_MultiState(t *testing.T) {
	f := census.StateToCountyFipsMap()
	var wg sync.WaitGroup
	wg.Add(len(f))
	index := 0
	for ss := range f {
		go func(sfips string) {
			defer wg.Done()
			GetByFipsStream(sfips, func(str NsiFeature) {
				index++
			})
			fmt.Println("Completed " + sfips)
		}(ss)
	}
	wg.Wait()
	if index != 109406858 {
		t.Errorf("GetByFipsStream(%s) yeilded %d structures; expected 109,406,858", "all states", index)
	} else {
		fmt.Println("Completed 109,406,858 structures")
	}
}
func TestNsi_FL_FoundationTypes(t *testing.T) {
	f := []string{"12"}
	foundationTypes := make(map[string]int64)
	for _, ss := range f {
		GetByFipsStream(ss, func(str NsiFeature) {
			val, ok := foundationTypes[str.Properties.FoundType]
			if ok {
				v2 := val + 1
				foundationTypes[str.Properties.FoundType] = v2
			} else {
				foundationTypes[str.Properties.FoundType] = 1
			}

		})
	}
	fmt.Println(foundationTypes)
}
func TestNsiByBbox(t *testing.T) {
	var bbox string = "-81.58418,30.25165,-81.58161,30.26939,-81.55898,30.26939,-81.55281,30.24998,-81.58418,30.25165"
	structures := GetByBbox(bbox)
	if len(structures.Features) != 1959 {
		t.Errorf("GetByBox(%s) yeilded %d structures; expected 1959", bbox, len(structures.Features))
	}
}
func TestNsiStatsByBbox(t *testing.T) {
	var bbox string = "-81.58418,30.25165,-81.58161,30.26939,-81.55898,30.26939,-81.55281,30.24998,-81.58418,30.25165"
	stats := GetStatsByBbox(bbox)
	diff := stats.SumStructVal - 953459824.285892
	if math.Abs(diff) > 0.0000009 {
		t.Errorf("GetByBox(%s) yeilded structure value of %f; expected 953459824.285892", bbox, stats.SumStructVal)
	}
}
func TestNsiByBboxStream(t *testing.T) {
	var bbox string = "-81.58418,30.25165,-81.58161,30.26939,-81.55898,30.26939,-81.55281,30.24998,-81.58418,30.25165"
	index := 0
	GetByBboxStream(bbox, func(str NsiFeature) {
		index++
	})
	if index != 1959 {
		t.Errorf("GetByBoxStream(%s) yeilded %d structures; expected 1959", bbox, index)
	}
}
func TestNSI_FIPS_CA_ERRORS(t *testing.T) {
	f := census.StateToCountyFipsMap()
	var wg sync.WaitGroup
	counties := f["06"]
	fails := make([]string, 0)
	wg.Add(len(counties))
	for _, ccc := range counties {
		go func(county string) {
			defer wg.Done()
			structures := GetByFips(county)
			if len(structures.Features) == 0 {
				fails = append(fails, county)
				//t.Errorf("GetByFips(%s) yeilded %d structures; expected more than zero", county, len(structures))
			}
		}(ccc)
	}
	wg.Wait()
	if len(fails) > 0 {
		s := "Counties: "
		for _, f := range fails {
			s += f + ", "
		}
		t.Errorf("There were %d failures of %d total counties, failed counties were: %s", len(fails), len(counties), s)
	}
}
