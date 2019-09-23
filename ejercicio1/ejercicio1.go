package main

import (
	"ejercicio1/sum"
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) > 2 {
		fmt.Println(fmt.Errorf("more than 1 parameter"))
	}
	p, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err)
	}

	suma := sum.Getsum(p)

	fmt.Println("sum is", suma)
}
