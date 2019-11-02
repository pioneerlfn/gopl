package gopl

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestWordCounter(t *testing.T) {
	var tests = []struct {
		raw  string
		want int
	}{
		{"hello, world!\nhello, lfn!\n", 4},
		{"hello, i am learning test in golang!\n", 7},
	}
	for _, test := range tests {
		file, err := ioutil.TempFile(".", "test")
		if err != nil {
			log.Fatal(err)
		}
		_, err = file.WriteString(test.raw)
		if err != nil {
			log.Fatal(err)
		}
		file.Seek(0, io.SeekStart)
		data := make([]byte, 128)
		rn, err := file.Read(data)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
		}
		name := file.Name()
		file.Close()
		os.Remove(name)
		var wc WordCounter
		got, err := wc.Counter(data[:rn])
		if err != nil {
			log.Fatal(err)
		}
		if got != test.want {
			t.Errorf("want = %d, got = %d\n", test.want, got)
		}
	}
}

func TestLineCounter(t *testing.T) {
	var tests = []struct {
		raw  string
		want int
	}{
		{"hello, world!\nhello, lfn!\n", 2},
		{"hello, i am learning test in golang!\n", 1},
	}
	for _, test := range tests {
		file, err := ioutil.TempFile(".", "test")
		if err != nil {
			log.Fatal(err)
		}
		_, err = file.WriteString(test.raw)
		if err != nil {
			log.Fatal(err)
		}
		file.Seek(0, io.SeekStart)
		data := make([]byte, 128)
		rn, err := file.Read(data)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
		}
		name := file.Name()
		file.Close()
		os.Remove(name)
		var lc LineCounter
		got, err := lc.Counter(data[:rn])
		if err != nil {
			log.Fatal(err)
		}
		if got != test.want {
			t.Errorf("want = %d, got = %d\n", test.want, got)
		}
	}
}
