package go_sat_solver

import "math"

type orClause []VariableAssignment

type simpleSatCnf struct {
	clauses []orClause
}

func (f simpleSatCnf) IsSatisfied() bool {
	return len(f.clauses) == 0
}

func (f simpleSatCnf) IsContradiction() bool {
	for _, c := range f.clauses {
		if len(c) == 0 {
			return true
		}
	}
	return false
}

type variableAssignment map[Variable]bool

// putAll adds all the assignments to the map
func (t variableAssignment) putAll(toAdd map[Variable]bool) {
	m := map[Variable]bool(t)
	for k, v := range toAdd {
		m[k] = v
	}
}

func (f simpleSatCnf) UnitPropagate() (CnfFormula, map[Variable]bool, []Variable) {
	assignments := variableAssignment(make(map[Variable]bool))
	formula := f
	for {
		var a map[Variable]bool
		var conflicts []Variable
		formula, a, conflicts = formula.unitPropagate1()
		if len(a) == 0 || len(conflicts) != 0 {
			return formula, map[Variable]bool(assignments), conflicts
		} else {
			assignments.putAll(a)
		}
	}
}

// unitPropagate1 assigns singleton clauses to the values that they must be.
// Unlike UnitPropagate, this method does not recurse and assign unit variables again
func (f simpleSatCnf) unitPropagate1() (simpleSatCnf, map[Variable]bool, []Variable) {
	assignments, conflicts := f.getUnitPropagateAssignments()
	return f.assignVariables(assignments), assignments, conflicts
}

// getUnitPropagateAssignments figures out what variables have been pinned down
// into unit clauses and need some assignment to work.
func (f simpleSatCnf) getUnitPropagateAssignments() (assignments map[Variable]bool, conflicts []Variable) {
	assignments = make(map[Variable]bool)
	for _, c := range f.clauses {
		if len(c) == 1 {
			assignedVar := c[0]
			variable := assignedVar.Variable()
			prevAssignment, ok := assignments[variable]
			if ok  {
				if prevAssignment == assignedVar.Value() && !containsVariable(conflicts, variable) {
					conflicts = append(conflicts, variable)
				}
			} else {
				assignments[variable] = assignedVar.Value()
			}
		}
	}
	// TODO: sort conflicts
	for _, con := range conflicts {
		delete(assignments, con)
	}
	return assignments, conflicts
}

// containsVariable returns true if the second argument is in the first slice
func containsVariable(variables []Variable, v Variable) bool {
	for _, s := range variables {
		if s == v {
			return true
		}
	}
	return false
}

func (f simpleSatCnf) PureVariableAssign() (CnfFormula, map[Variable]bool) {
	// TODO: implement.  Current implementation will produce correct result, but slower
	return f, nil
}

// Selects a variable to guess on.
// Precondition: the clause is not unsatisfyible or emtpy (satisfied)
func (f simpleSatCnf) ChooseVariable() Variable {
	// The first var of the first clause that has the fewest number of variables
	minLength := math.MaxInt64
	var v Variable
	for _, c := range f.clauses {
		if 0 < len(c) && len(c) < minLength {
			minLength = len(c)
			v = c[0].Variable()
		}
	}
	return v
}

func (f simpleSatCnf) AssignVariables(assignments map[Variable]bool) CnfFormula {
	return f.assignVariables(assignments)
}

func (f simpleSatCnf) assignVariables(assignments map[Variable]bool) simpleSatCnf {
	if len(assignments) == 0 {
		return f
	}
	newClauses := make([]orClause, 0, len(f.clauses))
clauseLoop:
	for _, c := range f.clauses {
		var varsToRemoveFromClause []Variable
		for _, negVar := range c {
			variable := negVar.Variable()
			if ass, ok := assignments[variable]; ok {
				if ass == negVar.Value() {
					// clause satified.  Don't add clause to the assignment
					continue clauseLoop
				} else {
					varsToRemoveFromClause = append(varsToRemoveFromClause, variable)
				}
			}
		}
		// TODO: sort varsToRemoveFromClause
		newOr := removeVars(c, varsToRemoveFromClause)
		newClauses = append(newClauses, newOr)
	}
	return simpleSatCnf{newClauses}
}

// Creates a copy of an or clause with variables removed.
func removeVars(or orClause, remove []Variable) orClause {
	if len(remove) == 0 {
		return or
	}
	resultClause := make([]VariableAssignment, 0, len(or)-len(remove))
	// TODO: do this with merging sorted lists
	set := make(map[Variable]VariableAssignment)
	for _, v := range or {
		set[v.Variable()] = v
	}
	for _, v := range remove {
		delete(set, v)
	}
	for _, v := range set {
		resultClause = append(resultClause, v)
	}
	// todo: sort or
	return orClause(resultClause)
}
