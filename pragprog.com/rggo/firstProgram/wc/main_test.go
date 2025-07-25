package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4
	res := count(b, false, false) // false indicates we want to count words

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word4")
	exp := 3
	res := count(b, true, false) // true indicates we want to count lines

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word4")
	exp := 35
	res := count(b, false, true) // true indicates we want to count bytes

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
