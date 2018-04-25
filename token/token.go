package token

type (
	Type int
)

const (
	Illegal Type = iota+1
	LParen
	RParen
	Decimal
	Binary
	Hexadecimal
)

func (i Type) String() string {
	switch i {
	case Illegal:
		return "ILLEGAL"
	case Decimal:
		return "DECIMAL"
	case LParen:
		return "LPAREN"
	case RParen:
		return "RPAREN"
	}
	panic("invalid token")
	return ""
}