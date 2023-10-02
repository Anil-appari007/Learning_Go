package main

import "fmt"

func main() {
	var a = 45

	switch a {
	case 100:
		fmt.Println("a equals 100")
	case 1:
		fmt.Println("a = 1")
	default:
		fmt.Println("No condition matched.")
	}

	switch a {
	case 2, 4:
		fmt.Println("Matched 2, 4")
	case 1, 45:
		fmt.Println("matched 1, 45")
	default:
		fmt.Println("no match")
	}

	// conditionals with switch case
	switch {
	case a == 45:
		fmt.Println("Match 1")
	case a < 100:
		fmt.Println("Match 1")
	case a > 4:
		fmt.Println("Match 1")
	}

	// Fallthrough in switch
	switch {
	case a == 45:
		fmt.Println("Match 1, running fallthrough")
		fallthrough
	case a < 100:
		fmt.Println("Match 2 reached bcoz of fallthrough")
		fallthrough
	case a > 4:
		fmt.Println("Match 3 reached bcoz of fallthrough")
		fallthrough
	default:
		fmt.Println("Default reached bcoz of fallthrough")
	}

}
