package main

func main() {
	expenses := Expenses{}

	expenses.add("Lunch", 20)
	expenses.add("Dinner", 10)

	expenses.summary()

	expenses.list()

	expenses.delete(1)

	expenses.list()

	expenses.update(0, "New Lunch Desc", 15)

	expenses.list()

	expenses.summary()
}
