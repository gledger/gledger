package ledger

type journal struct {
	balanceReader journalBalanceReader
	budgetReader  journalBudgetReader
}

func NewJournal(reader journalReader) journal {
	return journal{
		balanceReader: reader,
		budgetReader:  reader,
	}
}
