package parse

import (
	"os"
	"testing"
)

func TestATags(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{"tests/ex1.html"}},
		{"test2", args{"tests/ex2.html"}},
		{"test3", args{"tests/ex3.html"}},
		{"test4", args{"tests/ex4.html"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.Open(tt.args.filename)
			//TODO: compare std out to expected
			ATags(file)
		})
	}
}
