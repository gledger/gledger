# go-ledger

[![Build Status](https://travis-ci.org/zombor/go-ledger.svg)](https://travis-ci.org/zombor/go-ledger)

A golang toolkit and rest API to interface with a ledger journal file

This project aims to satisfy my usage patterns with ledger.
For example, there are many parts of the journal format that I do not use (like check numbers).
I don't plan on supporting these functions, but if you wish to implement them, feel free to submit a pull request.

This project is in very heavy development!

## Toolkit

The `ledger` package wraps go types around specific `ledger-cli` command line interfaces.

The main object is `journal`, which is constructed in the following way:

To construct one, you need to first construct a type that satisfies the `journalReader` interface. We provide a basic file reader in this package. Then you can create a journal object:

```go
fileJournal := NewFileReader("/path/to/ledger/journal.dat")

journal := NewJournal(fileJournal)
```

The `journal` type has the following functions on it:

### Balances

This returns a slice of account balances, not including budget accounts.

```go
balances := j.Balances()
// []Balance{Balance{Name:"Assets:Checking:Test Account",Value:"$123.45"}}
```

### Budgets

This returns a slice of budgets in the journal, which are defined as an account which has a `Budget:Assets` prefix.

```go
budgets := j.Budgets()
// []Budget{Budget{Name:"Budget:Assets:Rent",Value:"$123.45"}}
```

### AccountTransactions

This returns a slice of transactions for an account.

```go
trans := j.AccountTransactions("Assets:Checking:Test Account")
// []AccountTransaction{AccountTransaction{...}}
```

### AddTransaction

This adds a transaction to the ledger file

```go
t := j.AddTransaction(Transaction{...})
```

### DeleteTransaction

This deletes a transaction from the ledger file, specified by the transaction id

```go
err := j.DeleteTransaction("some-id")
```

## HTTP API

Also included in the `api` directory is an http json api server. This uses json input/output over http to interface with your ledger file using the toolkit methods above.

### Supported Endpoints

 - `GET /`
 - `GET /transactions?account=some-account-name`
 - `POST /transactions`
 - `DELETE /transactions?id=a-transaction-id`
