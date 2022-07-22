package newboard

import (
	"fmt"
	"time"
)

// putcard -> skillDoPassive
// attackcard -> skillDo
// befor turn start -> buffDoActive (after pre player click next)
// after turn end -> buffDoPassive (after player click next)

// Active skill
func (pca *deckCardstr) skillDo(pct *deckCardstr) {
	ska := pca.card.skill
	skt := pct.card.skill
	//pa := &player[pca.pNum]
	pt := &player[pct.pNum]

	switch ska {
	case "poison":
		if !pct.checkcbuf(ska) {
			pca.addBuff(pct)
			pca.addPlayerBuff(pct)
		}
	case "vampire":
		if pct.card.hp <= pca.card.dp {
			go func() {
				<-attackChan
				if !pct.cardOn {
					pca.card.hp += 3
				}
			}()
		}
	case "flying":
		for _, buff := range pct.debuf {
			if buff == "lneck" {
				return
			}
		}

		pt.pHP = pt.pHP - pca.card.dp
		pca.used = true
	case "tentacle":
		if !pct.checkcbuf(ska) {
			pca.addBuff(pct)
			pca.addPlayerBuff(pct)
		}

		go func() {
			var killed bool = false
			ptdc := &deckCard[pct.pNum][pct.deckNum]
			pca.addBuff(ptdc)
			for {
				//wait for turnChannel (when player click next, turnChannel will be opened "control.go")
				<-turnChan

				if !killed {
					if pct.cardOn {
						continue
					} else {
						killed = true
					}
				} else if killed {
					if pct.cardOn && ptdc.checkcbuf(ska) {
						ptdc.remBuff(ska)
						pca.addBuff(pct)
						pca.addPlayerBuff(pct)
						return
					}
				}
			}
		}()
	case "threat":
		if !pct.checkcbuf(ska) {
			pct.card.dp--
			pca.addBuff(pct)
		}
	case "fire":
		pt.pHP -= pca.card.dp
	case "strategy":
		pNumn := nexP(pct.pNum)
		if pNumn == pca.pNum {
			pNumn = nexP(pNumn)
		}
		if pNumn != pct.pNum && player[pNumn].pHP > 0 {
			if cardBoard[pNumn][pct.bNum].card != nil {
				pctn := cardBoard[pNumn][pct.bNum].card
				pctn.card.hp = pctn.card.hp - pca.card.dp
				if pctn.card.hp <= 0 {
					pctn.offCard()
				}
			} else {
				pt.pHP -= pca.card.dp
			}
		}
	case "dolphin":
		pca.addBuff(pct)
		pca.addPlayerBuff(pct)
	case "sonic":
		pca.addBuff(pct)
		pca.addPlayerBuff(pct)
	case "selfdestruct":
		go func() {
			<-attackChan

			pca.offCard()
		}()
	case "mistaken":
		go func() {
			<-attackChan

			if !pct.cardOn {
				pcapn, pcadn := pca.pNum, pca.deckNum
				pctpn, pctdn := pct.pNum, pct.deckNum
				pca.offCard()

				pcaDeckC := deckCard[pcapn][pcadn]
				deckCard[pcapn][pcadn] = deckCard[pctpn][pctdn]
				deckCard[pcapn][pcadn].pNum = pcapn
				deckCard[pcapn][pcadn].deckNum = pcadn

				deckCard[pctpn][pctdn] = pcaDeckC
				deckCard[pctpn][pctdn].deckNum = pctdn
				deckCard[pctpn][pctdn].pNum = pctpn

			}
		}()
	case "climb":
		go func() {
			<-attackChan

			if !pct.cardOn {
				if cardBoard[pct.pNum][preB(pct.bNum)].card != nil {
					pca.attackCard(cardBoard[pct.pNum][preB(pct.bNum)].card)
				}
				if cardBoard[pct.pNum][nexB(pct.bNum)].card != nil {
					pca.attackCard(cardBoard[pct.pNum][nexB(pct.bNum)].card)
				}
			}
		}()
	case "rush":
		go func() {
			<-attackChan
			if pct.cardOn && pca.cardOn {
				pca.card.hp--
				if pca.card.hp <= 0 {
					pca.offCard()
				}
			}
		}()
	case "incision":
		go func() {
			<-attackChan
			if !pct.cardOn {
				pdc := &deckCard[pct.pNum][pct.deckNum]
				pdc.debuf = append(pdc.debuf, ska)
				pca.addPlayerBuff(pct)
			}
		}()
	case "trans":
		goto AFTERSKT
	case "huge":
		if npct := cardBoard[pct.pNum][nexB(pct.bNum)].card; pct.bNum != 3 && pca.bNum == pct.bNum && npct != nil {
			pca.attackCard(npct)
		}

	}

	switch skt {
	case "revenge":
		pca.card.hp--
		if pca.card.hp <= 0 {
			pca.offCard()
		}
	case "wit":
		go func(pctCard deckCardstr, plb []buffStr) {
			<-attackChan

			if !pct.cardOn && cardBoard[pctCard.pNum][preB(pctCard.bNum)].card == nil {
				pct.putCard(&pctCard, preB(pctCard.bNum))
				deckCard[pct.pNum][pct.deckNum].bNum = preB(pctCard.bNum)
				deckCard[pct.pNum][pct.deckNum].cardOn = true

				playerBuff = plb
			}
		}(*pct, playerBuff)
	case "dolphin":
		pct.addPlayerBuff(pct)
	case "water":
		pct.addPlayerBuff(pct)
	case "company":
		go func() {
			<-attackChan
			if !pct.cardOn {
				cardS.pn = playerNow.pn
				playerNow = &player[pct.pNum]
				cardS.bn = pct.bNum
				cardS.sel = true
				cardS.skill = "company"

				bmsg = "추가할 카드를 선택하세요"
			}
		}()
	case "shell":
		go func() {
			<-attackChan
			if pct.cardOn {
				pct.addBuff(pct)
			}
		}()

	}
AFTERSKT:
}

