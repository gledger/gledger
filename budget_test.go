package ledger

import (
	"testing"

	"github.com/zombor/go-ledger/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/zombor/go-ledger/Godeps/_workspace/src/github.com/stretchr/testify/require"
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

type funcBudgetReader struct {
	budget func() string
}

func (f funcBudgetReader) Budget() string {
	return f.budget()
}
