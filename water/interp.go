package water

import (
	"fmt"
	"github.com/madlambda/water/ast"
)

type (
	Interp struct {
		Environ ast.SExpr
	}
)

var Nil ast.SExpr

func NewInterp() *Interp {
	return &Interp{
		Environ: setupBuiltins(),
	}
}

func setupBuiltins() ast.SExpr {
	env := make(ast.Map)
	env["quote"] = Quote
	//env["lambda"] = Λ(Lambda)
	env["apply"] = Λ(Apply)
	env["car"] = Λ(Car)
	env["cdr"] = Λ(Cdr)
	env["cons"] = Λ(Cons)
	env["display"] = Λ(Display)
	return ast.SExpr{
		Map: env,
	}
}

// Quote prevents a list from being evaluated.
// Why it's builtin? I'm not sure why..
// (lambda (x) x)
var Quote = ast.NewList(
	ast.NewSExprAtom(ast.NewSym("lambda")),
	ast.NewAtomList(ast.NewSym("x")),
	ast.NewSExprAtom(ast.NewSym("x")),
)

//func IsAtom(_ ast.Map, args ...ast.SExpr) (ast.SExpr, error) {
//	if len(args) != 1 {
///		return Nil, fmt.Errorf("atom? expects one arg")
//	}
//
//	a := args[0]
//	if a.Atom != nil || a == Nil {
//		return ast.True, nil
//	}
//	return ast.Nil, nil
//}

//func Lambda(_ Environ, args ...ast.SExpr) (ast.SExpr, error) {
//	if len(args) != 2 {
//		return nil, fmt.Errorf("LAMBDA expects two arguments")
//	}
//
//	params := args[0]
//	body := args[1]
//}

func Display(_ ast.Map, args ...ast.SExpr) (ast.SExpr, error) {
	for _, s := range args {
		fmt.Printf("%s ", s)
	}
	return ast.SExpr{}, nil
}

func Apply(env ast.Map, args ...ast.SExpr) (ast.SExpr, error) {
	if len(args) < 2 {
		return ast.SExpr{}, fmt.Errorf("apply: requires at least 2 argument")
	}

	fn := args[0].Atom.(ast.Sym)
	args = args[1:]
	fnenv := make(Environ)
	for k, v := range env {
		fnenv[k] = v
	}

	if obj, ok := env[fn.Value]; ok {
		params := Car(env, obj)
		body := Cdr(env, obj)

		for i := 0; i < len(args); i++ {
			p := Car(env, params)
			psym := p.Atom.(Sym).Value
			fnenv[psym] = args[i]
			params = Cdr(env, params)
		}

		return Eval(fnenv, body)
	}

	return ast.SExpr{}, fmt.Errorf("Symbol '%s' not found",
		fn.Value)
}

// Cons constructs a list (SICP implementation)
// (define cons 
// 		(lambda (p q) 
//			(lambda (m) (m p q))))
// AST of code above
var Cons = ast.NewList(
	ast.NewSExprAtom(ast.NewSym("lambda")),
	ast.NewAtomList(ast.NewSym("p"), ast.NewSym("q")),
	ast.NewList(
		ast.NewSExprAtom(ast.NewSym("lambda")),
		ast.NewAtomList(ast.NewSym("m")),
		ast.NewAtomList(
			ast.NewSym("m"),
			ast.NewSym("p"),
			ast.NewSym("q"),
		),
	),
)

// Car gets the head of the list
// (define car 
// 		(lambda (alist)
// 			(alist (lambda (p q) p))))
var Car = ast.NewList(
	ast.NewSExprAtom(ast.NewSym("lambda")),
	ast.NewAtomList(ast.NewSym("alist")),
	ast.NewList(
		ast.NewSExprAtom(ast.NewSym("lambda")),
		ast.NewAtomList(ast.NewSym("p"), ast.NewSym("q")),
		ast.NewSExprAtom(ast.NewSym("p")),
	),
)

// Cdr gets the tail of list
// (lambda (z)
// 		(z (lambda (p q) q)))
var Cdr = ast.NewList(
	ast.NewSExprAtom(ast.NewSym("lambda")),
	ast.NewAtomList(ast.NewSym("z")),
	ast.NewList(
		ast.NewSExprAtom(ast.NewSym("z")),
		ast.NewList(
			ast.NewSExprAtom(ast.NewSym("lambda")),
			ast.NewAtomList(ast.NewSym("p"), ast.NewSym("q")),
			ast.NewSExprAtom(ast.NewSym("q")),
		),
	),
)

func Eval(env Environ, args ...ast.SExpr) (ast.SExpr, error) {
	if len(args) != 1 {
		return Nil, fmt.Errorf("eval expects 1 argument")
	}

	code := args[0]
	if code.Atom != nil {
		at := code.Atom
		if at.Type() == ast.SymAtom {
			varname := at.(Sym).Value
			obj, ok := env[varname]
			if !ok {
				return Nil, fmt.Errorf("variable '%s' not found", varname)
			}
			return o
		}
	}

	fns := *code.Pair.Head
	if fns.Atom == nil {
	}
}