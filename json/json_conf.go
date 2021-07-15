package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ConnectionString struct {
	Driver string
	Server string
	Database string
	UID string
	PWD string
	Trusted string
}

type JStruct struct {
	device map[string]ConnectionString
}

func getJSONData(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}

func main() {
	data := getJSONData("json/redcheck.json")
	//fmt.Println(string(data))
	// var s JStruct
	var s interface{}
	if err := json.Unmarshal(data, &s); err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("%v", s)
	//for _, i := range s {
	//
	//}
	//fmt.Println(s["suis"])
}
