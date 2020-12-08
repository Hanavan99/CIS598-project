package optimizer

import (
	"fmt"
	"strings"
	"math"
	"gonum.org/v1/gonum/optimize"
)

type OptimizerFunction struct {
	tree ParseTreeRoot
	env map[string]interface{}
	args []string
	prop string
}

func (o OptimizerFunction) Evaluate(X []float64) float64 {
	tmpmap := make(map[string]interface{})
	copy(o.env, &tmpmap)
	for i := 0; i < len(o.args); i++ {
		tmpmap[o.args[i]] = X[i]
	}
	val, _ := Evaluate(o.tree, o.prop, Environment{tmpmap, o.tree})
	return val
}


func Solve(tree ParseTreeRoot, env Environment) error {

	// get all unbound properties
	unboundProps := GetUnboundProperties(tree)

	// figure out which ones are enums


	for _, v := range tree.Children() {
		//log.Printf("%s\n", v)
		solve, ok := v.(ParseTreeSolve)
		if ok {
			// do the solve operation
			_, err := resolveProperty(tree, tree, solve.Parameter, env)
			if err != nil {
				return err
			}
			
			DebugLogger.Printf("using strategy \"%s\"\n", solve.Strategy)
			switch solve.Strategy {
			case "minimize":
				DebugLogger.Printf("using minimize function\n")
			    /*err :=*/ minimize3(tree, env, solve.Parameter, unboundProps, 1e-4)
				if err != nil {
					return err
				}
				break
			case "maximize":
				DebugLogger.Printf("using maximize function\n")
			    /*err :=*/ maximize(tree, env, solve.Parameter, unboundProps, 1e-4)
				if err != nil {
					return err
				}
				break
			}
			return nil // quit out of loop
		}
	}
	return fmt.Errorf("no solve expression found")
}

func linesearch(f Function, X []float64, dir []float64, scale float64) (float64, []float64) {
	a, b := 0.0, 1.0
	X0 := X

	y := f.Evaluate(X)

	tmp := make([]float64, len(X))
	X1 := make([]float64, len(X))

	for {
		a, b = b, a + b
		//a++
		//b += 100
		

		alpha := -b * scale

		vector_scale(dir, alpha, &tmp)
		vector_add(X0, tmp, &X1)

		y1 := f.Evaluate(X1)
		fmt.Printf("%f; %f\n", alpha, y1)

		if (y1 < y) {
			y = y1
			X = X1
		} else {
			break
		}
	}

	return y, X
}

func gradient(f Function, X []float64, step float64) []float64 {
	var f0, f1, x_i float64
	G := make([]float64, len(X))
	f0 = f.Evaluate(X)
	for i := range X {
		x_i = X[i]
		X[i] += step
		f1 = f.Evaluate(X)
		G[i] = (f1 - f0) / step
		DebugLogger.Printf("(%f - %f) / %f = %f\n", f1, f0, step, G[i])
		X[i] = x_i
	}
	return G
}

func descent(f Function, X []float64, step float64, scale float64) []float64 {
	min := f.Evaluate(X)

	for {
		grad := gradient(f, X, step)
		DebugLogger.Printf("gradient=%s\n", grad)
		y, X1 := linesearch(f, X, grad, scale)
		if y < min {
			min = y
			X = X1
		} else {
			break
		}
	}

	return X
}

func vector_add(a []float64, b []float64, result* []float64) {
	length := int(math.Min(float64(len(a)), math.Min(float64(len(b)), float64(len(*result)))))
	for i := 0; i < length; i++ {
		(*result)[i] = a[i] + b[i]
	}
}

func vector_scale(a []float64, b float64, result* []float64) {
	length := int(math.Min(float64(len(a)), float64(len(*result))))
	for i := 0; i < length; i++ {
		(*result)[i] = a[i] * b
	}
}

func minimize2(tree ParseTreeRoot, env Environment, prop string, args []string, step float64, scale float64) {
	f := OptimizerFunction{tree, env.env, args, prop}
	X := make([]float64, len(args))

	for i, v := range args {
		X[i] = env.Get(v).(float64)
	}

	X1 := descent(f, X, step, scale)

	for i, v := range args {
		env.Put(v, X1[i])
	}
}

