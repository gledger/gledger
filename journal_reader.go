package ledger

type journalReader interface {
	journalBalanceReader
	journalBudgetReader
}

type journalBalanceReader interface {
	Balance() string
}

type journalBudgetReader interface {
	Budget() string
}
