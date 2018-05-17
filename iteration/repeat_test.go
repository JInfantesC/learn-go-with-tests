package iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if got != expected {
		t.Errorf("expected '%s' but got '%s'", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeated := Repeat("A", 5)
	fmt.Println(repeated)
	// Output: AAAAA
}
