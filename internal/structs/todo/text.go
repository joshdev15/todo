package todo

import "fmt"

type Text string

func (t *Text) ToString() string {
	return string(*t)
}

func (t *Text) Elipsis() string {
	finalString := string(*t)
	maxLength := int(10)
	runeText := []rune(finalString)

	if len(runeText) > maxLength {
		finalString = fmt.Sprintf("%v...", string(runeText[0:maxLength]))
	}

	return finalString
}
