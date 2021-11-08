// @file: main.go
// @date: 2021/10/29

package main

import "fmt"

func main() {
	persons := Persons{
		{"Robert", "male", "single"},
		{"John", "male", "married"},
		{"Laura", "female", "married"},
		{"Diana", "female", "single"},
		{"Mike", "male", "single"},
		{"Bobby", "male", "single"},
	}
	maleFilter := MaleFilter{}
	femaleFilter := FemaleFilter{}
	singleFilter := SingleFilter{}
	singleMaleFilter := NewAndFilter(singleFilter, maleFilter)
	singleOrFemaleFilter := NewOrFilter(singleFilter, femaleFilter)

	fmt.Printf("Males:\n%s", maleFilter.Do(persons))
	fmt.Printf("Females:\n%s", femaleFilter.Do(persons))
	fmt.Printf("Single Males:\n%s", singleMaleFilter.Do(persons))
	fmt.Printf("Single or Females:\n%s", singleOrFemaleFilter.Do(persons))
}
