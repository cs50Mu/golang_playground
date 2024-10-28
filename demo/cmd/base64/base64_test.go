package main

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	var err error
	b := NewBase64()

	fmt.Printf("Character at index 28: %c\n", b.charAt(28))

	// s := "Hi"
	s := "Encode to Base64 format"

	got := b.Encode([]byte(s))
	want := "RW5jb2RlIHRvIEJhc2U2NCBmb3JtYXQ="
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}

	// decode
	var dgot []byte
	dgot, err = b.Decode(want)
	if err != nil {
		panic(err)
	}
	got = string(dgot)
	want = s
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}

	s = "RG8geW91IGhhdmUgdG8gZGVhbCB3aXRoIEJhc2U2NCBmb3JtYXQ/IFRoZW4gdGhpcyBzaXRlIGlzIHBlcmZlY3QgZm9yIHlvdSEgVXNlIG91ciBzdXBlciBoYW5keSBvbmxpbmUgdG9vbCB0byBlbmNvZGUgb3IgZGVjb2RlIHlvdXIgZGF0YS4="
	dgot, err = b.Decode(s)
	if err != nil {
		panic(err)
	}
	got = string(dgot)
	want = "Do you have to deal with Base64 format? Then this site is perfect for you! Use our super handy online tool to encode or decode your data."
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
