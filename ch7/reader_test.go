package gopl

import (
	"bytes"
	"io"
	"testing"
)

func TestStringReader(t *testing.T) {
	var tests = []struct {
		source string
		want   string
	}{
		{"hello, world\n", "hello, world\n"},
		{"", ""},
		{"\nlfn", "\nlfn"},
	}
	for _, test := range tests {
		t.Run(test.source, func(t *testing.T) {
			r := NewReader(test.source)
			var buf bytes.Buffer
			dst := make([]byte, 2)
			for {
				n, err := r.Read(dst)
				if err == io.EOF {
					break
				} else if err != nil {
					t.Logf(err.Error())
				}
				buf.Write(dst[:n])
			}
			t.Logf("buf:%s", buf.String())
			if got := buf.String(); got != test.want {
				t.Errorf("want = %s, got = %s", test.want, got)
			}
		})
	}
}
