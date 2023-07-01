package people

import (
	"github.com/Julia1505/SafeboardGo/pkg/file"
	"html/template"
	"os"
)

type PeopleData struct {
	Name     string
	Address  string
	Postcode string
	Mobile   string
	Limit    string
	Birthday string
}

type DataForTemplate struct {
	OldFileName string
	FileName    string
	Headers     []string
	Data        []PeopleData
	Count       int
}

func NewTemplate(filename string) *DataForTemplate {
	return &DataForTemplate{
		OldFileName: filename,
		FileName:    file.NewFileExtension(filename, "html"),
		Data:        make([]PeopleData, 0, 5),
	}
}

type CreatorHTML struct {
	TemplateModel interface{}
}

func (b *CreatorHTML) Create(prefix, tempFile, newFile string) error {
	file, err := os.Create(prefix + newFile)
	defer file.Close()
	if err != nil {
		return err
	}

	temp := template.Must(template.New(tempFile).ParseFiles("./templates/" + tempFile))
	err = temp.Execute(file, b.TemplateModel)
	if err != nil {
		return err
	}
	return nil
}
