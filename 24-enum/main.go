package main

import (
	"errors"
	"fmt"
)

/*-- Form 1 --*/
type Weekday int

const (
	Monday    Weekday = 0
	Tuesday   Weekday = 1
	Wednesday Weekday = 2
	Thursday  Weekday = 3
	Friday    Weekday = 4
	Saturday  Weekday = 5
	Sunday    Weekday = 6
)

func (d Weekday) String() string {
	names := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	return names[d]
}

/*-- Form 2 --*/
type Operator string

const (
	EQUAL    Operator = "="
	GT       Operator = ">"
	GTE      Operator = ">="
	LT       Operator = "<"
	LTE      Operator = "<="
	CONTAINS Operator = "CONTAINS"
)

type FilterOperator struct {
	value Operator
}

func (f *FilterOperator) Value() string {
	return string(f.value)
}

func isValidOperator(operator Operator) bool {
	switch operator {
	case EQUAL, GT, GTE, LT, LTE, CONTAINS:
		return true
	}
	return false
}

func NewFilterOperator(value string) (*FilterOperator, error) {
	operator := Operator(value)
	if !isValidOperator(operator) {
		return nil, errors.New("invalid filter operator")
	}
	return &FilterOperator{value: operator}, nil
}

func main() {
	day := Monday
	fmt.Println(day)

	operator, err := NewFilterOperator("=")
	if err != nil {
		panic(err)
	}
	fmt.Println(operator.Value())
}
