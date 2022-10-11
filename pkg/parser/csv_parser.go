package parser

import (
	"github.com/Julia1505/SafeboardGo/pkg/people"
)

type CSVParser struct {
	typePar string
}

func NewCSVParser() Parser {
	return &CSVParser{typePar: CSVFile}
}

func (p *CSVParser) GetType() string {
	return p.typePar
}

func splitByColumns(row string) []string {
	res := make([]string, 0, 6)
	isInStringField := false
	word := ""

	for _, sim := range row {
		if sim == '"' {
			if !isInStringField {
				isInStringField = true
				continue
			} else {
				isInStringField = false
				continue
			}
		}

		if sim == ',' && !isInStringField {
			res = append(res, word)
			word = ""
			continue
		}

		word += string(sim)
	}
	res = append(res, word)

	return res
}

func (p *CSVParser) Parse(in <-chan string) (*people.DataForTemplate, error) {
	dataTemp := &people.DataForTemplate{Headers: make([]string, 0, 6), Data: make([]people.PeopleData, 0, 5)}
	isHeader := true
	for record := range in {
		rows := splitByColumns(record)
		if len(rows) != 6 {
			return nil, BadFormatFile
		}

		if isHeader {
			dataTemp.Headers = rows
			isHeader = false
		} else {
			newPeople := &people.PeopleData{
				Name:     rows[0],
				Address:  rows[1],
				Postcode: rows[2],
				Mobile:   rows[3],
				Limit:    rows[4],
				Birthday: rows[5],
			}
			dataTemp.Data = append(dataTemp.Data, *newPeople)
		}
	}
	return dataTemp, nil
}
