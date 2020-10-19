package main

import (
	"fmt"
	"io/ioutil"
	"optimizer"
)

func main() {

	dat, _ := ioutil.ReadFile("../sample_program2.txt")
	//fmt.Println(dat)
	tokens := optimizer.Tokenize(string(dat))

	fmt.Println("========== TOKENIZER OUTPUT ==========")
	token := tokens.Front()
	for token != nil {
		print(token.Value.(optimizer.Token).ID, "\t", token.Value.(optimizer.Token).Position, "\t\"", token.Value.(optimizer.Token).Content, "\"\n")
		token = token.Next()
	}

	fmt.Println("==========  PARSER  OUTPUT  ==========")
	_, err := optimizer.Parse(tokens)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}

	fmt.Println("==========       DONE       ==========")

}
