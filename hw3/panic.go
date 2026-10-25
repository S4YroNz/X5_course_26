package main

import "fmt"

func level3(catchHere bool) {
	defer fmt.Println("level3: defer выполнен")
	if catchHere {
		defer func() {
			if value := recover(); value != nil {
				fmt.Println("level3: перехвачено:", value)
			}
		}()
	}
	fmt.Println("level3: вызываем panic")
	panic("ошибка в level3")
}

func level2(catchHere bool, catchInLevel3 bool) {
	defer fmt.Println("level2: defer выполнен")
	if catchHere {
		defer func() {
			if value := recover(); value != nil {
				fmt.Println("level2: перехвачено:", value)
			}
		}()
	}
	level3(catchInLevel3)
	fmt.Println("level2: после вызова level3")
}

func level1(catchInLevel2 bool, catchInLevel3 bool) {
	defer fmt.Println("level1: defer выполнен")
	defer func() {
		if value := recover(); value != nil {
			fmt.Println("level1: перехвачено:", value)
		}
	}()
	level2(catchInLevel2, catchInLevel3)
	fmt.Println("level1: после вызова level2")
}

func main() {
	fmt.Println("--- Перехват в level3 ---")
	level1(false, true)
	fmt.Println("--- Перехват в level2 ---")
	level1(true, false)
	fmt.Println("--- Перехват в level1 ---")
	level1(false, false)
	fmt.Println("main: программа продолжает работу")
}
