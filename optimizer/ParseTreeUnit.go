package optimizer

type ParseTreeUnit struct {
	//parent ParseTree
	Name       string
	Multiplier float64
	Units      ParseTreeExpression
}

func (ptu ParseTreeUnit) Children() []ParseTree {
	return nil
}

//func (ptu ParseTreeUnit) Parent() ParseTree {
//	return ptu.parent
//}

// CompareTo compares one unit to another to check for equivalence.
func (ptu ParseTreeUnit) CompareTo(other ParseTreeUnit) bool {
	return true
}
