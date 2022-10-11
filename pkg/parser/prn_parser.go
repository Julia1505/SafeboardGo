package parser

import (
	"github.com/Julia1505/SafeboardGo/pkg/people"
	"unicode"
)

type PRNParser struct {
	typePar string
}

func NewPRNParser() Parser {
	return &PRNParser{typePar: PRNFile}
}

func (p *PRNParser) GetType() string {
	return p.typePar
}

func parseHeader(row string) ([]string, []int) {
	index := make([]int, 0, 6)
	prev := ' '
	res := make([]string, 0, 6)
	word := ""
	for i, sim := range row {
		if unicode.IsUpper(sim) {
			index = append(index, i)
		}

		if sim == ' ' && prev == ' ' {
			continue
		}

		if sim == ' ' && i == len(row) - 1 {
			res = append(res, word)
			break
		}

		if sim == ' ' && row[i+1] == ' ' {
			res = append(res, word)
			word = ""
			prev = sim
			continue
		}

		word += string(sim)
		prev = sim
	}
	res = append(res, word)
	return res, index
}

func parseRecord(row string) []string {

}

func (p *PRNParser) Parse(in <-chan string) (*people.DataForTemplate, error) {

	isHeader := true
	var indexHeader []int
	dataForTemp := &people.DataForTemplate{Data: make([]people.PeopleData,0,5), Headers: make([]string,0,6)}

	for record := range in {
		if isHeader {
			dataForTemp.Headers, indexHeader = parseHeader(record)
			if len(dataForTemp.Headers) != 6 {
				return nil,
			}
			isHeader = false
		} else {

		}
	}
}
