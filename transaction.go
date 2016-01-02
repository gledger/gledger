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

type Ledger []Transaction

func (l Ledger) Len() int           { return len(l) }
func (l Ledger) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l Ledger) Less(i, j int) bool { return l[i].Date.Before(l[j].Date) }