// Active skill when attack castle
func (pca *deckCardstr) skillDoCastle(pt *Players) {
	ska := pca.card.skill

	switch ska {
	case "strategy":
		pNumn := nexP(pt.pn)
		if pNumn == pca.pNum {
			pNumn = nexP(pNumn)
		}
		if pNumn != pt.pn && player[pNumn].pHP > 0 {
			if cardBoard[pNumn][pca.bNum].card != nil {
				pctn := cardBoard[pNumn][pca.bNum].card
				pctn.card.hp = pctn.card.hp - pca.card.dp
				if pctn.card.hp <= 0 {
					pctn.offCard()
				}
			} else {
				pt.pHP -= pca.card.dp
			}
		}
	case "selfdestruct":
		go func() {
			time.Sleep(300 * time.Millisecond)

			pca.offCard()
		}()
	case "huge":
		if npct := cardBoard[pt.pn][nexB(pca.bNum)].card; pca.bNum != 3 && npct != nil {
			pca.attackCard(npct)
		}

	}
}

// Passive skill, Activated when card is put on board
func (pca *deckCardstr) skillDoPassive() {
	ska := pca.card.skill
	//skt := pct.card.skill
	//pa := &player[pca.pNum]
	//pt := &player[pct.pNum]

	switch ska {
	case "run":
		if !pca.checkcbuf(ska) {
			pca.addBuff(pca)
			pca.addPlayerBuff(pca)
		}
	case "eggC":
		pca.addPlayerBuff(pca)
	case "cute":
		bn := pca.bNum
		prepc := cardBoard[preP(pca.pNum)][bn].card
		nexpc := cardBoard[nexP(pca.pNum)][bn].card
		if prepc != nil && prepc.card.dp != 0 && !prepc.checkcbuf(ska) {
			prepc.card.dp--
			pca.addBuff(prepc)
			pca.addPlayerBuff(prepc)
		}
		if nexpc != nil && nexpc.card.dp != 0 && !nexpc.checkcbuf(ska) {
			nexpc.card.dp--
			pca.addBuff(nexpc)
			pca.addPlayerBuff(nexpc)
		}
		pca.addPlayerBuff(pca)
	case "dolphin":
		pca.addBuff(pca)
	case "water":
		pca.addBuff(pca)
	case "company":
		pca.addBuff(pca)
	case "survive":
		pca.addBuff(pca)
		pca.addPlayerBuff(pca)
	case "helios":
		pca.addBuff(pca)
		pca.addPlayerBuff(pca)
	case "spacesupply":
		pca.addBuff(pca)
		pca.addPlayerBuff(pca)
	case "sCoin":
		pca.addBuff(pca)
		pca.addPlayerBuff(pca)
	case "totem":
		// add passive buff "lneck" to itself
		pca.addBuff(pca)

		// add buff "leader" to side card
		bn := pca.bNum
		prepc := cardBoard[pca.pNum][preB(bn)].card
		nexpc := cardBoard[pca.pNum][nexB(bn)].card
		if prepc != nil && !prepc.checkcbuf(ska) {
			prepc.card.dp++
			pca.addBuff(prepc)
			pca.addPlayerBuff(prepc)
		}
		if nexpc != nil && !nexpc.checkcbuf(ska) {
			nexpc.card.dp++
			pca.addBuff(nexpc)
			pca.addPlayerBuff(nexpc)
		}
		pca.addPlayerBuff(pca)
	case "lneck":
		pca.addBuff(pca)
	case "jump":
		if !pca.checkcbuf(ska) {
			pca.addBuff(pca)
			pca.addPlayerBuff(pca)
		}
	case "airproof":
		pca.addBuff(pca)
	case "diving":
		pca.addBuff(pca)

	}

}

