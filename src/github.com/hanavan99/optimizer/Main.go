package main

import (
	"container/list"
	"io/ioutil"
)

//	"strconv"

// TokenIdentifier indicates the token is an identifier
const TokenIdentifier = 1

// TokenNumber indicates the token is a number
const TokenNumber = 2
const TokenOperatorEquals = 3
const TokenOperatorAdd = 4
const TokenOperatorSubtract = 5
const TokenOperatorMultiply = 6
const TokenOperatorDivide = 7
const TokenOperatorExponent = 8
const TokenBraceOpen = 9
const TokenBraceClose = 10
const TokenParenthesisOpen = 11
const TokenParenthesisClose = 12
const TokenBracketOpen = 13
const TokenBracketClose = 14
const TokenStatementTerminator = 15
const TokenComma = 16
const TokenSpace = 17
const TokenComment = 18

// Token a token from the program
type Token struct {
	id       int
	position int
	content  string
}

func tokenize(str string) *list.List {

	tokens := list.New()

	var tokenType = 0
	var curToken = ""
	var curTokenPos = 0
	for i := 0; i < len(str); i++ {
		var c = str[i]

		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_' {
			if tokenType == 0 {
				tokenType = TokenIdentifier
				curToken += string(c)
				curTokenPos = i
			} else if tokenType == TokenIdentifier {
				curToken += string(c)
			} else if tokenType == TokenNumber {
				if c == 'e' || c == 'E' {
					curToken += string(c)
				} else {
					// this must be unit information
					if tokenType != 0 {
						tokens.PushBack(Token{tokenType, curTokenPos, curToken})
					}
					tokenType = TokenIdentifier
					curToken += string(c)
				}
			} else {
				panic("Don't know what happens here")
			}
		} else if (c >= '0' && c < '9') || c == '.' {
			if tokenType == 0 {
				tokenType = TokenNumber
				curTokenPos = i
			}
			curToken += string(c)
		} else if c == '-' {
			// could be negative number or subtraction
		} else if c == '=' {
			if tokenType != 0 {
				tokens.PushBack(Token{tokenType, curTokenPos, curToken})
			}
			tokens.PushBack(Token{TokenOperatorEquals, i, string(c)})
			curToken = ""
			tokenType = 0
		} else if c == '+' {
			if tokenType != 0 {
				tokens.PushBack(Token{tokenType, curTokenPos, curToken})
			}
			tokens.PushBack(Token{TokenOperatorAdd, i, string(c)})
			curToken = ""
			tokenType = 0
		} else if c == '*' {
			if tokenType != 0 {
				tokens.PushBack(Token{tokenType, curTokenPos, curToken})
			}
			tokens.PushBack(Token{TokenOperatorMultiply, i, string(c)})
			curToken = ""
			tokenType = 0
		} else if c == '/' {
			if tokenType != 0 {
				tokens.PushBack(Token{tokenType, curTokenPos, curToken})
			}
			tokens.PushBack(Token{TokenOperatorDivide, i, string(c)})
			curToken = ""
			tokenType = 0
		} else if c == '^' {
			if tokenType != 0 {
				tokens.PushBack(Token{tokenType, curTokenPos, curToken})
			}
			tokens.PushBack(Token{TokenOperatorExponent, i, string(c)})
			curToken = ""
			tokenType = 0
		} else if c == '{' {
			if tokenType != 0 {
				tokens.PushBack(Token{tokenType, curTokenPos, curToken})
			}
			tokens.PushBack(Token{TokenBraceOpen, i, string(c)})
			curToken = ""
			tokenType = 0
		} else if c == '}' {
			if tokenType != 0 {
				tokens.PushBack(Token{tokenType, curTokenPos, curToken})
			}
			tokens.PushBack(Token{TokenBraceClose, i, string(c)})
			curToken = ""
			tokenType = 0
		} else if c == '(' {
			if tokenType != 0 {
				tokens.PushBack(Token{tokenType, curTokenPos, curToken})
			}
			tokens.PushBack(Token{TokenParenthesisOpen, i, string(c)})
			curToken = ""
			tokenType = 0
		} else if c == ')' {
			if tokenType != 0 {
				tokens.PushBack(Token{tokenType, curTokenPos, curToken})
			}
			tokens.PushBack(Token{TokenParenthesisClose, i, string(c)})
			curToken = ""
			tokenType = 0
		} else if c == '[' {
			if tokenType != 0 {
				tokens.PushBack(Token{tokenType, curTokenPos, curToken})
			}
			tokens.PushBack(Token{TokenBracketOpen, i, string(c)})
			curToken = ""
			tokenType = 0
		} else if c == ']' {
			if tokenType != 0 {
				tokens.PushBack(Token{tokenType, curTokenPos, curToken})
			}
			tokens.PushBack(Token{TokenBracketClose, i, string(c)})
			curToken = ""
			tokenType = 0
		} else if c == ',' {
			if tokenType == TokenNumber {
				curToken += string(c)
			} else {
				if tokenType != 0 {
					tokens.PushBack(Token{tokenType, curTokenPos, curToken})
				}
				tokens.PushBack(Token{TokenComma, i, string(c)})
				curToken = ""
				tokenType = 0
			}
		} else if c == ';' {
			if tokenType != 0 {
				tokens.PushBack(Token{tokenType, curTokenPos, curToken})
			}
			tokens.PushBack(Token{TokenStatementTerminator, i, string(c)})
			curToken = ""
			tokenType = 0
		} else {
			if tokenType != 0 {
				tokens.PushBack(Token{tokenType, curTokenPos, curToken})
			}
			curToken = ""
			tokenType = 0
		}
	}

	return tokens

}

func main() {

	dat, _ := ioutil.ReadFile("sample_program.txt")
	tokens := tokenize(string(dat))

	token := tokens.Front()
	for token.Next() != nil {
		print(token.Value.(Token).id, ", ", token.Value.(Token).position, ", \"", token.Value.(Token).content, "\"\n")
		token = token.Next()
	}

}
