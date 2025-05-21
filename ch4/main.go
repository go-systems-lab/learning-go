package main

import "fmt"

func main() {
	// shadowing variables
	// x := 10
	// if x > 5 {
	// 	fmt.Println(x)
	// 	x := 5
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)

	// shadowing with multiple assignment
	// x := 10
	// if x > 5 {
	// 	x, y := 5, 20
	// 	fmt.Println(x, y)
	// }
	// fmt.Println(x)

	// shadowing package names
	// x := 10
	// fmt.Println(x)
	// fmt := "oops"
	// fmt.Println(fmt)

	// shadowing true
	// fmt.Println(true)
	// true := 10
	// fmt.Println(true)

	// if and else
	// n := rand.Intn(10)
	// if n == 0 {
	// 	fmt.Println("That's too low")
	// } else if n > 5 {
	// 	fmt.Println("That's too big:", n)
	// } else {
	// 	fmt.Println("That's a good number:", n)
	// }

	// scoping a variable to an if statement
	// if n := rand.Intn(10); n == 0 {
	// 	fmt.Println("That's too low")
	// } else if n > 5 {
	// 	fmt.Println("That's too big:", n)
	// } else {
	// 	fmt.Println("That's a good number:", n)
	// }

	// out of scope
	// if n := rand.Intn(10); n == 0 {
	// 	fmt.Println("That's too")
	// } else if n > 5 {
	// 	fmt.Println("That's too big:", n)
	// } else {
	// 	fmt.Println("That's a good number:", n)
	// }
	// fmt.Println(n)

	// A complete for statement
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// condition only for statement
	// i := 1
	// for i < 100 {
	// 	fmt.Println(i)
	// 	i = i * 2
	// }

	// Infinite looping nostalgia
	// for {
	// 	fmt.Println("Hello")
	// }

	// Confusing code
	// for i := 1; i < 100; i++ {
	// 	if i%3 == 0 {
	// 		if i%5 == 0 {
	// 			fmt.Println("FizzBuzz")
	// 		} else {
	// 			fmt.Println("Fizz")
	// 		}
	// 	} else if i%5 == 0 {
	// 		fmt.Println("Buzz")
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }

	// Using continue
	// for i := 1; i < 100; i++ {
	// 	if i%3 == 0 && i%5 == 0 {
	// 		fmt.Println("FizzBuzz")
	// 		continue
	// 	}
	// 	if i%3 == 0 {
	// 		fmt.Println("Fizz")
	// 		continue
	// 	}
	// 	if i%5 == 0 {
	// 		fmt.Println("Buzz")
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for range
	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for i, v := range evenVals {
	// 	fmt.Println(i, v)
	// }

	// ignore slice index
	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for _, v := range evenVals {
	// 	fmt.Println(v)
	// }

	// map iteration order varies
	// m := map[string]int{
	// 	"a": 1,
	// 	"c": 3,
	// 	"b": 2,
	// }

	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Loop", i)
	// 	for k, v := range m {
	// 		fmt.Println(k, v)
	// 	}
	// }

	// iterating over strings
	// samples := []string{"hello", "apple_#!"}
	// for _, sample := range samples {
	// 	for i, r := range sample {
	// 		fmt.Println(i, r, string(r))
	// 	}
	// 	fmt.Println()
	// }

	// modifying the value doesn't modify the source
	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for _, v := range evenVals {
	// 	v *= 2
	// }
	// fmt.Println(evenVals)

	// Labels
	// 	samples := []string{"hello", "apple_#!"}
	// outer:
	// 	for _, sample := range samples {
	// 		for i, r := range sample {
	// 			fmt.Println(i, r, string(r))
	// 			if r == 'l' {
	// 				continue outer
	// 			}
	// 		}
	// 		fmt.Println()
	// 	}

	// switch
	// words := []string{"a", "cow", "smile", "goopher", "octopus", "anthropologist"}
	// for _, word := range words {
	// 	switch size := len(word); size {
	// 	case 1, 2, 3, 4:
	// 		fmt.Println(word, "is a short word!")
	// 	case 5:
	// 		wordLen := len(word)
	// 		fmt.Println(word, "is exactly the right length:", wordLen)
	// 	case 6, 7, 8, 9:
	// 	default:
	// 		fmt.Println(word, "is a long word!")
	// 	}
	// }

	// break label
	// loop:
	// 	for i := 0; i < 10; i++ {
	// 		switch i {
	// 		case 0, 2, 4, 6:
	// 			fmt.Println(i, "is even")
	// 		case 3:
	// 			fmt.Println(i, "is divisible by 3 but not 2")
	// 		case 7:
	// 			fmt.Println("exit the loop!")
	// 			break loop
	// 		default:
	// 			fmt.Println(i, "is boring")
	// 		}
	// 	}

	// blank switch
	words := []string{"hi", "salutations", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}
}
