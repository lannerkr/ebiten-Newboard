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
// func (pca *deckCardstr) skillDo(pct *deckCardstr) {
// 	ska := pca.card.skill
// 	skt := pct.card.skill
// 	//pa := &player[pca.pNum]
// 	pt := &player[pct.pNum]
// 	sktcp := false

// 	switch ska {
// 	case "poison":
// 		if !pct.checkcbuf(ska) {
// 			pca.addBuff(pct)
// 			pca.addPlayerBuff(pct)
// 		}
// 	case "vampire":
// 		if pct.card.hp <= pca.card.dp {
// 			go func() {
// 				<-attackChan
// 				if !pct.cardOn {
// 					pca.card.hp += 3
// 				}
// 			}()
// 		}
// 	case "flying":
// 		pca.flyingFunc(pct)
// 		// for _, buff := range pct.debuf {
// 		// 	if buff == "lneck" {
// 		// 		return
// 		// 	}
// 		// }
// 		// pct.card.hp += pca.card.dp
// 		// pt.pHP -= pca.card.dp

// 	case "tentacle":
// 		if !pct.checkcbuf(ska) {
// 			pca.addBuff(pct)
// 			pca.addPlayerBuff(pct)
// 		}

// 		go func() {
// 			var killed bool = false
// 			ptdc := &deckCard[pct.pNum][pct.deckNum]
// 			pca.addBuff(ptdc)
// 			for {
// 				//wait for turnChannel (when player click next, turnChannel will be opened "control.go")
// 				<-turnChan

// 				if !killed {
// 					if pct.cardOn {
// 						continue
// 					} else {
// 						killed = true
// 					}
// 				} else if killed {
// 					if pct.cardOn && ptdc.checkcbuf(ska) {
// 						ptdc.remBuff(ska)
// 						pca.addBuff(pct)
// 						pca.addPlayerBuff(pct)
// 						return
// 					}
// 				}
// 			}
// 		}()
// 	case "threat":
// 		if !pct.checkcbuf(ska) {
// 			pct.card.dp--
// 			pca.addBuff(pct)
// 		}
// 	case "fire":
// 		pt.pHP -= pca.card.dp
// 	case "strategy":
// 		pNumn := nexP(pct.pNum)
// 		if pNumn == pca.pNum {
// 			pNumn = nexP(pNumn)
// 		}
// 		if pNumn != pct.pNum && player[pNumn].pHP > 0 {
// 			if cardBoard[pNumn][pct.bNum].card != nil {
// 				pctn := cardBoard[pNumn][pct.bNum].card
// 				pctn.card.hp = pctn.card.hp - pca.card.dp
// 				if pctn.card.hp <= 0 {
// 					pctn.offCard()
// 				}
// 			} else {
// 				pt.pHP -= pca.card.dp
// 			}
// 		}
// 	case "dolphin":
// 		pca.addBuff(pct)
// 		pca.addPlayerBuff(pct)
// 	case "sonic":
// 		pca.addBuff(pct)
// 		pca.addPlayerBuff(pct)
// 	case "selfdestruct":
// 		go func() {
// 			<-attackChan

// 			pca.offCard()
// 		}()
// 	case "mistaken":
// 		go func() {
// 			<-attackChan

// 			if !pct.cardOn {
// 				pcapn, pcadn := pca.pNum, pca.deckNum
// 				pctpn, pctdn := pct.pNum, pct.deckNum
// 				pca.offCard()

// 				pcaDeckC := deckCard[pcapn][pcadn]
// 				deckCard[pcapn][pcadn] = deckCard[pctpn][pctdn]
// 				deckCard[pcapn][pcadn].pNum = pcapn
// 				deckCard[pcapn][pcadn].deckNum = pcadn

// 				deckCard[pctpn][pctdn] = pcaDeckC
// 				deckCard[pctpn][pctdn].deckNum = pctdn
// 				deckCard[pctpn][pctdn].pNum = pctpn

// 			}
// 		}()
// 	case "climb":
// 		go func() {
// 			<-attackChan

// 			if !pct.cardOn {
// 				if cardBoard[pct.pNum][preB(pct.bNum)].card != nil {
// 					pca.attackCard(cardBoard[pct.pNum][preB(pct.bNum)].card)
// 				}
// 				if cardBoard[pct.pNum][nexB(pct.bNum)].card != nil {
// 					pca.attackCard(cardBoard[pct.pNum][nexB(pct.bNum)].card)
// 				}
// 			}
// 		}()
// 	case "rush":
// 		go func() {
// 			<-attackChan
// 			if pct.cardOn && pca.cardOn {
// 				pca.card.hp--
// 				if pca.card.hp <= 0 {
// 					pca.offCard()
// 				}
// 			}
// 		}()
// 	case "incision":
// 		go func() {
// 			<-attackChan
// 			if !pct.cardOn {
// 				pdc := &deckCard[pct.pNum][pct.deckNum]
// 				pdc.debuf = append(pdc.debuf, ska)
// 				pca.addPlayerBuff(pct)
// 			}
// 		}()
// 	case "trans":
// 		goto AFTERSKT
// 	case "huge":
// 		if npct := cardBoard[pct.pNum][nexB(pct.bNum)].card; pct.bNum != 3 && pca.bNum == pct.bNum && npct != nil {
// 			pca.attackCard(npct)
// 		}
// 	case "wield":
// 		if cardBoard[pct.pNum][nexB(pct.bNum)].card == nil {
// 			cardBoard[pct.pNum][pct.bNum].card = nil
// 			cardBoard[pct.pNum][nexB(pct.bNum)].card = pct
// 			pct.bNum = nexB(pct.bNum)
// 		}
// 	case "charge":
// 		pca.addBuff(pca)
// 		pca.addPlayerBuff(pca)
// 	case "runover":
// 		go func() {
// 			<-attackChan
// 			if pct.cardOn && pca.cardOn {
// 				npn := nexP(pct.pNum)
// 				if npn == pca.pNum {
// 					npn = nexP(npn)
// 				}

