package permute

import (
	"testing"
)

func Test_getDerangement(t *testing.T) {
	a := GetDerangements(8)
	if (len(a) != 14833) {
		t.Fatalf("error in derangement calculation %d", len(a))
	} 
	t.Fatalf("")
}
