package easy

import "testing"

func Test_isValid(t *testing.T) {

	var (
		tests = []struct {
			str      string
			expected bool
		}{
			{"()[]{}", true},
			{"{[()]}", true},
			{"()[]}", false},
			{"(]", false},
			{"({)[{()}]", false},
			{"{[]}({})[]()", true},
			{"(((([[]])))", false},
			{"", false},
			{"()]]", false},
		}
		funcs = []func(string) bool{
			isValid,
			isValidOne,
			isValidTwo,
		}
	)
	for _, fn := range funcs {

		for _, test := range tests {
			t.Run(test.str, func(t *testing.T) {
				actual := fn(test.str)
				if actual != test.expected {
					t.Errorf("error occured: expected %v but got %v", test.expected, actual)
				}
			})
		}
	}

}
