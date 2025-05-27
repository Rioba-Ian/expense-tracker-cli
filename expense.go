package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Expense struct {
	Description string
	Amount      int
	CreatedAt   time.Time
}

type Expenses []Expense

func (expenses *Expenses) add(desc string, amount int) {
	expense := Expense{
		Description: desc,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}

	*expenses = append(*expenses, expense)
}

func (expenses *Expenses) validateIndex(index int) error {
	if index < 0 || index >= len(*expenses) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (expenses *Expenses) list() {
	table := table.New(os.Stdout)

	table.SetRowLines(false)

	table.SetHeaders("#", "ID", "Date", "Description", "Amount")

	for index, e := range *expenses {
		table.AddRow("#", strconv.Itoa(index), e.CreatedAt.Format(time.RFC1123), e.Description, strconv.Itoa(e.Amount))
	}

	table.Render()
}