// Active PlayerBuff // excute when turns end (when player clicks next)
func (p *Players) buffDoActive() {

	for _, buffs := range playerBuff {
		buff := buffs.buff
		apl, apc := buffs.apl, buffs.apc
		tpl, tpc := buffs.tpl, buffs.tpc

		if buff != "" {
			switch buff {
			case "run":
				if playerNow == apl && apc.cardOn {
					for i := 0; i < 4; i++ {
						cardBoard[apc.pNum][i].card = nil
					}
					for i := 0; i < 10; i++ {
						if playingCard[apc.pNum][i].cardOn {
							bn := playingCard[apc.pNum][i].bNum
							cardBoard[apc.pNum][preB(bn)].card = &playingCard[apc.pNum][i]
							playingCard[apc.pNum][i].bNum = preB(bn)
						}
					}
				}
			case "tentacle":
				if playerNow == apl { // && apc.cardOn {
					if tpc.cardOn {
						tpc.card.hp = tpc.card.hp - 1
						if tpc.card.hp <= 0 {
							tpc.offCard()
						}
					}
				}
			case "jump":
				if playerNow == apl && apc.cardOn {
					npn, nbn := apc.pNum, apc.bNum
					if cardBoard[npn][preB(nbn)].card == nil {
						cardBoard[npn][preB(nbn)].card = apc
						cardBoard[npn][nbn].card = nil
						apc.bNum = preB(nbn)
					} else if cardBoard[npn][nexB(nbn)].card == nil {
						cardBoard[npn][nexB(nbn)].card = apc
						cardBoard[npn][nbn].card = nil
						apc.bNum = nexB(nbn)
					}
				}
			case "sonic":
				if playerNow == tpl && apl != tpl && tpc.cardOn {
					tpc.remBuff(buff)
					apc.removePlayerBuff(tpc, buff)
				}

			}
		}
	}
	//	}
}

