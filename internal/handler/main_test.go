package handler

import "testing"

func TestDupa(t *testing.T) {
	var result = Dupa("a")
	if result != "aadupa" {
		t.Error("Fail")
	}
}
