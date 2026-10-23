package orders

import "errors"

type Item struct {
	Name  string
	Price int
}

type Order struct {
	ID     int
	Status string
	Items  []Item
}

var StatusNames = make(map[string]string)

func init() {
	StatusNames["new"] = "Новый"
	StatusNames["paid"] = "Оплачен"
	StatusNames["done"] = "Завершён"
}

func CreateOrder(id int) (Order, error) {
	if id <= 0 {
		return Order{}, errors.New("ID должен быть положительным")
	}

	return Order{
		ID:     id,
		Status: "new",
	}, nil
}

func AddItem(order *Order, name string, price int) error {
	if order == nil {
		return errors.New("заказ не передан")
	}
	if name == "" {
		return errors.New("имя позиции не должно быть пустым")
	}
	if price < 0 {
		return errors.New("цена не должна быть отрицательной")
	}

	order.Items = append(order.Items, Item{
		Name:  name,
		Price: price,
	})

	return nil
}

func Total(order Order) int {
	sum := 0

	for _, item := range order.Items {
		sum += item.Price
	}

	return sum
}
