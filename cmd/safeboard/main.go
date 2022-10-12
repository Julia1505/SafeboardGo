package main

import (
	"fmt"
	"github.com/Julia1505/SafeboardGo/pkg/decoder"
	"github.com/Julia1505/SafeboardGo/pkg/file"
	"github.com/Julia1505/SafeboardGo/pkg/parser"
	"github.com/Julia1505/SafeboardGo/pkg/people"
	"os"
	"sync"
)

var (
	CSVFile = "csv"
	PRNFile = "prn"
)

func main() {
	prefix := "./data/"
	filename := "data.csv"
	typeFile, err := file.GetFileExtension(filename)
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

	fileDecoder, err := decoder.NewDecoder("ISO8859_1")
	if err != nil {
		fmt.Printf("error happened: %v", err)
		return
	}

	wg := &sync.WaitGroup{}
	transport := make(chan string)
	wg.Add(2)

	go func() {
		defer wg.Done()
		fileDecoder.Decode(file, transport)
	}()

	parserFile, err := parser.NewParser(typeFile)
	if err != nil {
		fmt.Printf("error happened: %v", err)
		return
	}

	var headers []string
	var data []people.PeopleData
	go func() {
		defer wg.Done()
		headers, data, err = parserFile.Parse(transport)
		if err != nil {
			fmt.Printf("error happened: %v", err)
			return
		}
	}()

	wg.Wait()

	dataForTemp := people.NewTemplate(filename)
	dataForTemp.Count = len(headers)
	dataForTemp.Data = data
	dataForTemp.Headers = headers

	creatorHTML := &people.CreatorHTML{TemplateModel: dataForTemp}
	err = creatorHTML.Create(prefix, "template_for_csv.html", dataForTemp.FileName)
	if err != nil {
		fmt.Printf("error happend: %v", err)
		return
	}
}
