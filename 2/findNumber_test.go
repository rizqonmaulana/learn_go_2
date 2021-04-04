package main

import "testing"

func TestLoop(t *testing.T){
	loop1, _ := loop("12346789")
	if loop1 != 5 {
		t.Error("Result value should be 5. your result test = ", loop1)
	}

	loop2, _ := loop("7891011121415")
	if loop2 != 13 {
		t.Error("Result value should be 13. your result test = ", loop2)
	}

	loop3, _ := loop("2324252627282930313334")
	if loop3 != 32 {
		t.Error("Result value should be 32. your result test = ", loop3)
	}
}