package ledger

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"
)

type fileReader struct {
	path string
}

func NewFileReader(path string) fileReader {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(fmt.Sprintf(`Journal file "%s" does not exist!`, path))
	}

	return fileReader{path}
}

func (r fileReader) Balance() string {
	out, err := exec.Command("ledger", "balance", "not", "budget", "--flat", "-f", r.path).Output()
	if err != nil {
		panic(err)
	}

	return string(out)
}

func (r fileReader) Budget() string {
	out, err := exec.Command("ledger", "balance", "Budget:Assets", "--flat", "-f", r.path).Output()
	if err != nil {
		panic(err)
	}

	return string(out)
}

func (r fileReader) AccountTransaction(account string) string {
	out, err := exec.Command("ledger", "register", "not", "budget", "and", account, "-S", "d", "-y", `%Y-%m-%d`, "--meta", "ID", "--meta-width", "40", "-f", r.path).Output()
	if err != nil {
		panic(err)
	}

	return string(out)
}

func (r fileReader) Print() string {
	out, err := exec.Command("ledger", "print", "-f", r.path).Output()
	if err != nil {
		panic(err)
	}

	return string(out)
}

func (r fileReader) WriteJournal(trans []Transaction) error {
	t := template.Must(template.New("journal").Parse(journalTemplate))

	var b bytes.Buffer
	err := t.Execute(&b, trans)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(r.path, b.Bytes(), 0644)
}
