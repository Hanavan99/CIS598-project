package optimizer

type ParseTreeProperty struct {
	name  string
	value ParseTreeExpression
}

func (ptp ParseTreeProperty) Children() []ParseTree {
	return []ParseTree{ptp.value}
}

func (ptp ParseTreeProperty) HasConstantValue() bool {
	return true
}
