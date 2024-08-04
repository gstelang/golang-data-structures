package common

import (
	"testing"
)

// go test -v ./common/common.go ./common/common_test.go
func TestIfPresent(t *testing.T) {

	arr := []string{"abc", "xyz"}
	includes := Includes[string](arr, "abc")

	if !includes {
		t.Errorf("Does not abc when it is present in the string collection")
	}
}

func TestIfAbsent(t *testing.T) {

	arr := []string{"abc", "xyz"}
	includes := Includes[string](arr, "mno")

	if includes {
		t.Errorf("should not include mno when it is NOT present in the string collection")
	}
}
