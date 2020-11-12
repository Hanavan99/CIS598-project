package optimizer

type ParseTreeSolve struct {
	Strategy string
	Parameter       string
}

func (pts ParseTreeSolve) Children() []ParseTree {
	return nil
}