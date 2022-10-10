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

	day := 1
	printFirstWeek(&day, today)
	printOtherWeeks(day, today)
}

func printHeader(month time.Month, year int) {
	c1 := color.New(color.FgHiBlue).Add(color.Bold)
	c2 := color.New(color.FgHiCyan).Add(color.Bold)
	c1.Printf("%s %d", months[month-1], year)
	c1.Println("")
	c2.Println(strings.Join(days[:], " "))
}

func printFirstWeek(day *int, today time.Time) {
	f := beginningOfMonth(today).Weekday()
	found := false

	for i, v := range days {
		if f.String()[0:3] == v {
			printDay(*day, i, today)
			*day++
			found = true
			continue
		} else {
			if !found {
				fmt.Print("    ")
			}
		}

		if found {
			printDay(*day, i, today)
		}
	}

	fmt.Println("")
}

func printOtherWeeks(day int, today time.Time) {
	e := endOfMonth(today)
	idx := 0
	for day <= e.Day() {
		printDay(day, idx, today)
		idx++

		if idx >= len(days) {
			fmt.Println("")
			idx = 0
		}

		day++
	}
}

func printDay(day int, idx int, today time.Time) {
	workdayColor := color.New(color.FgWhite).Add(color.Bold)
	holidayColor := color.New(color.FgWhite)
	currentDayColor := color.New(color.FgHiRed).Add(color.Bold)

	if day > 9 {
		if day == today.Day() {
			currentDayColor.Printf(" %d ", day)	
		} else if idx == 5 || idx == 6 {
			holidayColor.Printf(" %d ", day)
		} else {
			workdayColor.Printf(" %d ", day)
		}
	} else {
		if day == today.Day() {
			currentDayColor.Printf("  %d ", day)
		} else if idx == 5 || idx == 6 {
			holidayColor.Printf("  %d ", day)
		} else {
			workdayColor.Printf("  %d ", day)
		}
	}
}

func beginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

func endOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}