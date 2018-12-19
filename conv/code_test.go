package conv

import "testing"

func TestWhiteToRgb(t *testing.T) {
	h := "FFFFFF"

	expected := [3]uint8{255, 255, 255}
	actual, _ := ToRgb(h)

	if expected != actual {
		t.Fatalf("failed test. expected is %v, actual is %v", expected, actual)
	}
}

func TestWhiteToRgb2(t *testing.T) {
	h := "FFF"

	expected := [3]uint8{255, 255, 255}
	actual, _ := ToRgb(h)

	if expected != actual {
		t.Fatalf("failed test. expected is %v, actual is %v", expected, actual)
	}
}

func TestGreenToRgb(t *testing.T) {
	h := "2ECC18"

	expected := [3]uint8{46, 204, 24}
	actual, _ := ToRgb(h)

	if expected != actual {
		t.Fatalf("failed test. expected is %v, actual is %v", expected, actual)
	}
}
