package normal

import (
	"image"
	"image/png"
	"io"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	"github.com/golang/freetype/truetype"
	"github.com/ryomak/ogpgo"
)

const (
	ratio         = 3
	defaultWidth  = 400 * ratio
	defaultHeight = 210 * ratio
)

type NormalRequest struct {
	img    *image.RGBA
	width  int
	height int
}

func NewNormalRequest() *NormalRequest {
	return &NormalRequest{
		// OGP画像の土台を作成
		img:    image.NewRGBA(image.Rect(0, 0, defaultWidth, defaultHeight)),
		width:  defaultWidth,
		height: defaultHeight,
	}
}

func (r *NormalRequest) Normal(w io.Writer, text string) error {

	// fontの初期化
	koruriFont, err := ogpgo.KoruriBold()
	if err != nil {
		return err
	}

	// Drawerの初期化
	fontSize := expansion(14)
	whiteColor, _ := ogpgo.ParseHexColor("#ffffff")
	dr := &font.Drawer{
		Dst: r.img,
		Src: image.NewUniform(whiteColor),
		Face: truetype.NewFace(koruriFont, &truetype.Options{
			Size: float64(fontSize),
		}),
		Dot: fixed.Point26_6{},
	}

	// 描画
	paddingLeft := expansion(40)
	centerY := expansion(105)
	centerX := expansion(200)
	splitLines := ogpgo.SplitByMeasureWidth(text, defaultWidth-paddingLeft*2, dr)
	lineHeight := 1.5

	y := ogpgo.TextCenterYPosition(centerY, fontSize, lineHeight, len(splitLines))
	for _, line := range splitLines {
		x := ogpgo.TextCenterXPosition(centerX, line, fontSize)
		dr.Dot.X = fixed.I(x)
		dr.Dot.Y = fixed.I(y)
		dr.DrawString(line)

		// 次の開始位置を計算
		y += int(float64(fontSize) * lineHeight)
	}

	if err := png.Encode(w, r.img); err != nil {
		return err
	}

	return nil
}

func expansion(size int) int {
	return size * ratio
}
