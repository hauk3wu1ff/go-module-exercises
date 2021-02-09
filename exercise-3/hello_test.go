package hello

import "testing"

func TestHello(t *testing.T) {
	want := "I can eat glass and it doesn't hurt me."
	if got := Hello(); want != got {
		t.Errorf("Hello() = %q, want = %q", got, want)
	}
}
func TestOpt(t *testing.T) {
	want := "If a program is too slow, it must have a loop."
	if got := Opt(); want != got {
		t.Errorf("Opt() = %q, want = %q", got, want)
	}
}
