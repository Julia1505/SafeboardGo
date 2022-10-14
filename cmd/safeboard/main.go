package main

import (
	"fmt"
	"github.com/Julia1505/SafeboardGo/pkg/decoder"
	"github.com/Julia1505/SafeboardGo/pkg/file"
	"github.com/Julia1505/SafeboardGo/pkg/parser"
	"github.com/Julia1505/SafeboardGo/pkg/people"
	"os"
)

func main() {
	filename := "data.csv" // этот файл конвертируется

	prefix := "./data/"
	//var filename string
	//fmt.Scan(&filename)

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

	transport := make(chan string)
	go func() {
		fileDecoder.Decode(file, transport)
	}()

	parserFile, err := parser.NewParser(typeFile)
	if err != nil {
		fmt.Printf("error happened: %v", err)
		return
	}

	parseCh := make(chan []string)
	go func() {
		defer close(parseCh)
		err = parserFile.Parse(transport, parseCh)
		if err != nil {
			fmt.Printf("error happened: %v", err)
			return
		}
	}()

	dataForTemp := people.NewTemplate(filename)
	isHeader := true
	for elem := range parseCh {
		if isHeader {
			dataForTemp.Headers = elem
			dataForTemp.Count = len(elem)
			isHeader = false
		} else {
			dataForTemp.Data = append(dataForTemp.Data, people.PeopleData{
				Name:     elem[0],
				Address:  elem[1],
				Postcode: elem[2],
				Mobile:   elem[3],
				Limit:    elem[4],
				Birthday: elem[5],
			})
		}
	}
	//fmt.Println(dataForTemp.Data)

	creatorHTML := &people.CreatorHTML{TemplateModel: dataForTemp}
	err = creatorHTML.Create(prefix, "template.html", dataForTemp.FileName)
	if err != nil {
		fmt.Printf("error happend: %v", err)
		return
	}
}
