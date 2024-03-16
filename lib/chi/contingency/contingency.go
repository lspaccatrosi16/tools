package contingency

import (
	"github.com/lspaccatrosi16/go-cli-tools/logging"
	"github.com/lspaccatrosi16/tools/lib/chi/types"
)

func makeExpectedFrequencies(table *types.Table, totals *types.Table) *types.Table {
	expectedFreq := table.Copy(0, 0)

	xS, yS := table.Len()
	for x := 0; x < xS; x++ {
		for y := 0; y < yS; y++ {
			xF := (totals.Get(x, yS) * totals.Get(xS, y)) / totals.Get(xS, yS)
			expectedFreq.Add(x, y, xF)
		}
	}

	return expectedFreq

}

func GetContingencyTable(observed *types.Table, totals *types.Table) *types.Table {
	logger := logging.GetLogger()

	expected := makeExpectedFrequencies(observed, totals)

	logger.LogDivider()
	logger.Log("Expected Frequencies Table")

	logger.Log(expected.String())

	return expected
}