// 				if cardBoard[npn][pct.bNum].card == nil {
// 					pca.addBuff(pct)
// 					pca.addPlayerBuff(pct)

// 					cardBoard[pct.pNum][pct.bNum].card = nil
// 					cardBoard[npn][pct.bNum].card = pct
// 					pct.pNum = npn

// 				}
// 			}
// 		}()

// 	}

// SKTCOPY:
// 	switch skt {
// 	case "revenge":
// 		fmt.Println("021]")
// 		pca.card.hp--
// 		if pca.card.hp <= 0 {
// 			pca.offCard()
// 		}
// 	case "wit":
// 		fmt.Println("022]")
// 		go func(pctCard deckCardstr, plb []buffStr) {
// 			<-attackChan

// 			if !pct.cardOn && cardBoard[pctCard.pNum][preB(pctCard.bNum)].card == nil {
// 				pct.putCard(&pctCard, preB(pctCard.bNum))
// 				deckCard[pct.pNum][pct.deckNum].bNum = preB(pctCard.bNum)
// 				deckCard[pct.pNum][pct.deckNum].cardOn = true

// 				playerBuff = plb
// 			}
// 		}(*pct, playerBuff)
// 	case "dolphin":
// 		fmt.Println("023]")
// 		pct.addPlayerBuff(pct)
// 	case "water":
// 		fmt.Println("024]")
// 		pct.addPlayerBuff(pct)
// 	case "company":
// 		fmt.Println("025]")
// 		go func() {
// 			<-attackChan
// 			if !pct.cardOn {
// 				cardS.pn = playerNow.pn
// 				playerNow = &player[pct.pNum]
// 				cardS.bn = pct.bNum
// 				cardS.sel = true
// 				cardS.skill = "company"

// 				bmsg = "추가할 카드를 선택하세요"
// 			}
// 		}()
// 	case "shell":
// 		fmt.Println("026]")
// 		go func() {
// 			<-attackChan
// 			if pct.cardOn {
// 				pct.addBuff(pct)
// 			}
// 		}()
// 	case "airproof":
// 		fmt.Println("027]")
// 		if pca.card.dp != 2 {
// 			if pca.card.dp > 2 {
// 				pct.card.hp += pca.card.dp - 2
// 			} else {
// 				pct.card.hp -= 2 - pca.card.dp
// 			}
// 		}
// 	case "copy":
// 		fmt.Println("000]")
// 		if ppc := cardBoard[preP(pct.pNum)][pct.bNum].card; ppc != nil && !sktcp {

// 			skt = ppc.card.skill
// 			pct.card.skill = skt
// 			fmt.Println("001]", pct.card.skill)
// 			sktcp = true
// 			goto SKTCOPY
// 		} else if npc := cardBoard[nexP(pct.pNum)][pct.bNum].card; npc != nil {
// 			fmt.Println("002]")
// 			skt = npc.card.skill
// 			pct.card.skill = skt
// 			sktcp = false
// 			goto SKTCOPY
// 		} else {
// 			fmt.Println("003]")
// 			pct.card.skill = "copy"
// 			sktcp = false
// 		}
// 		//pca.copySkillDoCopy(pct)

