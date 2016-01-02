package ledger

import (
	"regexp"
	"strings"
	"time"
)

func ParseLedger(input string) []Transaction {
	splitter := regexp.MustCompile(`\n\n+`)
	transactionLines := splitter.Split(input, -1)

	transactions := make([]Transaction, 0)

	for i := 0; i < len(transactionLines); i++ {
		lines := strings.Split(transactionLines[i], "\n")
		date, name := parseTitleLine(lines[0])
		accounts := parseAccounts(lines[1:])

		transactions = append(transactions, Transaction{
			Date:     date,
			Name:     name,
			Accounts: accounts,
		})
	}

	return transactions
}

func parseTitleLine(input string) (time.Time, string) {
	parts := strings.SplitN(input, " ", 2)

	t, _ := time.Parse("2006/01/02", parts[0])

	return t, parts[1]
}

func parseAccounts(input []string) []Account {
	accounts := make([]Account, 0)

	splitter := regexp.MustCompile(`\s\s+`)

	for i := 0; i < len(input); i++ {
		var name, amount string

		parts := splitter.Split(input[i], 2)
		name = parts[0]
		if len(parts) == 2 {
			amount = parts[1]
		}

		accounts = append(accounts, Account{
			Name:   name,
			Amount: amount,
		})
	}

	return accounts
}
