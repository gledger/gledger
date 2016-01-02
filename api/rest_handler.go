package main

import (
	"encoding/json"
	"net/http"

	"github.com/zombor/ledger"
)

type journalReader interface {
	Balances() []ledger.Balance
	Budgets() []ledger.Budget
}

type restHandler struct {
	journal journalReader
}

func (h restHandler) Root(res http.ResponseWriter, req *http.Request) {
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

	json.NewEncoder(res).Encode(output)
}
