package knapsack

import (
	"github.com/yourbasic/bit"
	"golang.org/x/exp/constraints"
)

// Number is a numeric primitive constraint.
type Number interface {
	constraints.Integer | constraints.Float
}

// Get01MaxValue solves the 0-1 knapsack problem
// and returns the maximum value achievable with the specified constraints.
// maxWeight is the weight capacity of the knapsack.
// items is the list of items the algorithm can put into the knapsack.
// getWeight is a callback that returns the weight of a given item.
// getValue is a callback that returns the value of a given item.
// Both callbacks MUST be pure functions.
//
// This function runs in O(len(items) * maxWeight) time
// and uses O(maxWeight) space.
func Get01MaxValue[T any, Weight constraints.Unsigned, Value Number](maxWeight Weight, items []T, getWeight func(*T) Weight, getValue func(*T) Value) Value {
	// Perform the dynamic programming 0-1 knapsack algorithm.
	maxValue := make([]Value, maxWeight+1)
	for m := range items {
		itemWeight := getWeight(&items[m])
		itemValue := getValue(&items[m])
		for weight := maxWeight; weight >= itemWeight; weight-- {
			maxValueWithItem := itemValue + maxValue[weight-itemWeight]
			if maxValueWithItem > maxValue[weight] {
				maxValue[weight] = maxValueWithItem
			}
		}
	}
	return maxValue[maxWeight]
}

// max01Value is the maximum value currently achievable
// at a particular weight capacity.
type max01Value[Value Number] struct {
	// maxValue is the current maximum achievable value.
	maxValue Value

	// selectedItems is a bit set
	// containing the indexes of the items that yield maxValue.
	selectedItems bit.Set
}

// Get01Solution solves the 0-1 knapsack problem
// and returns a list of items that yields the maximum value
// given the specified constraints.
// maxWeight is the weight capacity of the knapsack.
// items is the list of items the algorithm can put into the knapsack.
// getWeight is a callback that returns the weight of a given item.
// getValue is a callback that returns the value of a given item.
// Both callbacks MUST be pure functions.
//
// This function runs in O(len(items) * maxWeight) time
// and uses O(len(items) * maxWeight) space.
func Get01Solution[T any, Weight constraints.Unsigned, Value Number](maxWeight Weight, items []T, getWeight func(*T) Weight, getValue func(*T) Value) (selection []T) {
	// Perform the dynamic programming 0-1 knapsack algorithm.
	maxValue := make([]max01Value[Value], maxWeight+1)
	for m := range items {
		itemWeight := getWeight(&items[m])
		itemValue := getValue(&items[m])
		for weight := maxWeight; weight >= itemWeight; weight-- {
			maxValueWithItem := itemValue + maxValue[weight-itemWeight].maxValue
			if maxValueWithItem > maxValue[weight].maxValue {
				maxValue[weight].maxValue = maxValueWithItem
				maxValue[weight].selectedItems.Set(&maxValue[weight-itemWeight].selectedItems).Add(m)
			}
		}
	}
	selection = make([]T, maxValue[maxWeight].selectedItems.Size())[:0]
	maxValue[maxWeight].selectedItems.Visit(func(index int) bool {
		selection = append(selection, items[index])
		return false
	})
	return
}
