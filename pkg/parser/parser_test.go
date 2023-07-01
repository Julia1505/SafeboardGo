package parser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

func TestPRNParser_Parse(t *testing.T) {
	type TestCase struct {
		inputData  []string
		expectData [][]string
		isError    bool
	}

	tests := []TestCase{
		{inputData: []string{"First name      Address               Postcode Mobile               Limit Birthday",
			"Oliver          Via Archimede, 103-91 2343aa   000 1119381        6000000 19570101",
			"Harry           Leonardo da Vinci 1   4532 AA  010 1118986       10433301 19751203"},
			expectData: [][]string{{"First name", "Address", "Postcode", "Mobile", "Limit", "Birthday"},
				{"Oliver", "Via Archimede, 103-91", "2343aa", "000 1119381", "6000000", "19570101"},
				{"Harry", "Leonardo da Vinci 1", "4532 AA", "010 1118986", "10433301", "19751203"}},
			isError: false,
		},
		{inputData: []string{""}, expectData: [][]string{{""}}, isError: true},
	}

	for caseNum, test := range tests {
		testName := fmt.Sprintf("test:%v", caseNum)

		t.Run(testName, func(t *testing.T) {
			in := make(chan string)
			go func() {
				defer close(in)
				for _, elem := range test.inputData {
					in <- elem
				}
			}()

			par, err := NewParser(PRNFile)
			out := make(chan []string)

			go func() {
				defer close(out)
				err = par.Parse(in, out)
				if err != nil && !test.isError {
					t.Errorf("unexpected error: %v", err)
				}
			}()

			var count int
			for elem := range out {
				assert.Equal(t, test.expectData[count], elem)
				count++
			}
		})
	}
}

func TestCSVParser_Parse(t *testing.T) {
	type TestCase struct {
		inputData  []string
		expectData [][]string
		isError    bool
	}

	tests := []TestCase{
		{inputData: []string{"Name,Address,Postcode,Mobile,Limit,Birthday",
			"\"Oliver, El\",\"Via Archimede, 103-91\",2343aa,000 1119381,6000000,01/01/1999"},
			expectData: [][]string{{"Name", "Address", "Postcode", "Mobile", "Limit", "Birthday"},
				{"Oliver, El", "Via Archimede, 103-91", "2343aa", "000 1119381", "6000000", "01/01/1999"}},
			isError: false},
		{inputData: []string{""}, expectData: [][]string{{""}}, isError: true},
	}

	for caseNum, test := range tests {
		testName := fmt.Sprintf("test:%v", caseNum)

		t.Run(testName, func(t *testing.T) {
			in := make(chan string)
			go func() {
				defer close(in)
				for _, elem := range test.inputData {
					in <- elem
				}
			}()

			par, err := NewParser(CSVFile)
			out := make(chan []string)

			go func() {
				defer close(out)
				err = par.Parse(in, out)
				if err != nil && !test.isError {
					t.Errorf("unexpected error: %v", err)
				}
			}()

			var count int
			for elem := range out {
				assert.Equal(t, test.expectData[count], elem)
				count++
			}
		})
	}

}
