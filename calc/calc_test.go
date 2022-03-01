package calc

import (
	"fmt"
	"testing"
)

// go test -timeout 30s -run ^TestDot$ go-vectors/calc
func TestDot(t *testing.T) {

	got := Dot([]float64{1.0, 2.0}, []float64{1.0, 2.0})
	want := 5.

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

// go test -bench ^BenchmarkDot$ go-vectors/calc
func BenchmarkDot(b *testing.B) {
	A, B := make([]float64, 1024), make([]float64, 1024)
	sum := 0.0
	for i := 0; i < b.N; i++ {
		sum += Dot(A, B)
	}
	fmt.Println("Finished Calculations for ", b.N, " runs sum: ", sum) // use result
}
