# Expense Tracker CLI

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/Rioba-Ian/expense-tracker-cli/go.yml)

![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/Rioba-Ian/expense-tracker-cli/main)

Expense Tracker CLI is a command-line interface for managing expenses.
It allows users to add, view, and delete expenses. It also provides a summary of expenses.

## Features

1. Users can add an expense with a description and amount.
2. Users can view all expenses.
3. Users can delete an expense by its ID.
4. Users can update an expense.
5. Users can view a summary of all expenses.
6. Users can view a summary of expenses for a specific month (of current year).

### Additional Features

Some additional features we intend to add in the future

1. Add expense categories and allow users to filter expenses by category.
2. Allow users to set a budget for each month and show a warning when the user exceeds the budget.
3. Allow users to export expenses to a CSV file.✅

### CLI tool usage

```sh
expense-tracker add --description="Lunch" --amount=10.00
expense-tracker list
# ID Date Description Amount
# 1 2023-01-01 Lunch 10.00
# 2 2023-01-02 Dinner 20.00
expense-tracker summary
# Total expenses: $10.00
expense-tracker delete --id=1
# Expense deleted successfully.
expense-tracker summary --month="May"
# Total expenses for August: $10.00

expense-tracker update --id 2 --description="Dinner" --amount=25.00
expense-tracker list
# ID Date Description Amount
# 1 2023-01-01 Lunch 10.00
# 2 2023-01-02 Dinner 25.00
```
