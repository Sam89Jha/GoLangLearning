package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "hello from Go!"
	file, err := os.Create("./fromString.txt")

	checkError(err)

	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters", length)

	defer file.Close()
	defer readFile("./fromString.txt")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fielName string) {
	data, err := ioutil.ReadFile(fielName)
	checkError(err)
	fmt.Println("\nText read from file", string(data))
}
