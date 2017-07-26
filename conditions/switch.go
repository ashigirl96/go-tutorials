package conditions

import (
	"fmt"
	"runtime"
	"time"
)

func SwitchMain() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s", os)
	}
}

func SwitchEvaluationOrder() {
	fmt.Println("When's Sunday?")
	today := time.Now().Weekday()
	switch time.Sunday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tommorow")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away!!,", today - 1)
	}
}

func printWorld() string {
	fmt.Print("皆さんおはこんばんにち, ")
	return " World!!!!"
}


func DeferMain() {
	defer fmt.Println(printWorld())

	fmt.Print("Hello, ")
}
