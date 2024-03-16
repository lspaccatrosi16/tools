package pipes

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

var oloc = flag.String("o", "", "path to input")

func DoOutput(pipeAllowed bool, out []byte) {
	if !flag.Parsed() {
		flag.Parse()
	}

	if *oloc != "" {
		dst, err := os.Create(*oloc)
		if err != nil {
			panic(err)
		}
		defer dst.Close()
		buf := bytes.NewBuffer(out)
		io.Copy(dst, buf)
	} else if !pipeAllowed {
		fmt.Println("expected -o")
		os.Exit(1)
	} else {
		os.Stdout.Write(out)
	}
}
