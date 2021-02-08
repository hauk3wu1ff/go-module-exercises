package hello

import "testing"

func TestHello(t *testing.T) {
	want := "I can eat glass and it doesn't hurt me."
	if got := Hello(); want != got {
		t.Errorf("Hello() = %q, want = %q", got, want)
	}
}
