package general

import "testing"

func Test_TwoEqualsTwo(t *testing.T) {
	if !TwoEqualsTwo() {
		t.Error("Two should equal two!!")
	}

}

func TwoEqualsTwo() bool {
	return 2 == 2
}
