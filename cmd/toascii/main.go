package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lspaccatrosi16/go-cli-tools/input"
	"github.com/lspaccatrosi16/go-cli-tools/logging"
	"github.com/lspaccatrosi16/tools/lib/pipes"
)

func main() {
	replacements := map[string]string{}
	logger := logging.GetLogger()

	if pipes.PipeOut() {
		logger.SetDisable(true)
	}

	src := pipes.GetInput(false)
	buf := bytes.NewBuffer(src)

	fmt.Println("got input")

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
			parsed.WriteString(repl)
		} else {
			parsed.WriteByte(ch[0])
		}
	}

	pipes.DoOutput(false, parsed.Bytes())
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

func getReplace(ch []byte) string {
	fmt.Printf("%s was found in your input.\n", ch)
inputChar:
	repl := input.GetInput("Replacement string")

	if len(repl) != 1 {
		fmt.Println("expected a single character")
		goto inputChar
	}

	return repl
}
