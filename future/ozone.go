package future

import "fmt"

func Addition(a int, b int) int {
result := a + b
return result
}

func Subtraction(a int , b int) int {
	result:= a - b
	return result
}

func Multiplication(a int, b int) int {
	result:= a * b
	return result
}

func Twinky () {
 result := Addition (20, 39)
 fmt.Printf("The result value is >> %d", result)

 result = Subtraction(40, 30)
	fmt.Printf("your result is >> %d", result)

	result = Subtraction(40, 30)
	fmt.Printf("your result is >> %d", result)
}