package lexer

import (
	"unicode"
	"unicode/utf8"
	"strings"

	"github.com/madlambda/one/token"
)

const eof = -1

type (
	Token struct {
		Type  token.Type
		Value string

		Pos int
	}

	Lexer struct {
		input     string
		pos       int
		start     int
		lastWidth int

		tokens chan Token
	}

	stateFn func(*Lexer) stateFn
)

func ReadTokens(code string) <-chan Token {
	l := &Lexer{
		input:  code,
		tokens: make(chan Token),
	}

	go run(l)
	return l.tokens
}

func run(l *Lexer) {
	for state := startState; state != nil; {
		state = state(l)
	}

	close(l.tokens)
}

func (l *Lexer) peek() rune {
	r := l.next()
	if r == eof {
		return eof
	}
	l.backup()
	return r
}

func (l *Lexer) next() rune {
	if l.pos >= len(l.input) {
		l.lastWidth = 0
		return eof
	}
	r, width := utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += width
	l.lastWidth = width
	return r
}

func (l *Lexer) backup() {
	l.pos -= l.lastWidth
}

func (l *Lexer) emit(typ token.Type) {
	l.tokens <- Token{
		Type:  typ,
		Value: l.input[l.start:l.pos],
		Pos:   l.start,
	}
	l.start = l.pos
}

func (l *Lexer) acceptDigit() {
	l.accept("0123456789")
}

func (l *Lexer) acceptDigits() {
	l.acceptRun("0123456789")
}

func (l *Lexer) accept(chars string) bool {
	return strings.IndexRune(chars, l.peek()) != -1
}

func (l *Lexer) acceptRun(chars string) {
	for strings.IndexRune(chars, l.next()) >= 0 {
	}
	l.backup()
}

func startState(l *Lexer) stateFn {
	r := l.peek()
	switch {
	case unicode.IsDigit(r):
		return lexNumber
	case r == eof:
		return nil
	}
	l.emit(token.Illegal)
	return nil
}

func lexNumber(l *Lexer) stateFn {
	r := l.peek()
	if r == '0' {
		l.next()
		if l.accept("xXx") {
			l.next()
			return lexHex
		}

		if l.accept("bBb") {
			l.next()
			return lexBin
		}
	}

	l.acceptDigits()
	l.emit(token.Decimal)
	return startState
}

func lexHex(l *Lexer) stateFn {
	hexChars := "123456789aAbBcCdDeEfF"
	if !l.accept(hexChars) {
		l.emit(token.Illegal)
		return nil
	}
	l.acceptRun(hexChars)
	l.emit(token.Hexadecimal)
	return startState
}

func lexBin(l *Lexer) stateFn {
	binChars := "01"
	if !l.accept(binChars) {
		l.emit(token.Illegal)
		return nil
	}
	l.acceptRun(binChars)
	l.emit(token.Binary)
	return startState
}