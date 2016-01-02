package ledger

import (
	"os/exec"
)

type fileReader struct {
	path string
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

func NewFileReader(path string) fileReader {
	return fileReader{path}
}
