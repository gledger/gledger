package ledger

import (
	"os/exec"
)

type fileReader struct {
	path string
}

func NewFileReader(path string) fileReader {
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
	out, err := exec.Command("ledger", "register", "not", "budget", "and", account, "-S", "-d", "-y", `%Y-%m-%d`, "-f", r.path).Output()
	if err != nil {
		panic(err)
	}

	return string(out)
}
