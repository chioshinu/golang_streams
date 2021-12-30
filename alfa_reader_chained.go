package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type alphaReaderChaned struct {
	reader io.Reader
}

func newAlphaReaderChaned(reader io.Reader) io.Reader {
	return &alphaReaderChaned{reader: reader}
}

func (a *alphaReaderChaned) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}

	copy(p, buf)
	return n, nil
}

func alfa_reader_chained_example() {
	// use an io.Reader as source for alphaReader
	reader := newAlphaReaderChaned(strings.NewReader("Hello! It's 9am, where is the sun?"))
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}

func alfa_reader_chained_file() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	reader := newAlphaReaderChaned(file)
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}
