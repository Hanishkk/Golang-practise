package conditional

import "fmt"

func PrintCases(x int) {
	switch x {
	case 1:
		{
			fmt.Println("Hurray")
		}

	case 2:
		{
			fmt.Println("SMILE")
		}

	case 3:
		{
			fmt.Println("hungry")
		}

	case 4:
		{
			fmt.Println("Gay")
		}

	default:
		{
			fmt.Println("Invalid input")
		}
	}
}
