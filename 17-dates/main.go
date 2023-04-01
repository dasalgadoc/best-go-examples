package main

import (
	"fmt"
	"time"
)

func main() {
	// Timezones: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	fmt.Println("Loca time: ", time.Now())

	// Timezone Sao Paulo
	location, _ := time.LoadLocation("America/Sao_Paulo")
	fmt.Println("Sao Paulo time: ", time.Now().In(location))

	// Truncate dates
	hourTruncate := time.Now().Truncate(time.Millisecond).In(location)
	fmt.Println("Sao Paulo hour: ", hourTruncate)
	dateTruncate := time.Now().Truncate(time.Hour).In(location)
	fmt.Println("Sao Paulo day: ", dateTruncate)

	// Custom dates
	lowerRange := time.Date(2023, 1, 1, 0, 0, 0, 0, location)
	maxRange := time.Date(2023, 1, 1, 23, 59, 59, 999999999, location)
	fmt.Println("start: ", lowerRange, " end: ", maxRange)

	fmt.Println(dateTruncate.After(lowerRange))
	fmt.Println(dateTruncate.Before(lowerRange))

	// Adding
	now := time.Now()
	newTime := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, now.Location())
	nightlyLte := time.Date(now.Year(), now.Month(), now.Day()+1, 6, 0, 0, 0, now.Location())
	fmt.Println(newTime)
	fmt.Println(nightlyLte)
	fmt.Printf("%T", nightlyLte)
}
