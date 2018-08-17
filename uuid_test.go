package uuid

import (
	"testing"
)

func TestUUIDLength(t *testing.T) {
	u, err := UUID()
	if err != nil {
		t.Fatal(err)
	}
	if len(u) != 36 {
		t.Errorf("want length 36, got length %v", len(u))
	}
}

// TODO: calculate the natural failure rate of this test
func TestUUIDRandomness(t *testing.T) {
	hist := make(map[rune]int64)
	var n int = 1e6
	for i := 0; i < n; i++ {
		u, err := UUID()
		if err != nil {
			t.Fatal(err)
		}
		for _, r := range u {
			hist[r]++
		}
	}
	zeroes := hist['0']
	total := int64(n) * 32
	t.Logf("%v zeroes of %v hex characters (pecentage: %v%%)", zeroes, total, float64(zeroes*100)/float64(total))
	if zeroes > total/4 {
		t.Errorf("expected less than one quarter of the generated hex digits for large sample to be zero, got too many zeroes: %v", zeroes)
	}
}
