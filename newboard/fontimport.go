package newboard

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	arcadeFont  font.Face
	arcadeFontS font.Face
	arcadeFontB font.Face
)

func fontimport() {
	ttp, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	tt, err := opentype.Parse(NanumGothic_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const (
		arcadeFontSize  = 14
		arcadeFontSizeS = 16
		arcadeFontSizeB = 20
		dpi             = 72
	)
	arcadeFont, err = opentype.NewFace(ttp, &opentype.FaceOptions{
		Size:    arcadeFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	arcadeFontS, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    arcadeFontSizeS,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	arcadeFontB, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    arcadeFontSizeB,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

}
