package main

import "fmt"

func minmaxsum(inp []int) {
	if len(inp) == 0 {
		fmt.Println("пустой срез")
		return
	}
	min := inp[0]
	max := inp[0]
	sum := 0
	for _, v := range inp {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		sum += v
	}
	fmt.Printf("минимум %v, максимум %v, сумма %v, среднее %v\n", min, max, sum, sum/len(inp))
}
