package main

import (
	"fmt"
)

type MyTree struct {
	Left  *MyTree
	Value int
	Right *MyTree
}

func insert(t *MyTree, v int) *MyTree {
	if t == nil {
		return &MyTree{nil, v, nil}
	}
	if t.Value < v {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func Walk(t *MyTree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Walker(t *MyTree) chan int {
	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

func compare(t1, t2 *MyTree) bool {
	ch1, ch2 := Walker(t1), Walker(t2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	var t *MyTree
	var t2 *MyTree
	t = insert(t, 6)
	t = insert(t, 10)
	t = insert(t, 1)
	t = insert(t, 3)
	t = insert(t, 5)
	t = insert(t, 4)
	t = insert(t, 9)

	t2 = insert(t2, 6)
	t2 = insert(t2, 10)
	t2 = insert(t2, 1)
	t2 = insert(t2, 3)
	t2 = insert(t2, 7)
	t2 = insert(t2, 4)
	t2 = insert(t2, 9)

	ch := Walker(t)
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	fmt.Println(compare(t, t2))
}
