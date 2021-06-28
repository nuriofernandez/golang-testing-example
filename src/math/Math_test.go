package math

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(1, 2)
	if total != 3 {
		t.Errorf("Sum operation failed got '%d' instead of '%d'.", total, 10)
	}
}
