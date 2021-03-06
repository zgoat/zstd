package zio

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestDumpReader(t *testing.T) {
	cases := []struct {
		in   io.ReadCloser
		want string
	}{
		{
			io.NopCloser(strings.NewReader("Hello")),
			"Hello",
		},
		{
			io.NopCloser(strings.NewReader("لوحة المفاتيح العربية")),
			"لوحة المفاتيح العربية",
		},
		{
			http.NoBody,
			"",
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			outR1, outR2, err := DumpReader(tc.in)
			if err != nil {
				t.Fatal(err)
			}

			out1 := mustRead(t, outR1)
			out2 := mustRead(t, outR2)

			if out1 != tc.want {
				t.Errorf("out1 wrong\nout:  %#v\nwant: %#v\n", out1, tc.want)
			}
			if out2 != tc.want {
				t.Errorf("out2 wrong\nout:  %#v\nwant: %#v\n", out2, tc.want)
			}
		})
	}
}

func mustRead(t *testing.T, r io.Reader) string {
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	return string(out)
}

func TestExists(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{".", true},               // Dir
		{"zio.go", true},          // File
		{"/dev/null", true},       // Device
		{"/proc/1/environ", true}, // Not readable
		{"/etc/localtime", true},  // Symlink

		{"/nonexistent-path", false},
		{"/nonexistent/path", false},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			out := Exists(tt.in)
			if out != tt.want {
				t.Errorf("\nout:  %#v\nwant: %#v\n", out, tt.want)
			}
		})
	}
}
