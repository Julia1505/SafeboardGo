package parser

import (
	"errors"
	"github.com/Julia1505/SafeboardGo/pkg/people"
)

var (
	CSVFile = "csv"
	PRNFile = "prn"
)

var (
	TypeNotExist  = errors.New("This type doesn't exist")
	BadFormatFile = errors.New("Bad format in file")
)

type Parser interface {
	GetType() string
	Parse(in <-chan string) ([]string, []people.PeopleData, error)
}

func NewParser(typeParser string) (Parser, error) {
	switch typeParser {
	case CSVFile:
		return NewCSVParser(), nil
	case PRNFile:
		return NewPRNParser(), nil
	default:
		return nil, TypeNotExist
	}
}
