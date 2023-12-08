package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("got '%d but want %d", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(3, 4)
	fmt.Println(sum)
	// Output: 7
}
