package optimizer

type ParseTreeEnum struct {
	name   string
	props  []ParseTreeProperty
	values map[string]ParseTreeEnumValue
}

func (pte ParseTreeEnum) Children() []ParseTree {
	return nil
}
