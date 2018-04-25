package ast

import (
	"github.com/madlambda/exact"
)

type (
	Rat struct {
		Value exact.Rat
	}

	Sym struct {
		Value string
	}
)



func NewSym(a string) Sym {
	return Sym{
		Value: a,
	}
}

func (a Sym) String() string { return a.Value }
func (a Sym) Type() TypeAtom { return SymAtom }
func (a Sym) Eq(b Atom) bool {
	if b.Type() != SymAtom {
		return false
	}

	bs := b.(Sym)
	return string(a.Value) == string(bs.Value)
}

func NewInt(i int64) Rat {
	var r Rat
	if i < 0 {
		r.Value = exact.NewNegRat(uint64(-i), 1)
	}
	r.Value = exact.NewRat(uint64(i), 1)
	return r
}

func (r Rat) Type() TypeAtom {
	return RatAtom
}

func (r Rat) String() string {
	return r.Value.String()
}

func (r Rat) Eq(b Atom) bool {
	if b.Type() != RatAtom {
		return false
	}

	br := b.(Rat)
	return exact.Cmp(r.Value, br.Value)
}
