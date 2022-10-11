package main

import (
	"errors"
	"fmt"
	"github.com/Julia1505/SafeboardGo/pkg/decoder"
	"github.com/Julia1505/SafeboardGo/pkg/parser"
	"github.com/Julia1505/SafeboardGo/pkg/people"
	"html/template"
	"os"
	"sync"
)

var (
	CSVFile = "csv"
	PRNFile = "prn"
)

var (
	NoExtension = errors.New("File doesn't have extension")
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func fileExtension(filename string) (string, error) {
	runeString := []rune(filename)
	l := len(runeString)
	ext := ""
	isComma := false
	for i := l - 1; i >= 0; i-- {
		if runeString[i] == '.' {
			isComma = true
			break
		}
		ext += string(runeString[i])
	}
	if ext == "" || isComma == false {
		return "", NoExtension
	}
	ext = reverse(ext)
	return ext, nil
}

func newFileName(filename string) string {
	runeString := []rune(filename)
	l := len(runeString)
	var ind int
	for i := l - 1; i >= 0; i-- {
		if runeString[i] == '.' {
			ind = i
			break
		}
	}

	runeString = runeString[:ind+1]
	res := string(runeString) + "html"

	return res
}

func main() {
	prefix := "./data/"
	filename := "data.csv"
	typeFile, err := fileExtension(filename)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}

	path := prefix + filename

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	defer file.Close()

	wg := &sync.WaitGroup{}

	transport := make(chan string)
	wg.Add(2)
	fileDecoder, err := decoder.NewDecoder("ISO8859_1")
	if err != nil {
		fmt.Printf("error happened: %v", err)
		return
	}

	go func() {
		defer wg.Done()
		fileDecoder.Decode(file, transport)
	}()

	parserFile, err := parser.NewParser(typeFile)
	if err != nil {
		fmt.Printf("error happened: %v", err)
		return
	}

	dataForTemp := &people.DataForTemplate{}
	go func() {
		defer wg.Done()
		dataForTemp, err = parserFile.Parse(transport)
		if err != nil {
			fmt.Printf("error happened: %v", err)
			return
		}

	}()

	wg.Wait()
	dataForTemp.OldFileName = filename
	dataForTemp.FileName = newFileName(filename)
	dataForTemp.Count = len(dataForTemp.Data)
	fmt.Println(dataForTemp.Headers)
	fmt.Println(dataForTemp.Data)

	newFile, err := os.Create(prefix + "newFile.html")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}

	temp, err := template.New("").ParseFiles("./templates/template_for_csv.html")
	temp = template.Must(template.New("template_for_csv.html").ParseFiles("./templates/template_for_csv.html"))
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Println(dataForTemp)
	err = temp.Execute(newFile, dataForTemp)
	if err != nil {
		fmt.Printf("error happened: %v\n", err)
		return
	}

}
