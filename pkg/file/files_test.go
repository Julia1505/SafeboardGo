package file

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFileExtension(t *testing.T) {
	type TestCase struct {
		in    string
		out   string
		isErr bool
	}

	tests := []TestCase{
		{in: "data.csv", out: "csv", isErr: false},
		{in: "newFile.dfsa.html", out: "html", isErr: false},
		{in: "gg.", out: "", isErr: true},
		{in: "fdsafads", out: "", isErr: true},
	}

	for caseNum, test := range tests {
		testName := fmt.Sprintf("in:%v out:%v error:%v", test.in, test.out, test.isErr)

		t.Run(testName, func(t *testing.T) {
			res, err := GetFileExtension(test.in)

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

func TestNewFileExtension(t *testing.T) {
	type TestCase struct {
		in  string
		ex  string
		out string
	}

	tests := []TestCase{
		{in: "data.csv", ex: "html", out: "data.html"},
		{in: "newFile.dfsa.html", ex: "go", out: "newFile.dfsa.go"},
	}

	for caseNum, test := range tests {
		testName := fmt.Sprintf("in:%v ex:%v out:%v", test.in, test.ex, test.out)

		t.Run(testName, func(t *testing.T) {
			res := NewFileExtension(test.in, test.ex)
			assert.Equal(t, test.out, res, "[%d] wrong results: got %+v, expected %+v", caseNum, res, test.out)
		})
	}
}
