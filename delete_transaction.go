package ledger

func (j journal) DeleteTransaction(id string) error {
	newTransactions := make([]Transaction, 0)

	for _, t := range j.parsePrintedTransactions(j.journalPrinter.Print()) {
		if t.Id != id {
			newTransactions = append(newTransactions, t)
		}
	}

	return j.journalWriter.WriteJournal(newTransactions)
}
