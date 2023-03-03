package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func filter_script(inputFileWithFormat string, filterBy string, outputFileWithFormat string) {
	byteJsonFile := get_json(inputFileWithFormat)
	data := filtered_json(byteJsonFile, filterBy)
	generate_json_file(outputFileWithFormat, data)
}

func get_json(jsonFileLocation string) []byte {
	fmt.Println("azs-json-sort")
	jsonFile, err := os.Open(jsonFileLocation)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened json")
	defer jsonFile.Close()

	byteJsonFile, _ := io.ReadAll(jsonFile)
	return byteJsonFile

}

func filtered_json(byteJsonFile []byte, filterBy string) []byte {
	var x []*Mark

	err := json.Unmarshal(byteJsonFile, &x)
	if err != nil {
		panic(err)
	}

	var newMarks []*Mark
	for _, mark := range x {
		if mark.Properties.HintContent == filterBy {
			newMarks = append(newMarks, mark)
		}
	}
	data, err := json.MarshalIndent(newMarks, "", "")
	fmt.Println(string(data), err)
	return (data)
}

func generate_json_file(outputFileWithFormat string, data []byte) {
	os.WriteFile(outputFileWithFormat, data, 0644)
}
