package sum

import "testing"

func TestSum(t *testing.T) {
	actual := Sum(1, 2)
	if actual != 3 {
		t.Errorf("expected 3, but got %d", actual)
	}
}
