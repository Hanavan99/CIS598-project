package optimizer

type ParseTreeRoot struct {
	children []ParseTree
}

func CreateParseTreeRoot() ParseTreeRoot {
	return ParseTreeRoot{[]ParseTree{}}
}

func (ptr ParseTreeRoot) Children() []ParseTree {
	return ptr.children
}

//func (ptr ParseTreeRoot) Parent() ParseTree {
//	return nil
//}

func (ptr *ParseTreeRoot) AddUnit(unit ParseTreeUnit) {
	ptr.children = append(ptr.children, unit)
}

func (ptr *ParseTreeRoot) AddEnum(enum ParseTreeEnum) {
	ptr.children = append(ptr.children, enum)
}
