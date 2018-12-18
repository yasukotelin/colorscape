package conv

import "testing"

func TestWhiteToRgb(t *testing.T) {
	h := "FFFFFF"

	example := [3]uint8{255, 255, 255}
	actual, _ := ToRgb(h)

	if example != actual {
		t.Fatalf("failed test. example is %v, actual is %v", example, actual)
	}
}

func TestGreenToRgb(t *testing.T) {
	h := "2ECC18"

	example := [3]uint8{46, 204, 24}
	actual, _ := ToRgb(h)

	if example != actual {
		t.Fatalf("failed test. example is %v, actual is %v", example, actual)
	}
}
