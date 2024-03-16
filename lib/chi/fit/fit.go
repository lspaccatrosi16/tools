package fit

import (
	"fmt"

	"github.com/lspaccatrosi16/go-cli-tools/input"
	"github.com/lspaccatrosi16/go-cli-tools/logging"
	"github.com/lspaccatrosi16/tools/lib/chi/common"
	"github.com/lspaccatrosi16/tools/lib/chi/types"
	"gonum.org/v1/gonum/stat/distuv"
)

func GetFitTable(observed *types.Table, totals *types.Table) *types.Table {
	logger := logging.GetLogger()

	logger.LogDivider()
	logger.Log("Distribution Test Fit")

	xS, yS := observed.Len()

	total := totals.Get(xS, 0)

	options := []input.SelectOption{
		{Name: "Binomial", Value: "b"},
		{Name: "Poisson", Value: "p"},
		{Name: "Uniform", Value: "u"},
		{Name: "Custom", Value: "c"},
	}

	chosen, err := input.GetSelection("Distribution Type", options)

	if err != nil {
		panic(err)
	}

	expected := types.MakeTable(xS, yS)

	xopts := []input.SelectOption{
		{Name: "=", Value: "EQ"},
		{Name: ">=", Value: "GTE"},
		{Name: "<=", Value: "LTE"},
	}

	switch chosen {
	case "b":
		trials := common.ParseFloat(input.GetInput("Trials (n)"))
		prob := common.ParseFloat(input.GetInput("Probability (p)"))
		bDist := distuv.Binomial{N: trials, P: prob}
		for x := 0; x < xS; x++ {
			sel, err := input.GetSelection("Value Type", xopts)
			if err != nil {
				panic(err)
			}

			value := common.ParseInt(input.GetInput("Value"))

			switch sel {
			case "EQ":
				expected.Add(x, 0, bDist.Prob(float64(value))*total)
			case "LTE":
				expected.Add(x, 0, bDist.CDF(float64(value))*total)
			case "GTE":
				expected.Add(x, 0, 1-bDist.CDF(float64(value-1))*total)
			}
		}

	case "p":
		mean := common.ParseFloat(input.GetInput("Probability (lambda)"))
		pDist := distuv.Poisson{Lambda: mean}
		for x := 0; x < xS; x++ {
			sel, err := input.GetSelection("Value Type", xopts)
			if err != nil {
				panic(err)
			}

			value := common.ParseFloat(input.GetInput("Value"))

			switch sel {
			case "EQ":
				expected.Add(x, 0, pDist.Prob(value)*total)
			case "LTE":
				expected.Add(x, 0, pDist.CDF(value)*total)
			case "GTE":
				expected.Add(x, 0, (1-pDist.CDF(value-1))*total)
			}
		}

	case "u":
		probability := 1.0 / float64(xS)
		for x := 0; x < xS; x++ {
			expected.Add(x, 0, probability*total)
		}
	case "c":
		logger.Log("Enter Custom Probabilities")
		for x := 0; x < xS; x++ {
			p := common.ParseFloat(input.GetInput(fmt.Sprintf("x_%d", x)))
			expected.Add(x, 0, p*total)
		}
		common.GetTableValues(expected)
	}

	return expected

}
