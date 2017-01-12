package eval

import (
	"fmt"
	"testing"
)

func TestPrecedence(t *testing.T) {
	fmt.Printf("%s\n", MathPrecedences.Tree().TreeString())
}
