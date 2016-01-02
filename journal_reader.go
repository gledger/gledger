package ledger

type journalReader interface {
	journalBalanceReader
	journalBudgetReader
	journalAccountTransactionReader
}

type journalBalanceReader interface {
	Balance() string
}

type journalBudgetReader interface {
	Budget() string
}

type journalAccountTransactionReader interface {
	AccountTransaction(string) string
}
