package optimizer

import "fmt"

type ParseTreeEnum struct {
	name   string
	props  []ParseTreeProperty
	values map[string]ParseTreeEnumValue
}

func (pte ParseTreeEnum) Children() []ParseTree {
	var children []ParseTree = make([]ParseTree, 0)
	for _, v := range pte.props {
		children = append(children, v)
	}
	for _, v := range pte.values {
		children = append(children, v)
	}
	return children
}

func (pte ParseTreeEnum) GetValue(valueName string, propertyName string) (ParseTreeProperty, error) {
	value, ok := pte.values[valueName]
	if ok {
		ordinal, err := pte.GetPropertyOrdinal(propertyName)
		if err == nil {
			return value.Properties[ordinal], nil
		}
		return ParseTreeProperty{}, err
	}
	return ParseTreeProperty{}, fmt.Errorf("enum \"%s\" does not contain value \"%s\"", pte.name, valueName)
}

func (pte ParseTreeEnum) GetPropertyOrdinal(propertyName string) (int, error) {
	for i, v := range pte.props {
		if v.Name == propertyName {
			return i, nil
		}
	}

	return -1, fmt.Errorf("enum \"%s\" does not have property \"%s\"", pte.name, propertyName)
}