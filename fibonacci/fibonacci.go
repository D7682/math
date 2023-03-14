package fibonacci

type Fibonacci interface {
	Calculate(n int) int
	List(n int) []int
}
