package optimizer

type ParseTreeProperty struct {
	Name  string
	Value ParseTreeExpression
	Units ParseTreeExpression
}

func (ptp ParseTreeProperty) Children() []ParseTree {
	return []ParseTree{ptp.Value}
}
