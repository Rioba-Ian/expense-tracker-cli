package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type PosArg struct {
	first string
}

type AddCmdFlags struct {
	Description string
	Amount      int
}

type UpdateCmdFlags struct {
	Id          int
	Description string
	Amount      int
}

type DeleteCmdFlags struct {
	Id int
}

type SummaryCmdFlags struct {
	Month string
}

func NewCmdFlags() *PosArg {
	firstArg := PosArg{}

	flag.Parse()

	if flag.NArg() > 0 {
		firstArg = PosArg{
			first: flag.Arg(0),
		}
		fmt.Println("First positional argument:", firstArg.first)

	}

	return &firstArg
}

func (arg *PosArg) Exec(expenses *Expenses) {
	firstArg := NewCmdFlags()
	switch {
	case firstArg.first == "add":
		NewAddCmdFlags().AddExec(expenses)
	case firstArg.first == "delete":
		NewDelCmdFlags().DeleteExec(expenses)
	case firstArg.first == "summary":
		secondSummaryArg := flag.Arg(1)
		fmt.Println("length of second summary:", len(secondSummaryArg))
		if len(secondSummaryArg) < 1 {
			expenses.summary()
		} else {
			NewSummaryCmdFlags().SummaryExec(expenses)
		}
	case firstArg.first == "update":
		NewUpdateCmdFlags().UpdateExec(expenses)
	case firstArg.first == "list":
		expenses.list()
	default:
		fmt.Println("Invalid Command.")
	}
}

/* Add commands */
func NewAddCmdFlags() *AddCmdFlags {
	addCf := AddCmdFlags{}

	fmt.Println("Positional arguments: ", flag.Args())
	fmt.Println("flag arguments: ", flag.Arg(1))
	fmt.Println("flag arguments: ", flag.Arg(2))

	desc := strings.Split(flag.Arg(1), "=")[1]
	amt := strings.Split(flag.Arg(2), "=")[1]
	convertedAmount, _ := strconv.Atoi(amt)

	addCf = AddCmdFlags{
		Description: desc,
		Amount:      convertedAmount,
	}

	return &addCf
}

func (add *AddCmdFlags) AddExec(expenses *Expenses) {
	expenses.add(add.Description, add.Amount)
}

/* Delete commands */

func NewDelCmdFlags() *DeleteCmdFlags {
	delCfg := DeleteCmdFlags{}
	delStr := strings.Split(flag.Arg(1), "=")[1]
	convertedInt, _ := strconv.Atoi(delStr)

	delCfg = DeleteCmdFlags{
		Id: convertedInt,
	}

	return &delCfg
}

func (del *DeleteCmdFlags) DeleteExec(expenses *Expenses) {
	expenses.delete(del.Id)
}

/* Summary By Month */

func NewSummaryCmdFlags() *SummaryCmdFlags {
	summaryCfg := SummaryCmdFlags{}

	summaryMonth := strings.Split(flag.Arg(1), "=")[1]

	summaryCfg = SummaryCmdFlags{
		Month: summaryMonth,
	}

	return &summaryCfg
}

func (summary *SummaryCmdFlags) SummaryExec(expenses *Expenses) {
	expenses.summaryByMonth(summary.Month)
}

// Update Expense
func NewUpdateCmdFlags() *UpdateCmdFlags {
	updateCfg := UpdateCmdFlags{}

	updateArgs := flag.Args()[1:]
	fmt.Println("update args all>>", updateArgs)
	id, desc, amount := strings.Split(updateArgs[0], "=")[1], strings.Split(updateArgs[1], "=")[1], strings.Split(updateArgs[2], "=")[1]
	fmt.Printf("id: %s, desc: %s, amount: %s\n", id, desc, amount)

	amountToInt, _ := strconv.Atoi(amount)
	idToInt, _ := strconv.Atoi(id)

	updateCfg = UpdateCmdFlags{
		Id:          idToInt,
		Description: desc,
		Amount:      amountToInt,
	}

	return &updateCfg
}

func (update *UpdateCmdFlags) UpdateExec(expenses *Expenses) {
	expenses.update(update.Id, update.Description, update.Amount)
}
