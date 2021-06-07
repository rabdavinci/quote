package main

import "testing"

func TestFillAndDelete(t *testing.T) {
	var qs Quotes
	qs.FillWithTestData(5)

	if len(qs) != 5 {
		t.Errorf("Expected Quotes length of 5 but got %d", len(qs))
	}

	qs.DeleteByIndex(1)
	if len(qs) != 4 {
		t.Errorf("Expected Quotes length of 4 but got %d", len(qs))
	}
}
