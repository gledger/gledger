package main

import (
	"encoding/json"
	"net/http"

	"github.com/zombor/go-ledger"
)

type accountTransactionReader interface {
	AccountTransactions(string) []ledger.AccountTransaction
}

type transactionWriter interface {
	AddTransaction(ledger.Transaction) ledger.Transaction
}

type transactionsHandler struct {
	journalReader accountTransactionReader
	journalWriter transactionWriter
}

func (h transactionsHandler) Get(req *http.Request) (int, interface{}, http.Header) {
	values := req.Form

	if values.Get("account") == "" {
		return 400, map[string]string{"message": "Must pass `account` query parameter."}, http.Header{"Content-Type": {"application/vnd.error+json"}}
	}

	accountTransactions := h.journalReader.AccountTransactions(values.Get("account"))

	return 200, accountTransactions, http.Header{"Content-Type": {"application/json"}}
}

func (h transactionsHandler) Post(req *http.Request) (int, interface{}, http.Header) {
	var trans ledger.Transaction

	err := json.NewDecoder(req.Body).Decode(&trans)

	if err != nil {
		return 400, map[string]string{"message": err.Error()}, http.Header{"Content-Type": {"application/vnd.error+json"}}
	}

	t := h.journalWriter.AddTransaction(trans)

	return 200, t, http.Header{"Content-Type": {"application/json"}}
}
