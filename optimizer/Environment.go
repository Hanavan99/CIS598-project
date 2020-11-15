package optimizer

import "log"

// Environment associates variables with values, as well as a way to look up values in the tree
type Environment struct {
	env map[string]interface{}
	tree ParseTreeRoot
}

func CreateEnvironment(tree ParseTreeRoot) Environment {
	return Environment{make(map[string]interface{}), tree}
}

func (env Environment) Exists(key string) bool {
	_, exists := env.env[key]
	if !exists {
		_, err := resolveProperty(env.tree, env.tree, key, env)
		if err == nil {
			return true
		}
	}
	return exists
}

func (env Environment) Get(key string) interface{} {
	log.Printf("looking up key \"%s\" in environment\n", key)
	ret, exists := env.env[key]
	if !exists {
		prop, err := resolveProperty(env.tree, env.tree, key, env)
		if err == nil {
			return prop.Value.Evaluate(env)
		}
		return nil
	}
	return ret
}

func (env Environment) Put(key string, value interface{}) {
	env.env[key] = value
}
