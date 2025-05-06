package main

import (
	"bytes"
	"testing"
)

func TestEncodeToString(t *testing.T) {

	tests := []struct {
		input    []byte
		expected string
	}{
		{[]byte("Hello"), "48656C6C6F"},
		{[]byte("123"), "313233"},
		{[]byte{1, 2, 3}, "010203"},
		{[]byte{'a', 'q', 'n', 'o', 'x'}, "61716E6F78"},
	}

	for _, i := range tests {
		result := EncodeToString(i.input)
		if result != i.expected {
			t.Errorf("EncodeToString(%v) = %s; want %s", i.input, result, i.expected)
		}
	}
}

func TestDecodeString(t *testing.T) {

	tests := []struct {
		input            string
		expected         []byte
		expectedErrorMsg string
	}{
		{"48656C6C6F", []byte("Hello"), ""},
		{"313233", []byte("123"), ""},
		{"010203", []byte{1, 2, 3}, ""},
		{"61716E6F78", []byte{'a', 'q', 'n', 'o', 'x'}, ""},
		{"aqnox", nil, "string is not a hex"},
	}

	for _, i := range tests {
		result, err := DecodeString(i.input)

		if i.expectedErrorMsg != "" {
			if err == nil || i.expectedErrorMsg != err.Error() {
				t.Errorf("DecodString(%s) error = %v; want %s", i.input, err, i.expectedErrorMsg)
			}
		} else {
			if !bytes.Equal(result, i.expected) {
				t.Errorf("DecodeString(%s) = %v; want %v", i.input, result, i.expected)
			}

			if err != nil {
				t.Errorf("DecodeString(%s) unexpected error: %v", i.input, err)
			}
		}
	}
}
