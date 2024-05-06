package normal

import (
	"os"
	"strconv"
	"testing"

	"github.com/ryomak/ogpgo"
)

func TestNormal(t *testing.T) {
	ogpgo.KouriBoldPath = "./../" + ogpgo.KouriBoldPath

	type testCase struct {
		text string
	}
	testCases := []testCase{
		{
			text: "こんにちは。テストOGPです",
		},
		{
			text: "長い文章長い文章長い文章長い文章長い文章長い文章長い文章長い文章長い",
		},
		{
			text: "長い文章長い文章長い文章長い文章長い文章長い文章長い文章長い文章長い文章長い文章長い文章長い文章長い文章長い文章長い文章長い文章",
		},
	}
	for i, tt := range testCases {
		file, err := os.Create("generated-ogp-normal-" + strconv.Itoa(i) + ".png")
		if err != nil {
			t.Errorf("Create() error = %v", err)
			return
		}

		if err := Normal(file, tt.text); err != nil {
			t.Errorf("Create() error = %v", err)
			return
		}
	}

}
