package main

import "fmt"

func main() {
	n := 23
	if n%2 == 1 {
		n = n*3 + 1
	} else {
		n = n / 2
	}

	fmt.Println(n)
}
