package main

import (
	"fmt"
	"io/ioutil"
	"optimizer"
	"log"
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
		log.Fatalf("parse error: %s\n", err.Error())
	}

	for i, v := range tree.Children() {
		unit, ok := v.(optimizer.ParseTreeUnit)
		if ok {
			fmt.Printf("%d: %s\n", i, unit.Name)
		}

	}

	fmt.Println("==========  SOLVER  OUTPUT  ==========")

	env := optimizer.CreateEnvironment(tree)
	env.Put("pi", 3.1415926535)
	env.Put("rocket.nosecone.length", 20.0)
	env.Put("rocket.nosecone.mat", "ABS")
	val, err := optimizer.Evaluate(tree, "rocket.nosecone.mass", env)
	if err != nil {
		log.Fatalf("evaluation error: %s\n", err.Error())
	} else {
		fmt.Printf("result: %f\n", val)
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
