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
	Amount      string
}

type DeleteCmdFlags struct {
	Id int
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
		expenses.summary()
	case firstArg.first == "list":
		expenses.list()
	default:
		fmt.Println("Add not entered.")
	}
	fmt.Println("List commands")
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

/* Summary */
