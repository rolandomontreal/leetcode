package main

import "testing"

func Test_romanToInt(t *testing.T) {
	s := "I"
	result := romanToInt(s)
	expected := 1
	if result != expected {
		t.Errorf("'%s' was not correctly converted to %d, instead it was converted to %d", s, expected, result)
	}
}