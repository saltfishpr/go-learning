package main

import (
	"io"
	"testing"
)

func TestLexer_Lex(t *testing.T) {
	input := `
[
  {
    "precision": "zip",
    "Latitude": 37.7668,
    "Longitude": -122.3959,
    "Address": "",
    "City": "SAN FRANCISCO",
    "State": "CA",
    "Zip": "94107",
    "Country": "US"
  },
  {
    "precision": "zip",
    "Latitude": 37.371991,
    "Longitude": -122.02602,
    "Address": "",
    "City": "SUNNYVALE",
    "State": "CA",
    "Zip": "94085",
    "Country": "US"
  }
]`

	l := NewLexer([]byte(input))

	for {
		token, err := l.Lex()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatal(err)
		}
		t.Log(token)
	}
}
