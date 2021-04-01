package main

import (
	"bufio"
	"bytes"
	"io"
	"os"

	"sample/osc1337"
)

func main() {
	fname := "image.png"
	body, err := importFile(fname)
	if err != nil {
		panic(err.Error())
	}

	var buf bytes.Buffer
	e := osc1337.NewEncoder()
	e.W = &buf
	e.Name = []byte(fname)
	e.Width = "50%"
	e.Inline = true
	e.Type = "image/png"

	e.Encode(body)

	os.Stdout.Write(buf.Bytes())
	os.Stdout.Sync()
}

func importFile(fname string) ([]byte, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var r io.Reader = bufio.NewReader(f)

	body, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return body, nil
}
