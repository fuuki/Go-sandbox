package checker

import (
	"fmt"
	"runtime"
	"testing"
)

func callerInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if ok {
		// fname := filepath.Dir(file)
		return fmt.Sprintf("%s:%d", file, line)
	}
	return ""
}

// Equal は値が等しいか確認する
func Equal(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Errorf("\n[Unexpected] %s: actual: %v, expected: %v", callerInfo(2), actual, expected)
	}
}
