package main

import "golang.org/x/tour/tree"
import "fmt"

func _Walk(t *tree.Tree, ch chan int) {
	if(t != nil) {
		_Walk(t.Left, ch)
		ch <- t.Value
		_Walk(t.Right, ch)
	}
}


// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	_Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v1 := range ch1 {
		v2 := <- ch2
		if v1 != v2 {
			return false	
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for v := range ch {
		fmt.Println(v)
	}

	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("Correct")
	}
	
	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("Correct")
	}
}
