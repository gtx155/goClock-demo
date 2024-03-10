package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	day := t.Day()
	month := t.Month()
	year := t.Year()
	hour := t.Hour()
	min := t.Format("15:04:05")
	minPart := min[3:5]

	fmt.Println("Date Format: (DD/MM/YY)")
	fmt.Println("Date:", day, month, year)
	fmt.Println("Time:", hour, ":", minPart)
}
