package main

import (
	"flag"
	"os"

	"github.com/lspaccatrosi16/go-cli-tools/logging"
	"github.com/lspaccatrosi16/tools/lib/pipes"
	"github.com/lspaccatrosi16/tools/lib/ts-go/generator"
	"github.com/lspaccatrosi16/tools/lib/ts-go/parser"
	"github.com/lspaccatrosi16/tools/lib/ts-go/util"
)

var verbose = flag.Bool("v", false, "Display Verbose Logging")
var help = flag.Bool("h", false, "Shows the help message")
var showXml = flag.Bool("x", false, "Print intermediate XML representation of type system")
var useDefaults = flag.Bool("d", false, "Use default go types")

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	logger := logging.GetLogger()
	logger.SetVerbose(*verbose)

	if pipes.PipeOut() {
		logger.SetDisable(true)
	}

	src := pipes.GetInput(true)

	tree := parser.ParseInput(string(src))

	if *showXml {
		logger.Log(util.FormatIr(tree))
	}

	settings := generator.NewSettings(*useDefaults)

	generated := generator.Generate(settings, tree)

	pipes.DoOutput(true, generated.Bytes())
}
