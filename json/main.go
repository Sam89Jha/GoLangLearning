package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Resposne type %T\n", resp)
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	content := string(bytes)
	//fmt.Println(content)
	tours := ToursFromJson(content)
	for _, val := range tours {
		fmt.Println(val.Name)
	}
}

type Tour struct {
	Name, Price string
}

func ToursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}
	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}
		tours = append(tours, tour)
	}
	return tours
}
