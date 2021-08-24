package main

import (
	"fmt"
	"time"
)

func main() { 
	PrintDay()
	printType := func(i interface{}) { 
		switch t := i.(type) {
		case bool:
			fmt.Println("bool value: ", t) 
		case int:
			fmt.Println("int value: ", t)
		case float64:
			fmt.Println("float64 value: ", t) 
		case string: 
			fmt.Println("string value: ", t)
		}
	}
	printType(true)
	printType(1)
	printType(114.514)
	printType("Try this")
}

func PrintDay() {
	switch time.Now().Weekday() {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	}
}