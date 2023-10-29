package utils

func Sum(a, b int) int {
	return a + b
}

func Diff(a, b int) int {
	return a - b
}

func Bark() string {
	return "Woff!!"
}

func Barks() string {
	return "Woff Woff Woff!!"
}

func Sqr(a int) int {
	return a * a
}

func PowerOf(a, b int) int {
	if b == 0 {
		return 1
	} else if b == 1 {
		return a
	} else {
		return a * PowerOf(a, b-1)
	}
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}
