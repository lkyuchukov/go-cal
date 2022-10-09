package main

import (
	"fmt"
	"strings"
	"time"
	"github.com/fatih/color"
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

	printHeader(month, year)

	firstDay := beginningOfMonth(today).Weekday()
	lastDay := endOfMonth(today)

	day := 1

	printFirstWeek(firstDay, &day)

	fmt.Println("")

	printOtherWeeks(day, lastDay)
}

func printHeader(month time.Month, year int) {
	c1 := color.New(color.FgHiBlue).Add(color.Bold)
	c2 := color.New(color.FgHiMagenta).Add(color.Bold)
	c1.Printf("%s %d", months[month-1], year)
	c1.Println("")
	c2.Println(strings.Join(days[:], " "))
}

func printFirstWeek(firstDay time.Weekday, day *int) {
	found := false
	c1 := color.New(color.FgHiRed).Add(color.Bold)
	for _, v := range days {
		if firstDay.String()[0:3] == v {
			c1.Printf("  %d ", *day)
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