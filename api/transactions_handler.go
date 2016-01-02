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

type transactionDeleter interface {
	DeleteTransaction(string) error
}

type transactionsHandler struct {
	journalReader  accountTransactionReader
	journalWriter  transactionWriter
	journalDeleter transactionDeleter
}

func (h transactionsHandler) Get(req *http.Request) (int, interface{}, http.Header) {
	values := req.Form

	if values.Get("account") == "" {
		return http.StatusBadRequest, map[string]string{"message": "Must pass `account` query parameter."}, http.Header{"Content-Type": {"application/vnd.error+json"}}
	}

	accountTransactions := h.journalReader.AccountTransactions(values.Get("account"))

	return http.StatusOK, accountTransactions, http.Header{"Content-Type": {"application/json"}}
}

func (h transactionsHandler) Post(req *http.Request) (int, interface{}, http.Header) {
	var trans ledger.Transaction

	err := json.NewDecoder(req.Body).Decode(&trans)

	if err != nil {
		return http.StatusBadRequest, map[string]string{"message": err.Error()}, http.Header{"Content-Type": {"application/vnd.error+json"}}
	}

	t := h.journalWriter.AddTransaction(trans)

	return http.StatusOK, t, http.Header{"Content-Type": {"application/json"}}
}

func (h transactionsHandler) Delete(req *http.Request) (int, interface{}, http.Header) {
	values := req.Form

	if values.Get("id") == "" {
		return http.StatusBadRequest, map[string]string{"message": "Must pass `id` query parameter."}, http.Header{"Content-Type": {"application/vnd.error+json"}}
	}

	err := h.journalDeleter.DeleteTransaction(values.Get("id"))

	if err != nil {
		return http.StatusInternalServerError, map[string]string{"message": err.Error()}, http.Header{"Content-Type": {"application/vnd.error+json"}}
	}

	return http.StatusOK, nil, http.Header{}
}
