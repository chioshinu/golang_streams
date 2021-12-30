package main

import (
	"fmt"
	"io"
	"os"
)

func write_file() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}
	file, err := os.Create("./test.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		err := file.Close()
		fmt.Printf("Error ob defer: %v\n", err)
	}()

	for _, p := range proverbs {
		n, err := file.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
	fmt.Println("file write done")
	err = file.Close()
	fmt.Printf("%v\n", err)
}

func read_file() {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	p := make([]byte, 4)
	for {
		n, err := file.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}
}
