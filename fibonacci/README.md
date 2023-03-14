# Fibonacci Package
This package provides functions for generating the Fibonacci sequence.

## Usage
To use the `fibonacci` package in your Go program, import it as follows:

```go
import "github.com/D7682/math/fibonacci"
```

Then you can call the `Fibonacci` function to generate the sequence. For example:
```go
// Get the 10th fibonacci number in the sequence using the cached version
cachedNumber := NewCached().Calculate(10)

// Get the 10th fibonacci number in the sequence using the uncached version
uncachedNumber := NewUncached().Calculate(10)

fmt.Println(cachedNumber)
fmt.Println(uncachedNumber)
```
This will generate and print the first 10 fibonacci numbers

# Cached vs. Uncached
The fibonacci package provides two implementations of the Fibonacci sequence: a cached version and an uncached version. The cached version stores previously computed values in a cache to avoid redundant computation, while the uncached version computes each value from scratch.

To use the cached version, call the CachedFibonacci function. To use the uncached version, call the UncachedFibonacci function. For example:

```go 
// Generate the first 10 Fibonacci numbers using the cached version.
cachedNumbers := NewCached().List(10)

// Generate the first 10 Fibonacci numbers using the uncached version.
uncachedNumbers := fibonacci.UncachedFibonacci(10)

fmt.Println(uncachedNumbers)
fmt.Println(cachedNumbers)
```

# Tests

## Concurrent Testing
To run the tests for the fibonacci package concurrently, use the -parallel flag when running the go test command. The -parallel flag specifies the maximum number of parallel test functions to run at once. For example:

```sh
go test -p 4 ./fibonacci
```

This command will run up to 4 test functions in parallel. If the -parallel flag is not specified, Go will automatically run tests in parallel if there are enough available cores on the machine.

Note that not all tests can be run in parallel, especially if they modify global state or rely on shared resources. However, for tests that are independent and self-contained, concurrent testing can significantly speed up the test suite.

## Testing Cached vs. Uncached

To run the tests for the Fibonacci implementation, use the following command in the terminal:

```sh
go test ./math/fibonacci -v
```

This will run all the tests in the package and print the results to the console. If you want to run only the cached tests or only the uncached tests, you can use the following commands:

To run only the cached tests:

```sh
go test -run TestFibonacciCached ./math/fibonacci -v
```

To run only the uncached tests:

```sh
go test -run TestFibonacciUncached ./math/fibonacci -v
```
Note that the `./math/fibonacci` argument specifies the directory containing the package, which should be relative to the current working directory.

# Performance

The Fibonacci sequence grows very quickly, so computing large numbers in the sequence can be computationally expensive. The cached implementation provided in this package can greatly improve performance by storing previously computed values and reusing them when possible. However, even with caching, the performance of the Fibonacci function may degrade quickly for very large inputs.