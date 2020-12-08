package main

import (
	"fmt"
	"io/ioutil"
	"optimizer"
	"log"
	"os"
)

/*
============================== TODO ==============================
1. Implement GetUnboundProperties() function          DONE
2. Get multivariate Newton's Method implemented
3. Make optimizer obey summarize statements
4. Get gradient descent/ascent implemented            DONE
5. Get unit conversions implemented
6. Implement error checking on parse tree
7. Accept file name from command line                 DONE
8. Allow for non fully qualified name resolution
9. Fix expression parsing to not require ";"
10. Allow user to adjust parameters to GD algorithm
11. Fix tokenizer to handle negative numbers
============================== ==== ==============================
*/

func main() {

	// initialize logger
	optimizer.InitLoggers()

	// get command line arguments
	args := os.Args[1:]

	// check for file name argument
	if len(args) >= 1 {

		// read in file
		dat, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}

		fmt.Println("========== TOKENIZER OUTPUT ==========")

		// convert file contents into linked list of tokens
		tokens := optimizer.Tokenize(string(dat))

		// define expression function arities
		funcDefs := make(map[string]int)
		funcDefs["sin"] = 1

		
		/*token := tokens.Front()
		for token != nil {
			print(token.Value.(optimizer.Token).ID, "\t", token.Value.(optimizer.Token).Position, "\t\"", token.Value.(optimizer.Token).Content, "\"\n")
			token = token.Next()
		}*/

		fmt.Println("==========  PARSER  OUTPUT  ==========")
		tree, err := optimizer.Parse(tokens, funcDefs)
		if err != nil {
			fmt.Printf("parse error: %s\n", err.Error())
			os.Exit(1)
		}

		/*for i, v := range tree.Children() {
			unit, ok := v.(optimizer.ParseTreeUnit)
			if ok {
				fmt.Printf("%d: %s\n", i, unit.Name)
			}

		}*/

		//log.SetOutput(ioutil.Discard)

		fmt.Println("==========  SOLVER  OUTPUT  ==========")

		props := optimizer.GetUnboundProperties(tree)
		fmt.Printf("Unbound properties: %s\n", props)

		env := optimizer.CreateEnvironment(tree)
		err = optimizer.Solve(tree, env)
		if err != nil {
			log.Fatal(err)
		}
		/*for k, v := range env.GetMap() {
			fmt.Printf("property \"%s\" = \"%.20f\"\n", k, v)
		}*/

		fmt.Println("==========       DONE       ==========")
	} else {
		fmt.Printf("not enough arguments! usage: ./optimizer file_name\n")
	}

}
