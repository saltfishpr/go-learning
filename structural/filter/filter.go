// @file: filter.go
// @date: 2021/10/29

package main

import "reflect"

type Filter interface {
	Do(Persons) Persons
}

type MaleFilter struct{}

func (MaleFilter) Do(persons Persons) Persons {
	var res Persons
	for i := 0; i < len(persons); i++ {
		if persons[i].gender == "male" {
			res = append(res, persons[i])
		}
	}
	return res
}

type FemaleFilter struct{}

func (FemaleFilter) Do(persons Persons) Persons {
	var res Persons
	for i := 0; i < len(persons); i++ {
		if persons[i].gender == "female" {
			res = append(res, persons[i])
		}
	}
	return res
}

type SingleFilter struct{}

func (SingleFilter) Do(persons Persons) Persons {
	var res Persons
	for i := 0; i < len(persons); i++ {
		if persons[i].maritalStatus == "single" {
			res = append(res, persons[i])
		}
	}
	return res
}

type AndFilter struct {
	filter, otherFilter Filter
}

func NewAndFilter(filter Filter, otherFilter Filter) *AndFilter {
	return &AndFilter{filter: filter, otherFilter: otherFilter}
}

func (f AndFilter) Do(persons Persons) Persons {
	return f.otherFilter.Do(f.filter.Do(persons))
}

type OrFilter struct {
	filter, otherFilter Filter
}

func NewOrFilter(filter Filter, otherFilter Filter) *OrFilter {
	return &OrFilter{filter: filter, otherFilter: otherFilter}
}

func (f OrFilter) Do(persons Persons) Persons {
	persons1 := f.filter.Do(persons)
	persons2 := f.otherFilter.Do(persons)
	for i := 0; i < len(persons2); i++ {
		if !contains(persons1, persons2[i]) {
			persons1 = append(persons1, persons2[i])
		}
	}
	return persons1
}

func contains(persons Persons, person *Person) bool {
	for _, v := range persons {
		if reflect.DeepEqual(v, person) {
			return true
		}
	}
	return false
}
