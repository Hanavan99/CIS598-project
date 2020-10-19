package optimizer

import (
	"container/list"
	"errors"
	"fmt"
	"strconv"
)

func tokenID(token *list.Element) int {
	return token.Value.(Token).ID
}

func tokenPos(token *list.Element) int {
	return token.Value.(Token).Position
}

func tokenContent(token *list.Element) string {
	return token.Value.(Token).Content
}

// Parse parses a list of tokens into a parse tree
func Parse(tokens *list.List) (*ParseTree, error) {

	token := tokens.Front()
	tree := ParseTree{nil, nil}

	for token != nil {
		if tokenID(token) != TokenIdentifier {
			return nil, errors.New("unit, assembly, enum, summarize, or solve expected at position " + strconv.Itoa(tokenPos(token)) + " " + tokenContent(token))
		}

		switch tokenContent(token) {
		case "unit":
			_token, err := parseUnit(token.Next(), tree)
			token = _token
			if err != nil {
				return &tree, err
			}
			break
		}
		token = token.Next()
	}

	return &tree, nil

}

func parseUnit(token *list.Element, tree ParseTree) (*list.Element, error) {
	if tokenID(token) != TokenIdentifier {
		return token, errors.New("identifier expected but \"" + tokenContent(token) + "\" given")
	}

	name := tokenContent(token)
	token = token.Next()

	if tokenID(token) == TokenStatementTerminator {
		fmt.Println("declared unit \"" + name + "\"")
		return token, nil
	} else if tokenID(token) == TokenOperatorEquals {
		return token, errors.New("unit conversions for unit \"" + name + "\" NYI")
	}

	return token, errors.New("unit \"")
}
