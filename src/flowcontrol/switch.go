package main

// https://go.dev/tour/flowcontrol/9
import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// example 1
	// This is a basic example of a switch statement in Go. The variable i is assigned the value of 2.
	// The switch statement then checks the value of i and runs the corresponding case.
	// In this case, the output would be "i is two". The default case will run if none of the other cases match.
	i := 2
	switch i {
	case 1:
		fmt.Println("i is one")
	case 2:
		fmt.Println("i is two")
	case 3:
		fmt.Println("i is three")
	default:
		fmt.Println("i is something else")
	}

	//example 2
	fmt.Print("Go runs on ")
	// unlike the other C,C++,Java,JavaScript,PHP,Go does not have a break statement in the switch,
	// In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS x.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// example 3
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	//example 4
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm bool")
		case int:
			fmt.Println("I'm int")
		case float64:
			fmt.Println("I'm float64")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(false)
	whatAmI(2.0)
	whatAmI(2)
	whatAmI("example")

	// example 5
	fmt.Println("When's Wednesday?")
	today := time.Now().Weekday()
	switch time.Wednesday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

}
