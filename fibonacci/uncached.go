package fibonacci

type Uncached struct{}

func NewUncached() *Uncached {
	return &Uncached{}
}

func (u *Uncached) Calculate(n int) int {
	if n <= 1 {
		return n
	}

	return u.Calculate(n-1) + u.Calculate(n-2)
}

/*
List
will create an array of size `n`, and then loop through the first `n` fibonacci numbers and calculate each one using the `Calculate` method.
The result is stored in the `result` array and returned as the final output.
*/
func (u *Uncached) List(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		result[i] = u.Calculate(i)
	}
	return result
}
