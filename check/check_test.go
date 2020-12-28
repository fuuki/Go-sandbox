package checker

import (
	"testing"
)

func TestEqual(t *testing.T) {
	Equal(t, "sample-text", "sample-text")
	Equal(t, 123, 123)
	// checker(t, "actual-value", "expected-value") // FAIL
}
