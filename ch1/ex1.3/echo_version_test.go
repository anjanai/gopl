package echo

import (
	"testing"
)

var testCases = []struct {
	args    []string
	expects string
}{
	{[]string{"test1", "1", "2", "3"}, "1 2 3"},
	{[]string{"test2", "a", "b", "c"}, "a b c"},
}

func TestEcho1(t *testing.T) {
	for _, c := range testCases {
		got := Echo1(c.args)
		if got != c.expects {
			t.Fatalf("FAIL: Echo1\nExpected %s, got %s", c.expects, got)
		}
	}
}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			Echo1(c.args)
		}

	}
}

func TestEcho2(t *testing.T) {
	for _, c := range testCases {
		got := Echo2(c.args)
		if got != c.expects {
			t.Fatalf("FAIL: Echo2\nExpected %s, got %s", c.expects, got)
		}
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			Echo2(c.args)
		}

	}
}

func TestEcho3(t *testing.T) {
	for _, c := range testCases {
		got := Echo3(c.args)
		if got != c.expects {
			t.Fatalf("FAIL: Echo2\nExpected %s, got %s", c.expects, got)
		}
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			Echo3(c.args)
		}

	}
}
