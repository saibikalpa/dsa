package easy

import "testing"

func Test_isAnagram(t *testing.T) {
	funcs := []func(string, string) bool{
		isAnagram,
		isAnagramOne,
	}

	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}

	for _, fn := range funcs {
		for _, test := range tests {
			t.Run(test.s, func(t *testing.T) {
				actual := fn(test.s, test.t)
				if actual != test.want {
					t.Errorf("wanted %v, but got %v", test.want, actual)
				}
			})
		}
	}

}
