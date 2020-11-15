package optimizer

type ParseTreeEnumValue struct {
	Name   string
	Properties []ParseTreeProperty
}

func (ptev ParseTreeEnumValue) Children() []ParseTree {
	children := make([]ParseTree, 0)
	for _, v := range ptev.Properties {
		children = append(children, v)
	}
	return children
}
