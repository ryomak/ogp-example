package ogpgo

import (
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

var (
	cacheKoruiBold *truetype.Font

	KouriBoldPath = "static/font/Koruri-Bold.ttf"
)

func KoruriBold() (*truetype.Font, error) {
	if cacheKoruiBold != nil {
		return cacheKoruiBold, nil
	}
	f, err := newFont(KouriBoldPath)
	if err != nil {
		return nil, err
	}
	cacheKoruiBold = f
	return f, nil
}

func newFont(path string) (*truetype.Font, error) {
	fntBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return freetype.ParseFont(fntBytes)
}
