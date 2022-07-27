package newboard

import (
	"fmt"
	"time"
)

// putcard -> skillDoPassive
// attackcard -> skillDo
// befor turn start -> buffDoActive (after pre player click next)
// after turn end -> buffDoPassive (after player click next)

// Attack skill
func (pca *deckCardstr) skillDoA(pct *deckCardstr, ska string) {
	//ska := pca.card.skill
	//skt := pct.card.skill
	//pa := &player[pca.pNum]
	pt := &player[pct.pNum]

	switch ska {
	case "poison":
		if !pct.checkcbuf(ska) {
			pca.addBuff(pct, ska)
			pca.addPlayerBuff(pct, ska)
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
		pca.flyingFunc(pct)
		if pca.used {
			return
		}
	case "tentacle":
		if !pct.checkcbuf(ska) {
			pca.addBuff(pct, ska)
			pca.addPlayerBuff(pct, ska)
		}

		go func() {
			var killed bool = false
			ptdc := &deckCard[pct.pNum][pct.deckNum]
			pca.addBuff(ptdc, ska)
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
						pca.addBuff(pct, ska)
						pca.addPlayerBuff(pct, ska)
						return
					}
				}
			}
		}()
	case "threat":
		if !pct.checkcbuf(ska) {
			pct.card.dp--
			pca.addBuff(pct, ska)
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
		pca.addBuff(pct, "sonic")
		pca.addPlayerBuff(pct, ska)
	case "sonic":
		pca.addBuff(pct, ska)
		pca.addPlayerBuff(pct, ska)
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
				pca.addPlayerBuff(pct, ska)
			}
		}()
	case "trans":
		return
	case "huge":
		if npct := cardBoard[pct.pNum][nexB(pct.bNum)].card; pct.bNum != 3 && pca.bNum == pct.bNum && npct != nil {
			pca.attackCard(npct)
		}
	case "wield":
		if cardBoard[pct.pNum][nexB(pct.bNum)].card == nil {
			cardBoard[pct.pNum][pct.bNum].card = nil
			cardBoard[pct.pNum][nexB(pct.bNum)].card = pct
			pct.bNum = nexB(pct.bNum)
		}
	case "charge":
		pca.addBuff(pca, ska)
		pca.addPlayerBuff(pca, ska)
	case "runover":
		go func() {
			<-attackChan
			if pct.cardOn && pca.cardOn {
				npn := nexP(pct.pNum)
				if npn == pca.pNum {
					npn = nexP(npn)
				}

				if cardBoard[npn][pct.bNum].card == nil {
					pca.addBuff(pct, ska)
					pca.addPlayerBuff(pct, ska)

					cardBoard[pct.pNum][pct.bNum].card = nil
					cardBoard[npn][pct.bNum].card = pct
					pct.pNum = npn

				}
			}
		}()
	case "copy":
		fmt.Println("000]")
		if ppc := cardBoard[preP(pca.pNum)][pca.bNum].card; ppc != nil {

			pska := ppc.card.skill
			//pct.card.skill = pskt
			fmt.Println("001]", pca.card.skill)
			pca.skillDoA(pct, pska)

		}
		if npc := cardBoard[nexP(pca.pNum)][pca.bNum].card; npc != nil {
			nska := npc.card.skill
			//pct.card.skill = nskt
			fmt.Println("002]", pca.card.skill)
			pca.skillDoA(pct, nska)

		}

	}

	// target skill
	if pca.card.skill != "copy" || pct.card.skill != "copy" {
		pca.skillDoT(pct, pct.card.skill)
	}
}
func (pca *deckCardstr) skillDoT(pct *deckCardstr, skt string) {
	//skt := pct.card.skill
	//pa := &player[pca.pNum]
	//pt := &player[pct.pNum]

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
		pct.addPlayerBuff(pct, skt)
	case "water":
		pct.addPlayerBuff(pct, skt)
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
				pct.addBuff(pct, skt)
			}
		}()
	case "airproof":
		if pca.card.dp != 2 {
			if pca.card.dp > 2 {
				pct.card.hp += pca.card.dp - 2
			} else {
				pct.card.hp -= 2 - pca.card.dp
			}
		}
	case "copy":
		fmt.Println("000]")
		if ppc := cardBoard[preP(pct.pNum)][pct.bNum].card; ppc != nil {

			pskt := ppc.card.skill
			//pct.card.skill = pskt
			fmt.Println("001]", pct.card.skill)
			pca.skillDoT(pct, pskt)

		}
		if npc := cardBoard[nexP(pct.pNum)][pct.bNum].card; npc != nil {
			nskt := npc.card.skill
			//pct.card.skill = nskt
			fmt.Println("002]", pct.card.skill)
			pca.skillDoT(pct, nskt)

		}

	}
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
	case "charge":
		pca.addBuff(pca, ska)
		pca.addPlayerBuff(pca, ska)

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
			pca.addBuff(pca, ska)
			pca.addPlayerBuff(pca, ska)
		}
	case "eggC":
		pca.addPlayerBuff(pca, ska)
	case "cute":
		bn := pca.bNum
		prepc := cardBoard[preP(pca.pNum)][bn].card
		nexpc := cardBoard[nexP(pca.pNum)][bn].card
		if prepc != nil && prepc.card.dp != 0 && !prepc.checkcbuf(ska) {
			prepc.card.dp--
			pca.addBuff(prepc, ska)
			pca.addPlayerBuff(prepc, ska)
		}
		if nexpc != nil && nexpc.card.dp != 0 && !nexpc.checkcbuf(ska) {
			nexpc.card.dp--
			pca.addBuff(nexpc, ska)
			pca.addPlayerBuff(nexpc, ska)
		}
		pca.addPlayerBuff(pca, ska)
	case "dolphin":
		pca.addBuff(pca, ska)
	case "water":
		pca.addBuff(pca, ska)
	case "company":
		pca.addBuff(pca, ska)
	case "survive":
		pca.addBuff(pca, ska)
		pca.addPlayerBuff(pca, ska)
	case "helios":
		pca.addBuff(pca, ska)
		pca.addPlayerBuff(pca, ska)
	case "spacesupply":
		pca.addBuff(pca, ska)
		pca.addPlayerBuff(pca, ska)
	case "sCoin":
		pca.addBuff(pca, ska)
		pca.addPlayerBuff(pca, ska)
	case "totem":
		// add passive buff "lneck" to itself
		pca.addBuff(pca, "lneck")

		// add buff "leader" to side card
		bn := pca.bNum
		prepc := cardBoard[pca.pNum][preB(bn)].card
		nexpc := cardBoard[pca.pNum][nexB(bn)].card
		if prepc != nil && !prepc.checkcbuf(ska) {
			prepc.card.dp++
			pca.addBuff(prepc, "leader")
			pca.addPlayerBuff(prepc, ska)
		}
		if nexpc != nil && !nexpc.checkcbuf(ska) {
			nexpc.card.dp++
			pca.addBuff(nexpc, "leader")
			pca.addPlayerBuff(nexpc, ska)
		}
		pca.addPlayerBuff(pca, ska)
	case "lneck":
		pca.addBuff(pca, ska)
	case "jump":
		if !pca.checkcbuf(ska) {
			pca.addBuff(pca, ska)
			pca.addPlayerBuff(pca, ska)
		}
	case "airproof":
		pca.addBuff(pca, ska)
	case "diving":
		pca.addBuff(pca, ska)
	case "submersion":
		pca.addPlayerBuff(pca, ska)
	case "mistery":
		pca.addBuff(pca, ska)
	case "grow":
		pca.addBuff(pca, ska)
		pca.addPlayerBuff(pca, ska)
	case "power":
		pca.addBuff(pca, ska)
	case "copy":
		cp := copyBookStr{
			pca,
			[2]tcStr{
				{nil, ""},
				{nil, ""},
			},
		}
		copybook = append(copybook, cp)
		pca.addPlayerBuff(pca, ska)
		pca.OnboardSkillCopy()

	}

}

