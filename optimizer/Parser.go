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
func Parse(tokens *list.List, funcDefs map[string]int) (ParseTreeRoot, error) {

	token := tokens.Front()
	tree := ParseTreeRoot{make([]ParseTree, 0)}

	for token != nil {
		if tokenID(token) != TokenIdentifier {
			return tree, errors.New("unit, assembly, enum, summarize, or solve expected at position " + strconv.Itoa(tokenPos(token)) + " " + tokenContent(token))
		}

		switch tokenContent(token) {
		case "unit":
			_token, unit, err := parseUnit(token.Next(), tree, funcDefs)
			token = _token
			if err != nil {
				return tree, err
			}
			tree.AddUnit(unit)
			break
		}
		token = token.Next()
	}

	return tree, nil

}

func parseUnit(token *list.Element, tree ParseTreeRoot, funcDefs map[string]int) (*list.Element, ParseTreeUnit, error) {
	if tokenID(token) != TokenIdentifier {
		return token, ParseTreeUnit{}, errors.New("identifier expected but \"" + tokenContent(token) + "\" given")
	}

	name := tokenContent(token)
	token = token.Next()

	if tokenID(token) == TokenStatementTerminator {
		fmt.Println("declared unit \"" + name + "\"")
		return token, ParseTreeUnit{name, 1, nil}, nil
	} else if tokenID(token) == TokenOperatorEquals {
		token = token.Next()
		if tokenID(token) == TokenNumber {
			multiplier, err := strconv.ParseFloat(tokenContent(token), 64)
			if err != nil {
				return nil, ParseTreeUnit{name, 1, nil}, errors.New("\"" + tokenContent(token) + "\" is not a valid float64")
			}
			token, exp, err := ParseExpression(token.Next(), tree, funcDefs)
			fmt.Printf("declared unit \"%s\" = %f %s\n", name, multiplier, exp.ToString())
			return token, ParseTreeUnit{name, multiplier, exp}, err
		}
		return token, ParseTreeUnit{name, 1, nil}, errors.New("unit \"" + name + "\" missing conversion factor")
	}

	return token, ParseTreeUnit{}, errors.New("unit \"")
}

// BuildExpression builds an expression tree from a stack containing tokens in Reverse Polish Notation, as well as function definitions that map function names to argument counts.
func BuildExpression(outputQueue *Stack, funcDefs map[string]int) (ParseTreeExpression, error) {
	//for i, v := range outputQueue.Slice() {
	//	fmt.Printf("%d!: %s\n", i, v.(Token).Content)
	//}

	token := outputQueue.Pop().(Token)
	//fmt.Printf("token ID=%d Content=%s\n", token.ID, token.Content)
	switch token.ID {
	case TokenIdentifier:
		//fmt.Printf("identifier %s\n", token.Content)
		_, exists := funcDefs[token.Content]
		if exists {
			// return a function expression
		} else {
			return ParseTreeExpressionVariable{token.Content}, nil
		}
		break
	case TokenNumber:
		//fmt.Printf("number %s\n", token.Content)
		val, err := strconv.ParseFloat(token.Content, 64)
		if err != nil {
			return nil, err
		}
		return ParseTreeExpressionConstant{val}, nil
	case TokenOperatorAdd:
		//fmt.Printf("token +\n")
		right, err := BuildExpression(outputQueue, funcDefs)
		if err != nil {
			return nil, err
		}
		left, err := BuildExpression(outputQueue, funcDefs)
		if err != nil {
			return nil, err
		}
		return ParseTreeExpressionAdd{left, right}, nil
	}
	return nil, errors.New("invalid stack")
}

