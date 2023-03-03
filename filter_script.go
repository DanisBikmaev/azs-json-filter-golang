package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func filter_script(JsonFileName string, filterBy string) {
	byteJsonFile := get_json(JsonFileName)
	data := filtered_json(byteJsonFile, filterBy)
	generate_json_file(JsonFileName, data)
}

func get_json(inputJsonFileName string) []byte {
	fmt.Println("azs-json-sort")
	jsonFile, err := os.Open("json_input/" + inputJsonFileName + ".json")
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

func generate_json_file(JsonFileName string, data []byte) {
	os.WriteFile("json_output/"+JsonFileName+".json", data, 0644)
}
