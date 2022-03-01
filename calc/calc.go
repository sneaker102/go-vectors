package calc

func Dot(A, B []float64) float64 {
	dot := 0.0
	for i := range A {
		dot += A[i] * B[i]
	}
	return dot
}
