package ledger

import "time"

type journalReader interface {
	journalBalanceReader
	journalBudgetReader
	journalAccountTransactionReader
	journalPrinter
	journalWriter
}

type journalBalanceReader interface {
	Balance() string
}

type journalBudgetReader interface {
	Budget() string
	BudgetInDate(time.Time, time.Time) string
}

type journalAccountTransactionReader interface {
	AccountTransaction(string) string
}

type journalPrinter interface {
	Print() string
}

type journalWriter interface {
	WriteJournal([]Transaction) error
}
