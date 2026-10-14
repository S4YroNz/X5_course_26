package main

import "fmt"

func getCategory(price int) string {
	switch {
	case price < 0:
		return fmt.Sprintf("invalid price: %v", price)
	case price < 10000:
		return "small"
	case price < 50000:
		return "regular"
	default:
		return "large"
	}
}
