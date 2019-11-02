package gopl

import (
	"bufio"
	"bytes"
)

type LineCounter int

func (lc *LineCounter) Counter(data []byte) (int, error) {
	r := bytes.NewReader(data)
	scanner := bufio.NewScanner(r)
	// scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*lc++
		/*if scanner.Err() != nil {
			return int(*wc), scanner.Err()
		}*/
	}
	return int(*lc), nil
}
