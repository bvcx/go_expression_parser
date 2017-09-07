package go_sat_solver

import "fmt"

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

type Literal struct {
	value bool
}

func (e And) Children() []Expr { return e.clauses }
func (e Or) Children() []Expr { return e.clauses }
func (e Not) Children() []Expr { return e.child[:] }
func (e Symbol) Children() []Expr { return nil }
func (e Literal) Children() []Expr { return nil }

func toExpr(o interface{}) (r Expr, ok bool) {
	if e, ok := o.(Expr); ok {
		return e, true
	} else if s, ok := o.(string); ok {
		return Symbol{s}, true
	} else if b, ok := o.(bool); ok {
		return Literal{b}, true
	} else {
		return nil, false
	}
}

func toExprOrPanic(o interface{}) Expr {
	if expr, ok := toExpr(o); ok {
		return expr
	} else {
		panic(fmt.Sprintf("Can't convert to an propositional expression: %#v", o))
	}
}

func toExprs(args []interface{}) []Expr {
	result := make([]Expr, len(args), len(args))
	for i, o := range args {
		result[i] = toExprOrPanic(o)
	}
	return result
}

func and(args ...interface{}) Expr {
	exprs := toExprs(args)
	switch len(exprs) {
	case 0:
		return Literal{true}
	case 1:
		return exprs[0]
	default:
		return And{exprs}
	}
}

func or(args ...interface{}) Expr {
	exprs := toExprs(args)
	switch len(exprs) {
	case 0:
		return Literal{false}
	case 1:
		return exprs[0]
	default:
		return Or{exprs}
	}
}

func not(o interface{}) Expr {
	expr := toExprOrPanic(o)
	return Not{[1]Expr{expr}}
}
