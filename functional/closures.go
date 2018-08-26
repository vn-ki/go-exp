package functional

func Adder(x int) func(int) int {
	foo := func(a int) int {
		return x + a
	}
	x = 3
	return foo
}
