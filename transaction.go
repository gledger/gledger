package ledger

import (
	"time"
)

type Transaction struct {
	Id       string    `json:"id"`
	Date     time.Time `json:"date"`
	Payee    string    `json:"payee"`
	Accounts []Account `json:"accounts"`
}

func (t Transaction) JournalDate() string {
	return t.Date.Format("2006/01/02")
}

type Account struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
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
