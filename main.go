package main

import (
	"fmt"
)

type Tree struct {
	Left  *Tree
	Val   int
	Right *Tree
}

func (t *Tree) Add(v int) *Tree {
	if t == nil {
		t = &Tree{nil, v, nil}
		return t
	}
	if t.Val < v {
		t.Left = t.Left.Add(v)
		return t
	}
	t.Right = t.Right.Add(v)
	return t
}

func (t *Tree) Print() {
	if t == nil {
		return
	}
	t.Left.Print()
	fmt.Println(t.Val)
	t.Right.Print()


}

func (t* Tree) Walker()chan int {
	ch := make(chan int)
	go func (){
		t.Walk(ch)
		close(ch)
	}()
	return ch
}

func (t *Tree)Walk(ch chan<- int) {
	if t == nil {
		return
	}
	t.Left.Walk(ch)
	ch<-t.Val
	t.Right.Walk(ch)
}

func (t1 *Tree)Compare(t2 *Tree) bool {
	ch1, ch2 := t1.Walker(), t2.Walker()

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
	var t1, t2 *Tree

	t1 = t1.Add(6)
	fmt.Println(t1)
	t1.Add(8)
	t1.Add(9)
	t1.Add(4)
	t1.Add(10)
	
	t2 = t2.Add(6)
	t2.Add(8)
	t2.Add(9)
	t2.Add(4)
	t2.Add(10)
	t1.Print()
	t2.Print()
	fmt.Println(t1.Compare(t2))
}
