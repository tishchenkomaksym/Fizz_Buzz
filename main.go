package main

import (
	"fmt"
)

func main() {

	for i := 1; i < 100; i++ {
		fizzBuzz(i)
	}

}

func fizzBuzz(i int) {
	buzz := "Buzz"
	fizz := "Fizz"

	if i%3 == 0 && i%5 == 0 {
		fmt.Println(fizz, buzz)
	} else if i%3 == 0 {
		fmt.Println(fizz)
	} else if i%5 == 0 {
		fmt.Println(buzz)
	} else {
		fmt.Println(i)
	}

}
