package optimizer

import (
	"container/list"
)

//	"strconv"

// TokenIdentifier indicates the token is an identifier
const TokenIdentifier int = 1

// TokenNumber indicates the token is a number
const TokenNumber int = 2
const TokenOperatorEquals int = 3
const TokenOperatorAdd int = 4
const TokenOperatorSubtract int = 5
const TokenOperatorMultiply int = 6
const TokenOperatorDivide int = 7
const TokenOperatorExponent int = 8
const TokenBraceOpen int = 9
const TokenBraceClose int = 10
const TokenParenthesisOpen int = 11
const TokenParenthesisClose int = 12
const TokenBracketOpen int = 13
const TokenBracketClose int = 14
const TokenStatementTerminator int = 15
const TokenComma int = 16
const TokenSpace int = 17
const TokenComment int = 18
const TokenUnitSeparator int = 19

// Token a token from the program
type Token struct {
	// ID the ID of the token
	ID int
	// Position the location of the first character in the source code
	Position int
	// Content the content of the token
	Content string
}

// Tokenize converts input code into a list of tokens
func Tokenize(str string) *list.List {

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
		} else if (c >= '0' && c <= '9') || c == '.' {
			if tokenType == 0 {
				tokenType = TokenNumber
				curTokenPos = i
			}
			curToken += string(c)
		} else if c == '-' {
			// could be negative number or subtraction
			if len(curToken) > 1 && curToken[len(curToken) - 1] == 'e' {
				curToken += string(c)
			} else {
				if tokenType != 0 {
					tokens.PushBack(Token{tokenType, curTokenPos, curToken})
				}
				tokens.PushBack(Token{TokenOperatorSubtract, i, string(c)})
				curToken = ""
				tokenType = 0
			}
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
			/*if tokenType == TokenNumber {
				curToken += string(c)
			} else*/{
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
		} else if c == ':' {
			if tokenType != 0 {
				tokens.PushBack(Token{tokenType, curTokenPos, curToken})
			}
			tokens.PushBack(Token{TokenUnitSeparator, i, string(c)})
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

// IsOperator returns whether or not the given token ID is an operator
func IsOperator(token int) bool {
	switch token {
	case TokenOperatorAdd, TokenOperatorSubtract, TokenOperatorMultiply, TokenOperatorDivide, TokenOperatorExponent:
		return true
	}
	return false
}
