package newboard

import (
	"image"
	"image/color"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func gameMenu(screenN *ebiten.Image) {

	width := ScreenWidth
	height := ScreenHeight

	restart := images["n_start"]

	gameover := images["n_exit"]

	win := images["win"]

	player2 := images["n_2playerN"]

	menu := images["n_menu"]
	baw, bah = menu.Size()
	opback := &ebiten.DrawImageOptions{}
	opback.GeoM.Translate(0, 0)
	opback.GeoM.Scale(float64(width)/float64(baw), float64(height)/float64(bah))

	opre := &ebiten.DrawImageOptions{}
	opre.GeoM.Translate(float64(width/2-320), float64(height/2-200))
	opgo := &ebiten.DrawImageOptions{}
	opgo.GeoM.Translate(float64(width/2), float64(height/2-200))
	opwin := &ebiten.DrawImageOptions{}
	opwin.GeoM.Translate(0, float64(height/5))
	op2p := &ebiten.DrawImageOptions{}
	op2p.GeoM.Translate(float64(width/2-260), float64(height/2+20))

	// selWindow := ebiten.NewImage(width, height)
	// selWindow.Fill(color.Black)
	// sop := &ebiten.DrawImageOptions{}
	// screenN.DrawImage(selWindow, sop)

	screenN.DrawImage(menu, opback)

	screenN.DrawImage(restart, opre)
	screenN.DrawImage(gameover, opgo)
	screenN.DrawImage(player2, op2p)

	if gameWin {
		screenN.DrawImage(win, opwin)
	}

	homeButton := images["touch"].SubImage(image.Rect(240, 0, 320, 64)).(*ebiten.Image)
	ophome := &ebiten.DrawImageOptions{}
	ophome.GeoM.Translate(float64(buttonPos[0].sx), float64(buttonPos[0].sy))
	screenN.DrawImage(homeButton, ophome)

	text.Draw(screenN, version, arcadeFontB, 100, 100, color.Black)

}

func menuselecting(mx int) {
	if mx == 1 {
		twoplay = false
		newGame()
		menu = false
		pickBools = true
		//go pickNumber(ji)
	} else if mx == 2 {
		os.Exit(0)
	} else if mx == 3 {
		twoplay = true
		newGame()
		player[2].pHP = 0
		looseplayer[2] = true
		winPoint = 1
		menu = false
		pickBools = true
		//go pickNumber(ji)
	} else if mx == 5 {
		touchedHome = true
		go func() {
			time.Sleep(200 * time.Millisecond)
			touchedHome = false
		}()
		menu = false
	}
}

func newGame() {

	newShuffle()
	//shuffleCard()
	playerinit()
	copybook = []copyBookStr{}
	winPoint = 0

	gameWin = false
	menu = false

	jipick = 0
	pickplayer = 0
}

func gamewin() {

	for i := 0; i < 3; i++ {
		if player[i].pHP <= 0 && !looseplayer[i] {
			winPoint = winPoint + 1
			looseplayer[i] = true
			for b := 0; b < 4; b++ {
				if cardBoard[i][b].card != nil {
					cardBoard[i][b].card.offCard()
				}
			}
		}
	}

	if winPoint >= 2 {
		gameWin = true
		menu = true

	}

}
