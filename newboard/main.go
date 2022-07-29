package newboard

import (
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// version 1.0
// mobile porting done
// version 1.01
// func cruel add
// version 1.02
// bug fix and assets add
// version 1.03
// background change
// cards add // orc ...
// 다람쥐 특성변경
// version 1.1
// optimizing
// used card display
// version 1.2
// copy skill implement
// version 1.3
// redesign
// cruelbools control modified (attackControl())
// version 1.4
// cardPick add
// version 1.4.1
// copy playerBuff bug fix
// pickBool touch control fix
// deckInit() functin add
// 2p mouseposition fix
// version 1.4.2
// dealer skill implement

const (
	ScreenWidth  int = 800 //800 1280 1440 //
	ScreenHeight int = 600 //600 960 720 //
	cardTotal    int = 56

	version string = "version 1.4.2"
)

type Game struct {
	touches     []ebiten.TouchID
	canvasImage *ebiten.Image
}

var (
	backImage *ebiten.Image = ebiten.NewImage(ScreenWidth, ScreenHeight)
	backop    *ebiten.DrawImageOptions

	bmsg string = "공격자 카드를 선택하세요"
	amsg string = ""

	menu                     bool = true
	touchedHome, touchedNext bool = false, false

	gameWin     bool    = false
	winPoint    int     = 0
	looseplayer [3]bool = [3]bool{false, false, false}
	twoplay     bool    = false

	attacker *deckCardstr = nil
	target   *deckCardstr = nil

	bigbugCard Cardstr

	pickBools  bool = true
	jipick     int  = 0
	ji         *int = &jipick
	pickplayer int  = 0
	pickChan        = make(chan string)
	//pickNChan       = make(chan string)

	dealBool, dealerS        bool = false, false
	dOffer, dReciever, dealP int  = 0, 0, 0
	dOfferD, dRecieverD      int  = 0, 0
)

func init() {
	LoadImages()
	fontimport()
	fillBack(backImage)
	cardimport()
	deckInit()
	newShuffle()
	//shuffleCard()
	playerinit()

	//fmt.Println(bigbugCard)
}

func (g *Game) Update() error {
	// g.touches = ebiten.AppendTouchIDs(g.touches[:0])
	g.touchControl()
	control()

	gamewin()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.DrawImage(backImage, backop)
	drawPlayerStat(screen)

	// if !cardS.sel && !menu && !pickBools {
	// 	drawCard(screen)
	// } else if cardS.sel && !menu && !pickBools {
	// 	playerNow.drawCardSelect(screen)
	// } else
	if menu {
		gameMenu(screen)
	} else if cardS.sel {
		playerNow.drawCardSelect(screen)
	} else if pickBools {
		pickingCard(screen, ji)
	} else if dealBool {
		drawDealer(screen)
	} else {
		drawCard(screen)
	}

	if touchedHome {
		drawTouchedHome(screen)
	}
	if touchedNext {
		drawTouchedNext(screen)
	}

	// ebitenutil.DebugPrint(screen, msg)
	mx, my := ebiten.CursorPosition()
	msg := fmt.Sprintf("(%d, %d)", mx, my)
	pbuff := playerBuff
	var msgbuff string
	for _, mb := range pbuff {
		msgbuff += fmt.Sprintf("%v:%v, ", mb.apl.pn, mb.buff)
	}

	msg += fmt.Sprintf("\n %v", msgbuff)
	msg += fmt.Sprintf("\n %v", dealBool)
	for _, t := range g.touches {
		x, y := ebiten.TouchPosition(t)
		msg += fmt.Sprintf("\n(%d, %d) touch %d", x, y, t)
	}
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	g := &Game{
		canvasImage: ebiten.NewImage(ScreenWidth, ScreenHeight),
	}
	return g

}
