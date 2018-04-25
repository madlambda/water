package lexer_test

import (
	"testing"
	"github.com/madlambda/one/lexer"
	"github.com/madlambda/one/token"
)

func consumeTokens(tokens <-chan lexer.Token) []lexer.Token {
	var values []lexer.Token
	for t := range tokens {
		values = append(values, t)
	}
	return values
}

func TestLexer(t *testing.T) {
	for _, tc := range []struct {
		code     string
		expected []lexer.Token
	}{
		{
			code: "9",
			expected: []lexer.Token{
				{
					Type:  token.Decimal,
					Value: "9",
				},
			},
		},
		{
			code: "100",
			expected: []lexer.Token{
				{
					Type:  token.Decimal,
					Value: "100",
				},
			},
		},
		{
			code: "242342342342423534645654",
			expected: []lexer.Token{
				{
					Type:  token.Decimal,
					Value: "242342342342423534645654",
				},
			},
		},
		{
			code: "0xff",
			expected: []lexer.Token{
				{
					Type:  token.Hexadecimal,
					Value: "0xff",
				},
			},
		},
		{
			code: "0b10010100",
			expected: []lexer.Token{
				{
					Type:  token.Binary,
					Value: "0b10010100",
				},
			},
		},
	} {
		tokens := lexer.ReadTokens(tc.code)

		got := consumeTokens(tokens)

		if len(got) != len(tc.expected) {
			t.Logf("got(%#v)", got)
			t.Logf("exp(%#v)", tc.expected)
			t.Fatalf("element size mismatched: %d != %d", len(got), len(tc.expected))
		}

		for i := 0; i < len(got); i++ {
			if got[i].Type != tc.expected[i].Type {
				t.Fatalf("type differ: %s != %s",
					got[i].Type,
					tc.expected[i].Type)
			}

			if got[i].Value != tc.expected[i].Value {
				t.Fatalf("value differ: %s != %s",
					got[i].Value,
					tc.expected[i].Value)
			}
		}
	}
}