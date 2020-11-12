package optimizer

type ParseTreeSummarize struct {
	Parameter  string
}

func (ptp ParseTreeSummarize) Children() []ParseTree {
	return nil
}
