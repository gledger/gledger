package ledger

type journalReader interface {
	journalBalanceReader
}

type journalBalanceReader interface {
	Balance() string
}
