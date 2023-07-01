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

func TestISODecode(t *testing.T) {
	type TestCase struct {
		filename string
		out      []string
	}

	tests := []TestCase{
		{filename: "testdata/data.csv", out: []string{"Name,Address,Postcode,Mobile,Limit,Birthday", "\"Oliver, El\",\"Via Archimede, 103-91\",2343aa,000 1119381,6000000,01/01/1999", "\"Harry\",Leonardo da Vinci 1,4532 AA,010 1118986,343434,31/12/1965", "\"Jack\",\"Via Rocco Chinnici 4d\",3423 ba,0313-111475,22,05/04/1984", "\"Noah\",\"Via Giannetti, 4-32\",2340 CC,28932222,434,03/10/1964", "\"Charlie\",\"Via Aldo Moro, 7\",3209 DD,30-34563332,343.8,04/10/1954", "\"Mia\",\"Via Due Giugno, 12-1\",4220 EE,43433344329,6343.6,10/08/1980", "\"Lilly\",Arcisstraße 21,12343,+44 728 343434,34342.3,20/10/1997", ""}},
		{filename: "testdata/empty_data.csv", out: []string{"", ""}},
	}

	for caseNum, test := range tests {
		testName := fmt.Sprintf("file:%v", test.filename)

		t.Run(testName, func(t *testing.T) {
			decoder := &ISO1Decoder{}
			file, err := os.Open(test.filename)
			if err != nil {
				panic("can't open test file")
			}
			defer file.Close()

			testChan := make(chan string)
			go func() {
				decoder.Decode(file, testChan)
			}()

			var count int
			for res := range testChan {
				assert.Equal(t, test.out[count], res)
				count++
			}

			if count+1 != len(test.out) {
				t.Errorf("[%d] expected_n: %v result_n: %v", caseNum, len(test.out), count)
			}

		})
	}
}
