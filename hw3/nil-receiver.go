package main

import "fmt"

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func validateBad(age int) error {
	var err *ValidationError
	if age < 0 {
		err = &ValidationError{Message: "возраст не может быть отрицательным"}
	}
	return err
}

func validateGood(age int) error {
	if age < 0 {
		return &ValidationError{Message: "возраст не может быть отрицательным"}
	}
	return nil
}

func (e *ValidationError) IsEmpty() bool {
	return e == nil
}

func showNilReceiver() {
	defer func() {
		if value := recover(); value != nil {
			fmt.Println("Перехвачена паника при обращении к полю nil-receiver")
		}
	}()

	var pointer *ValidationError
	var err error = pointer
	fmt.Println("Указатель равен nil:", pointer == nil)
	fmt.Println("Интерфейс равен nil:", err == nil)
	fmt.Println("Метод с проверкой nil:", pointer.IsEmpty())
	fmt.Println(err.Error())
}

func main() {
	fmt.Println("validateBad(20) == nil:", validateBad(20) == nil)
	fmt.Println("validateGood(20) == nil:", validateGood(20) == nil)
	fmt.Println("validateGood(-1):", validateGood(-1))
	showNilReceiver()
	fmt.Println("После showNilReceiver программа продолжает работу")
}
