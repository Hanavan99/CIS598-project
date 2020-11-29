package optimizer

type Function interface {
	Evaluate(X []float64) float64
}

