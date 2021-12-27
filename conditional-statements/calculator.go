package conditional

import "fmt"

func Calculator(x int, y int, operation string) int {
	switch operation {
	case "Addition":
		{
			result := x + y
			return result
		}

	case "Subtract":
		{
			result := x - y
			return result
		}

	case "multiplication":
		{
			result := x * y
			return result
		}

	case "division":
		{
			result := x / y
			return result
		}

	case "modulus":
		{
			result := x % y
			return result
		}

	default:
		{
			fmt.Println("invalid operation please fuckk off")
			return 0
		}
	}
}
