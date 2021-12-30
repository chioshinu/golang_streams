package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func read_full_file() {
	bytes, err := ioutil.ReadFile("./streams")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s", bytes)
}
