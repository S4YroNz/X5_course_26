package main

import (
	"errors"
	"fmt"
)

var ErrEmptyMessage = errors.New("пустое сообщение")

type SendError struct {
	Channel string
	Err     error
}

func (e *SendError) Error() string {
	return fmt.Sprintf("канал %s: %v", e.Channel, e.Err)
}

// Возвращаем причину, чтобы errors.Is мог пройти по цепочке.
func (e *SendError) Unwrap() error {
	return e.Err
}

// Считаем ошибки отправки одинаковыми, если совпадает канал.
// Причину здесь не проверяем: её проверит errors.Is через Unwrap.
func (e *SendError) Is(target error) bool {
	other, ok := target.(*SendError)
	if !ok || other == nil {
		return false
	}
	return e.Channel == other.Channel
}

func send(message string) error {
	if message == "" {
		return &SendError{Channel: "email", Err: ErrEmptyMessage}
	}
	return nil
}

func main() {
	err := send("")
	if err != nil {
		fmt.Println("Ошибка:", err)
		fmt.Println("Причина:", errors.Unwrap(err))
		fmt.Println("Пустое сообщение:", errors.Is(err, ErrEmptyMessage))
		fmt.Println("Канал email:", errors.Is(err, &SendError{Channel: "email"}))
		fmt.Println("Канал sms:", errors.Is(err, &SendError{Channel: "sms"}))
		wrapped := fmt.Errorf("отправка уведомления: %w", err)
		fmt.Println("Причина через обёртку:", errors.Is(wrapped, ErrEmptyMessage))
		fmt.Println("Канал через обёртку:", errors.Is(wrapped, &SendError{Channel: "email"}))
	}
	fmt.Println("Отправка непустого сообщения:", send("Привет"))
}
