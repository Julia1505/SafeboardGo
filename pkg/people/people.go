package people

type PeopleData struct {
	//Id       uint
	Name     string `csv:"Name"`
	Address  string `csv:"Address"`
	Postcode string `csv:"Postcode"`
	Mobile   string `csv:"Mobile"`
	Limit    string `csv:"Limit"`
	Birthday string `csv:"Birthday"`
}

type DataForTemplate struct {
	OldFileName string
	FileName    string
	Headers     []string
	Data        []PeopleData
	Count       int
}
