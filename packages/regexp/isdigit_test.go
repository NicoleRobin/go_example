package main

import "testing"

type TestCase struct {
	user   string
	result bool
}

func TestIsDigit(t *testing.T) {
	testCases := []TestCase{
		{"123456@xunlei.net", true},
		{"zjw@xunlei.net", false},
		{"ygl@xunlei.net", true},
	}

	for _, testCase := range testCases {
		if IsDigit(testCase.user) != testCase.result {
			t.Error("failed")
		}
	}
}
