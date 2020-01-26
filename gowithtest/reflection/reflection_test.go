package main

import (
	"testing"
)

type testCases struct {
	Input  interface{}
	Expect []string
}

func test(t *testing.T) {

	var cases []testCases

	case1 := struct {
		Input  interface{}
		Expect []string
	}{Input: "Husnu", Expect: []string{"Husnu"}}

	case2 := struct {
		Input  interface{}
		Expect []string
	}{Input: Person{"Husnu", 19}, Expect: []string{"Husnu"}}

	cases = append(cases, case1)
	cases = append(cases, case2)

	for _, test := range cases {
		var got []string
		reflection(test.Input, func(input string) {
			got = append(got, input)
		})
	}

}

type Person struct {
	Name string
	Age  int
}
