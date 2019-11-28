package main

import "testing"

func TestHello(t *testing.T) {
	got := say()
	actual := "hello"
	if got != actual {
		t.Fatal("no hello")
	}
}
