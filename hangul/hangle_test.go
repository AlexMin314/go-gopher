package hangul

import (
	"fmt"
	"testing"
)

func TestHasConstantSuffixs(t *testing.T) {
	got := HasConstantSuffixs("하핳")
	want := true

	if got != want {
		t.Errorf("test failed")
		// t.Errorf("got %q want %q", got, want)
	}
}

func ExampleHasConstantSuffixs() {
	fmt.Println(HasConstantSuffixs("하핳"))
	// Output:
	// true
}
