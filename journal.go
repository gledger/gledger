package ledger

type journal struct {
	balanceReader journalBalanceReader
}

func NewJournal(reader journalReader) journal {
	return journal{
		balanceReader: reader,
	}
}
