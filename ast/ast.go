package ast

import (
	"fmt"
)

type (
	TypeAtom int
	Atom     interface {
		Type() TypeAtom
		String() string
		Eq(a Atom) bool
	}

	Map map[string]SExpr

	Pair struct {
		Head *SExpr
		Tail *SExpr
	}

	SExpr struct {
		Pair Pair
		Atom Atom
		Map Map
	}
)

const (
	SymAtom TypeAtom = iota + 1
	RatAtom
)

var (
	Nil  SExpr
	True SExpr = NewSExprAtom(NewSym("t"))
)

func (sexpr SExpr) String() string {
	if sexpr.Atom == nil {
		if sexpr.Pair.Head == nil {
			return "()"
		}

		return fmt.Sprintf("(%s %s)",
			sexpr.Pair.Head,
			sexpr.Pair.Tail,
		)
	}

	return sexpr.Atom.String()
}

func NewSExprAtom(a Atom) SExpr {
	return SExpr{
		Atom: a,
	}
}

func NewSExprPair(a, b SExpr) SExpr {
	return SExpr{
		Pair: Pair{
			Head: &a,
			Tail: &b,
		},
	}
}

func NewList(elems ...SExpr) SExpr {
	var last = Nil
	for i := len(elems)-1; i >= 0; i-- {
		last = NewSExprPair(elems[i], last)
	}
	return last
}

func NewAtomList(atoms ...Atom) SExpr {
	var last SExpr = Nil
	for i := len(atoms) - 1; i >= 0; i-- {
		last = NewSExprPair(NewSExprAtom(atoms[i]), last)
	}
	return last
}

func Eq(a, b *SExpr) bool {
	if a == b {
		return true
	}

	if a.Atom != nil && b.Atom != nil {
		if a.Atom.Type() != b.Atom.Type() {
			fmt.Printf("here3")
			return false
		}

		return a.Atom.Eq(b.Atom)
	} else if a.Atom != nil || b.Atom != nil {
		return false
	}

	return Eq(a.Pair.Head, b.Pair.Head) &&
		Eq(a.Pair.Tail, b.Pair.Tail)
}