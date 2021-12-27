package conditional

import "fmt"

func Condition() {
	if 69 > 72 {
		fmt.Printf("is smaller")
	} else {
		fmt.Printf("is greater")
	}

	if 99 == 1 {
		fmt.Printf("is worthless")
	} else {
		fmt.Printf("\nis worst")
	}

	if num := 88; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 99 {
		fmt.Println("\nhas 10 digit")
	} else {
		fmt.Println("nothing is here")
	}
}

func ConditonWithParameter(x int, y float64) {
	if x >= 5 {
		fmt.Printf("then %d is greater than 5 ", x)
	} else {
		fmt.Printf("then %d is smaller than 5", x)
	}

	//Checking if the value of Y is equal to the given number//
	if y == 3.14 {
		fmt.Printf("then %f is equal to 3.14", y)
	} else {
		fmt.Printf("then %f is not equal to 3.14", y)
	}

}

func ConditionsWithMultipleParameters(x int, y int, z float32, w float32) {
	if x > y {
		fmt.Printf("%d is greater than %d", x, y)
	} else if x == y {
		fmt.Printf("\n%d is equal to %d", x, y)
	} else {
		fmt.Printf("\n%d is smaller than %d", x, y)
	}

	if z > w {
		fmt.Printf("\n%f is greater than %f", z, w)
	} else if z == w {
		fmt.Printf("\n%f is equal to %f", z, w)
	} else {
		fmt.Printf("\n%f is smaller than %f", z, w)
	}
}
