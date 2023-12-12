package string_type

import "testing"

func TestStringBuilder(t *testing.T) {
	tests := []struct {
		testName string
		slice    []string
		str      string
	}{
		{
			"test1",
			[]string{"Hello", "World"},
			"HelloWorld",
		}, {
			"test2",
			[]string{"Hello", "World"},
			"HelloWorld",
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			str := stringBuilder(test.slice)
			if str != test.str {
				t.Errorf("error")
			}
		})
	}
}
