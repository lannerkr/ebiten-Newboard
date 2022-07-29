package newboard

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (pca *deckCardstr) attpc(pct *deckCardstr) {
	pct.card.hp -= pca.card.dp
	pca.used = true

}
func (pca *deckCardstr) attpl(pct *deckCardstr) {
	player[pct.pNum].pHP -= pca.card.dp
	pca.used = true

}

/////////////////////////////////////////////////////////////////////
// flying
func (pca *deckCardstr) flyingFunc(pct *deckCardstr) {
	for _, buff := range pct.debuf {
		if buff == "lneck" {
			return
		}
	}
	pca.attpl(pct)
}

/////////////////////////////////////////////////////////////////////
// copy
type copyBookStr struct {
	pc *deckCardstr
	tc [2]tcStr
}
type tcStr struct {
	tpc  *deckCardstr
	buff string
}

var copybook []copyBookStr

func (pca *deckCardstr) OnboardSkillCopy() {
	ppc := &cardBoard[preP(pca.pNum)][pca.bNum]
	npc := &cardBoard[nexP(pca.pNum)][pca.bNum]

	for i, cpb := range copybook {
		if cpb.pc == pca {
			if cpb.tc[0].tpc != ppc.card {
				if cpb.tc[0].buff != "" {
					pca.remBuff(cpb.tc[0].buff)
					pca.removePlayerBuff(pca, cpb.tc[0].buff)
				}
				if ppc.card == nil {
					copybook[i].tc[0] = tcStr{nil, ""}

				} else {
					copybook[i].tc[0].tpc = ppc.card
					pca.card.skill = ppc.card.card.skill
					bl0 := len(pca.debuf)
					pca.skillDoPassive()
					bl1 := len(pca.debuf)
					if bl0 != bl1 {
						copybook[i].tc[0].buff = pca.debuf[bl1-1]
					}
					pca.card.skill = "copy"
				}
			}

			if cpb.tc[1].tpc != npc.card {
				if cpb.tc[1].buff != "" {
					pca.remBuff(cpb.tc[1].buff)
					pca.removePlayerBuff(pca, cpb.tc[1].buff)
				}
				if npc.card == nil {
					copybook[i].tc[1] = tcStr{nil, ""}
				} else {
					copybook[i].tc[1].tpc = npc.card
					pca.card.skill = npc.card.card.skill
					bl0 := len(pca.debuf)
					pca.skillDoPassive()
					bl1 := len(pca.debuf)
					if bl0 != bl1 {
						copybook[i].tc[1].buff = pca.debuf[bl1-1]
					}
					pca.card.skill = "copy"
				}
			}

		}
	}
}

/////////////////////////////////////////////////////////////////////
// cruel
type cruelStr struct {
	pc    *deckCardstr
	bn    int
	bools bool
}

var cruel cruelStr = cruelStr{nil, 20, false}

func (pc *deckCardstr) cruelty(cbn int) {
	cards := false
	for i := 0; i < 4; i++ {
		if cardBoard[pc.pNum][i].card != nil && cardBoard[pc.pNum][i].card.card.price >= 2 {
			cards = true
			break
		}
	}
	if !cards {
		bmsg = "희생할 $2 이상 카드가 없습니다."
		cardS.sel = false
		return
	}

	cardS.sel = false
	cruel.bools = true
	cruel.pc = pc
	cruel.bn = cbn
	bmsg = "희생할 카드를 선택하세요."

}

/////////////////////////////////////////////////////////////////////
// dealer
func drawDealer(screen *ebiten.Image) {
	dealerP := &player[dealP]
	if dealerS {
		dealerP.drawCardSelect(screen)
	} else {
		var plimg, plused [3]*ebiten.Image
		var upop [3]*ebiten.DrawImageOptions

		plimg[0] = images["picking_p0"]
		plimg[1] = images["picking_p1"]
		plimg[2] = images["picking_p2"]
		plused[0] = images["picked_p0"]
		plused[1] = images["picked_p1"]
		plused[2] = images["picked_p2"]

		for i := 0; i < 3; i++ {
			if i == 2 && twoplay {
				break
			}
			p := i * 4
			upop[i] = &ebiten.DrawImageOptions{}
			upop[i].GeoM.Scale(float64(cardw)/160, float64(cardh)/180)
			upop[i].GeoM.Translate(bC[p].bx, bC[p].by)
			screen.DrawImage(plimg[i], upop[i])
		}

		nop := &ebiten.DrawImageOptions{}
		nop.GeoM.Scale(float64(cardw)/160, float64(cardh)/180)
		nop2 := &ebiten.DrawImageOptions{}
		nop2.GeoM.Scale(float64(cardw)/160, float64(cardh)/180)

		///////////////////////////////////////////////////////////////////////////

		// deal offer user select
		//dOffer = 0
		nop.GeoM.Translate(bC[dOffer*4].bx, bC[dOffer*4].by)
		screen.DrawImage(plused[dOffer], nop)

		// deal reciever select
		//dReciever = 1
		nop2.GeoM.Translate(bC[dReciever*4+1].bx, bC[dReciever*4+1].by)
		screen.DrawImage(plused[dReciever], nop2)

		// offer card select
		screen.DrawImage(deckCard[dOffer][dOfferD].cardImg, cardop[dOffer][2])

		// recievre card select
		screen.DrawImage(deckCard[dReciever][dRecieverD].cardImg, cardop[dReciever][3])

	}
}

func dealerCardChange() {
	//fmt.Println(deckCard[dOffer][dOfferD], deckCard[dReciever][dRecieverD])

	offerCard := deckCard[dOffer][dOfferD]
	recieverCard := deckCard[dReciever][dRecieverD]

	deckCard[dOffer][dOfferD] = recieverCard
	deckCard[dOffer][dOfferD].pNum, deckCard[dOffer][dOfferD].deckNum = offerCard.pNum, offerCard.deckNum
	deckCard[dReciever][dRecieverD] = offerCard
	deckCard[dReciever][dRecieverD].pNum, deckCard[dReciever][dRecieverD].deckNum = recieverCard.pNum, recieverCard.deckNum

	//fmt.Println(deckCard[dOffer][dOfferD], deckCard[dReciever][dRecieverD])
}
