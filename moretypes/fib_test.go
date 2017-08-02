package moretypes

import (
	"testing"
	"github.com/ashigirl96/moretypes"
)

func TestFib(t *testing.T) {
	actual := moretypes.Fib(9)
	expected := 55
	if actual != expected {
		t.Errorf("Error: Got %v\n want %v", actual, expected)
	}
	actual = moretypes.Fib(-5)
	expected = 1
	if actual != expected {
		t.Errorf("Error: Got %v\n want %v", actual, expected)
	}
}
