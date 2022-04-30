package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Words struct {
	Words []string `json:"words"`
}
type Kata struct {
	Kata []string `json:"kata"`
}

func parseFromFile(tipe int) (arr []string) {
	arr = []string{}
	if tipe == 0 {
		fileContent, err := os.Open("./words.json")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer fileContent.Close()
		byteResult, _ := ioutil.ReadAll(fileContent)
		var word Words
		json.Unmarshal(byteResult, &word)
		arr = word.Words
	} else {
		fileContent, err := os.Open("./kata.json")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer fileContent.Close()
		byteResult, _ := ioutil.ReadAll(fileContent)
		var kata Kata
		json.Unmarshal(byteResult, &kata)
		arr = kata.Kata
	}
	return

}

func updateFile(tipe int, arr []string) {
	if tipe == 0 {
		fileContent, err := os.Open("./words.json")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer fileContent.Close()
		var word Words
		word.Words = arr
		file, _ := json.MarshalIndent(word, "", " ")
		ioutil.WriteFile("./words.json", file, 0644)
	} else {
		fileContent, err := os.Open("./kata.json")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer fileContent.Close()
		var kata Kata
		kata.Kata = arr
		file, _ := json.MarshalIndent(kata, "", " ")
		ioutil.WriteFile("./kata.json", file, 0644)
	}
}
func insertArray(inserted []string, arr *[]string) {
	for i := 0; i < len(inserted); i++ {
		if !contains(inserted[i], *arr) {
			*arr = append(*arr, inserted[i])
		}
	}
}
func contains(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}
func deleteString(deleted string, arr *[]string) {
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i] == deleted {
			*arr = append((*arr)[:i], (*arr)[i+1:]...)
			i--
		}
	}
}
func insertString(inserted string, arr *[]string) {
	if !contains(inserted, *arr) {
		*arr = append(*arr, inserted)
	}
}
func main() {
	// var arr = []string{}
	// arr = parseFromFile(1)
	// insertString("sheva", &arr)
	// updateFile(1, arr)
}
