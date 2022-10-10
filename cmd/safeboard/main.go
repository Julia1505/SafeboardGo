package main

import (
	"encoding/csv"
	"fmt"
	_ "github.com/Julia1505/SafeboardGo/pkg/people"
	"os"
	"sync"
)

//func MakePeopleData(record []string) (people.PeopleData, error) {
//	if len(record) !=
//}

func main() {
	prefix := "./data/"
	filename := prefix + "data.csv"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	//var data []people.PeopleData

	transport := make(chan []string)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer close(transport)
		defer wg.Done()

		for {
			record, err := reader.Read()
			if err != nil {
				break
			}
			fmt.Println(record)
			//transport <- record
		}
	}()

	go func() {
		defer wg.Done()
		for record := range transport {

		}
	}()

	wg.Wait()
}
