package matrix

import "fmt"

type Matrix[T any] struct {
	data [][]T
	rows int
	cols int
}

func New[T any](rows, cols int) *Matrix[T] {
	data := make([][]T, rows)
	for i := range data {
		data[i] = make([]T, cols)
	}
	
	return &Matrix[T]{
		data: data,
		rows: rows,
		cols: cols,
	}
}

func (m *Matrix[T]) Set(row, col int, val T) {
	if row >= 0 && row < m.rows && col >= 0 && col < m.cols {
		m.data[row][col] = val
	}
}

func (m *Matrix[T]) Get(row, col int) (T, bool) {
	if row >= 0 && row < m.rows && col >= 0 && col < m.cols {
		return m.data[row][col], true
	}
	var zero T
	return zero, false
}

func (m *Matrix[T]) Print() {
	for _, row := range m.data {
		fmt.Println(row)
	}
}