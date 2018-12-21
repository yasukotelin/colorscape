package conv

import "testing"

func TestWhiteToColorCode(t *testing.T) {
	rgb := []uint8{255, 255, 255}

	expected := "FFFFFF"

	actual, _ := ToColorCode(rgb[0], rgb[1], rgb[2])

	if expected != actual {
		t.Fatalf("failed test. expected is %v, actual is %v", expected, actual)
	}
}

func TestGreenToColorCode(t *testing.T) {
	rgb := []uint8{46, 204, 24}

	expected := "2ECC18"

	actual, _ := ToColorCode(rgb[0], rgb[1], rgb[2])

	if expected != actual {
		t.Fatalf("failed test. expected is %v, actual is %v", expected, actual)
	}
}
