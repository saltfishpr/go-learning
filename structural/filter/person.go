// @file: person.go
// @date: 2021/10/29

package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name          string
	gender        string
	maritalStatus string
}

type Persons []*Person

func (p Persons) String() string {
	var builder strings.Builder
	for _, person := range p {
		builder.WriteString(fmt.Sprintf("\t%#v\n", person))
	}
	return builder.String()
}
