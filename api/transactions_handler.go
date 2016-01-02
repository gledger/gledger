package main

import (
	"net/http"
	"net/url"

	"github.com/zombor/go-ledger"
)

type accountTransactionReader interface {
	AccountTransactions(string) []ledger.AccountTransaction
}

type transactionsHandler struct {
	journal accountTransactionReader
}

func (h transactionsHandler) Get(values url.Values, headers http.Header) (int, interface{}, http.Header) {
	if values.Get("account") == "" {
		return 400, map[string]string{"message": "Must pass `account` query parameter."}, http.Header{"Content-Type": {"application/vnd.error+json"}}
	}

	accountTransactions := h.journal.AccountTransactions(values.Get("account"))

	return 200, accountTransactions, http.Header{"Content-Type": {"application/json"}}
}