// 	}
// 	if sktcp {
// 		fmt.Println("004]")
// 		skt = "copy"
// 		goto SKTCOPY
// 	}
// AFTERSKT:
// }
func (pca *deckCardstr) skillDoA(pct *deckCardstr) {
	ska := pca.card.skill
	//skt := pct.card.skill
	//pa := &player[pca.pNum]
	pt := &player[pct.pNum]
	//sktcp := false

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
		pca.flyingFunc(pct)
		if pca.used {
			return
		}
		// for _, buff := range pct.debuf {
		// 	if buff == "lneck" {
		// 		return
		// 	}
		// }
		// pct.card.hp += pca.card.dp
		// pt.pHP -= pca.card.dp

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
		pca.addBuff(pca)
		pca.addPlayerBuff(pca)
	case "runover":
		go func() {
			<-attackChan
			if pct.cardOn && pca.cardOn {
				npn := nexP(pct.pNum)
				if npn == pca.pNum {
					npn = nexP(npn)
				}

				if cardBoard[npn][pct.bNum].card == nil {
					pca.addBuff(pct)
					pca.addPlayerBuff(pct)

					cardBoard[pct.pNum][pct.bNum].card = nil
					cardBoard[npn][pct.bNum].card = pct
					pct.pNum = npn

				}
			}
		}()

	}

	pca.skillDoT(pct)
}
func (pca *deckCardstr) skillDoT(pct *deckCardstr) {
	skt := pct.card.skill
	//pa := &player[pca.pNum]
	//pt := &player[pct.pNum]
	sktcp := false

SKTCOPY:
	switch skt {
	case "revenge":
		fmt.Println("021]")
		pca.card.hp--
		if pca.card.hp <= 0 {
			pca.offCard()
		}
	case "wit":
		fmt.Println("022]")
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
		fmt.Println("023]")
		pct.addPlayerBuff(pct)
	case "water":
		fmt.Println("024]")
		pct.addPlayerBuff(pct)
	case "company":
		fmt.Println("025]")
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
		fmt.Println("026]")
		go func() {
			<-attackChan
			if pct.cardOn {
				pct.addBuff(pct)
			}
		}()
	case "airproof":
		fmt.Println("027]")
		if pca.card.dp != 2 {
			if pca.card.dp > 2 {
				pct.card.hp += pca.card.dp - 2
			} else {
				pct.card.hp -= 2 - pca.card.dp
			}
		}
	case "copy":
		fmt.Println("000]")
		if ppc := cardBoard[preP(pct.pNum)][pct.bNum].card; ppc != nil && !sktcp {

			skt = ppc.card.skill
			pct.card.skill = skt
			fmt.Println("001]", pct.card.skill)
			sktcp = true
			goto SKTCOPY
		} else if npc := cardBoard[nexP(pct.pNum)][pct.bNum].card; npc != nil {
			fmt.Println("002]")
			skt = npc.card.skill
			pct.card.skill = skt
			sktcp = false
			goto SKTCOPY
		} else {
			fmt.Println("003]")
			pct.card.skill = "copy"
			sktcp = false
		}
		//pca.copySkillDoCopy(pct)

	}
	if sktcp {
		fmt.Println("004]")
		skt = "copy"
		goto SKTCOPY
	}
	// AFTERSKT:
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
		pca.addBuff(pca)
		pca.addPlayerBuff(pca)

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
	case "submersion":
		pca.addPlayerBuff(pca)
	case "mistery":
		pca.addBuff(pca)
	case "grow":
		pca.addBuff(pca)
		pca.addPlayerBuff(pca)
	case "power":
		pca.addBuff(pca)
	case "copy":
		cp := copyBookStr{
			pca,
			[2]tcStr{
				{nil, ""},
				{nil, ""},
			},
		}
		copybook = append(copybook, cp)
		pca.addPlayerBuff(pca)
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
					apc.addBuff(apc)
				} else if playerNow == apl && aturn+2 == apl.turn {
					apc.remBuff(buff)
					apc.removePlayerBuff(apc, buff)
					apc.addPlayerBuff(apc)
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
	ppc := cardBoard[preP(pca.pNum)][pca.bNum].card
	npc := cardBoard[nexP(pca.pNum)][pca.bNum].card

	for i, cpb := range copybook {
		if cpb.pc == pca {
			pc := ppc
			for k := 0; k < 2; k++ {
				if cpb.tc[k].tpc != pc {
					if cpb.tc[k].buff != "" {
						pca.remBuff(cpb.tc[k].buff)
					}
					if pc == nil {
						cpb.tc[k].tpc = nil
					} else {
						copybook[i].tc[k].tpc = pc
						pca.card.skill = pc.card.skill
						bl0 := len(pca.debuf)
						pca.skillDoPassive()
						bl1 := len(pca.debuf)
						if bl0 != bl1 {
							copybook[i].tc[k].buff = pca.debuf[bl1-1]
						}
						pca.card.skill = "copy"
					}
				}
				pc = npc
			}
		}
	}

	// if ppc != nil && !pca.checkcbuf(ppc.card.skill) {
	// 	pca.card.skill = ppc.card.skill
	// 	pca.skillDoPassive()
	// 	pca.card.skill = "copy"

	// }

	// if npc != nil && !pca.checkcbuf(npc.card.skill) {
	// 	pca.card.skill = npc.card.skill
	// 	pca.skillDoPassive()
	// 	pca.card.skill = "copy"
	// }
}

// func (pc *deckCardstr) copySkillCheck() {

// 	var pskill, nskill string
// 	if ppc := cardBoard[preP(pc.pNum)][pc.bNum].card; ppc != nil {
// 		pskill = ppc.card.skill
// 	}
// 	if npc := cardBoard[nexP(pc.pNum)][pc.bNum].card; npc != nil {
// 		nskill = npc.card.skill
// 	}

// RECHECK:
// 	for _, buff := range pc.debuf {
// 		if buff != pskill && buff != nskill {
// 			pc.remBuff(buff)
// 			goto RECHECK
// 		}
// 	}
// }

func (pca *deckCardstr) copySkillDoCopy(pct *deckCardstr) {

}
