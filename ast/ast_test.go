package ast_test

import (
	"testing"

	"github.com/madlambda/water/ast"
	"github.com/madlambda/water/ast/atom"
)

func TestFmtSExpr(t *testing.T) {
	for _, tc := range []struct{
		Root ast.SExpr
		Fmt string
	} {
		{
			Root: ast.SExpr{
				Atom: atom.NewInt(0),
			},
			Fmt: "0/1",
		},
		{
			Root: ast.SExpr{
				Atom: atom.NewInt(100),
			},
			Fmt: "100/1",
		},
		{
			Root: ast.SExpr{
				Pair: ast.Pair{
					Head: &ast.SExpr{
						Atom: atom.NewInt(100),
					},
					Tail: &ast.SExpr{
						Atom: atom.NewInt(1),
					},
				},
			},
			Fmt: "(100/1 . 1/1)",
		},
		{
			Root: ast.SExpr{
				Atom: atom.NewSym("lambda"),
			},
			Fmt: "lambda",
		},
	} {
		if tc.Fmt != tc.Root.String() {
			t.Fatalf("differs: '%s' != '%s'",
				tc.Fmt, tc.Root.String())
		}
	}
}