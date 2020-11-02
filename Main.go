package main

import (
	"fmt"
	"io/ioutil"
	"optimizer"
)

func main() {

	dat, _ := ioutil.ReadFile("../sample_program.txt")
	//fmt.Println(dat)
	tokens := optimizer.Tokenize(string(dat))

	funcDefs := make(map[string]int)
	funcDefs["sin"] = 1

	fmt.Println("========== TOKENIZER OUTPUT ==========")
	token := tokens.Front()
	for token != nil {
		print(token.Value.(optimizer.Token).ID, "\t", token.Value.(optimizer.Token).Position, "\t\"", token.Value.(optimizer.Token).Content, "\"\n")
		token = token.Next()
	}

	fmt.Println("==========  PARSER  OUTPUT  ==========")
	tree, err := optimizer.Parse(tokens, funcDefs)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}

	for i, v := range tree.Children() {
		fmt.Printf("%d: %s\n", i, v.(optimizer.ParseTreeUnit).Name)
	}

	fmt.Println("==========       DONE       ==========")

	/*s := optimizer.Stack{}
	s.Push(3)
	s.Push(6)
	//s.Pop()
	s = s.Reverse()
	for i, v := range s.Slice() {
		fmt.Printf("%d: %d\n", i, v)
	}

	_, _, err := optimizer.ParseExpression(tokens.Front(), optimizer.ParseTreeRoot{}, funcDefs)
	if err != nil {
		fmt.Println(err.Error())
	}*/

}
