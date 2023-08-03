package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Classroom struct {
	Teacher  string   `json:"teacher"`
	Subjects []string `json:"subjects"`
	Day      Weekday  `json:"day"`
}

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

var toID = map[string]Weekday{
	"Monday":    Monday,
	"Tuesday":   Tuesday,
	"Wednesday": Wednesday,
	"Thursday":  Thursday,
	"Friday":    Friday,
	"Saturday":  Saturday,
	"Sunday":    Sunday,
}

func (c *Weekday) UnmarshalJSON(b []byte) error {
	var t string
	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}
	d, ok := toID[t]
	if !ok {
		return errors.New("invalid weekday")
	}
	*c = d

	return nil
}

func main() {
	var goodCass Classroom
	goodSource := `{"teacher": "Mr. Smith","subjects": ["Math", "English"],"day": "Monday"}`
	err := json.Unmarshal([]byte(goodSource), &goodCass)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", goodCass)

	var badCass Classroom
	badSource := `{"teacher": "Mr. Johnson","subjects": ["Math", "English"],"day": "CrazyDay"}`
	err = json.Unmarshal([]byte(badSource), &badCass)
	if err != nil {
		fmt.Println("I can't unmarshal this JSON because: ", err)
	}
	// note how the badCass is still initialized with the default value (Monday)
	fmt.Printf("%+v \n", badCass)
}
