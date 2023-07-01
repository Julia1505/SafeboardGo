package decoder

import (
	"bufio"
	_ "golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io"
)

type ISO1Decoder struct {
	typeDec string
}

func NewISO1Decoder() Decoder {
	return &ISO1Decoder{typeDec: ISO_1}
}

func (d *ISO1Decoder) GetType() string {
	return d.typeDec
}

// файлы PRN CSV кодируются ISO
func (d *ISO1Decoder) Decode(file io.Reader, out chan<- string) {
	decoded := transform.NewReader(bufio.NewReader(file), charmap.ISO8859_1.NewDecoder())
	scanner := bufio.NewScanner(decoded)
	defer close(out)

	for scanner.Scan() {
		record := scanner.Text()
		out <- record
	}
}
