package other

import (
	"testing"
)

// go test -v ./other/rle.go ./other/rle_test.go
func TestAllUnique(t *testing.T) {
	output := rle("abcde")
	if output != "1a1b1c1d1e" {
		t.Errorf("failed for all unique characters")
	}
}

func TestAllSame(t *testing.T) {
	output := rle("aaaaa")
	if output != "5a" {
		t.Errorf("failed to match all same chars")
	}
}
