package main

import (
	"fmt"
	"os"

	"github.com/lspaccatrosi16/go-cli-tools/args"
	"github.com/lspaccatrosi16/go-libs/fs"
)

func init() {
	args.RegisterEntry(args.NewStringEntry("compression", "c", "archive format", "tar.gz"))
	args.RegisterEntry(args.NewStringEntry("input", "i", "input path", ""))
	args.RegisterEntry(args.NewStringEntry("output", "o", "output path", ""))
}

func main() {
	err := args.ParseOpts()
	handle(err)

	compression, err := args.GetFlagValue[string]("compression")
	handle(err)
	var usedComp fs.CompressionType
	switch compression {
	case "tar.gz":
		usedComp = fs.TarGz
	case "zip":
		usedComp = fs.Zip
	default:
		handle(fmt.Errorf("unknown compression type: %s", compression))
	}

	input, err := args.GetFlagValue[string]("input")
	handle(err)
	if input == "" {
		handle(fmt.Errorf("no input path provided"))
	}

	output, err := args.GetFlagValue[string]("output")
	handle(err)
	if output == "" {
		handle(fmt.Errorf("no output path provided"))
	}

	cmd := args.GetArgs()
	if len(cmd) == 0 {
		handle(fmt.Errorf("no command provided"))
	}

	switch cmd[0] {
	case "compress":
		err = compress(input, output, usedComp)
	case "extract":
		err = decompress(input, output, usedComp)
	default:
		err = fmt.Errorf("unknown command: %s", cmd[0])
	}

	handle(err)
}

func handle(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func compress(input, output string, compression fs.CompressionType) error {
	f, err := os.Create(output)
	if err != nil {
		return err
	}

	defer f.Close()
	return fs.Compress(input, f, compression)
}

func decompress(input, output string, compression fs.CompressionType) error {
	f, err := os.Open(input)
	if err != nil {
		return err
	}

	defer f.Close()
	return fs.Decompress(f, output, compression)
}
