package newboard

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var cardNimg [5]*ebiten.Image
var cardN int

func chimgMenu(screenN *ebiten.Image) {
	width := ScreenWidth
	height := ScreenHeight

	allcardN := images["card_all"]
	allbackN := images["back_all"]
	// cardNimg[1] = images["card_2"]
	// cardNimg[2] = images["card_3"]
	// cardNimg[3] = images["card_4"]
	// cardNimg[4] = images["card_5"]

	back1 := images["back_1"]
	// back2 := images["back_2"]
	// back3 := images["back_3"]
	// back4 := images["back_4"]

	cardN = 0
	samplecard := deckCardstr{allCard[0], theCardimg(allCard[0]), 0, 0, 0, false, false, nil}

	backN := back1
	baw, bah = backN.Size()
	opback := &ebiten.DrawImageOptions{}
	opback.GeoM.Translate(0, 0)
	opback.GeoM.Scale(float64(width)/float64(baw), float64(height)/float64(bah))

	screenN.DrawImage(backN, opback)
	screenN.DrawImage(samplecard.cardImg, cardop[0][0])
	samplecard.playingStat(screenN)

	screenN.DrawImage(allcardN, cardop[1][0])
	screenN.DrawImage(allbackN, cardop[2][0])

	nextButton := images["touch"].SubImage(image.Rect(80, 0, 160, 64)).(*ebiten.Image)
	homeButton := images["touch"].SubImage(image.Rect(240, 0, 320, 64)).(*ebiten.Image)

	ophome := &ebiten.DrawImageOptions{}
	ophome.GeoM.Translate(float64(buttonPos[0].sx), float64(buttonPos[0].sy))
	opnext := &ebiten.DrawImageOptions{}
	opnext.GeoM.Translate(float64(buttonPos[1].sx), float64(buttonPos[1].sy))

	screenN.DrawImage(nextButton, opnext)
	screenN.DrawImage(homeButton, ophome)

}
