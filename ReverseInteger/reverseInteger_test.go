package ReverseInteger

import (
	"reflect"
	"testing"
)

type testData struct {
	Input  int
	Output int
}

func generateTestData() []testData {
	data := []testData{
		testData{Input: 123, Output: 321},
		{Input: -123, Output: -321},
		{Input: 120, Output: 21},
		{Input: 0, Output: 0},
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

func Test_reverse2(t *testing.T) {
	data := generateTestData()
	for _, v := range data {
		actual := reverse2(v.Input)
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

func Test_reverse3(t *testing.T) {
	data := generateTestData()
	for _, v := range data {
		actual := reverse3(v.Input)
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
