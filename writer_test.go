package writer

import "testing"

func testInput(t *testing.T) {
	want := ""
	if got := Input(); got != want {
		t.ErrorF("Input() = %q, want %q ", got, want)
	}
}
