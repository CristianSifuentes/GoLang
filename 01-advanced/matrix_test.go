package main

import (
	"testing"
)

func TestMulConcurrent_SmallKnown(t *testing.T) {
	// a, _ := FromSlice[float64](2, 3, []float64{
	// 	1, 2, 3,
	// 	4, 5, 6,
	// })

	// b, _ := FromSlice[float64](3, 2, []float64{
	// 	7, 8,
	// 	9, 10,
	// 	11, 12,
	// })

	// Expected:
	// [ 1*7+2*9+3*11 , 1*8+2*10+3*12 ] = [ 58, 64
	// [ 4*7+5*9+6*11 , 4*8+5*10+6*12 ] = [ 139, 154 ]
	// c, err := MulConcurrent(context.Background(), a, b, context.WithWorkers(2), WithBlockRows(1))
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// if got := c.At(0, 0); got != 58 {
	// 	t.Fatalf("c(0,0) = %v, want 58", got)
	// }
	// if got := c.At(0, 1); got != 64 {
	// 	t.Fatalf("c(0,1) = %v, want 64", got)
	// }
	// if got := c.At(1, 0); got != 139 {
	// 	t.Fatalf("c(1,0) = %v, want 139", got)
	// }
	// if got := c.At(1, 1); got != 154 {
	// 	t.Fatalf("c(1,1) = %v, want 154", got)
	// }
}

// func BenchmarkMulConcurrent_512() {
// 	a, _ := NewMatrix[Float](512, 512)
// 	cA := FillDeterministic(a, 1.0)
// 	bm, _ := NewMatrix[float64](512, 512)
// 	cB := FillDeterministic(bm, 2.0)

// }
