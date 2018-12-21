package conv

import (
	"strconv"
	"strings"
)

// ToColorCode はr, g, bをカラーコード表記文字列に変換する
func ToColorCode(r uint8, g uint8, b uint8) (string, error) {
	rHex := strconv.FormatInt(int64(r), 16)
	gHex := strconv.FormatInt(int64(g), 16)
	bHex := strconv.FormatInt(int64(b), 16)

	return strings.ToUpper(rHex + gHex + bHex), nil
}
