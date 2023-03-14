package fibonacci

type Cached struct {
	cache map[int]int
}

/*
List
Calls the `Calculate` method for each number up to `n` and adds it to an array that is returned as the result.
Note that the implementation of `List` is specific to the `CachedFibonacci` struct, so if you have a different implementation of the
`Fibonacci` interface you will need to implement `List` accordingly.
*/
func (c *Cached) List(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		result[i] = c.Calculate(i)
	}
	return result
}

func NewCached() *Cached {
	return &Cached{cache: make(map[int]int)}
}

func (c *Cached) Calculate(n int) int {
	if n <= 1 {
		return n
	}

	if val, ok := c.cache[n]; ok {
		return val
	}

	val := c.Calculate(n-1) + c.Calculate(n-2)
	c.cache[n] = val

	return val
}
