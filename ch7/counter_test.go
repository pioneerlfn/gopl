package gopl

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestWordCounter(t *testing.T) {
	file, err := ioutil.TempFile(".", "test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	defer file.Close()
	_, err = file.WriteString("hello, world!\nhello, lfn!\n")
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
	var wc WordCounter
	n, err := wc.Counter(data[:rn])
	if err != nil {
		log.Fatal(err)
	}
	if n != 4 {
		t.Fatal("counter error")
	}
}

func TestLineCounter(t *testing.T) {
	file, err := ioutil.TempFile(".", "test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	defer file.Close()
	_, err = file.WriteString("hello, world!\nhello, lfn!\n")
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
	var lc LineCounter
	n, err := lc.Counter(data[:rn])
	if err != nil {
		log.Fatal(err)
	}
	if n != 2 {
		t.Error("counter error: n = ", n)
	}
}
