package main

import (
	"fmt"
)

type Weekday int

// Declare typed constants each with type of Weekday
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

//Weekday type

// String returns the name of the day
// func (day Weekday) String() string {
//     names := [...]string{"Sunday", "Monday", "Tuesday", "Wednesday",
//         "Thursday", "Friday", "Saturday"}

//     if day < Sunday || day > Saturday {
//       return "Unknown"
//     }

//     return names[day]
// }

// Weekend return true for a weekend day
// func (day Weekday) Weekend() bool {
// 	switch day {
// 	case Sunday, Saturday:
// 		return true
// 	default:
// 		return false
// 	}
// }

func main() {
	// Which day it is? Sunday
	fmt.Printf("Which day it is? %s\n", Sunday)

	// Is Saturday a weekend day? true
	//fmt.Printf("Is Saturday a weekend day? %t\n", Saturday.Weekend())

	// Is Monday a weekend day? false
	//fmt.Printf("Is Monday a weekend day? %t\n", Monday.Weekend())
}
