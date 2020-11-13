package writer

import "testing"

func testInput(t *testing.T) {
	want := ""
	if got := Input(); got != want {
		t.Errorf("Input() = %q, want %q ", got, want)
	}
}
