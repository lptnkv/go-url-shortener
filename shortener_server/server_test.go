package main

import "testing"

func HelloTest(t *testing.T) {
	got := foo()
	if got != 42 {
		t.Errorf("bullshit")
	}
}
