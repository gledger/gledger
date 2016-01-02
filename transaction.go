package ledger

import (
	"time"
)

type Transaction struct {
	Date     time.Time
	Name     string
	Accounts []Account
}

type Account struct {
	Name, Amount string
}

type TransactionsByDate []Transaction

func (l TransactionsByDate) Len() int           { return len(l) }
func (l TransactionsByDate) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l TransactionsByDate) Less(i, j int) bool { return l[i].Date.Before(l[j].Date) }

type AccountTransaction struct {
	Date           time.Time `json:"date"`
	Payee          string    `json:"payee"`
	Account        string    `json:"account"`
	Value          string    `json:"value"`
	RunningBalance string    `json:"running_balance"`
}
