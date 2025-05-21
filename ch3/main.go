package main

import "fmt"

func main() {
	// Understanding capacity
	// var x []int
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 10)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 20)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 40)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 50)
	// fmt.Println(x, len(x), cap(x))

	// Slice whuch might stay nil
	// var data []int

	// slice with default values
	// data := []int{2, 4, 6, 8}

	// Slicing slices
	// x := []string{"a", "b", "c", "d"}
	// y := x[:2]
	// z := x[1:]
	// d := x[1:3]
	// e := x[:]
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)
	// fmt.Println("d:", d)
	// fmt.Println("e:", e)

	// Slices with overlapping storage
	// x := []string{"a", "b", "c", "d"}
	// y := x[:2]
	// z := x[1:]
	// x[1] = "y"
	// y[0] = "x"
	// z[1] = "z"
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// append makes overlapping slices more confusing
	// x := []string{"a", "b", "c", "d"}
	// y := x[:2]
	// fmt.Println(cap(x), cap(y))
	// y = append(y, "z")
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)

	// even more confusing slices
	// x := make([]string, 0, 5)
	// x = append(x, "a", "b", "c", "d")
	// y := x[:2]
	// z := x[2:]
	// fmt.Println(cap(x), cap(y), cap(z))
	// y = append(y, "i", "j", "k")
	// x = append(x, "x")
	// z = append(z, "y")
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// full slice expression protects against append
	// x := make([]string, 0, 5)
	// x = append(x, "a", "b", "c", "d")
	// y := x[:2:2]
	// z := x[2:4:4] //last value is last position in parent slice's capacity
	// fmt.Println(cap(x), cap(y), cap(z))
	// y = append(y, "i", "j", "k")
	// x = append(x, "x")
	// z = append(z, "y")
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// converting strings to slices
	// var s string = "Hello, V"
	// var bs []byte = []byte(s)
	// var rs []rune = []rune(s)
	// fmt.Println(bs)
	// fmt.Println(rs)

	// Using a map
	// totalWins := map[string]int{}
	// totalWins["Orcas"] = 1
	// totalWins["Lions"] = 2
	// fmt.Println(totalWins["Orcas"])
	// fmt.Println(totalWins["Kittens"])
	// totalWins["Kittens"]++
	// fmt.Println(totalWins["Kittens"])
	// totalWins["Lions"] = 3
	// fmt.Println(totalWins["Lions"])

	// Using a map as a set
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in set")
	}
}
