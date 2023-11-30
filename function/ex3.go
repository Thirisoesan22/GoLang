package main

import (
	"errors"
	"fmt"
)

func divideWithRemainder(q int, d int) (int, int, error) {
	quotient := q / d
	remainder := q % d
	return quotient, remainder, errors.New("This is error")
}

func main() {
	q, d, err := divideWithRemainder(13, 10)
	if err != nil {
		fmt.Println(err)
	}
	{
		fmt.Printf("The quotient is %d and remainder is %d \n", q, d)
	}
}
