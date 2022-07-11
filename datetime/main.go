package main

import (
	"fmt"
	"time"
)

func main() {

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Value 1: ")
	// input1, _ := reader.ReadString('\n')
	// float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic("Value 1 must be a number")
	// }
	// fmt.Print("Value 2: ")
	// fmt.Println("A simple calculator")

	// panic("This is  my panic")

	t := time.Now()
	fmt.Print(t.Format("m/d/yyyy"))
}
