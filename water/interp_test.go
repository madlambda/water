package water_test

import (
	"github.com/madlambda/water/ast"
	"github.com/madlambda/water/water"
	"testing"
)

func noerr(a ast.SExpr, _ error) ast.SExpr { return a }

func newFuncallCode(funcname string, a ast.SExpr) ast.SExpr {
	// constructed by: (cons '<funcname> (cons '<a> '()))
	return noerr(
		water.Cons(nil, ast.NewSExprAtom(ast.NewSym(funcname)),
			noerr(water.Cons(nil, a, water.Nil))))
}

func TestQuote(t *testing.T) {
	for _, tc := range []struct {
		code ast.SExpr
		res  ast.SExpr
		err  error
	}{
		{
			// (quote a)
			code: ast.NewSExprAtom(ast.NewSym("a")),
			res:  ast.NewSExprAtom(ast.NewSym("a")),
		},
		{
			// (quote (a b c))
			code: ast.NewAtomList(
				ast.NewSym("a"),
				ast.NewSym("b"),
				ast.NewSym("c"),
			),
			res: ast.NewAtomList(
				ast.NewSym("a"),
				ast.NewSym("b"),
				ast.NewSym("c"),
			),
		},
	} {
		res, err := water.Quote(nil, tc.code)
		if err != nil {
			t.Fatal(err)
		}

		if !ast.Eq(&res, &tc.res) {
			t.Fatalf("sexpr differs: %s != %s", res, tc.res)
		}
	}
}

func TestIsAtom(t *testing.T) {
	for _, tc := range []struct {
		code ast.SExpr
		res  ast.SExpr
		err  error
	}{
		{
			// (atom? 'a) == t
			code: ast.NewSExprAtom(ast.NewSym("a")),
			res:  ast.True,
		},
		{
			// (atom? '()) == t
			code: ast.Nil,
			res:  ast.True,
		},
		{
			// (atom? '(a)) == '()
			code: ast.NewAtomList(
				ast.NewSym("a"),
			),
			res: ast.Nil,
		},
	} {
		res, err := water.IsAtom(nil, tc.code)
		if err != nil {
			t.Fatal(err)
		}

		if !ast.Eq(&res, &tc.res) {
			t.Fatalf("sexpr differs: %s != %s", res, tc.res)
		}
	}
}