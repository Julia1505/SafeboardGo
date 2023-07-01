package parser

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

func (p *CSVParser) Parse(in <-chan string, out chan<- []string) error {
	headers := make([]string, 0, 6)
	isHeader := true
	var cols int
	for record := range in {
		rows := splitByColumns(record)
		if isHeader {
			cols = len(rows)
		}

		if len(rows) != cols {
			return BadFormatFile
		}

		if isHeader {
			headers = rows
			isHeader = false
			out <- headers
		} else {
			out <- rows
		}
	}
	return nil
}
