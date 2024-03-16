package types

import (
	"bytes"
	"fmt"
	"strings"
)

type Table [][]float64

func (t *Table) Add(x, y int, v float64) {
	if t == nil {
		t = new(Table)
	}

	if (*t)[x] == nil {
		(*t)[x] = []float64{}
	}
	(*t)[x][y] = v
}

func (t *Table) Get(x, y int) float64 {
	if t == nil {
		return 0
	}

	if x < len(*t) {
		if y < len((*t)[x]) {
			return (*t)[x][y]
		}
	}
	return 0
}

func (t *Table) Len() (int, int) {
	if t == nil {
		return 0, 0
	}
	if len(*t) == 0 {
		return 0, 0
	}

	return len(*t), len((*t)[0])
}

func (t *Table) String() string {
	if t == nil {
		return ""
	}
	buf := bytes.NewBuffer(nil)

	xS, yS := t.Len()

	values := make([][]float64, yS)
	for i := 0; i < yS; i++ {
		values[i] = make([]float64, xS)
	}

	for x, c := range *t {
		for y, f := range c {
			values[y][x] = f
		}
	}

	for _, l := range values {
		entries := []string{}
		for _, f := range l {

			entries = append(entries, fmt.Sprintf(" %10.5f ", f))
		}

		fmt.Fprintln(buf, "|"+strings.Join(entries, "|")+"|")

	}

	return buf.String()
}

func (t *Table) Copy(expX, expY int) *Table {
	xS, yS := t.Len()
	newT := MakeTable(xS+expX, yS+expY)

	for x := 0; x < xS; x++ {
		for y := 0; y < yS; y++ {
			f := t.Get(x, y)
			newT.Add(x, y, f)
		}
	}

	return newT
}

func MakeTable(x, y int) *Table {
	t := make(Table, x)
	for i := 0; i < x; i++ {
		yA := make([]float64, y)
		for j := 0; j < y; j++ {
			yA[j] = 0
		}
		t[i] = yA
	}
	return &t
}
