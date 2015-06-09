package cms

import (
	"fmt"
	"math/big"
	"testing"
)

func TestSizedImagePath(t *testing.T) {
	i := SizedImage{Key: "Portrait2013", Width: 250, Height: 250}
	if i.GetPath() != "Portrait2013/s250x250" {
		t.Errorf("Expected '%s', got '%s'", "Portrait2013/s250x250", i.GetPath())
	}
}

func Test16By9(t *testing.T) {
	x := big.NewInt(32)
	y := big.NewInt(18)
	z := big.NewInt(0).GCD(nil, nil, x, y).Uint64()
	x2 := x.Uint64() / z
	y2 := y.Uint64() / z

	fmt.Printf("x: %d", x2)
	fmt.Printf("y: %d", y2)
	fmt.Printf("z: %d ", z)

	if 2 != z {
		t.Errorf("Expected %d, got %d", 2, z)
	}
}
