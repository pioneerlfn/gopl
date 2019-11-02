package gopl

import (
	"bufio"
	"bytes"
)

type WordCounter int

func (wc *WordCounter) Counter(data []byte) (int, error) {
	r := bytes.NewReader(data)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*wc++
		/*if scanner.Err() != nil {
			return int(*wc), scanner.Err()
		}*/
	}
	return int(*wc), nil
}
