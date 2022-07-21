package newboard

import (
	"fmt"
	"image/color"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func drawCard(screen *ebiten.Image) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if cardBoard[i][j].card != nil {
				screen.DrawImage(cardBoard[i][j].card.cardImg, cardop[i][j])
				cardBoard[i][j].card.playingStat(screen)
			}
		}
	}

	if attacker != nil {
		bCn := 4*attacker.pNum + int(attacker.bNum)
		attackerop := &ebiten.DrawImageOptions{}
		attackerop.GeoM.Translate(bC[bCn].bx+4, bC[bCn].by+4)

		screen.DrawImage(attacker.cardImg, attackerop)
		attacker.playingStat(screen)
	}
}

func (p *Players) drawCardSelect(screen *ebiten.Image) {

	d := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if !deckCard[p.pn][d].cardOn {
				screen.DrawImage(deckCard[p.pn][d].cardImg, cardop[i][j])
				if dbuf := fmt.Sprintf("%v", deckCard[p.pn][d].debuf); dbuf != "[]" {
					px, py := bC[d].bx, bC[d].by
					text.Draw(screen, dbuf, arcadeFontS, cardw/17*11+int(px), cardh/40*13+int(py), color.RGBA{0xff, 0, 0, 0xff})
				}
			}
			d++
			if d >= 10 {
				break
			}
		}
	}
}

func theCardimg(dc Cardstr) *ebiten.Image {

	cardimg := images["playingcard"] //.SubImage(image.Rect(0, 0, 160, 180))
	img := ebiten.NewImage(cardw, cardh)
	imgop := &ebiten.DrawImageOptions{}
	imgop.GeoM.Scale(float64(cardw)/160, float64(cardh)/180)
	img.DrawImage(cardimg, imgop)

	sx := cardw
	sy := cardh
	healthPoint := "HP"
	damagePoint := "DP"
	unitPrice := "$"
	vPC := strconv.Itoa(dc.price)
	vHP := strconv.Itoa(dc.hp)
	vDP := strconv.Itoa(dc.dp)

	text.Draw(img, string(dc.name), arcadeFontS, sx/10*1, sy/35*11, color.Black)
	text.Draw(img, unitPrice, arcadeFont, sx/16*10, sy/6, color.Black)
	text.Draw(img, vPC, arcadeFont, sx/13*10, sy/6, color.Black)
	text.Draw(img, dc.char, arcadeFont, sx/17*10, sy/4, color.RGBA{0xff, 0, 0, 0xff})
	text.Draw(img, healthPoint, arcadeFont, sx/8, sy/12*10, color.RGBA{0, 0, 0xff, 0xff})
	text.Draw(img, damagePoint, arcadeFont, sx/15*10, sy/12*10, color.RGBA{0xff, 0, 0, 0xff})
	text.Draw(img, string(dc.capa), arcadeFontS, sx/133*10, sy/18*10, color.Black)
	text.Draw(img, vHP, arcadeFont, sx/53*10, int(math.Round(float64(sy)*9.4/10)), color.Black)
	text.Draw(img, vDP, arcadeFont, sx/53*36, int(math.Round(float64(sy)*9.4/10)), color.Black)

	cop := &ebiten.DrawImageOptions{}
	cop.GeoM.Translate(float64(sx/8), float64(sy/20))
	c2op := &ebiten.DrawImageOptions{}
	c2op.GeoM.Translate(float64(sx/26*10), float64(sy/3))
	img.DrawImage(dc.img, cop)
	img.DrawImage(dc.cImg, c2op)

	return img
}

func (pc *deckCardstr) playingStat(screen *ebiten.Image) {
	sx := cardw
	sy := cardh
	bCn := 4*pc.pNum + int(pc.bNum)
	px, py := bC[bCn].bx, bC[bCn].by
	HPboard := images["HPboard"]
	hpBw, hpBh := HPboard.Size()
	hpop := &ebiten.DrawImageOptions{}
	hpop.GeoM.Scale(0.9*float64(sx)/float64(hpBw), 0.12*float64(sy)/float64(hpBh))
	hpop.GeoM.Translate(px+float64(sx)/20, py+float64(sy/10)*8.4)

	screen.DrawImage(HPboard, hpop)
	vHP := strconv.Itoa(pc.card.hp)
	text.Draw(screen, vHP, arcadeFont, sx/53*10+int(px), int(math.Round(float64(sy)*9.4/10))+int(py), color.Black)
	vDP := strconv.Itoa(pc.card.dp)
	text.Draw(screen, vDP, arcadeFont, sx/53*36+int(px), int(math.Round(float64(sy)*9.4/10))+int(py), color.Black)
	dbuf := fmt.Sprintf("%v", pc.debuf)
	text.Draw(screen, dbuf, arcadeFontS, sx/17*11+int(px), sy/40*13+int(py), color.RGBA{0xff, 0, 0, 0xff})
}

func drawPlayerStat(screen *ebiten.Image) {
	width := ScreenWidth
	//height := ScreenHeight
	p1y, p2y, p3y := int(bC[0].by), int(bC[4].by), int(bC[8].by)

	text.Draw(screen, strconv.Itoa(player[0].pHP), arcadeFont, width/12*10, p1y+90, color.RGBA{0, 0, 0xff, 0xff})
	text.Draw(screen, strconv.Itoa(player[1].pHP), arcadeFont, width/12*10, p2y+90, color.RGBA{0, 0, 0xff, 0xff})
	text.Draw(screen, strconv.Itoa(player[2].pHP), arcadeFont, width/12*10, p3y+90, color.RGBA{0, 0, 0xff, 0xff})

	text.Draw(screen, strconv.Itoa(player[0].pMoney), arcadeFont, width/12*11, p1y+90, color.Black)
	text.Draw(screen, strconv.Itoa(player[1].pMoney), arcadeFont, width/12*11, p2y+90, color.Black)
	text.Draw(screen, strconv.Itoa(player[2].pMoney), arcadeFont, width/12*11, p3y+90, color.Black)

	if playerNow.pn == 0 {
		text.Draw(screen, bmsg, arcadeFontB, width/120*103, p1y+120, color.Black)
	} else if playerNow.pn == 1 {
		text.Draw(screen, bmsg, arcadeFontB, width/120*103, p2y+120, color.Black)
	} else if playerNow.pn == 2 {
		text.Draw(screen, bmsg, arcadeFontB, width/120*103, p3y+120, color.Black)
	}

	text.Draw(screen, amsg, arcadeFontB, int(sepw), p3y+cardh+10, color.Black)

}
