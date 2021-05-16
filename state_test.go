package script

import (
	"bytes"
	"testing"
)

func TestGetToken(t *testing.T) {
	tests := [][]byte{[]byte("1"), []byte(" 1"), []byte("1 ")}
	for _, test := range tests {
		var token []byte
		token, test = GetToken(test)
		if !bytes.Equal(token, []byte("1")) {
			t.Error("Token not parsed correctly")
		}
		token, test = GetToken(test)
		if token != nil || test != nil {
			t.Error("Found more than one token")
		}
	}

}
