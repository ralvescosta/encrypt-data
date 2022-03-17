package io

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var inputFolder = "input"
var inputFileName = "data.json"
var outputFolder = "output"

var osGetwd = os.Getwd
var readFile = ioutil.ReadFile
var writeFile = ioutil.WriteFile
var jsonUnmarshal = json.Unmarshal
var jsonMarshalIndent = json.MarshalIndent

type InputData struct {
	PublicKey string `json:"public_key"`
	Payload   map[string]interface{}
}

func getInputFile() (string, error) {
	p, err := osGetwd()
	if err != nil {
		log.Println("[IO::getInputFile] get path error")
		return "", err
	}

	return fmt.Sprintf("%s%s%s%s%s", p, string(os.PathSeparator), inputFolder, string(os.PathSeparator), inputFileName), nil
}

func getOutputFile() (string, error) {
	p, err := osGetwd()
	if err != nil {
		log.Println("[IO::getOutputFile] get path error")
		return "", err
	}

	return fmt.Sprintf("%s%s%s%s%s", p, string(os.PathSeparator), "output", string(os.PathSeparator), time.Now().Format(time.RFC3339)+".json"), nil
}

func ReadInput() (*InputData, error) {
	log.Println("[IO::ReadInput] stating ReadInput...")
	path, err := getInputFile()
	if err != nil {
		return nil, err
	}
	log.Printf("[IO::ReadInput] reading file %s\n", path)

	file, err := readFile(path)
	if err != nil {
		log.Println("[IO::ReadInput] open file")
		return nil, err
	}

	var data *InputData
	err = jsonUnmarshal(file, data)
	if err != nil {
		log.Println("[IO::ReadInput] parsing data")
		return nil, err
	}

	log.Println("[IO::ReadInput] file read successfully")
	return data, nil
}

type OutputData struct {
	EncryptedData string `json:"encrypted_data"`
}

func WriteOutput(encrypted []byte) error {
	log.Println("[IO::WriteOutput] writing output...")

	fileName, err := getOutputFile()
	if err != nil {
		return err
	}

	fileData, err := jsonMarshalIndent(OutputData{string(encrypted)}, "", " ")
	if err != nil {
		log.Println("[IO::WriteOutput] parsing error")
		return err
	}

	err = writeFile(fileName, fileData, 0644)
	if err != nil {
		log.Println("[IO::WriteOutput] write file error")
	}

	log.Printf("[IO::WriteOutput] file written: %s", fileName)
	return err
}
