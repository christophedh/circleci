package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	expect := hello("hello")
	if expect != "hello you" {
		t.Errorf("hello(hello); want \"hello you\" got \"%s\"", expect)
	}
}
