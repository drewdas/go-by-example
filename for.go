package main

import "fmt"

func main() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Println("Start Loop")
	x := 1
	for {
		fmt.Println(x)
		if x == 10 {
			print("End Loop \n")
			break
		}
		x = x + 1
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