// Active PlayerBuff // excute when turns end (when player clicks next)
func (p *Players) buffDoActive() {

	for _, buffs := range playerBuff {
		buff := buffs.buff
		apl, apc := buffs.apl, buffs.apc
		tpl, tpc := buffs.tpl, buffs.tpc
		aturn := buffs.aturn

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
			case "submersion":
				if playerNow == apl && aturn+1 == apl.turn {
					apc.addBuff(apc, buff)
				} else if playerNow == apl && aturn+2 == apl.turn {
					apc.remBuff(buff)
					apc.removePlayerBuff(apc, buff)
					apc.addPlayerBuff(apc, buff)
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
					apc.addBuff(prepc, buff)
					apc.addPlayerBuff(prepc, buff)
				}
				if nexpc != nil && nexpc.card.dp != 0 && !nexpc.checkcbuf(buff) {
					nexpc.card.dp--
					apc.addBuff(nexpc, buff)
					apc.addPlayerBuff(nexpc, buff)
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
					apc.addBuff(prepc, buff)
					apc.addPlayerBuff(prepc, buff)
				}
				if nexpc != nil && !nexpc.checkcbuf(buff) {
					nexpc.card.dp++
					apc.addBuff(nexpc, buff)
					apc.addPlayerBuff(nexpc, buff)
				}
			case "charge":
				if playerNow == apl && playerNow.turn > aturn+1 {
					apc.remBuff(buff)
					apc.removePlayerBuff(apc, buff)
				} else if playerNow == apl {
					apc.used = true
				}
			case "grow":
				if playerNow == apl && playerNow.turn >= aturn+3 {
					bigbugCardN := deckCardstr{bigbugCard, theCardimg(bigbugCard), apc.pNum, 20, apc.deckNum, false, false, nil}
					nbn := apc.bNum
					apc.offCard()
					apc.putCard(&bigbugCardN, nbn)
					deckCard[apc.pNum][apc.deckNum].cardOn = true
					//apc.removePlayerBuff(apc, buff)
				}
			case "copy":
				//apc.copySkillCheck()
				apc.OnboardSkillCopy()

			}
		}
	}
}
