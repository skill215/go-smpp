package pdutext

import (
	"bytes"
	"testing"
)

func TestBinary2Encoder(t *testing.T) {
	text := []byte("Olá mundão")
	want := text
	s := Binary2(text)
	if s.Type() != 0x04 {
		t.Fatalf("Unexpected data type; want 0x04, have %d", s.Type())
	}
	have := s.Encode()
	if !bytes.Equal(want, have) {
		t.Fatalf("Unexpected text; want %q, have %q", want, have)
	}
}

func TestBinary2Decoder(t *testing.T) {
	text := []byte("Olá mundão")
	want := text
	s := Binary2(text)
	if s.Type() != 0x04 {
		t.Fatalf("Unexpected data type; want 0x04, have %d", s.Type())
	}
	have := s.Decode()
	if !bytes.Equal(want, have) {
		t.Fatalf("Unexpected text; want %q, have %q", want, have)
	}
}
