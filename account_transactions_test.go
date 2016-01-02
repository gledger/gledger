package ledger

import (
	"testing"
	"time"

	"github.com/zombor/go-ledger/Godeps/_workspace/src/github.com/stretchr/testify/require"
)

func Test_AccountTransactions_Parses_AccountTransactionsOutput(t *testing.T) {
	j := journal{
		accountTransactionReader: funcAccountTransactionReader{
			accountTransactions: func(string) string {
				return `2015-12-29 Opening Balances   Liabilities:American Express   $-700.95           $-700.95
2015-12-31 Amazon         Liabilities:American Express   $-14.32           $-715.27
2016-01-01 Amazon         Liabilities:American Express   $-51.81           $-767.08
2016-01-01 Amazon         Liabilities:American Express   $-18.90           $-785.98
2016-01-01 Fandango.com   Liabilities:American Express   $-40.68           $-826.66
`
			},
		},
	}

	transactions := j.AccountTransactions("anything")

	require.Equal(t, 5, len(transactions))
	require.Equal(t, AccountTransaction{
		Date:           time.Date(2015, time.December, 29, 0, 0, 0, 0, time.UTC),
		Payee:          "Opening Balances",
		Account:        "Liabilities:American Express",
		Value:          "$-700.95",
		RunningBalance: "$-700.95",
	}, transactions[0])
	require.Equal(t, AccountTransaction{
		Date:           time.Date(2015, time.December, 31, 0, 0, 0, 0, time.UTC),
		Payee:          "Amazon",
		Account:        "Liabilities:American Express",
		Value:          "$-14.32",
		RunningBalance: "$-715.27",
	}, transactions[1])
	require.Equal(t, AccountTransaction{
		Date:           time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		Payee:          "Amazon",
		Account:        "Liabilities:American Express",
		Value:          "$-51.81",
		RunningBalance: "$-767.08",
	}, transactions[2])
	require.Equal(t, AccountTransaction{
		Date:           time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		Payee:          "Amazon",
		Account:        "Liabilities:American Express",
		Value:          "$-18.90",
		RunningBalance: "$-785.98",
	}, transactions[3])
	require.Equal(t, AccountTransaction{
		Date:           time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		Payee:          "Fandango.com",
		Account:        "Liabilities:American Express",
		Value:          "$-40.68",
		RunningBalance: "$-826.66",
	}, transactions[4])
}

type funcAccountTransactionReader struct {
	accountTransactions func(string) string
}

func (f funcAccountTransactionReader) AccountTransaction(a string) string {
	return f.accountTransactions(a)
}
