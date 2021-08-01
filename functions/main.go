package main

func div(a float64, b float64) float64 {
	return a / b
}

// TODO 1: implement [div] function with error handling and refactor to have one type signature for both parameters
func div1(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

// TODO 2: variadic parameters and slices
func sum(nums ...int) int {
	var s int
	for _, num := range nums {
		s += num
	}
	return s
}

// TODO 3: multiple return values
func sum2(nums ...int) (int, int) {
	var s int
	for _, num := range nums {
		s += num
	}
	return s, len(nums)
}

// TODO 4: ignoring returned values
var _, _ = sum2(1, 2, 3)

// TODO 5: named return values
func sum3(nums ...int) (sum int, n int) {
	n = len(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	return
}

// TODO 7: function type declarations
type MyFunc func(int) string
