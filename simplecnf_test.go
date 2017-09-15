package go_sat_solver

import "testing"
import "reflect"

var sampleSat1 *simpleSatCnf = &simpleSatCnf{[]orClause{
	orClause{1, 2, 3},
	orClause{-1, -2},
	orClause{-1, -3},
	orClause{1},
	}}
var sampleSat2 *simpleSatCnf = &simpleSatCnf{[]orClause{
	orClause{1, 2, 3},
	orClause{-1, -2},
	orClause{-1, -3},
	orClause{1},
	orClause{2}}}

func TestSimpleSatCnf(t *testing.T) {
	cases := []struct {
		problem *simpleSatCnf
		assignment    map[Variable]bool
		satisfiable bool
	}{
		{sampleSat1, map[Variable]bool{1:true, 2:false, 3:false}, true},
		{sampleSat2, nil, false},
	}
	for _, c := range cases {
		formula, assignment, sat := Dpll(c.problem)
		if c.satisfiable {
			if !sat {
				t.Errorf("No solution found for %v.  Got %v and assignment %v", c.problem, formula, assignment)
			}
			// assuming a unique solution for now
			if !reflect.DeepEqual(map[Variable]bool(assignment), c.assignment) {
				t.Errorf("Different solution found for %v.  Got %v and assignment %#v, wanted %v#", c.problem, formula, assignment, c.assignment)
			}
		} else {
			if sat {
				t.Errorf("Got assignment for unsatisfiable %v: %v", c.problem, assignment)
			}
		}
	}
}
