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

func (ptr *ParseTreeRoot) AddAssembly(assembly ParseTreeAssembly) {
	ptr.children = append(ptr.children, assembly)
}

func (ptr *ParseTreeRoot) AddSummarize(summarize ParseTreeSummarize) {
	ptr.children = append(ptr.children, summarize)
}

func (ptr *ParseTreeRoot) AddSolve(solve ParseTreeSolve) {
	ptr.children = append(ptr.children, solve)
}