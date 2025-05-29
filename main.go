package main

func main() {
	expenses := Expenses{}
	storage := NewStorage[Expenses]("expenses.json")
	storage.Load(&expenses)

	expenses.list()

	expenses.delete(1)

	expenses.list()

	expenses.update(0, "Lunch", 15)

	expenses.list()

	expenses.summary()

	storage.Save(expenses)
}