func minimize3(tree ParseTreeRoot, env Environment, prop string, args []string, step float64) {
	/*gd := optimize.GradientDescent{}
	gd.Init(len(args), 1)
	opchan := make(chan optimize.Task)
	resultchan := make(chan optimize.Task)
	gd.Run(opchan, resultchan, nil)
	 
	for  {
		status, _ := gd.Status()
		if status != optimize.NotTerminated {
			break
		}
		op := <- opchan
	}*/

	optimizefunc := func(x []float64) float64 {
		tmpmap := make(map[string]interface{})
		copy(env.env, &tmpmap)
		for i, v := range args {
			tmpmap[v] = x[i]
		}
		val, _ := Evaluate(tree, prop, Environment{tmpmap, tree})
		return val
	}

	gradfunc := func(grad, x []float64) {
		cx := make([]float64, len(x))
		for i := 0; i < len(x); i++ {
			cx[i] = x[i]
		}
		for i := 0; i < len(x); i++ {
			cx[i] += step
			grad[i] = (optimizefunc(cx) - optimizefunc(x)) / step
			cx[i] -= step
		}
	}

	p := optimize.Problem {
		Func: optimizefunc,
		Grad: gradfunc,
	}

	x := make([]float64, len(args))
	for i, _ := range x {
		x[i] = 1
	}

	result, err := optimize.Minimize(p, x, nil, nil)
	if err != nil {
		DebugLogger.Fatal(err)
	}
	for i, v := range args {
		fmt.Printf("%s = %f\n", v, result.X[i])
	}
	fmt.Printf("minimum value = %f\n", result.F)
}

func maximize(tree ParseTreeRoot, env Environment, prop string, args []string, step float64) {

	optimizefunc := func(x []float64) float64 {
		tmpmap := make(map[string]interface{})
		copy(env.env, &tmpmap)
		for i, v := range args {
			tmpmap[v] = x[i]
		}
		val, _ := Evaluate(tree, prop, Environment{tmpmap, tree})
		return -val
	}

	gradfunc := func(grad, x []float64) {
		cx := make([]float64, len(x))
		for i := 0; i < len(x); i++ {
			cx[i] = x[i]
		}
		for i := 0; i < len(x); i++ {
			cx[i] += step
			grad[i] = (optimizefunc(cx) - optimizefunc(x)) / step
			cx[i] -= step
		}
	}

	p := optimize.Problem {
		Func: optimizefunc,
		Grad: gradfunc,
	}

	x := make([]float64, len(args))
	for i, _ := range x {
		x[i] = 1
	}

	result, err := optimize.Minimize(p, x, nil, nil)
	if err != nil {
		DebugLogger.Fatal(err)
	}
	for i, v := range args {
		fmt.Printf("%s = %f\n", v, result.X[i])
	}
	fmt.Printf("maximum value = %f\n", -result.F)
}


func minimize(tree ParseTreeRoot, env Environment, prop string, args []string, step float64, t float64) error {

	var x map[string]interface{} = make(map[string]interface{})
	var x_prev map[string]interface{} = make(map[string]interface{})
	var x_next map[string]interface{} = make(map[string]interface{})
	var Dx map[string]interface{} = make(map[string]interface{})
	var g map[string]interface{} = make(map[string]interface{})
	var g_prev map[string]interface{} = make(map[string]interface{})
	var Dg map[string]interface{} = make(map[string]interface{})
	var tmp map[string]interface{} = make(map[string]interface{})
	var dx float64 = math.Inf(1)
	var alpha float64
	var err error

	copy(env.env, &x)
	zero(&x_prev, args)
	zero(&g_prev, args)

	for dx > t {
		DebugLogger.Printf("dx=%f; t=%f\n", dx, t)

		err = Gradient(tree, prop, args, Environment{x, tree}, step, &g)
		if err != nil {
			return err
		}

		sub(g, g_prev, &Dg, args)

		sub(x, x_prev, &Dx, args)

		// compute ratio of dot products
		alpha = dot(Dx, Dg, args) / dot(Dg, Dg, args)

		mult(g, alpha, &tmp, args)
		sub(x, tmp, &x_next, args)

		dx = dist(x_next, x, args)

		copy(x, &x_prev)
		copy(x_next, &x)
		DebugLogger.Printf("g_prev=%s; g=%s; Dg=%s; alpha=%f, Dx=%s; x_prev=%s; x=%s; x_next=%s; dx=%f\n", g_prev, g, Dg, alpha, Dx, x_prev, x, x_next, dx)
	}
	DebugLogger.Printf("done minimizing, dx=%f; t=%f\n", dx, t)

	env.env = x
	return nil

}

