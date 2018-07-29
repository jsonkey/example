package main

import (
	"fmt"

	jutils "github.com/jsonkey/example/src/utils"
)

type PersonInfo struct {
	ID   string
	Name string
	Age  int16
}

func main() {

	fmt.Println(jutils.IsFileExist("/anc"))

	var person1 map[string]PersonInfo

	person1 = make(map[string]PersonInfo, 2)

	person1["123"] = PersonInfo{"123", "tommy", 30}
	fmt.Println("你好", person1["123"])

	person, ok := person1["23"]

	// ok return bool
	if ok {
		fmt.Println("Found person", person)
	} else {
		fmt.Println("Not Found")
	}

	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.23456
	MyPrintf(v1, v2, v3, v4)

}

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is a int type")
		case string:
			fmt.Println(arg, "is string type")
		case int64:
			fmt.Println(arg, "is int64 type")
		default:
			fmt.Println(arg, "is default type")
		}
	}
}
