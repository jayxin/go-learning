package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)
	personDB["12345"] = PersonInfo{"12345", "Tom", "China"}
	personDB["1"] = PersonInfo{"1", "Jack", "China"}

	person, ok := personDB["1234"]

	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}

	var myMap map[string]PersonInfo
	myMap = make(map[string]PersonInfo, 100)
	myMap = map[string]PersonInfo{
		"1234": {"1", "Jack", "China"},
		//"1234": PersonInfo{"1", "Jack", "China"},
	}
	fmt.Println(myMap)
	delete(myMap, "1234")
	delete(myMap, "key")
}
