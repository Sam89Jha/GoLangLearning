package main

import "fmt"

func main() {

	poodle := Dog{"Poodle", 10, "Woff!!"}
	fmt.Printf("%+v\n", poodle)
	poodle.Speak()
	poodle.Sound = "Arf!"
	poodle.Speak()
	poodle.speakThreeTimes()
	poodle.speakThreeTimes()
}

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d Dog) speakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %vb %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
