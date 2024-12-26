package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Year  int
	Email string
}

func marshal() {
	student := Student{
		Name:  "Ansar",
		Age:   17,
		Year:  2007,
		Email: "kelesansar62@gmail.com",
	}
	jsonData, _ := json.Marshal(student)
	fmt.Println(string(jsonData))
}
func main() {
	marshal()
}
