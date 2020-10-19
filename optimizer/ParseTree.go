package optimizer

import "container/list"

// ParseTree data structure for the parsed input code
type ParseTree struct {
	Children *list.List
	Value    interface{}
}

// AddChild adds a child to this parse tree
func (tree ParseTree) AddChild(child ParseTree) {
	tree.Children.PushBack(child)
}
