package ledger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Balances_Parses_BalanceOutput(t *testing.T) {
	j := journal{
		balanceReader: funcBalanceReader{
			balance: func() string {
				return `$3,456.80  Assets:Checking:My Account
          $12,876.60  Assets:Retirement:401K
           $5,678.00  Assets:Savings:Emergency
          $-7,737.72  Equity:Opening Balances
              $32.03  Expenses:Groceries
          $-2,715.69  Income:Salary
          $-1,858.69  Liabilities:American Express
--------------------
                   0
`
			},
		},
	}

	balances := j.Balances()

	require.Equal(t, 7, len(balances))
	assert.Equal(t, Balance{Name: "Assets:Checking:My Account", Value: "$3,456.80"}, balances[0])
	assert.Equal(t, Balance{Name: "Assets:Retirement:401K", Value: "$12,876.60"}, balances[1])
	assert.Equal(t, Balance{Name: "Assets:Savings:Emergency", Value: "$5,678.00"}, balances[2])
	assert.Equal(t, Balance{Name: "Equity:Opening Balances", Value: "$-7,737.72"}, balances[3])
	assert.Equal(t, Balance{Name: "Expenses:Groceries", Value: "$32.03"}, balances[4])
	assert.Equal(t, Balance{Name: "Income:Salary", Value: "$-2,715.69"}, balances[5])
	assert.Equal(t, Balance{Name: "Liabilities:American Express", Value: "$-1,858.69"}, balances[6])
}

type funcBalanceReader struct {
	balance func() string
}

func (f funcBalanceReader) Balance() string {
	return f.balance()
}
