package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Penjumlahan salah, hasil: %d, seharusnya: 10", total)
	}
}
