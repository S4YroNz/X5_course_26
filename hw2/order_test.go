package orders

import "testing"

func TestCreateOrder(t *testing.T) {
	order, err := CreateOrder(1)
	if err != nil {
		t.Fatal(err)
	}

	if order.ID != 1 {
		t.Errorf("ожидали ID 1, получили %d", order.ID)
	}

	if Total(order) != 0 {
		t.Error("сумма пустого заказа должна быть 0")
	}

	for _, id := range []int{0, -1} {
		_, err := CreateOrder(id)
		if err == nil {
			t.Errorf("ожидали ошибку для ID %d", id)
		}
	}
}

func TestAddItem(t *testing.T) {
	order, err := CreateOrder(1)
	if err != nil {
		t.Fatal(err)
	}

	if err := AddItem(&order, "Хлеб", 5000); err != nil {
		t.Fatal(err)
	}
	if err := AddItem(&order, "Молоко", 8000); err != nil {
		t.Fatal(err)
	}

	if len(order.Items) != 2 {
		t.Fatalf("ожидали 2 позиции, получили %d", len(order.Items))
	}

	if Total(order) != 13000 {
		t.Errorf("ожидали сумму 13000, получили %d", Total(order))
	}
}

func TestAddItemErrors(t *testing.T) {
	if err := AddItem(nil, "Хлеб", 5000); err == nil {
		t.Error("ожидали ошибку для nil")
	}

	order := Order{ID: 1}

	if err := AddItem(&order, "", 5000); err == nil {
		t.Error("ожидали ошибку для пустого имени")
	}

	if err := AddItem(&order, "Хлеб", -1); err == nil {
		t.Error("ожидали ошибку для отрицательной цены")
	}

	if len(order.Items) != 0 {
		t.Error("некорректные позиции не должны добавляться")
	}
}

func TestOrderCopy(t *testing.T) {
	original := Order{
		ID:     1,
		Status: "new",
		Items:  []Item{{Name: "Хлеб", Price: 5000}},
	}

	copied := original

	copied.Items = make([]Item, len(original.Items))
	copy(copied.Items, original.Items)

	copied.ID = 2
	copied.Status = "paid"
	copied.Items[0].Price = 100

	if original.ID != 1 || original.Status != "new" {
		t.Error("исходный заказ изменился")
	}

	if original.Items[0].Price != 5000 {
		t.Error("позиция исходного заказа изменилась")
	}
}

func TestStatusNames(t *testing.T) {
	if StatusNames["new"] != "Новый" {
		t.Error("неверное название статуса")
	}
}
