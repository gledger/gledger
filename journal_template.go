package ledger

const journalTemplate = `{{range .}}{{.JournalDate}} {{.Payee}}
  ; ID: {{.Id}}
  {{range .Accounts}}{{.Name}}  {{.Amount}}
  {{end}}
{{end}}`
