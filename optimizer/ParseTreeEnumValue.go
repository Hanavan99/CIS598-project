package optimizer

type ParseTreeEnumValue struct {
	Name   string
	Values []float64
	Units  []ParseTreeExpression
}

func (ptev ParseTreeEnumValue) Children() []ParseTree {
	children := make([]ParseTree, 0)
	for _, v := range ptev.Units {
		children = append(children, v)
	}
	return children
}
