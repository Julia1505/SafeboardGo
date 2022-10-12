package parser

import (
	"fmt"
	"testing"
)

func TestNewParser(t *testing.T) {
	type TestCase struct {
		extension string
		typeDec   string
		isError   bool
	}

	tests := []TestCase{
		{extension: CSVFile, typeDec: CSVFile, isError: false},
		{extension: PRNFile, typeDec: PRNFile, isError: false},
		{extension: "go", isError: true},
	}

	for caseNum, test := range tests {
		testName := fmt.Sprintf("file:%v", test.extension)

		t.Run(testName, func(t *testing.T) {
			par, err := NewParser(test.extension)

			if test.isError && err == nil {
				t.Errorf("[%d] expected error, got nil", caseNum)
			}

			if !test.isError && err != nil {
				t.Errorf("[%d] unexpected error: %v", caseNum, err)
			}

			if test.isError == false && par.GetType() != test.typeDec {
				t.Errorf("[%d] expected: %v result: %v", caseNum, test.typeDec, par.GetType())
			}
		})
	}
}
