package gap

import (
	"testing"
)

type input struct {
	operation   string
	inputString string
	lposition   int
	rposiiton   int
	isFront     bool
}
type tests struct {
	inputs          []input
	expectedResults []string
}

func atoi(s string) int {
	var result int
	var negative bool
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0
		}
		result = result*10 + int(c-'0')
	}
	if negative {
		result *= -1
	}
	return result
}
func TestPrompts(t *testing.T) {
	tests := new(tests)
	tests.expectedResults = []string{"hello",
		"hello beautiful world",
		"hello our beautiful world",
		"hello our eautiful world",
		"hello oureautiful world",
		"hello ourautiful world",
		"hellorautiful world",
		"hellotiful world",
		"24",
		"40"}
	tests.inputs = []input{
		{"insert", " beautiful world", 5, 0, false},
		{"insert", "our ", 6, 0, false},
		{"delete", "", 0, 0, false},
		{"backspace", "", 0, 0, false},
		{"delete", "", 0, 0, false},
		{"select delete", "", 5, 8, true},
		{"select delete", "", 5, 8, false},
		{"gap length", "", 0, 0, false},
		{"buffer length", "", 0, 0, false},
	}
	gap := New(tests.expectedResults[0])
	t.Run("New gap buffer", func(t *testing.T) {
		if got := gap.GetString(); got != tests.expectedResults[0] {
			t.Errorf("TEST %v : Expected : %v\nGot : %v\n", 0, tests.expectedResults[0], got)
		}
	})
	for i := 0; i < len(tests.inputs); i++ {
		t.Run(tests.inputs[i].operation, func(t *testing.T) {
			switch tests.inputs[i].operation {
			case "insert":
				if got := gap.Insert([]rune(tests.inputs[i].inputString), tests.inputs[i].lposition).GetString(); got != tests.expectedResults[i+1] {
					t.Errorf("TEST %v : Expected : %v\nGot : %v\n", i+1, tests.expectedResults[i+1], got)
				}
			case "delete":
				if got := gap.Delete().GetString(); got != tests.expectedResults[i+1] {
					t.Errorf("TEST %v : Expected : %v\nGot : %v\n", i+1, tests.expectedResults[i+1], got)
				}
			case "backspace":
				if got := gap.Backspace().GetString(); got != tests.expectedResults[i+1] {
					t.Errorf("TEST %v : Expected : %v\nGot : %v\n", i+1, tests.expectedResults[i+1], got)
				}
			case "select delete":
				if got := gap.SelectDelete(tests.inputs[i].lposition, tests.inputs[i].rposiiton, tests.inputs[i].isFront).GetString(); got != tests.expectedResults[i+1] {
					t.Errorf("TEST %v : Expected : %v\nGot : %v\n", i+1, tests.expectedResults[i+1], got)
				}
			case "gap length":
				if got := gap.GapLength(); got != atoi(tests.expectedResults[i+1]) {
					t.Errorf("TEST %v : Expected : %v\nGot : %v\n", i+1, tests.expectedResults[i+1], got)
				}
			case "buffer length":
				if got := gap.BufferLength(); got != atoi(tests.expectedResults[i+1]) {
					t.Errorf("TEST %v : Expected : %v\nGot : %v\n", i+1, tests.expectedResults[i+1], got)
				}
			}
		})
	}
}
