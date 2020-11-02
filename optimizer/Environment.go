package optimizer

// associates variables with values
type Environment struct {
	env map[string]interface{}
}

func CreateEnvironment() Environment {
	return Environment{make(map[string]interface{})}
}

func (env Environment) Exists(key string) bool {
	_, exists := env.env[key]
	return exists
}

func (env Environment) Get(key string) interface{} {
	return env.env[key]
}

func (env Environment) Put(key string, value interface{}) {
	env.env[key] = value
}
