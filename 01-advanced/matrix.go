package main

import (
	"errors"
	"fmt"
	"runtime"
)

// ---------- Numeric constraints (no external deps) --------------
type Float interface {
	~float32 | ~float64
}

// ---------- Errors (sentinel  + type validation) -----------------
var (
	ErrDimemsionMismatch = errors.New("matrix dimension mismatch")
	ErrEmptyMatrix       = errors.New("matrix cannot be empty")
)

// ---------- Matrix type with methods (generics + value semantics) --------------
type Matrix[T Float] struct {
	r, c int
	data []T // row-major: data[i*c + j]
}

// NewMatrix allocates an RxCmatrix filled with zero values of T.
func NewMatrix[T Float](rows, cols int) (Matrix[T], error) {
	if rows <= 0 || cols <= 0 {
		return Matrix[T]{}, ErrEmptyMatrix
	}
	return Matrix[T]{r: rows, c: cols, data: make([]T, rows*cols)}, nil
}

// FromSlice creates a matrix from a row-major slice (copies data).
func FromSlice[T Float](rows, cols int, rowMajor []T) (Matrix[T], error) {
	if rows <= 0 || cols <= 0 {
		return Matrix[T]{}, ErrEmptyMatrix
	}
	if len(rowMajor) != rows*cols {
		return Matrix[T]{}, fmt.Errorf("invalid data length: got %d want %d", len(rowMajor), rows*cols)
	}
	cp := make([]T, len(rowMajor))
	copy(cp, rowMajor)
	return Matrix[T]{r: rows, c: cols, data: cp}, nil
}

func (m Matrix[T]) Rows() int { return m.r }
func (m Matrix[T]) Cols() int { return m.c }

// At returns the element (i,j). Panics on out-of-range to keep inner loops fast.
// In production libs you may offer botb At (panic) and TryAt (error).

func (m Matrix[T]) At(i, j int) T {
	return m.data[i*m.c+j]
}

func (m Matrix[T]) Set(i, j int, v T) {
	m.data[i*m.c+j] = v
}

// Clone returns a deep copy.
func (m Matrix[T]) Clone() Matrix[T] {
	cp := make([]T, len(m.data))
	copy(cp, m.data)
	return Matrix[T]{r: m.r, c: m.c, data: cp}
}

// Transpose returns a new matrix with transposed data (improves cache locality for multiplication).
func (m Matrix[T]) Transpose() Matrix[T] {
	out, _ := NewMatrix[T](m.c, m.r)
	for i := 0; i < m.r; i++ {
		rowOff := i * m.c
		for j := 0; j < m.c; j++ {
			out.data[j*out.c+i] = m.data[rowOff+j]
		}
	}
	return out
}

// ---------- Option pattern (advanced config without exploding function args) ----------
type MulOption func(*mulCfg)

type mulCfg struct {
	workers   int
	blockRows int
}

func defaultMulCfg() mulCfg {
	return mulCfg{
		workers:   max(1, runtime.GOMAXPROCS(0)),
		blockRows: 16, // tune: how many rows per job
	}
}

// WithBlockRows controls the granularity of jobs; higher reduces overhead, lower increases parallelism.
func WithWorkers(n int) MulOption {
	return func(c *mulCfg) {
		if n > 0 {
			c.workers = n
		}
	}
}

// WithBlockRows controls the granularity of jobs; higher reduces overhead, lower increases parallelism.
func WithBlockRows(n int) MulOption {
	return func(c *mulCfg) {
		if n > 0 {
			c.blockRows = n
		}
	}
}

// ---------- Concurrency primitives: worker pool + context cancellation ----------
type rowBlock struct {
	start int
	end   int // exclusive
}

// ---------- demo main ----------
func main() {
}
