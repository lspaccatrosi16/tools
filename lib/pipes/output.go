package pipes

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

var oloc = flag.String("o", "", "path to input")

func DoOutput(out []byte) {
	if !flag.Parsed() {
		flag.Parse()
	}

	if *oloc != "" {
		dst, err := os.Create(*oloc)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("input location not found")
				os.Exit(1)
			} else {
				panic(err)
			}
		}
		defer dst.Close()
		buf := bytes.NewBuffer(out)
		io.Copy(dst, buf)
	} else {
		os.Stdout.Write(out)
	}
}
