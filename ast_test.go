package go_sat_solver

import "testing"
import "reflect"

func TestChildrenCount(t *testing.T) {
	cases := []struct {
		expr Expr
		want int
	}{
		{And{[]Expr{Symbol{"x1"}, Symbol{"x2"}}}, 2},
		{Or{[]Expr{Symbol{"x1"}, Symbol{"x2"}, Symbol{"x3"}}}, 3},
		{Not{[1]Expr{Symbol{"x1"}}}, 1},
		{Symbol{"x1"}, 0},
		{Literal{true}, 0},
		{Literal{false}, 0},
	}
	for _, c := range cases {
		got := len(c.expr.Children())
		if got != c.want {
			t.Errorf("len(Children(%#v)) == %v, want %v", c.expr, got, c.want)
		}
	}
}

func TestAstConvenienceBuilders(t *testing.T) {
	cases := []struct {
		expr Expr
		want Expr
	}{
		{and("x1", "x2"), And{[]Expr{Symbol{"x1"}, Symbol{"x2"}}}},
		{and("x1"), Symbol{"x1"}},
		{and(), Literal{true}},
		{or("x1", "x2", "x3"), Or{[]Expr{Symbol{"x1"}, Symbol{"x2"}, Symbol{"x3"}}}},
		{or("x1"), Symbol{"x1"}},
		{or(), Literal{false}},
		{not("x1"), Not{[1]Expr{Symbol{"x1"}}}},
		{not(true), Not{[1]Expr{Literal{true}}}},
	}
	for _, c := range cases {
		got := c.expr
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Builder made %#v, want %#v", got, c.want)
		}
	}
}
