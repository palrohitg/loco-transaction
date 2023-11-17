//package main
//
//import (
//	"encoding/json"
//	"fmt"
//)
//
//type Person struct {
//	Name string `json:"name"`
//	Age  int    `json:"age"`
//}
//
//func main() {
//	jsonData := []byte(`{"name":"Bob","age":25}`)
//	var person Person
//	err := json.Unmarshal(jsonData, &person) // jsonData --> Golang Objects
//	if err != nil {
//		fmt.Println("error occurred here")
//	}
//	fmt.Print("Testing the construct here ")
//	fmt.Println(person)
//	fmt.Println("Final version to make the changes")
//	person1 := Person{
//		Name: "vikas",
//		Age:  24,
//	}
//	jsonData, err1 := json.Marshal(person1) // Golang DataStructure --> JSON objects
//	if err1 != nil {
//		fmt.Println(err1)
//	}
//	fmt.Println(jsonData)
//	return
//}

package main

import "fmt"

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Sample JSON data
	//jsonData := []byte(`{"name":"Bob","age":25}`)
	//
	//// Creating a Person struct to hold the unmarshaled data
	//var person Person
	//
	//// Unmarshaling the JSON into the struct
	//err := json.Unmarshal(jsonData, &person)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//
	//// Displaying the unmarshaled data
	//fmt.Println(person.Name)
	//fmt.Println(person.Age)
	ans := 14 / 7
	fmt.Println(ans)
}
