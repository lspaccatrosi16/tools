package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	case "compress", "extract":
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

		if cmd[0] == "compress" {
			err = compress(input, output, usedComp)
		} else {
			err = decompress(input, output, usedComp)
		}
	case "local":
		err = local(output)
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

type mfestEntry struct {
	AppName string `json:"appName"`
	Path    string `json:"cfgPath"`
	Replace bool   `json:"replace"`
}

func local(output string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	mfest, err := os.ReadFile(filepath.Join(wd, "manifest.json"))
	if err != nil {
		return err
	}

	var manifest []mfestEntry
	err = json.Unmarshal(mfest, &manifest)
	if err != nil {
		return err
	}

	found := false

	for _, ent := range manifest {
		if ent.AppName == output {
			found = true
			hd, err := os.UserHomeDir()
			if err != nil {
				return nil
			}
			newP := strings.ReplaceAll(ent.Path, "$HOME", hd)

			err = os.RemoveAll(ent.AppName)
			if err != nil {
				return err
			}
			err = fs.Copy(filepath.Join(wd, ent.AppName), newP)
			return err
		}
	}

	if !found {
		return fmt.Errorf("could not find config for %s in the manifest", output)
	}
	return nil
}
