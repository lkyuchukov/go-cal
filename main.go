package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

var months = [12]string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

var days = [7]string{
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
	"Sun",
}

func main() {
	today := time.Now()
	month := today.Month()
	year := today.Year()

	fmt.Printf("%s %d", months[month-1], year)
	fmt.Println("")
	fmt.Println(strings.Join(days[:], " "))

	firstDay := beginningOfMonth(today).Weekday()
	lastDay := endOfMonth(today)

	day := 1

	printFirstWeek(firstDay, &day)

	fmt.Println("")

	printOtherWeeks(day, lastDay)
}

func printFirstWeek(firstDay time.Weekday, day *int) {
	found := false
	for _, v := range days {


		if firstDay.String()[0:3] == v {
			fmt.Printf("  %d ", *day)
			*day++
			found = true
			continue
		} else {
			if !found {
				fmt.Print("    ")
			}
		}

		if found {
			fmt.Printf("  %d", *day)
			*day++
		}
	}
}

func printOtherWeeks(day int, e time.Time) {
	idx := 0

	for day <= e.Day() {
		if day > 9 {
			fmt.Printf(" %d ", day)
		} else {
			fmt.Printf("  %d ", day)
		}
		idx++

		if idx >= len(days) {
			fmt.Println("")
			idx = 0
		}

		day++
	}
}

func beginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

func endOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}
