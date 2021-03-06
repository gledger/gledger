package ledger

import (
	"strings"
	"time"
)

type Budget struct {
	Name, Value string
}

func (j journal) Budgets() []Budget {
	lines := strings.Split(j.budgetReader.Budget(), "\n")

	budgets := make([]Budget, 0)

	for i := 0; i < len(lines); i++ {
		line := strings.SplitN(strings.TrimLeft(lines[i], " "), "  ", 2)

		if len(line) == 2 {
			budgets = append(budgets, Budget{Name: line[1], Value: line[0]})
		}
	}

	return budgets
}

func (j journal) BudgetsInDate(start, end time.Time) []Budget {
	lines := strings.Split(j.budgetReader.BudgetInDate(start, end), "\n")

	budgets := make([]Budget, 0)

	for i := 0; i < len(lines); i++ {
		line := strings.SplitN(strings.TrimLeft(lines[i], " "), "  ", 2)

		if len(line) == 2 {
			budgets = append(budgets, Budget{Name: line[1], Value: line[0]})
		}
	}

	return budgets
}
