package main

import (
	"fmt"
	"math"
	"time"
)

//const aConst int = 64

func main() {
	// var aString string = "This is Go!"
	// fmt.Println(aString)
	// fmt.Printf("The variable type is %T\n", aString)

	// var aInteger int
	// fmt.Println(aInteger)

	// var defaultInt int
	// fmt.Println((defaultInt))

	// var anotherString = "This is another string"
	// fmt.Println(anotherString)
	// fmt.Printf("The variable type is %T\n", anotherString)

	// var anotherInt = 53
	// fmt.Println(anotherInt)
	// fmt.Printf("The variable type is %T\n", anotherInt)

	// myString := "This is also a string"
	// fmt.Println(myString)
	// fmt.Printf("The variable type is %T\n", myString)

	// fmt.Println(aConst)
	// fmt.Printf("The variable type is %T\n", aConst)

	/*reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text : ")

	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	fmt.Print("Enter a number")
	numInt, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInt), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of numnber :", aFloat)
	}
	*/

	//TypeCastPractice()

	workWithDateAndTimes()

}

func TypeCastPractice() {
	i1, i2, i3 := 12, 4566, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integrer sum is ", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum is ", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("Float sum is ", floatSum)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Cisrcumference: %.2f:\n", circumference)
}

func workWithDateAndTimes() {
	n := time.Now()
	fmt.Println("Recorded this training at ", n)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go Launch at ", t)

	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	fmt.Printf("Parsed time is %T", parsedTime)
}