// Passive PlayerBuff // excute when turns start (when previous player clicks next)
func (p *Players) buffDoPassive() {

	for _, buffs := range playerBuff {
		apl, aturn, apc := buffs.apl, buffs.aturn, buffs.apc
		tpl, tturn, tpc := buffs.tpl, buffs.tturn, buffs.tpc
		buff := buffs.buff

		if buff != "" {
			switch buff {
			case "poison":
				if playerNow == apl && playerNow.turn == aturn+1 {
					tpc.offCard()
				}
			case "eggC":
				if apc.cardOn {
					if playerNow == apl && playerNow.turn == aturn+2 {
						apc.offCard()
						playerNow.bn = apc.bNum
						cardS.sel = true
						cardS.skill = "eggC"
						bmsg = "추가할 카드를 선택하세요"
						amsg = fmt.Sprintf("0) cardselecting %v , %v\n", apc.pNum, apc.bNum)

					} else {
						continue
					}
				}
			case "cute":
				bn := apc.bNum
				prepc := cardBoard[preP(apc.pNum)][bn].card
				nexpc := cardBoard[nexP(apc.pNum)][bn].card
				if tpc != apc && tpc.bNum != apc.bNum && tpc.checkcbuf(buff) {
					tpc.remBuff(buff)
					apc.removePlayerBuff(tpc, buff)
				}

				if prepc != nil && prepc.card.dp != 0 && !prepc.checkcbuf(buff) {
					prepc.card.dp--
					apc.addBuff(prepc)
					apc.addPlayerBuff(prepc)
				}
				if nexpc != nil && nexpc.card.dp != 0 && !nexpc.checkcbuf(buff) {
					nexpc.card.dp--
					apc.addBuff(nexpc)
					apc.addPlayerBuff(nexpc)
				}
			case "sonic":
				if playerNow == tpl && apl != tpl && tpc.cardOn {
					tpc.used = true
				}
			case "water":
				if playerNow == apl && apl == tpl && apc.cardOn {
					apc.card.hp++
					apc.removePlayerBuff(apc, buff)
				}
			case "survive":
				if playerNow == apl && playerNow.turn == aturn+2 {
					apc.removePlayerBuff(apc, buff)
					apc.remBuff(buff)
				}
			case "helios":
				apc.card.dp += 2
			case "spacesupply":
				if playerNow == apl {
					nbn := nexB(apc.bNum)
					pbn := preB(apc.bNum)
					if cardBoard[apc.pNum][nbn].card != nil {
						cardBoard[apc.pNum][nbn].card.card.hp++
					}
					if cardBoard[apc.pNum][pbn].card != nil {
						cardBoard[apc.pNum][pbn].card.card.hp++
					}
				}
			case "sCoin":
				if playerNow == apl {
					playerNow.pMoney++
				}
			case "incision":
				if playerNow == tpl && playerNow.turn == tturn+4 {
					apc.removePlayerBuff(tpc, buff)
					pdc := &deckCard[tpc.pNum][tpc.deckNum]
					pdc.remBuff(buff)

				}
			case "leader":
				bn := apc.bNum
				prepc := cardBoard[apc.pNum][preB(bn)].card
				nexpc := cardBoard[apc.pNum][nexB(bn)].card
				if tpc != apc && (tpc.bNum != preB(apc.bNum) || tpc.bNum != nexB(apc.bNum)) && tpc.checkcbuf(buff) {
					tpc.remBuff(buff)
					apc.removePlayerBuff(tpc, buff)
				}

				if prepc != nil && !prepc.checkcbuf(buff) {
					prepc.card.dp++
					apc.addBuff(prepc)
					apc.addPlayerBuff(prepc)
				}
				if nexpc != nil && !nexpc.checkcbuf(buff) {
					nexpc.card.dp++
					apc.addBuff(nexpc)
					apc.addPlayerBuff(nexpc)
				}

			}
		}
	}
}
