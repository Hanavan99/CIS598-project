package optimizer

type ParseTreeAssembly struct {
	name          string
	subassemblies []ParseTreeAssembly
	props         []ParseTreeProperty
}

func (pta ParseTreeAssembly) Children() []ParseTree {
	var children []ParseTree = make([]ParseTree, 0)
	for _, v := range pta.props {
		children = append(children, v)
	}
	for _, v := range pta.subassemblies {
		children = append(children, v)
	}
	return children
}