func zero(result* map[string]interface{}, args []string) {
	for _, arg := range args {
		(*result)[arg] = 1e3
	}
}

func copy(src map[string]interface{}, dest* map[string]interface{}) {
	for k, v := range src {
		(*dest)[k] = v
	}
}

func dist(a map[string]interface{}, b map[string]interface{}, args []string) float64 {
	var result float64 = 0
	for _, arg := range args {
		var diff float64 = b[arg].(float64) - a[arg].(float64)
		result += diff * diff
	}
	return math.Sqrt(result)
}

func mult(a map[string]interface{}, b float64, result* map[string]interface{}, args []string) {
	for k, v := range a {
		if (contains(args, k)) {
			(*result)[k] = v.(float64) * b
		} else {
			(*result)[k] = v
		}
	}
}

func sub(a map[string]interface{}, b map[string]interface{}, result* map[string]interface{}, args []string) {
	for k, v := range a {
		if (contains(args, k)) {
			(*result)[k] = v.(float64) - b[k].(float64)
		} else {
			(*result)[k] = v
		}
	}
}

func dot(a map[string]interface{}, b map[string]interface{}, args []string) float64 {
	var result float64 = 0
	for _, arg := range args {
		result += a[arg].(float64) * b[arg].(float64)
	}
	return result
}

func contains(array []string, value string) bool {
	for i := 0; i < len(array); i++ {
		if (array[i] == value) {
			return true
		}
	}
	return false
}

func Gradient(tree ParseTreeRoot, solve string, vars []string, env Environment, step float64, result* map[string]interface{}) error {
	DebugLogger.Printf("evaluating function with current environment\n")
	cur, err := Evaluate(tree, solve, env)
	if err != nil {
		return err
	}
	for _, v := range vars {
		DebugLogger.Printf("computing numerical partial derivative with respect to \"%s\"\n", v)
		penv := env.Clone()
		DebugLogger.Printf("computing current value of \"%s\"\n", v)
		start, err := Evaluate(tree, v, env)
		if err != nil {
			return err
		}
		//printf("Value of start")
		penv.Put(v, start + step)
		DebugLogger.Printf("computing current value of \"%s\" + %f\n", v, step)
		cur2, err := Evaluate(tree, solve, penv)
		if err != nil {
			return err
		}
		(*result)[v] = (cur2 - cur) / step
	}
	return nil
}

func Evaluate(tree ParseTreeRoot, name string, env Environment) (float64, error) {
	prop, err := resolveProperty(tree, tree, name, env)
	if err != nil {
		return 0, err
	}
	DebugLogger.Printf("%s\n", prop.Value)
	val := prop.Value
	if val != nil {
		return val.Evaluate(env), nil
	}
	return env.Get(name).(float64), nil
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
	DebugLogger.Printf("looking for property/assembly \"%s\"\n", name[0])

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
				DebugLogger.Printf("found property \"%s\"\n", name[0])
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

func GetUnboundProperties(tree ParseTreeRoot) []string {
	return getUnboundProperties(tree, "")
}

func getUnboundProperties(tree ParseTree, qualifiedName string) []string {
	DebugLogger.Printf("searching for unbound properties in \"%s\"\n", qualifiedName)
	props := make([]string, 0)
	for _, v := range tree.Children() {
		assembly, ok := v.(ParseTreeAssembly)
		if ok {
			props = append(props, getUnboundProperties(assembly, qualifiedName + assembly.name + ".")...)
		}
		property, ok := v.(ParseTreeProperty)
		if ok && property.Value == nil {
			props = append(props, qualifiedName + property.Name)
		}
	}
	return props;
}