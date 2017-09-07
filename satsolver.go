package go_sat_solver

type Variable uint

type VariableAssignment int

func (s VariableAssignment) Variable() Variable {
	if s.Value() {
		return Variable(s)
	} else {
		return Variable(-s)
	}
}

func (s VariableAssignment) Value() bool {
	return s > 0
}

type CnfFormula interface {
	IsSatisfied() bool
	IsContradiction() bool
	UnitPropagate() (CnfFormula, map[Variable]bool, []Variable)
	PureVariableAssign() (CnfFormula, map[Variable]bool)
	ChooseVariable() Variable
	AssignVariables(assignement map[Variable]bool) CnfFormula
}
