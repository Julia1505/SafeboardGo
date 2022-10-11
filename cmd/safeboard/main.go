package main

import (
	"fmt"
	"github.com/Julia1505/SafeboardGo/pkg/decoder"
	_ "github.com/Julia1505/SafeboardGo/pkg/people"
	"os"
	"sync"
)

func main() {
	prefix := "./data/"
	filename := prefix + "data.prn"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	defer file.Close()

	wg := &sync.WaitGroup{}

	transport := make(chan string)
	wg.Add(1)
	fileDecoder, err := decoder.NewDecoder("ISO8859_1")
	if err != nil {
		fmt.Printf("error happened: %v", err)
		return
	}

	go func() {
		defer wg.Done()
		fileDecoder.Decode(file, transport)
	}()

	//parserFile, err := parser.NewParser("csv")
	//if err != nil {
	//	fmt.Printf("error happened: %v", err)
	//	return
	//}
	//
	//var dataForTemp people.DataForTemplate
	//go func() {
	//	defer wg.Done()
	//	dataForTemp, err = parserFile.Parse(transport)
	//	if err != nil {
	//		fmt.Printf("error happened: %v", err)
	//		return
	//	}
	//
	//}()
	//fmt.Println(dataForTemp)
	for elem := range transport {
		fmt.Println(elem)
	}

	wg.Wait()

	//newFile, err := os.Create(prefix + "newFile.html")
	//if err != nil {
	//	fmt.Printf("error:%v\n", err)
	//	return
	//}
	//
	//temp, err := template.New("").ParseFiles("./templates/template_for_csv.html")
	//temp = template.Must(template.New("template_for_csv.html").ParseFiles("./templates/template_for_csv.html"))
	//if err != nil {
	//	fmt.Printf("error: %v\n", err)
	//	return
	//}
	////
	////fmt.Println(dataForTemp)
	//err = temp.Execute(newFile, dataForTemp)
	//if err != nil {
	//	fmt.Printf("error: %v\n", err)
	//	return
	//}

}
