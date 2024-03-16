package common

import (
	"fmt"
	"math"
	"strconv"

	"github.com/lspaccatrosi16/go-cli-tools/input"
	"github.com/lspaccatrosi16/tools/lib/chi/types"
)

func MakeTotals(table *types.Table) *types.Table {
	xS, yS := table.Len()
	totals := table.Copy(1, 1)

	xtTotal := 0.0
	ytTotal := 0.0

	for y := 0; y < yS; y++ {
		t := 0.0
		for x := 0; x < xS; x++ {
			f := table.Get(x, y)
			t += f
		}
		xtTotal += t
		totals.Add(xS, y, t)
	}

	for x := 0; x < xS; x++ {
		t := 0.0
		for y := 0; y < yS; y++ {
			f := table.Get(x, y)
			t += f
		}
		ytTotal += t
		totals.Add(x, yS, t)
	}

	if xtTotal == ytTotal {
		totals.Add(xS, yS, xtTotal)
	} else {
		panic("xtotals and xtotals do not match")
	}
	return totals
}

func MakeChiStatistic(observed *types.Table, expected *types.Table) (*types.Table, float64) {
	chiStatistic := observed.Copy(0, 0)

	xS, yS := observed.Len()

	total := 0.0

	for x := 0; x < xS; x++ {
		for y := 0; y < yS; y++ {
			xf := expected.Get(x, y)
			chiStat := math.Pow(observed.Get(x, y)-xf, 2) / xf
			total += chiStat
			chiStatistic.Add(x, y, chiStat)

		}
	}

	return chiStatistic, total
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Errorf("cannot parse \"%s\"", s))
	}
	return i
}

func ParseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(fmt.Errorf("cannot parse \"%s\"", s))
	}
	return f
}

func GetTableValues(table *types.Table) {
	if table == nil {
		panic("nil table")
	}
	xS, yS := table.Len()

	for y := 0; y < yS; y++ {
		for x := 0; x < xS; x++ {
			msg := fmt.Sprintf("[%d %d]", x, y)
			val := ParseFloat(input.GetInput(msg))
			table.Add(x, y, val)
		}
	}
}