func ParseExpression(token *list.Element, tree ParseTreeRoot, funcDefs map[string]int) (*list.Element, ParseTreeExpression, error) {
	operatorStack := Stack{}
	outputQueue := Stack{}

	for tokenID(token) != TokenStatementTerminator {
		switch tokenID(token) {
		case TokenNumber:
			//num, _ := strconv.Atoi(tokenContent(token))
			outputQueue.Push(token.Value)
			//fmt.Printf("pushed number %s to output queue\n", tokenContent(token))
			break
		case TokenIdentifier:
			if tokenID(token.Next()) == TokenParenthesisOpen {
				operatorStack.Push(token.Value) // it is a function
				//fmt.Printf("pushed function \"%s\" to operator stack\n", tokenContent(token))
			} else {
				outputQueue.Push(token.Value) // it is a variable/unit name
				//fmt.Printf("pushed variable/unit \"%s\" to output queue\n", tokenContent(token))
			}
			break
		case TokenOperatorAdd, TokenOperatorSubtract, TokenOperatorMultiply, TokenOperatorDivide, TokenOperatorExponent:
			//fmt.Printf("handling operator %s\n", tokenContent(token))
			var tokenID = tokenID(token)
			//for operatorStack.Peek() != nil && (IsOperator(operatorStack.Peek().(Token).ID)) && ((operatorPrecedence(operatorStack.Peek().(Token).ID) > operatorPrecedence(tokenID)) || (operatorPrecedence(operatorStack.Peek().(Token).ID) == operatorPrecedence(tokenID) && isLeftAssociative(tokenID))) && (operatorStack.Peek().(Token).ID != TokenParenthesisOpen) {
			for operatorStack.Peek() != nil {
				var isop = IsOperator(operatorStack.Peek().(Token).ID) || operatorStack.Peek().(Token).ID == TokenIdentifier
				var greaterPrecedence = operatorPrecedence(operatorStack.Peek().(Token).ID) > operatorPrecedence(tokenID)
				var equalPrecedence = operatorPrecedence(operatorStack.Peek().(Token).ID) == operatorPrecedence(tokenID)
				var leftAssoc = isLeftAssociative(tokenID)
				var leftParen = operatorStack.Peek().(Token).ID == TokenParenthesisOpen
				//fmt.Printf("op=%d isOp=%t greaterPrecedence=%t equalPrecedence=%t leftAssoc=%t leftParen=%t\n", operatorStack.Peek().(Token).ID, isop, greaterPrecedence, equalPrecedence, leftAssoc, leftParen)
				if (isop) && ((greaterPrecedence) || (equalPrecedence && leftAssoc)) && (!leftParen) {
					var poppedToken = operatorStack.Pop()
					outputQueue.Push(poppedToken)
					//fmt.Printf("popped operator %s from stack and pushed to output queue\n", poppedToken.(Token).Content)
				} else {
					break
				}
			}

			operatorStack.Push(token.Value)
			//fmt.Printf("pushed operator %s to operator stack\n", tokenContent(token))
			break
		case TokenParenthesisOpen:
			operatorStack.Push(token.Value)
			//fmt.Printf("pushed ( to operator stack\n")
			break
		case TokenParenthesisClose:
			//fmt.Println("handling )")
			for operatorStack.Peek() != nil && operatorStack.Peek().(Token).ID != TokenParenthesisOpen {
				var poppedToken = operatorStack.Pop()
				outputQueue.Push(poppedToken)
				//fmt.Printf("popped operator %s from stack and pushed to output queue\n", poppedToken.(Token).Content)
			}
			//fmt.Printf("popped operator ( from stack\n")
			operatorStack.Pop() // remove open parenthesis
			// TODO check for mismatched parenthesis
			break
		case TokenComma:
			//ignore
			//fmt.Printf("ignoring comma\n")
			break
		}
		token = token.Next()
		//fmt.Printf("operator stack: ")
		//for _, v := range operatorStack.Slice() {
		//	fmt.Printf("%s ", v.(Token).Content)
		//}
		//fmt.Println()
		//fmt.Printf("next token is %s (%d)\n", tokenContent(token), tokenID(token))
	}

	for operatorStack.Length() > 0 {
		outputQueue.Push(operatorStack.Pop())
	}

	//for i, v := range outputQueue.Slice() {
	//fmt.Printf("Token %d: id=%d pos=%d content=\"%s\"\n", i, v.(Token).ID, v.(Token).Position, v.(Token).Content)
	//}

	exp, err := BuildExpression(&outputQueue, funcDefs)
	return token, exp, err
}

func operatorPrecedence(op int) int {
	switch op {
	case TokenOperatorAdd, TokenOperatorSubtract:
		return 2
	case TokenOperatorMultiply, TokenOperatorDivide:
		return 3
	case TokenOperatorExponent:
		return 4
	case TokenIdentifier:
		return 5
	}
	return -1
}

func isLeftAssociative(op int) bool {
	switch op {
	case TokenOperatorAdd, TokenOperatorSubtract, TokenOperatorMultiply, TokenOperatorDivide:
		return true
	case TokenOperatorExponent:
		return false
	}
	return false
}
