package ledger

import (
	"regexp"
	"strings"
	"time"
)

func (j journal) AccountTransactions(account string) []AccountTransaction {
	// TODO: This probably is bugged if accounts or payees have multiple spaces
	lineSplitter := regexp.MustCompile(`\s\s+`)

	lines := strings.Split(j.accountTransactionReader.AccountTransaction(account), "\n")

	accountTransactions := make([]AccountTransaction, 0)

	for _, l := range lines {
		// TODO: Sometimes a single transaction will have multiple lines and break this loop
		fields := lineSplitter.Split(l, -1)

		if len(fields) == 5 {
			datePayee := strings.SplitN(fields[1], " ", 2)
			date, _ := time.Parse("2006-01-02", datePayee[0])
			accountTransactions = append(accountTransactions, AccountTransaction{
				Id:             fields[0],
				Date:           date,
				Payee:          datePayee[1],
				Account:        fields[2],
				Value:          fields[3],
				RunningBalance: fields[4],
			})
		}
	}

	return accountTransactions
}
