package main

import "fmt"

// tokenKind determines type of a token.
type tokenKind int

const (
	tokenKindUndef      tokenKind = iota //
	tokenKindWhitespace                  // 0x20 0x09 0x0A 0x0D
	tokenKindDelim                       // { } [ ] : ,
	tokenKindString                      // A string literal, e.g. "abc\u1234"
	tokenKindNumber                      // Number literal, e.g. -1.5e-5
	tokenKindBoolean                     // Boolean literal: true false.
	tokenKindNull                        // null keyword.
)

// token describes a single token: type, position in the input and value.
type token struct {
	kind tokenKind

	literal []byte // Raw value of a token.
}

func (t token) String() string {
	return fmt.Sprintf("kind: %d, literal: %s", t.kind, string(t.literal))
}
