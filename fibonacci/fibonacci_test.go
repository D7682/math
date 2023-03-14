package fibonacci

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/*
TestCachedFibonacci

Tests the `Calculate` and `List` methods of the `Cached` implementation.
This function will also output the contents of the cache for the `Cached` implementation
after the tests run.
*/
func TestCachedFibonacci(t *testing.T) {
	fib := NewCached()

	// Test Calculate method
	unCached := NewUncached()
	for i := 0; i <= 50; i++ {
		expected := unCached.Calculate(i)
		actual := fib.Calculate(i)

		if actual != expected {
			t.Errorf("Incorrect result for %d: expected %d, but got %d", i, expected, actual)
		}
	}

	// Test List method
	expectedList := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025}
	actualList := fib.List(25)

	if !reflect.DeepEqual(actualList, expectedList) {
		t.Errorf("Incorrect list: expected %v, but got %v", expectedList, actualList)
	}

	fmt.Println("Cache contents:")
	keys := make([]int, len(fib.cache))
	i := 0
	for k := range fib.cache {
		keys[i] = k
		i++
	}

	sort.Ints(keys)

	// Access the map values in order
	for _, k := range keys {
		val := fib.cache[k]
		fmt.Printf("n: %d, value: %d\n", k, val)
	}
}

/*
TestUncachedFibonacci

Tests the `Calculate` and `List` methods of the `Uncached` implementation.
*/
func TestUncachedFibonacci(t *testing.T) {
	fib := NewUncached()

	// Test Calculate method
	unCached := NewUncached()
	for i := 0; i <= 50; i++ {
		expected := unCached.Calculate(i)
		actual := fib.Calculate(i)

		if actual != expected {
			t.Errorf("Incorrect result for %d: expected %d, but got %d", i, expected, actual)
		}
	}

	// Test List method
	expectedList := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025}
	actualList := fib.List(25)

	if !reflect.DeepEqual(actualList, expectedList) {
		t.Errorf("Incorrect list: expected %v, but got %v", expectedList, actualList)
	}
}
