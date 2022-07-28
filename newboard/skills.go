package newboard

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
