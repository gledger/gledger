package ledger

import (
	"regexp"
	"strings"
	"time"

	"github.com/zombor/go-ledger/Godeps/_workspace/src/github.com/twinj/uuid"
)

func (j journal) AddTransaction(t Transaction) Transaction {
	t.Id = uuid.NewV4().String()

	existingTransactions := j.parsePrintedTransactions(j.journalPrinter.Print())

	existingTransactions = append(existingTransactions, t)

	err := j.journalWriter.WriteJournal(existingTransactions)
	if err != nil {
		panic(err)
	}

	return t
}

func (j journal) parsePrintedTransactions(trans string) []Transaction {
	transactions := make([]Transaction, 0)
	splitter := regexp.MustCompile(`\s\s+`)

	if trans == "" {
		return transactions
	}

	lines := strings.Split(strings.TrimSpace(trans), "\n\n")

	for _, printedTrans := range lines {
		var t Transaction
		var err error
		t.Accounts = make([]Account, 0)

		lines := strings.Split(printedTrans, "\n")

		firstLine := strings.SplitN(lines[0], " ", 2)

		t.Payee = firstLine[1]
		parsedTime, err := time.Parse("2006/01/02", firstLine[0])
		if err != nil {
			panic(err)
		}
		t.Date = TransactionDate(parsedTime)

		for _, a := range lines[1:] {
			a = strings.TrimLeft(a, " ")

			if strings.HasPrefix(a, ";") {
				idFinder := regexp.MustCompile(`; ID: (?P<Id>.+)`)
				matches := idFinder.FindStringSubmatch(a)

				if len(matches) == 2 {
					t.Id = matches[1]
				}
			} else {
				var account Account

				parts := splitter.Split(a, 2)

				account.Name = parts[0]

				if len(parts) == 2 {
					account.Amount = parts[1]
				}

				t.Accounts = append(t.Accounts, account)
			}
		}

		transactions = append(transactions, t)
	}

	return transactions
}
