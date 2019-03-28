package instructions

import (
	"fmt"
	"testing"
)

func TestTurnLeft(t *testing.T) {
	fmt.Sprintf("turn left")
}

func BenchmarkTurnLeft(b *testing.B) {
	fmt.Sprintf("benchmark turn left")
}
