package main

import (
	"fmt"

	"github.com/lspaccatrosi16/go-cli-tools/input"
	"github.com/lspaccatrosi16/go-cli-tools/logging"
	"github.com/lspaccatrosi16/tools/lib/chi/common"
	"github.com/lspaccatrosi16/tools/lib/chi/contingency"
	"github.com/lspaccatrosi16/tools/lib/chi/fit"
	"github.com/lspaccatrosi16/tools/lib/chi/types"
)

func main() {
	var xS, yS int

	logger := logging.GetLogger()

	logger.Log("Chi Squared Table Producer")

	options := []input.SelectOption{
		{Name: "Contingency Table", Value: "c"},
		{Name: "Test For a Fit", Value: "f"},
	}

	choice, err := input.GetSelection("Type of Test", options)

	if err != nil {
		panic(err)
	}
	xStr := input.GetInput("Number of Columns")
	xS = common.ParseInt(xStr)

	if choice == "f" {
		yS = 1
	} else {
		yStr := input.GetInput("Number of Rows")
		yS = common.ParseInt(yStr)
	}

	observed := types.MakeTable(xS, yS)

	logger.LogDivider()
	logger.Log("Enter Table Values")

	common.GetTableValues(observed)
	totals := common.MakeTotals(observed)

	logger.LogDivider()
	logger.Log("Totals Table")

	logger.Log(totals.String())

	var expected *types.Table

	switch choice {
	case "c":
		expected = contingency.GetContingencyTable(observed, totals)
	case "f":
		expected = fit.GetFitTable(observed, totals)
	}

	logger.LogDivider()
	logger.Log("Expected Frequencies Table")
	logger.Log(expected.String())

	chiStatistic, totalChi := common.MakeChiStatistic(observed, expected)

	logger.LogDivider()
	logger.Log("Chi Squared Statistic Table")

	logger.Log(chiStatistic.String())

	logger.LogDivider()
	logger.Log(fmt.Sprintf("Total Chi Squared Statistic: %.10f", totalChi))
}
