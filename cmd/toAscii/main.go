package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/lspaccatrosi16/go-cli-tools/input"
)

var ipt = flag.String("i", "", "input path")
var opt = flag.String("o", "", "output path")

func main() {
	replacements := map[string]byte{}

	flag.Parse()

	if *ipt == "" {
		fmt.Println("must provide an input")
	}

	if *opt == "" {
		fmt.Println("must provide an output")
	}

	src, err := os.Open(*ipt)
	if err != nil {
		panic(err)
	}

	defer src.Close()

	dst, err := os.Create(*opt)
	if err != nil {
		panic(err)
	}

	defer dst.Close()

	buf := bytes.NewBuffer(nil)

	io.Copy(buf, src)

	parsed := bytes.NewBuffer(nil)

	for {
		b1, err := buf.ReadByte()

		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}

		ch, err := readChar(buf, b1)
		if err != nil {
			panic(err)
		}

		if len([]byte(ch)) > 1 {
			key := ""
			for _, b := range ch {
				key += fmt.Sprintf("%x", b)
			}

			repl, ok := replacements[key]

			if !ok {
				repl = getReplace(ch)
				replacements[key] = repl
			}
			parsed.WriteByte(repl)
		} else {
			parsed.WriteByte(ch[0])
		}
	}

	io.Copy(dst, parsed)

}

func readChar(buf *bytes.Buffer, b1 byte) ([]byte, error) {
	var ch []byte
	if b1&0b10000000 == 0 {
		//0xxxxxxx
		ch = []byte{b1}
	} else if b1&0b01000000 == 0 {
		//10xxxxxx
		return nil, fmt.Errorf("invalid utf8 char")
	} else if b1&0b00100000 == 0 {
		//110xxxxx
		b2, err := buf.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("unfinished utf8 char")
		}
		ch = []byte{b1, b2}
	} else if b1&0b00010000 == 0 {
		//1110xxxx
		b2, err := buf.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("unfinished utf8 char")
		}
		b3, err := buf.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("unfinished utf8 char")
		}
		ch = []byte{b1, b2, b3}
	} else if b1&0b00001000 == 0 {
		//11110xxx
		b2, err := buf.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("unfinished utf8 char")
		}
		b3, err := buf.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("unfinished utf8 char")
		}
		b4, err := buf.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("unfinished utf8 char")
		}
		ch = []byte{b1, b2, b3, b4}
	} else {
		return nil, fmt.Errorf("invalid utf8 char")
	}
	return ch, nil
}

func getReplace(ch []byte) byte {
	fmt.Printf("%s was found in your input.\n", ch)
inputChar:
	repl := input.GetInput("Replacement character")

	if len(repl) != 1 {
		fmt.Println("expected a single character")
		goto inputChar
	}

	return repl[0]
}
