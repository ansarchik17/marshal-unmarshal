package main

import (
	"encoding/json"
	"fmt"
)

type Teacher struct {
	Name  string
	Age   int
	Year  int
	Email string
}

func unmarshal() {
	jsonStr := `{"Name":"Ansar","Age":17,"Year":2007,"Email":"kelesansar62@gmail.com"}`
	var teacher Teacher

	err := json.Unmarshal([]byte(jsonStr), &teacher)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(teacher)
}
func main() {
	unmarshal()
}
