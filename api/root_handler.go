package main

import (
	"net/http"

	"github.com/zombor/go-ledger"
)

type journalReader interface {
	Balances() []ledger.Balance
	Budgets() []ledger.Budget
}

type rootHandler struct {
	journal journalReader
}

func (h rootHandler) Get(*http.Request) (int, interface{}, http.Header) {
	type account struct {
		Name    string `json:"name"`
		Balance string `json:"balance"`
	}
	type out struct {
		Accounts []account `json:"accounts"`
		Budgets  []account `json:"budgets"`
	}

	var output out
	output.Accounts = make([]account, 0)

	for _, b := range h.journal.Balances() {
		output.Accounts = append(output.Accounts, account{Name: b.Name, Balance: b.Value})
	}

	for _, b := range h.journal.Budgets() {
		output.Budgets = append(output.Budgets, account{Name: b.Name, Balance: b.Value})
	}

	return 200, output, http.Header{"Content-type": {"application/json"}}
}
