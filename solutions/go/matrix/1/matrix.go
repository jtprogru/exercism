package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (*Matrix, error) {
	var err error
	lines := strings.Split(s, "\n")
	m := make(Matrix, len(lines))
	for idx, line := range lines {
		ws := strings.Fields(line)
		if idx > 0 && len(ws) != len(m[0]) {
			return nil, errors.New("rows have unequal length")
		}
		row := make([]int, len(ws))
		for i, w := range ws {
			if row[i], err = strconv.Atoi(w); err != nil {
				return nil, errors.New("invalid int in element data")
			}
		}
		m[idx] = row
	}

	return &m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	mm := *m
	if len(mm) == 0 {
		return nil
	}
	c := make([][]int, len(mm[0]))
	for i := range c {
		col := make([]int, len(mm))
		for j := range col {
			col[j] = mm[j][i]
		}
		c[i] = col
	}
	return c
}

func (m *Matrix) Rows() [][]int {
	row := make([][]int, len(*m))
	for idx, mr := range *m {
		row[idx] = append([]int{}, mr...)
	}
	return row
}

func (m *Matrix) Set(row, col, val int) bool {
	mm := *m
	if row < 0 || row >= len(mm) || col < 0 {
		return false
	}
	if cols := len(mm[0]); col >= cols {
		return false
	}
	mm[row][col] = val
	return true
}
