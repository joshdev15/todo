package todo

import (
	"reflect"
	"testing"
)

func TestText(t *testing.T) {
	initialString := "This is a test"
	text := Text(initialString)

	value := reflect.TypeOf(text)

	if value.Kind() != reflect.String {
		t.Errorf("Expected %v, got %v", reflect.String, value.Kind())
	}

	if text.ToString() != initialString {
		t.Errorf("Expected %v, got %v", initialString, value.Kind())
	}

	if text.Elipsis() == "This is a..." {
		t.Errorf("Expected %v, got %v", initialString, value.Kind())
	}
}
