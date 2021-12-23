/**Addition of two numbers**/

package main

import "fmt"

/**addition**/
func doubleaddition(d int, e int) int {
	result := d + e
	return result
}

func subtraction(a int, b int) int {
	result := a - b
	return result
}

func division(a int, b int) int {
	result := a / b
	return result
}

func multiplication(a int, b int) int {
	result := a * b
	return result
}

func reminder(a int, b int) int {
	result := a % b
	return result
}

func main() {
	//new()
	result := addition(20, 30, 40)
	fmt.Printf("The result of addition is >> %d", result)

}

func new() {
	result := doubleaddition(20, 30)
	fmt.Printf("The result of addition is >> %d", result)

	//**substract**//
	result = subtraction(40, 30)
	fmt.Printf("your result is >> %d", result)

	//**division**//
	result = division(80, 10)
	fmt.Printf("\n your result is >> %d", result)

	//**multiplication**//
	result = multiplication(80, 10)
	fmt.Printf("\n your result is >> %d", result)

	//** reminder**//
	//result = reminder(80, 30)
	fmt.Printf("\nyour result is >> %d", result)

}
