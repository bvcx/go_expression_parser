package go_sat_solver

var possibleValues = [2]bool{false, true}

// Dpll finds a satisfiable assignmet for the boolean SAT formula, or returns nil if it is not satisfiable.
// This algorithm will never give up or restart.
func Dpll(c CnfFormula) (formula CnfFormula, assignments variableAssignment, ok bool) {
	if c.IsSatisfied() {
		return c, nil, true
	}
	if c.IsContradiction() {
		return c, nil, false
	}
	assignments = make(map[Variable]bool)
	c, a, conflicts := c.UnitPropagate()
	assignments.putAll(a)
	if len(conflicts) > 0 {
		return c, assignments, false
	}
	c, a = c.PureVariableAssign()
	assignments.putAll(a)
	s := c.ChooseVariable()
	for _, v := range possibleValues {
		result, a, ok := Dpll(c.AssignVariables(map[Variable]bool{s: v}))
		if ok {
			assignments.putAll(a)
			return result, assignments, true
		}
	}
	return c, assignments, false
}
