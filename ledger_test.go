package ledger

import (
	"testing"
	"time"
)

func checkTransaction(expected, actual Transaction, t *testing.T) {
	if actual.Date != expected.Date {
		t.Errorf("Expected Date to be %s, but was %s", expected.Date, actual.Date)
	}

	if actual.Payee != expected.Payee {
		t.Errorf("Expected Payee to be `%s`, but was `%s`", expected.Payee, actual.Payee)
	}

	if len(actual.Accounts) != len(expected.Accounts) {
		t.Errorf("Expected %d accounts, but got %d", len(expected.Accounts), len(actual.Accounts))
	} else {
		for i := 0; i < len(actual.Accounts); i++ {
			if actual.Accounts[i] != expected.Accounts[i] {
				t.Errorf("Expected account %d to be %#v, got %#v", i, expected.Accounts[i], actual.Accounts[i])
			}
		}
	}
}

func Test_Basic_Parse(t *testing.T) {
	input := `2016/01/01 Pacific Bell
Expenses:Utilities:Phone  $23.00
Assets:Checking`

	transactions := ParseLedger(input)

	if len(transactions) != 1 {
		t.Errorf("Expected 1 parsed transaction, got %d", len(transactions))
	} else {
		actual := transactions[0]

		checkTransaction(
			Transaction{
				Date:  time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
				Payee: "Pacific Bell",
				Accounts: []Account{
					Account{Name: "Expenses:Utilities:Phone", Amount: "$23.00"},
					Account{Name: "Assets:Checking", Amount: ""},
				},
			},
			actual,
			t,
		)
	}
}

func Test_Multiple_Parse(t *testing.T) {
	input := `2016/01/01 Pacific Bell
Expenses:Utilities:Phone  $23.00
Assets:Checking

2016/01/02 ComEd
Expenses:Utilities:Electricity  $54.32
Liabilities:American Express`

	transactions := ParseLedger(input)

	if len(transactions) != 2 {
		t.Errorf("Expected 2 parsed transaction, got %d", len(transactions))
	} else {
		actual1 := transactions[0]
		actual2 := transactions[1]

		checkTransaction(
			Transaction{
				Date:  time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
				Payee: "Pacific Bell",
				Accounts: []Account{
					Account{Name: "Expenses:Utilities:Phone", Amount: "$23.00"},
					Account{Name: "Assets:Checking", Amount: ""},
				},
			},
			actual1,
			t,
		)

		checkTransaction(
			Transaction{
				Date:  time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC),
				Payee: "ComEd",
				Accounts: []Account{
					Account{Name: "Expenses:Utilities:Electricity", Amount: "$54.32"},
					Account{Name: "Liabilities:American Express", Amount: ""},
				},
			},
			actual2,
			t,
		)
	}
}
