package ledger

type journal struct {
	balanceReader            journalBalanceReader
	budgetReader             journalBudgetReader
	accountTransactionReader journalAccountTransactionReader
	journalPrinter           journalPrinter
	journalWriter            journalWriter
}

func NewJournal(reader journalReader) journal {
	return journal{
		balanceReader:            reader,
		budgetReader:             reader,
		accountTransactionReader: reader,
		journalPrinter:           reader,
		journalWriter:            reader,
	}
}
