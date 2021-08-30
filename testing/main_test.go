package main

import "testing"

func TestSum(t *testing.T) {
	/* total := Sum(1, 2)
	if total != 3 {
		t.Error("Expected 3, got ", total)
	} */
	tables := []struct {
		a, b  int
		total int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{3, 2, 5},
	}

	for _, table := range tables {
		total := Sum(table.a, table.b)
		if total != table.total {
			t.Error("Expected ", table.total, ", got ", total)
		}
	}
}

func TestGetMax(t *testing.T) {
	items := []struct {
		a, b int
		max  int
	}{
		{1, 2, 2},
		{2, 2, 2},
		{3, 2, 3},
	}
	for _, item := range items {
		max := GetMax(item.a, item.b)
		if max != item.max {
			t.Error("Expected ", item.max, ", got ", max)
		}
	}
}

func TestFibo(t *testing.T) {
	items := []struct {
		n int
		f int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{7, 21},
		{8, 34},
		{9, 55},
		{10, 89},
		{24, 75025},
		{48, 7778742049},
	}
	for _, item := range items {
		fibonacci := Fibonacci(item.n)
		if fibonacci != item.f {
			t.Error("Expected ", item.f, ", got ", fibonacci)
		}
	}
}
