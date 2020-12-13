package main

import "testing"

func TestQueen(t *testing.T) {
}

func TestValid(t *testing.T) {
	b := Board{Queen{0, 0}}
	if !b.Valid() {
		t.Errorf("%b expected to be valid, but was not", b)
	}

	b = Board{Queen{0, 0}, Queen{0, 0}}
	if b.Valid() {
		t.Errorf("two queens on top of each other are not valid")
	}
}
