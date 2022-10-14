package parser

import (
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
	res := make([]string, 0, 6)
	word := make([]rune, 0, 10)
	prev := ' '

	runeString := []rune(row)
	for i, sim := range runeString {
		if unicode.IsUpper(sim) {
			index = append(index, i)
		}

		if i == len(runeString)-1 {
			if sim != ' ' || (sim == ' ' && prev != ' ') {
				word = append(word, sim)
				res = append(res, string(word))
			} else if sim == ' ' && prev == ' ' {
				break
			}
		}

		if sim == ' ' && prev == ' ' {
			continue
		}

		if sim == ' ' && (runeString[i+1] == ' ' || unicode.IsUpper(runeString[i+1])) {
			prev = sim
			res = append(res, string(word))
			word = make([]rune, 0, 10)
			continue
		}

		prev = sim
		word = append(word, sim)
	}
	return res, index
}

func parseRecord(row string, index []int) []string {
	res := make([]string, 0, 6)
	word := make([]rune, 0, 10)
	prev := ' '
	var counter int

	runeString := []rune(row)
	for i, sim := range runeString {
		if i == len(runeString)-1 {
			if sim != ' ' || (sim == ' ' && prev != ' ') {
				word = append(word, sim)
				res = append(res, string(word))
			}
			break
		}

		if prev != ' ' && sim == ' ' && (i+1) == index[counter+1] {
			prev = sim
			res = append(res, string(word))
			word = make([]rune, 0, 10)
			counter++
			continue
		}

		if sim == ' ' && prev == ' ' {
			continue
		}

		if sim == ' ' && runeString[i+1] == ' ' {
			prev = sim
			res = append(res, string(word))
			word = make([]rune, 0, 10)
			counter++
			continue
		}

		prev = sim
		word = append(word, sim)
	}

	return res
}

func (p *PRNParser) Parse(in <-chan string, out chan<- []string) error {
	isHeader := true
	var indexHeader []int
	headers := make([]string, 0, 6)
	var cols int

	for record := range in {
		if isHeader {
			headers, indexHeader = parseHeader(record)
			if len(headers) <= 0 {
				return BadFormatFile
			}
			cols = len(headers)
			out <- headers
			isHeader = false
		} else {
			data := parseRecord(record, indexHeader)
			if data == nil {
				continue
			}
			if len(data) != cols {
				return BadFormatFile
			}
			out <- data
		}
	}
	return nil
}
