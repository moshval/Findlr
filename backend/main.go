package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/moshval/Findlr/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Words struct {
	Words []string `json:"words"`
}
type GetRequest struct {
	Tipe int    `form:"type"`
	Word string `form:"word"`
	Bad  string `form:"bad"`
	Exc  string `form:"exc"`
}
type PostRequest struct {
	Tipe int    `json:"type"`
	Word string `json:"word"`
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
		var word Words
		json.Unmarshal(byteResult, &word)
		arr = word.Words
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
		file, _ := json.MarshalIndent(word, "", "    ")
		ioutil.WriteFile("./words.json", file, 0644)
	} else {
		fileContent, err := os.Open("./kata.json")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer fileContent.Close()
		var word Words
		word.Words = arr
		file, _ := json.MarshalIndent(word, "", "    ")
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

func deleteString(deleted string, arr *[]string) bool {
	var isDeleted = false
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i] == deleted {
			*arr = append((*arr)[:i], (*arr)[i+1:]...)
			i--
			isDeleted = true
		}
	}
	return isDeleted
}
func insertString(inserted string, arr *[]string) bool {
	if !contains(inserted, *arr) {
		*arr = append(*arr, inserted)
		return true
	}
	return false
}
func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return "localhost:" + port
}

func filterWords(tipe int, word string, exc string, bad string) []string {
	arr := parseFromFile(tipe)
	fmt.Println(exc)
	var filtered []string
	for _, v := range arr {
		// assume word len == bad len
		for i := range word {
			if (v[i] != word[i] && string(word[i]) != "_") || (bad[i] == v[i]) || (!strings.Contains(v, string(bad[i])) && string(bad[i]) != "_") {
				break
			}
			if i == len(word)-1 {
				filtered = append(filtered, v)
			}
		}
	}
	if exc != "" {
		for l := 0; l < len(exc); l++ {
			for k := 0; k < len(filtered); k++ {
				curr := filtered[k]
				if strings.Contains(curr, string(exc[l])) {
					filtered = append(filtered[:k], filtered[k+1:]...)
					k--
				}
			}
		}
	}
	return filtered
}

// getWords ... Get Words based on query
// @Summary Get Words Based on Query
// @Description Get Words Based on Query
// @Param type query int true "Language (0 : EN, 1 : ID)"
// @Param word query string true "Word"
// @Param bad query string true "Bad Word Placement"
// @Param exc query string false "Excluded Letters"
// @Tags Words
// @Success 200 {array} Words
// @Router /words [get]
func getWords(c *gin.Context) {
	var getreq GetRequest
	if (c.ShouldBind(&getreq)) == nil {
		var arr []string
		if getreq.Tipe == 0 {
			wrd := string(getreq.Word)
			bad := string(getreq.Bad)
			exc := string(getreq.Exc)
			arr = filterWords(0, wrd, exc, bad)
		} else {
			wrd := string(getreq.Word)
			bad := string(getreq.Bad)

			exc := string(getreq.Exc)
			arr = filterWords(1, wrd, exc, bad)
		}
		c.JSON(200, gin.H{
			"words": arr,
		})
	}
}

// deleteWord ... Delete Words
// @Summary Delete Words based on parameters
// @Description Delete Words based on parameters
// @Tags Words
// @Accept json
// @Param PostRequest body PostRequest true "Deleted Word" default(PostRequest)
// @Success 200 {string} string
// @Router /words/delete [post]
func deleteWord(c *gin.Context) {
	var getreq PostRequest
	if err := c.BindJSON(&getreq); err != nil {
		return
	}
	var arr []string
	var del bool
	if getreq.Tipe == 0 {
		wrd := string(getreq.Word)
		arr = parseFromFile(0)
		del = deleteString(wrd, &arr)
		updateFile(0, arr)
	} else {
		wrd := string(getreq.Word)
		arr = parseFromFile(1)
		del = deleteString(wrd, &arr)
		updateFile(1, arr)
	}
	if del {
		c.JSON(200, "Deleted "+getreq.Word)
	} else {
		c.JSON(200, "Not Existed")
	}
}

// addWord ... Add Words
// @Summary Add new Words based on parameters
// @Description Add new Words based on parameters
// @Tags Words
// @Accept json
// @Param PostRequest body PostRequest true "Added Word" default(PostRequest)
// @Success 200 {string} string
// @Router /words/add [post]
func addWord(c *gin.Context) {
	var getreq PostRequest
	if err := c.BindJSON(&getreq); err != nil {
		return
	}
	var arr []string
	var ins bool
	if getreq.Tipe == 0 {
		wrd := string(getreq.Word)
		arr = parseFromFile(0)
		ins = insertString(wrd, &arr)
		updateFile(0, arr)
	} else {
		wrd := string(getreq.Word)
		arr = parseFromFile(1)
		ins = insertString(wrd, &arr)
		updateFile(1, arr)
	}
	if ins {
		c.JSON(200, "Inserted "+getreq.Word)
	} else {
		c.JSON(200, "Already Exist")
	}
}

// @title Words API Documentation
// @version 1.0.0
// @host localhost:8080

func main() {
	r := gin.Default()
	r.GET("/words", getWords)
	r.POST("/words/add", addWord)
	r.POST("/words/delete", deleteWord)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(getPort())

}
