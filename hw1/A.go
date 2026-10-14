package main

import "fmt"

func main() {
	for i := range 30 {
		v := i + 1
		if v%5 == 0 && v%3 == 0 {
			fmt.Println("FizzBuzz")
		} else if v%5 == 0 {
			fmt.Println("Fizz")
		} else if v%3 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(v)
		}

	}
}
