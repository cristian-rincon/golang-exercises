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
