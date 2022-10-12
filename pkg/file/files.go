package file

import "errors"

var (
	NoExtension = errors.New("File doesn't have extension")
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func GetFileExtension(filename string) (string, error) {
	runeString := []rune(filename)
	l := len(runeString)
	ext := ""
	isComma := false
	for i := l - 1; i >= 0; i-- {
		if runeString[i] == '.' {
			isComma = true
			break
		}
		ext += string(runeString[i])
	}
	if ext == "" || isComma == false {
		return "", NoExtension
	}
	ext = reverse(ext)
	return ext, nil
}

func NewFileExtension(filename, ext string) string {
	runeString := []rune(filename)
	l := len(runeString)
	var ind int
	for i := l - 1; i >= 0; i-- {
		if runeString[i] == '.' {
			ind = i
			break
		}
	}

	runeString = runeString[:ind+1]
	res := string(runeString) + ext
	return res
}
