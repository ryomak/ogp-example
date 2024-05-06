package ogpgo

import (
	"errors"
	"image/color"

	"golang.org/x/image/font"
)

func SplitByMeasureWidth(text string, maxWidth int, dr *font.Drawer) []string {
	var (
		lines []string
		line  string
	)
	for _, v := range []rune(text) {
		vs := string(v)
		w := dr.MeasureString(line + vs).Round()
		switch {
		case maxWidth <= w:
			lines = append(lines, line)
			line = vs
		default:
			line = line + vs
		}
	}
	lines = append(lines, line)
	return lines
}

func ParseHexColor(s string) (c color.RGBA, err error) {
	errInvalidFormat := errors.New("invalid format")

	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}

func TextCenterXPosition(centerX int, text string, fontSize int) int {
	return centerX - len([]rune(text))*fontSize/2
}

func TextCenterYPosition(centerY int, fontSize int, lineHeight float64, lines int) int {
	paddingLength := lines - 1
	lineSize := (lineHeight - 1) * float64(fontSize) * float64(paddingLength)

	// 中央の座標 - (フォントサイズ * 行数 + 行間のサイズ) / 2 + (フォントサイズ * 行間のサイズ) / 2
	// 中央の座標 - lineすべてで使う高さ / 2 + 開始位置を指定するための高さ
	return centerY - (fontSize*lines+int(lineSize))/2 + int(float64(fontSize)*lineHeight)/2
}
