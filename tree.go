package main

import "golang.org/x/tour/tree"

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	go Walk(t1, c1)
	go Walk(t2, c2)
	ch := make(chan bool)
	go func() {
		a, b := 0, 0
		for i := 0; i < 10; i++ {
			if a, b = <-c1, <-c2; a != b {
				ch <- false
			}

		}
		ch <- true
	}()
	return <-ch
}
