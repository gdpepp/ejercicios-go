package main

import (
	"fmt"
	"os"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {
	var s1[] int
	var t *tree

	//convert args to int
	for _ , i := range os.Args[1:] {
		var err error
		var val int

		val , err = strconv.Atoi(i)
		//error checking
		if err != nil {
			panic("Could not convert to int")
		}

		s1 = append(s1, val)
	}

	fmt.Println(s1)

	for _ ,i := range s1 {
		t = add(t, i)
	}

	readTree(t)
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t := &tree{value, nil, nil}
		return t
	}
	if value < t.value {
		if t.left != nil {
			t.left = add(t.left, value)
		} else {
			t.left = &tree{value, nil, nil}
		}
	} else {
		if t.right != nil {
			t.right = add(t.right, value)
		} else {
			t.right = &tree{value, nil, nil}
		}
	}
	return t
}

func readTree(t *tree) {
	if t != nil {
		if t.left != nil {
			readTree(t.left)
		}
		fmt.Println(t.value)
		if t.right != nil {
			readTree(t.right)
		}
	}
}
