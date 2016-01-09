package ledger

import (
	"testing"
	"time"

	"github.com/gledger/gledger/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/gledger/gledger/Godeps/_workspace/src/github.com/stretchr/testify/require"
)

func Test_Budgets_Parses_BudgetOutput(t *testing.T) {
	j := journal{
		budgetReader: funcBudgetReader{
			budget: func() string {
				return `    $37.81  Budget:Assets:Phone
             $389.75  Budget:Assets:Rent
             $902.43  Budget:Assets:Groceries
             $459.32  Budget:Assets:Spending Money
--------------------
                   0
`
			},
		},
	}

	budgets := j.Budgets()

	require.Equal(t, 4, len(budgets))
	assert.Equal(t, Budget{Name: "Budget:Assets:Phone", Value: "$37.81"}, budgets[0])
	assert.Equal(t, Budget{Name: "Budget:Assets:Rent", Value: "$389.75"}, budgets[1])
	assert.Equal(t, Budget{Name: "Budget:Assets:Groceries", Value: "$902.43"}, budgets[2])
	assert.Equal(t, Budget{Name: "Budget:Assets:Spending Money", Value: "$459.32"}, budgets[3])
}

func Test_Budgets_Parses_BudgetInDateOutput(t *testing.T) {
	j := journal{
		budgetReader: funcBudgetReader{
			budgetInDate: func(time.Time, time.Time) string {
				return `    $37.81  Budget:Assets:Phone
             $389.75  Budget:Assets:Rent
             $902.43  Budget:Assets:Groceries
             $459.32  Budget:Assets:Spending Money
--------------------
                   0
`
			},
		},
	}

	budgets := j.BudgetsInDate(time.Now(), time.Now())

	require.Equal(t, 4, len(budgets))
	assert.Equal(t, Budget{Name: "Budget:Assets:Phone", Value: "$37.81"}, budgets[0])
	assert.Equal(t, Budget{Name: "Budget:Assets:Rent", Value: "$389.75"}, budgets[1])
	assert.Equal(t, Budget{Name: "Budget:Assets:Groceries", Value: "$902.43"}, budgets[2])
	assert.Equal(t, Budget{Name: "Budget:Assets:Spending Money", Value: "$459.32"}, budgets[3])
}

type funcBudgetReader struct {
	budget       func() string
	budgetInDate func(time.Time, time.Time) string
}

func (f funcBudgetReader) Budget() string {
	return f.budget()
}

func (f funcBudgetReader) BudgetInDate(start, end time.Time) string {
	return f.budgetInDate(start, end)
}
