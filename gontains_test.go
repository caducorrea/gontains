package gontains

import (
	"reflect"
	"testing"
)

type TestData struct {
	slc      []string
	str      string
	result   bool
	hasError bool
}

func TestAll(t *testing.T) {

	data := []TestData{
		{[]string{}, "test", false, false},
		{[]string{"test"}, "test", true, false},
	}

	for _, dt := range data {
		result := Gontains(dt.slc, dt.str)

		if !reflect.DeepEqual(dt.result, result) {
			t.Errorf("gontains() with args %v - %v : FAILED, expected %v but got value %v", dt.slc, dt.str, dt.result, result)
		}
	}
}
