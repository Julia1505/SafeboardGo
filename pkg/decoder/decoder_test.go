package decoder

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewDecoder(t *testing.T) {
	type TestCase struct {
		charset string
		typeDec string
		isError bool
	}

	tests := []TestCase{
		{charset: ISO_1, typeDec: ISO_1, isError: false},
		{charset: "utf-8", typeDec: "", isError: true},
	}

	for caseNum, test := range tests {
		testName := fmt.Sprintf("file:%v", test.charset)

		t.Run(testName, func(t *testing.T) {
			dec, err := NewDecoder(test.charset)

			if test.isError && err == nil {
				t.Errorf("[%d] expected error, got nil", caseNum)
			}

			if !test.isError && err != nil {
				t.Errorf("[%d] unexpected error: %v", caseNum, err)
			}

			if test.isError == false && dec.GetType() != test.typeDec {
				t.Errorf("[%d] expected: %v result: %v", caseNum, test.typeDec, dec.GetType())
			}
		})
	}

}

func TestGetFileExtension(t *testing.T) {
	type TestCase struct {
		filename string
		out      []string
	}

	tests := []TestCase{
		{filename: "data.csv", out: []string{"", ""}},
	}

	for caseNum, test := range tests {
		testName := fmt.Sprintf("file:%v", test.filename)

		t.Run(testName, func(t *testing.T) {
			file, err := os.Open(test.filename)
			if err != nil {
				panic("can't open test file")
			}

			if test.isErr && err == nil {
				t.Errorf("[%d] expected error, got nil", caseNum)
			}

			if !test.isErr && err != nil {
				t.Errorf("[%d] unexpected error: %v", caseNum, err)
			}

			assert.Equal(t, test.out, res, "[%d] wrong results: got %+v, expected %+v", caseNum, res, test.out)
		})
	}
}
