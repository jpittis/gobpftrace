package objdump

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Given a path to a binary, run objdump on it and parse out the addresses
// associated with each of the provided symbols.
func FindAddrsInObjdump(binaryPath string, symbols []string) (map[string]string, error) {
	cmd := exec.Command("objdump", "-t", binaryPath)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return findAddrs(out, symbols)
}

func findAddrs(out []byte, symbols []string) (map[string]string, error) {
	lines := bytes.Split(out, []byte("\n"))
	addrs := map[string]string{}
	for _, symbol := range symbols {
		addr, err := findAddr(lines, symbol)
		if err != nil {
			return nil, err
		}
		addrs[symbol] = addr
	}
	return addrs, nil
}

func findAddr(lines [][]byte, symbol string) (string, error) {
	for _, line := range lines {
		if bytes.Contains(line, []byte(symbol)) {
			return string(bytes.Split(line, []byte(" "))[0]), nil
		}
	}
	return "", fmt.Errorf("Symbol %s not found in objdump", symbol)
}
