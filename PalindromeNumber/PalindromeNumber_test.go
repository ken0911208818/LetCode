package PalindromeNumber

import (
	"reflect"
	"testing"
)

type testData struct {
	Input  int
	Output bool
}

func generateTestData() []testData {
	data := []testData{
		testData{Input: 121, Output: true},
		{Input: -121, Output: true},
		{Input: 10, Output: false},
		{Input: -101, Output: true},
	}
	return data
}
func Test_reverse1(t *testing.T) {
	data := generateTestData()
	for _, v := range data {
		actual := reverse1(v.Input)
		expected := v.Output
		if reflect.DeepEqual(expected, actual) {
			t.Log(true)
		} else {
			t.Error("not equal!!")
			t.Error("actual= ", actual)
			t.Error("expected= ", expected)
		}
	}
}
