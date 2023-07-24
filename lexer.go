package main

import (
	"fmt"
	"io"
	"unicode/utf8"
)

const (
	literalFalse = "false"
	literalTrue  = "true"
	literalNull  = "null"
)

type Lexer struct {
	src []byte

	tokenStart Position
	tokenEnd   Position
}

func NewLexer(src []byte) *Lexer {
	return &Lexer{
		src:        src,
		tokenStart: Position{Offset: 0, Line: 1, Column: 1},
		tokenEnd:   Position{Offset: 0, Line: 1, Column: 1},
	}
}

func (l *Lexer) Lex() (token, error) {
	r, err := l.next()
	if err != nil {
		return token{}, err
	}

	switch r {
	// delimiters
	case 0x5B, 0x7B, 0x5D, 0x7D, 0x3A, 0x2C:
		return l.emit(tokenKindDelim)

	// numbers
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
		if err := l.fetchNumber(r); err != nil {
			return token{}, err
		}
		return l.emit(tokenKindNumber)

	// strings
	case '"':
		var escapeMode bool
		for {
			r2, err := l.next()
			if err != nil {
				return token{}, l.errorf("read string error: %w", err)
			}

			if escapeMode {
				switch r2 {
				// ", \, /, b, f, n, r, t
				case 0x22, 0x5C, 0x2F, 0x62, 0x66, 0x6E, 0x72, 0x74:
					escapeMode = false

				// u
				case 0x75:
					for i := 0; i < 4; i++ {
						r3, err := l.next()
						if err != nil {
							return token{}, l.errorf("read string error: %w", err)
						}
						if !isHexDigit(r3) {
							return token{}, l.errorf("invalid string: not a hex digit follow \\u")
						}
					}
					escapeMode = false

				default:
					return token{}, l.errorf("invalid string")
				}
			} else {
				if r2 == '"' {
					break
				}
				if r2 == '\\' {
					escapeMode = true
				}
			}
		}
		return l.emit(tokenKindString)

	// false?
	case 'f':
		if err := l.fetchLiteral(literalFalse); err != nil {
			return token{}, err
		}
		return l.emit(tokenKindBoolean)

	// true?
	case 't':
		if err := l.fetchLiteral(literalTrue); err != nil {
			return token{}, err
		}
		return l.emit(tokenKindBoolean)

	// null?
	case 'n':
		if err := l.fetchLiteral(literalNull); err != nil {
			return token{}, err
		}
		return l.emit(tokenKindNull)

	}

	if isWhitespace(r) {
		for l.sniff(isWhitespace) {
			_, _ = l.next()
		}
		return l.emit(tokenKindWhitespace)
	}

	return token{}, l.errorf("invalid token")
}

func (l *Lexer) fetchNumber(first rune) error {
	const (
		state0 = iota // start
		state1        // -
		state2        // 0
		state3        // int
		state4        // .
		state5        // frac
		state6        // e
		state7
		state8
	)
	var state int
	switch first {
	case '-':
		state = state1
	case '0':
		state = state2
	case '1', '2', '3', '4', '5', '6', '7', '8', '9':
		state = state3
	}

LOOP:
	for {
		r, err := l.peek()
		if err != nil {
			return l.errorf("read number error: %w", err)
		}

		switch state {
		case state1:
			switch r {
			case '0':
				state = state2
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				state = state3
			default:
				return l.errorf("invalid number")
			}
		case state2:
			switch r {
			case '.':
				state = state4
			case 'e', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				return l.errorf("invalid number")
			default:
				break LOOP
			}
		case state3:
			switch r {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				state = state3
			case '.':
				state = state4
			case 'e':
				state = state6
			default:
				break LOOP
			}
		case state4:
			switch r {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				state = state5
			default:
				return l.errorf("invalid number")
			}
		case state5:
			switch r {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				state = state5
			case '.':
				return l.errorf("invalid number")
			case 'e':
				state = state6
			default:
				break LOOP
			}
		case state6:
			switch r {
			case '+', '-':
				state = state7
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				state = state8
			default:
				return l.errorf("invalid number")
			}
		case state7:
			switch r {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				state = state8
			default:
				return l.errorf("invalid number")
			}
		case state8:
			switch r {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				state = state8
			case '.', 'e':
				return l.errorf("invalid number")
			default:
				break LOOP
			}
		default:
			return l.errorf("invalid number: invalid state: %d", state)
		}
		_, _ = l.next()
	}

	return nil
}

func (l *Lexer) fetchLiteral(lit string) error {
	for _, v := range lit[1:] {
		r, err := l.peek()
		if err != nil {
			return err
		}
		if r != v {
			return l.errorf("literal value not match")
		}
		_, _ = l.next()
	}
	return nil
}

func (l *Lexer) sniff(wantFn func(rune) bool) bool {
	r, _, err := l.readRune()
	if err != nil {
		return false
	}
	return wantFn(r)
}

func (l *Lexer) peek() (rune, error) {
	r, _, err := l.readRune()
	return r, err
}

func (l *Lexer) next() (rune, error) {
	r, n, err := l.readRune()
	if err != nil {
		return r, err
	}
	if r == '\n' {
		l.tokenEnd.Line++
		l.tokenEnd.Column = 1
	} else {
		l.tokenEnd.Column++
	}
	l.tokenEnd.Offset += n
	return r, nil
}

func (l *Lexer) readRune() (rune, int, error) {
	r, n := utf8.DecodeRune(l.src[l.tokenEnd.Offset:])
	if n == 0 {
		return r, 0, io.EOF
	}
	if r == utf8.RuneError {
		return r, n, l.errorf("invalid utf-8")
	}
	return r, n, nil
}

func (l *Lexer) emit(t tokenKind) (token, error) {
	token := token{
		kind:    t,
		literal: l.src[l.tokenStart.Offset:l.tokenEnd.Offset],
	}
	l.tokenStart = l.tokenEnd
	return token, nil
}

func (l *Lexer) errorf(format string, args ...interface{}) error {
	return fmt.Errorf("position: %s, error: %w", l.tokenStart, fmt.Errorf(format, args...))
}

func isWhitespace(r rune) bool {
	return r == 0x20 ||
		r == 0x09 ||
		r == 0x0A ||
		r == 0x0D
}

func isHexDigit(r rune) bool {
	return '0' <= r && r <= '9' ||
		'a' <= r && r <= 'f' ||
		'A' <= r && r <= 'F'
}
