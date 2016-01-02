package ledger

type journal struct {
	balanceReader            journalBalanceReader
	budgetReader             journalBudgetReader
	accountTransactionReader journalAccountTransactionReader
}

func NewJournal(reader journalReader) journal {
	return journal{
		balanceReader:            reader,
		budgetReader:             reader,
		accountTransactionReader: reader,
	}
}
