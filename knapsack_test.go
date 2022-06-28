package knapsack

import (
	"testing"
)

type testItem struct {
	weight uint64
	value  int64
}

func getTestItemWeight(ti *testItem) uint64 {
	return ti.weight
}

func getTestItemValue(ti *testItem) int64 {
	return ti.value
}

func getTotalValue(ti []testItem) int64 {
	value := int64(0)
	for _, i := range ti {
		value += i.value
	}
	return value
}

func TestGet01MaxValue(t *testing.T) {
	items := []testItem{
		{weight: 10, value: 60},
		{weight: 20, value: 100},
		{weight: 30, value: 120},
	}
	if v := Get01MaxValue(uint64(50), items, getTestItemWeight, getTestItemValue); v != 220 {
		t.Errorf(`unexpected value: %v`, v)
	}

	if v := Get01MaxValue(uint64(0), items, getTestItemWeight, getTestItemValue); v != 0 {
		t.Errorf(`unexpected value: %v`, v)
	}
	if v := Get01MaxValue(uint64(1000), items[:0], getTestItemWeight, getTestItemValue); v != 0 {
		t.Errorf(`unexpected value: %v`, v)
	}
	if v := Get01MaxValue(uint64(9), items, getTestItemWeight, getTestItemValue); v != 0 {
		t.Errorf(`unexpected value: %v`, v)
	}
	if v := Get01MaxValue(uint64(10), items, getTestItemWeight, getTestItemValue); v != 60 {
		t.Errorf(`unexpected value: %v`, v)
	}
	if v := Get01MaxValue(uint64(19), items, getTestItemWeight, getTestItemValue); v != 60 {
		t.Errorf(`unexpected value: %v`, v)
	}
	if v := Get01MaxValue(uint64(20), items, getTestItemWeight, getTestItemValue); v != 100 {
		t.Errorf(`unexpected value: %v`, v)
	}
	if v := Get01MaxValue(uint64(30), items, getTestItemWeight, getTestItemValue); v != 160 {
		t.Errorf(`unexpected value: %v`, v)
	}

	items = []testItem{
		{weight: 1, value: 10},
		{weight: 2, value: 15},
		{weight: 3, value: 40},
	}
	if v := Get01MaxValue(uint64(6), items, getTestItemWeight, getTestItemValue); v != 65 {
		t.Errorf(`unexpected value: %v`, v)
	}

	items = []testItem{
		{weight: 30, value: 10},
		{weight: 10, value: 20},
		{weight: 40, value: 30},
		{weight: 20, value: 40},
	}
	if v := Get01MaxValue(uint64(40), items, getTestItemWeight, getTestItemValue); v != 60 {
		t.Errorf(`unexpected value: %v`, v)
	}

	items = []testItem{
		{weight: 95, value: 55},
		{weight: 4, value: 10},
		{weight: 60, value: 47},
		{weight: 32, value: 5},
		{weight: 23, value: 4},
		{weight: 72, value: 50},
		{weight: 80, value: 8},
		{weight: 62, value: 61},
		{weight: 65, value: 85},
		{weight: 46, value: 87},
	}
	if v := Get01MaxValue(uint64(269), items, getTestItemWeight, getTestItemValue); v != 295 {
		t.Errorf(`unexpected value: %v`, v)
	}

	items = []testItem{
		{weight: 92, value: 44},
		{weight: 4, value: 46},
		{weight: 43, value: 90},
		{weight: 83, value: 72},
		{weight: 84, value: 91},
		{weight: 68, value: 40},
		{weight: 92, value: 75},
		{weight: 82, value: 35},
		{weight: 6, value: 8},
		{weight: 44, value: 54},
		{weight: 32, value: 78},
		{weight: 18, value: 40},
		{weight: 56, value: 77},
		{weight: 83, value: 15},
		{weight: 25, value: 61},
		{weight: 96, value: 17},
		{weight: 70, value: 75},
		{weight: 48, value: 29},
		{weight: 14, value: 75},
		{weight: 58, value: 63},
	}
	if v := Get01MaxValue(uint64(878), items, getTestItemWeight, getTestItemValue); v != 1024 {
		t.Errorf(`unexpected value: %v`, v)
	}
}

func TestGet01Solution(t *testing.T) {
	items := []testItem{
		{weight: 30, value: 10},
		{weight: 10, value: 20},
		{weight: 40, value: 30},
		{weight: 20, value: 40},
	}
	expected := []testItem{items[1], items[3]}
	solution := Get01Solution(uint64(40), items, getTestItemWeight, getTestItemValue)
	if getTotalValue(solution) != 60 {
		t.Errorf(`unexpected value: %v`, getTotalValue(solution))
	} else if len(solution) != 2 {
		t.Errorf(`unexpected items: %v`, solution)
	} else if solution[0] != expected[0] || solution[1] != expected[1] {
		t.Errorf(`expected items %v but got %v`, expected, solution)
	}

	if solution = Get01Solution(uint64(0), items, getTestItemWeight, getTestItemValue); len(solution) != 0 {
		t.Errorf(`unexpected solution: %v`, solution)
	}

	if solution = Get01Solution(uint64(9), items, getTestItemWeight, getTestItemValue); len(solution) != 0 {
		t.Errorf(`unexpected solution: %v`, solution)
	}

	if solution = Get01Solution(uint64(10), items, getTestItemWeight, getTestItemValue); len(solution) != 1 {
		t.Errorf(`unexpected solution: %v`, solution)
	} else if getTotalValue(solution) != 20 {
		t.Errorf(`unexpected value: %v`, getTotalValue(solution))
	} else if solution[0] != items[1] {
		t.Errorf(`unexpected solution: %v`, solution)
	}
}
