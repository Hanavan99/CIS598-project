package optimizer

// ParseTreeExpression represents an expression specified by the user, typically arithmetic in nature
type ParseTreeExpression interface {
	ParseTree
	Evaluate(env Environment) float64
	HasKnownValue(env Environment) bool
	ToString() string
}
