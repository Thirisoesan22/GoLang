package main

import "fmt"

func greet(name string) (string, error) {
	return "Hello " + name + "...", nil
}

func calculateSquare(num int) int {
	return num * num
}

func main() {
	greet_msg, err := greet("Thiri")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(greet_msg)

	fmt.Println(calculateSquare(7))
}
