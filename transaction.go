package ledger

import (
	"strings"
	"time"
)

type TransactionDate time.Time

func (t *TransactionDate) UnmarshalJSON(buf []byte) error {
	tt, err := time.Parse("2006-01-02", strings.Trim(string(buf), `"`))
	if err != nil {
		return err
	}
	*t = TransactionDate(tt)
	return nil
}

func (t TransactionDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format("2006-01-02") + `"`), nil
}

type Transaction struct {
	Id       string          `json:"id"`
	Date     TransactionDate `json:"date"`
	Payee    string          `json:"payee"`
	Accounts []Account       `json:"accounts"`
}

func (t Transaction) JournalDate() string {
	return time.Time(t.Date).Format("2006/01/02")
}

type Account struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
}

type AccountTransaction struct {
	Id             string    `json:"id"`
	Date           time.Time `json:"date"`
	Payee          string    `json:"payee"`
	Account        string    `json:"account"`
	Value          string    `json:"value"`
	RunningBalance string    `json:"running_balance"`
}
