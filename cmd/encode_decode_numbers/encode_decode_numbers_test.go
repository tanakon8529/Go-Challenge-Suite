package main

import (
	"os"
	"reflect"
	"testing"
)

func TestDecodeInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
	}{
		{name: "Test 1", input: "LLRR=", expected: []int{2, 1, 0, 1, 2, 2}},
		{name: "Test 2", input: "==RLL", expected: []int{0, 0, 0, 2, 1, 0}},
		{name: "Test 3", input: "=LLRR", expected: []int{0, 0, 1, 0, 1, 2}},
		{name: "Test 4", input: "RRL=R", expected: []int{0, 1, 2, 0, 0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect the standard input to our test input
			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()
			r, w, _ := os.Pipe()
			os.Stdin = r
			w.Write([]byte(tt.input + "\n"))
			w.Close()

			if got := decodeInput(); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("decodeInput() = %v, want %v", got, tt.expected)
			}
		})
	}
}
