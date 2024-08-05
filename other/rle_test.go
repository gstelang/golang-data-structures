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

func TestExceedSameCharLimit(t *testing.T) {
	output := rle("aaaaaaaaaa")
	if output != "9a1a" {
		t.Errorf("failed to accomodate max char limit")
	}
}

func TestDecode(t *testing.T) {
	output := decodeRle("1a1b1c1d1e")
	if output != "abcde" {
		t.Errorf("Unable to decode")
	}
}

//	go test -v ./other/rle.go ./other/rle_test.go -fuzz=Fuzz
//
// TODO: need to add more checks with respect allowable chars for this to work correctly.
func FuzzRle(f *testing.F) {
	// Seed corpus
	testcases := []string{"", "a", "aa", "aaa", "aaaaaaaaaa", "abcde", "b", "aaaaa", ""}
	for _, tc := range testcases {
		// Provide seed inputs
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		encoded := rle(orig)
		decoded := decodeRle(encoded)

		if decoded != orig {
			t.Errorf("rle(%q) = %q, decode(%q) = %q", orig, encoded, encoded, decoded)
		}

	})
}
