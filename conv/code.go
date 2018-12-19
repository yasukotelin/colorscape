package conv

import "strconv"

// ToRgb は16進数カラーコード文字列から10進数RGB表記へと変換する関数
func ToRgb(code string) ([3]uint8, error) {
	var rgb [3]uint8

	var codes [3]string
	switch len(code) {
	case 3:
		codes = [3]string{
			code[0:1] + code[0:1],
			code[1:2] + code[1:2],
			code[2:3] + code[2:3],
		}
	case 6:
		codes = [3]string{
			code[0:2],
			code[2:4],
			code[4:6],
		}
	default:
		// todo: エラー出力
	}

	for i, c := range codes {
		d, err := strconv.ParseUint(c, 16, 8)
		if err != nil {
			return rgb, err
		}
		rgb[i] = uint8(d)
	}
	return rgb, nil
}
