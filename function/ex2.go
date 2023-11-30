package main

import "fmt"

func sumAndDifference(num1, num2 int) (int, int) {
	sum := num1 + num2
	difference := num1 - num2
	return sum, difference
}

func main() {
	sum, diff := sumAndDifference(5, 8)

	fmt.Println("Sum = ", sum)
	fmt.Println("Difference = ", diff)
}
