package pipes

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

var iloc = flag.String("i", "", "path to input")

func GetInput() []byte {
	if !flag.Parsed() {
		flag.Parse()
	}

	if *iloc != "" {
		src, err := os.Open(*iloc)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("input location not found")
				os.Exit(1)
			} else {
				panic(err)
			}
		}
		defer src.Close()
		buf := bytes.NewBuffer(nil)
		io.Copy(buf, src)
		return buf.Bytes()
	} else {
		stdin, err := io.ReadAll(os.Stdin)

		if err != nil {
			panic(err)
		}
		return stdin
	}
}

func PipeOut() bool {
	if !flag.Parsed() {
		flag.Parse()
	}
	return *oloc != ""
}
