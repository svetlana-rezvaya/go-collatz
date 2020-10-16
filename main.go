package main

import "fmt"

func main() {
	n := 23
	if false {
		n = n*3 + 1
	} else {
		n = n / 2
	}

	fmt.Println(n)
}
