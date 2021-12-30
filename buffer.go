package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type People struct {
	Name    string `json:"full_name,omitempty"`
	Age     int    `json:",omitempty"`
	Partner *People
}

func (p *People) SetPartner()

func json_example() {
	people := People{Name: "Vasya", Age: 10}
	second := People{Name: "Masha", Age: 13, Partner: &people}
	content, err := json.Marshal(second)
	if err != nil {
		log.Fatalf("Can't marshal\n%v\n", err)
		return
	}
	fmt.Printf("%s\n", content)
}

func buffer_example() {
	file, err := os.Open("./games.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		fmt.Print(line)
	}

}

func readInput(fileName string) ([]string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed to open")
		return nil, err
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	defer file.Close()

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)
	results := []string{}

	for scanner.Scan() {
		results = append(results, scanner.Text())
	}

	return results, nil
}
