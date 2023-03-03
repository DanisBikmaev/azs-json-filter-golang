package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func sort_script(json_file_location string) {
	fmt.Println("azs-json-sort")
	jsonFile, err := os.Open(json_file_location)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var marks Marks

	json.Unmarshal(byteValue, &marks)

	for i := 0; i < len(marks.Marks); i++ {
		fmt.Println(marks.Marks[i].Id)
	}
}
