package main

import "testing"

func TestGreeting(t *testing.T) {
	result := greeting("hello")
	expected := "<b>hello</b>"
	if result != expected {
		t.Errorf("greeting('hello') %s; want '<b>hello</b>'", result)
	}
}
