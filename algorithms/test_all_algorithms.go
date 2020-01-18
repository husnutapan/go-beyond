package main

import "fmt"

type testcase struct {
	input       string
	expectation bool
}

type testcase2 struct {
	input1      string
	input2      string
	expectation bool
}

func main() {
	//for string has same chars
	test1 := testcase2{
		input1:      "husnu",
		input2:      "usnhu",
		expectation: true,
	}
	test2 := testcase2{
		input1:      "abscsd",
		input2:      "abssad",
		expectation: true,
	}
	cases := []testcase2{test1, test2}

	for _, c := range cases {
		result := controlTwoString(c.input1, c.input2)
		if result != c.expectation {
			fmt.Printf("Input1 value %s,Input2 value %s. Expectation %t", c.input1, c.input2, !c.expectation)
		}
	}

	//for unique all chars
	testcase := []testcase{
		{"husnu", false},
		{"xyz", true},
		{"xyzzzz", true},
	}
	for _, test := range testcase {
		result := isUniqueCharacter(test.input)
		if result != test.expectation {
			fmt.Printf("Input value %s. Expectation %t", test.input, !test.expectation)
		}
	}
}
