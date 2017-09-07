package go_sat_solver

type Expr interface {
	Children() []Expr
}

type And struct {
	clauses []Expr
}

type Or struct {
	clauses []Expr
}

type Not struct {
	child [1]Expr
}

type Symbol struct {
	name string
}

func (e And) Children() []Expr { return e.clauses }
func (e Or) Children() []Expr { return e.clauses }
func (e Not) Children() []Expr { return e.child[:] }
func (e Symbol) Children() []Expr { return nil }
