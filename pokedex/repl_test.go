package main
import ("testing"
"fmt")


func TestCleanInput(t *testing.T) {
    cases := []struct {
	input    string
	expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello Pickachu world  ",
			expected: []string{"hello", "pickachu", "world"},
		},
		{
			input:    "Los pollos HeRmAnOs",
			expected: []string{"los","pollos", "hermanos"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The length don't match")
			fmt.Printf("\nActual: %s\nExpected: %s\n", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if len(word) != len(expectedWord){
				t.Errorf("The word don't match expectedWord")
				fmt.Printf("\nActual: %s\nExpected: %s\n", word, expectedWord)
				t.Fail()
			}
		}
	}
}