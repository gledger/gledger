package ledger

import (
	"strings"
)

type Balance struct {
	Name, Value string
}

func (j journal) Balances() []Balance {
	lines := strings.Split(j.balanceReader.Balance(), "\n")

	balances := make([]Balance, 0)

	for i := 0; i < len(lines); i++ {
		line := strings.SplitN(strings.TrimLeft(lines[i], " "), "  ", 2)

		if len(line) == 2 {
			balances = append(balances, Balance{Name: line[1], Value: line[0]})
		}
	}

	return balances
}
