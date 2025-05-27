package main

func main() {
	expenses := Expenses{}

	expenses.add("Lunch", 20)
	expenses.add("Dinner", 10)

	expenses.list()

}
