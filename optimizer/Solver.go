package optimizer

import (
	"fmt"
	"strings"
	"log"
)

func Solve(tree ParseTreeRoot, env Environment) error {

	for _, v := range tree.Children() {
		solve, ok := v.(ParseTreeSolve)
		if ok {
			// do the solve operation
			_, err := resolveProperty(tree, tree, solve.Parameter, env)
			if err != nil {
				return err
			}
		}
	}
	return fmt.Errorf("no solve expression found")
}

func Evaluate(tree ParseTreeRoot, name string, env Environment) (float64, error) {
	prop, err := resolveProperty(tree, tree, name, env)
	if err != nil {
		return 0, err
	}
	log.Printf("%s\n", prop.Value)
	return prop.Value.Evaluate(env), nil
}

func getPropertyEnum(tree ParseTreeRoot, property ParseTreeProperty) (ParseTreeEnum, bool) {
	// check if top level expression for property is a variable
	variable, ok := property.Units.(ParseTreeExpressionVariable)
	if ok {
		for _, v := range tree.Children() {
			enum, ok := v.(ParseTreeEnum)
			if ok && enum.name == variable.name {
				return enum, true
			}
		}
	}

	
	return ParseTreeEnum{}, false
}

func resolveProperty(root ParseTreeRoot, tree ParseTree, name string, env Environment) (ParseTreeProperty, error) {
	prop := strings.Split(name, ".")
	return resolveProperty2(root, tree, make([]string, 0), prop, env)
}

func resolveProperty2(root ParseTreeRoot, tree ParseTree, qualifiedName []string, name []string, env Environment) (ParseTreeProperty, error) {
	log.Printf("looking for property/assembly \"%s\"\n", name[0])

	// look through children of tree to find child with name[0]
	for _, v := range tree.Children() {

		// check if the child is an assembly
		assembly, ok := v.(ParseTreeAssembly)
		if ok && assembly.name == name[0] && len(name) > 1 {
			return resolveProperty2(root, assembly, append(qualifiedName, name[0]), name[1:], env)
		}

		// check if the child is a property
		property, ok := v.(ParseTreeProperty)
		if ok && property.Name == name[0] {
			if len(name) == 1 {
				log.Printf("found property \"%s\"\n", name[0])
				return property, nil
			}

			// check if this property is an enum type
			enum, ok := getPropertyEnum(root, property)
			if ok {
				qualifiedNameStr := strings.Join(append(qualifiedName, name[0]), ".")
				ok := env.Exists(qualifiedNameStr)
				if ok {
					valueName := env.Get(qualifiedNameStr).(string)
					return resolveProperty2(root, enum.values[valueName], append(qualifiedName, name[0]), name[1:], env)
				}
				return ParseTreeProperty{}, fmt.Errorf("no property found 2")
			}
			
		}
	}
	return ParseTreeProperty{}, fmt.Errorf("no property found")
}