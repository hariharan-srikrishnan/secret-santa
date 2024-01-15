package permute

import (
	"testing"
)

func TestGetDerangements(t *testing.T) {
	a := getValidDerangements(8)
	t.Logf("Number of derangements: %d", len(a))
	if (len(a) != 14833) {
		t.Fatalf("error in derangement calculation %d", len(a))
	} 
}
