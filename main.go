package main

import (
	"flag"
	"fmt"
)

func main() {
	nFlag := flag.Int("n", 23, "start number")
	flag.Parse()

	n := *nFlag
	fmt.Println(n)

	for n != 1 {
		if n%2 == 1 {
			n = n*3 + 1
		} else {
			n = n / 2
		}

		fmt.Println(n)
	}
}
