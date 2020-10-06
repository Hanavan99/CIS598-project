package main

import (
	list "container/list"
	"fmt"
)

func main() {

	var yay = tokenize("hello")
	var elem = yay.Front()
	for i := 0; i < yay.Len(); i++ {
		fmt.Println(elem)
		elem = elem.Next()
	}
}

func tokenize(code string) *list.List {
	var ret = list.New()
	ret.PushFront("testing")
	return ret
}
