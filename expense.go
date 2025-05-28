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
		amount := "$" + strconv.Itoa(e.Amount)
		table.AddRow("#", strconv.Itoa(index), e.CreatedAt.Format(time.RFC1123), e.Description, amount)
	}

	table.Render()
}

func (expenses *Expenses) delete(index int) error {
	e := *expenses

	if err := e.validateIndex(index); err != nil {
		return err
	}

	*expenses = append(e[:index], e[index+1:]...)

	return nil
}

func (expenses *Expenses) update(index int, desc string, amount int) error {
	e := *expenses

	if err := e.validateIndex(index); err != nil {
		return err
	}

	if desc == "" {
		err := errors.New("Description cannot be empty")
		return err
	}

	if amount < 0 {
		err := errors.New("Amount should be greater than 0")
		return err
	}

	e[index].Description = desc
	e[index].Amount = amount

	return nil
}
