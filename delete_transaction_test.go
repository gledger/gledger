package ledger

import (
	"testing"
	"time"

	"github.com/zombor/go-ledger/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func Test_DeleteTransaction_DeletesTransactionById(t *testing.T) {
	writeJournalCalled := false

	j := journal{
		journalPrinter: funcJournalPrinter{
			printJournal: func() string {
				return `2016/01/01 Amazon
    ; ID: abcdef
    Liabilities:American Express             $-14.32
    Expenses:Home Maintenance

2016/01/02 Amazon
    ; ID: bcdef
    Liabilities:American Express             $-4.67
    Expenses:Home Maintenance
`
			},
		},
		journalWriter: funcJournalWriter{
			writeJournal: func(ts []Transaction) error {
				writeJournalCalled = true

				assert.Equal(t,
					[]Transaction{
						Transaction{
							Id:    "bcdef",
							Date:  TransactionDate(time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC)),
							Payee: "Amazon",
							Accounts: []Account{
								Account{
									Name:   "Liabilities:American Express",
									Amount: "$-4.67",
								},
								Account{
									Name: "Expenses:Home Maintenance",
								},
							},
						},
					},
					ts,
				)

				return nil
			},
		},
	}

	j.DeleteTransaction("abcdef")

	assert.True(t, writeJournalCalled)
}

type funcJournalPrinter struct {
	printJournal func() string
}

func (f funcJournalPrinter) Print() string {
	return f.printJournal()
}

type funcJournalWriter struct {
	writeJournal func([]Transaction) error
}

func (f funcJournalWriter) WriteJournal(ts []Transaction) error {
	return f.writeJournal(ts)
}
