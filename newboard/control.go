package newboard

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) touchControl() {

	g.touches = inpututil.AppendJustPressedTouchIDs(g.touches[:0])

	for _, t := range g.touches {
		x, y := ebiten.TouchPosition(t)
		mpos.mp, mpos.mc = mousePosition(x, y)
		dn := 4*mpos.mp + mpos.mc

		if menu {
			mpos.mp = mousePositionNG(x, y)
			go menuselecting(mpos.mp)
		} else {
			controler(mpos.mp, mpos.mc, dn)
		}

		// else if mpos.mp == 5 {
		// 	if mpos.mc == 5 {
		// 		mcc5()
		// 	} else if mpos.mc == 6 && !menu {
		// 		mcc6()
		// 	}
		// } else if mpos.mp != 8 && mpos.mc != 8 && !cardS.sel && !pickBools && !dealBool {
		// 	attackControl(mpos.mp, mpos.mc)
		// } else if mpos.mp != 8 && mpos.mc != 8 && cardS.sel && 4*mpos.mp+mpos.mc <= 9 && !deckCard[playerNow.pn][dn].cardOn && !pickBools && !dealBool {
		// 	selectControl(mpos.mp, mpos.mc)
		// } else if mpos.mp != 8 && mpos.mc != 8 && pickBools {
		// 	pickControl(mpos.mp, mpos.mc)
		// } else if mpos.mp != 8 && mpos.mc != 8 && dealBool {
		// 	dealControl(mpos.mp, mpos.mc)
		// }
	}
}

func control() {

	if inpututil.IsKeyJustReleased(ebiten.KeyTab) {
		mcc6()
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {

		mpos.mp, mpos.mc = mousePosition(ebiten.CursorPosition())
		dn := 4*mpos.mp + mpos.mc

		if menu {
			mpos.mp = mousePositionNG(ebiten.CursorPosition())
			go menuselecting(mpos.mp)
		} else {
			controler(mpos.mp, mpos.mc, dn)
		}

		// if menu {
		// 	mpos.mp = mousePositionNG(ebiten.CursorPosition())
		// 	go menuselecting(mpos.mp)
		// } else if mpos.mp == 5 {
		// 	if mpos.mc == 5 {
		// 		mcc5()
		// 	} else if mpos.mc == 6 {
		// 		mcc6()
		// 	}
		// } else if mpos.mp != 8 && mpos.mc != 8 && !cardS.sel && !pickBools && !dealBool {
		// 	attackControl(mpos.mp, mpos.mc)
		// } else if mpos.mp != 8 && mpos.mc != 8 && cardS.sel && 4*mpos.mp+mpos.mc <= 9 && !deckCard[playerNow.pn][dn].cardOn && !pickBools && !dealBool {
		// 	selectControl(mpos.mp, mpos.mc)
		// } else if mpos.mp != 8 && mpos.mc != 8 && pickBools {
		// 	pickControl(mpos.mp, mpos.mc)
		// } else if mpos.mp != 8 && mpos.mc != 8 && dealBool {
		// 	dealControl(mpos.mp, mpos.mc)
		// }
	}
}

func controler(mp, mc, dn int) {
	if mp == 5 {
		if mc == 5 {
			mcc5()
		} else if mc == 6 {
			mcc6()
		}
	} else if mp != 8 && mc != 8 && !cardS.sel && !pickBools && !dealBool {
		attackControl(mp, mc)
	} else if mp != 8 && mc != 8 && cardS.sel && 4*mp+mc <= 9 && !deckCard[playerNow.pn][dn].cardOn && !pickBools && !dealBool {
		selectControl(mp, mc)
	} else if mp != 8 && mc != 8 && pickBools {
		pickControl(mp, mc)
	} else if mp != 8 && mc != 8 && dealBool {
		dealControl(mp, mc)
	}
}

func mcc5() {
	//fmt.Println("ji: ", *ji)
	touchedHome = true
	go func() {
		time.Sleep(200 * time.Millisecond)
		touchedHome = false
	}()

	if menu {
		menu = false
	} else if cardS.sel /*|| cruelbool || dealerFb*/ {
		cardS.sel = false
	} else if pickBools {
		pickBools = false
	} else if dealBool && dealerS {
		dealerS = false
	} else if dealBool {
		dealBool = false
	} else {
		menu = true
	}
}

func mcc6() {
	//fmt.Println(playerNow.pn)
	touchedNext = true
	go func() {
		time.Sleep(200 * time.Millisecond)
		touchedNext = false
	}()

	// whenever player click next, new channel function start
	// if previous turnChannel is closed, open new turnChannel.
	// if previous turnChannel is not closed, close opened channel(previous channel function will be end) and open new channel.
	go func() {
		select {
		case <-turnChan:
		default:
		}
		turnChan <- ""
	}()

	if pickBools {
		go pickNumber(ji)
	}

	if dealBool && !dealerS {
		dealerCardChange()
		dealBool = false
		return
	}

	if !cardS.sel && !pickBools && !dealBool && !dealerS {
		playerNow.buffDoActive()

		playerNow.pMoney = playerNow.pMoney + 3
		bmsg = "공격자 카드를 선택하세요"
		for i := 0; i < 10; i++ {
			playingCard[playerNow.pn][i].used = false
		}

		//fmt.Println(playerNow, playerNow.pn, playerNow.pMoney)
		playerNow = playerNow.next
		playerNow.turn++
		if playerNow.pHP <= 0 {
			playerNow = playerNow.next
			playerNow.turn++
		}
		//fmt.Println(playerNow, playerNow.pn, playerNow.pMoney)

		playerNow.buffDoPassive()

		//strL(cardBoard[0][0].card.debuf)
	}
}

func attackControl(mp, mc int) {

	// card selecting page call
	if playerNow.pn == mp && cardBoard[mp][mc].card == nil {
		attacker = nil
		cardS.sel = true
		playerNow.bn = mc
		bmsg = "추가할 카드를 선택하세요"
		amsg = fmt.Sprintf("0) cardselecting %v , %v\n", mp, mc)
		if cruel.bools {
			cruel = cruelStr{nil, 20, false}
		}

		//
		// cruel selcting
	} else if playerNow.pn == mp && cardBoard[mp][mc].card != nil && cruel.bools {
		pc := cruel.pc
		bn := cruel.bn
		dc := &deckCard[pc.pNum][pc.deckNum]
		attacker = nil

		if cardBoard[mp][mc].card.card.price >= 2 {
			cardBoard[mp][mc].card.offCard()
			pc.putCard(dc, bn)

			playerNow.bn = 20
			playerNow.pMoney -= dc.card.price
			bmsg = "공격자 카드를 선택하세요"

			cruel = cruelStr{nil, 20, false}
		} else {
			bmsg = "$2 이상의 희생카드를 선택하세요"
		}
		return

		//
		// attacker selcting
	} else if playerNow.pn == mp && cardBoard[mp][mc].card != nil && !cardBoard[mp][mc].card.used && attacker == nil {
		attacker = cardBoard[mp][mc].card
		if playerNow.turn == 1 && attacker.card.skill != "advance" {
			attacker = nil
			bmsg = "첫 턴은 공격할 수 없습니다."
			return
		}

		bmsg = "공격대상을 선택하세요"
		amsg = fmt.Sprintf("attack() is called. mp : %v , mc : %v\nattacker:%v , target:%v\n", mp, mc, attacker.card.name, target)

		//
		// attacker canceling
	} else if playerNow.pn == mp && attacker != nil {
		bmsg = "공격자 카드를 선택하세요"
		amsg = fmt.Sprintf("attacker is canceled. mp : %v , mc : %v\nattacker:%v , target:%v\n", mp, mc, attacker.card.name, target)
		attacker = nil

		//
		// target selcting if attacker was selected
	} else if playerNow.pn != mp && cardBoard[mp][mc].card != nil && attacker != nil && int(attacker.bNum) == mc {
		target = cardBoard[mp][mc].card
		bmsg = "공격자 카드를 선택하세요"
		amsg = fmt.Sprintf("attack() is called. mp : %v , mc : %v\nattacker:%v , target:%v\n", mp, mc, attacker.card.name, target.card.name)
		attacker.attackCard(target)
		attacker, target = nil, nil

		//
		// target selcting if attacker skill is horn and selected
	} else if playerNow.pn != mp && cardBoard[mp][mc].card != nil && attacker != nil && attacker.card.skill == "horn" && (preB(int(attacker.bNum)) == mc || nexB(int(attacker.bNum)) == mc) {
		target = cardBoard[mp][mc].card
		bmsg = "공격자 카드를 선택하세요"
		amsg = fmt.Sprintf("attack() is called. mp : %v , mc : %v\nattacker:%v , target:%v\n", mp, mc, attacker.card.name, target.card.name)
		attacker.attackCard(target)
		attacker, target = nil, nil

		//
		// Castle attacking if attacker was selected and target card is empty
	} else if playerNow.pn != mp && cardBoard[mp][mc].card == nil && attacker != nil && int(attacker.bNum) == mc && player[mp].pHP > 0 {
		targetPlayer := &player[mp]
		bmsg = "공격자 카드를 선택하세요"
		amsg = fmt.Sprintf("attackCastle() is called. mp : %v , mc : %v\nattacker:%v , target:%v\n", mp, mc, attacker.card.name, target)
		attacker.attackCastle(targetPlayer)
		attacker, target = nil, nil

		//
		// Castle attacking if attacker skill is horn and selected, and target card is empty
	} else if playerNow.pn != mp && cardBoard[mp][mc].card == nil && attacker != nil && attacker.card.skill == "horn" && (preB(int(attacker.bNum)) == mc || nexB(int(attacker.bNum)) == mc) && player[mp].pHP > 0 {
		targetPlayer := &player[mp]
		bmsg = "공격자 카드를 선택하세요"
		amsg = fmt.Sprintf("attackCastle() is called. mp : %v , mc : %v\nattacker:%v , target:%v\n", mp, mc, attacker.card.name, target)
		attacker.attackCastle(targetPlayer)
		attacker, target = nil, nil
	}
}

func selectControl(mp, mc int) {
	cardNum := 4*mp + mc
	pc := &playingCard[playerNow.pn][cardNum]
	dc := &deckCard[playerNow.pn][cardNum]
	//fmt.Println(cardNum, pc.pNum, pc.deckNum, dc.pNum, dc.deckNum)

	for _, debuffs := range dc.debuf {
		if debuffs == "incision" {
			bmsg = "자르기에 죽은 카드입니다.\n추가할 카드를 선택하세요."
			return
		}
	}

	if cardS.skill != "" {
		switch cardS.skill {
		case "eggC":
			playerNow.pMoney += dc.card.price
			cardS.skill = ""
		case "company":
			if dc.card.skill == "company" {
				bmsg = "같은카드를 선택할 수 없습니다.\n추가할 카드를 선택하세요."
				return
			}
			playerNow.pMoney += dc.card.price
			playerNow.bn = cardS.bn
		}
	}

	if dc.card.skill == "cruelty" && playerNow.pMoney >= dc.card.price {
		*pc = *dc
		pc.cruelty(playerNow.bn)
		return
	}

	if playerNow.pMoney >= dc.card.price {
		pc.putCard(dc, playerNow.bn)
		pc.skillDoPassive()

		playerNow.bn = 20
		cardS.sel = false
		playerNow.pMoney -= deckCard[playerNow.pn][cardNum].card.price
		bmsg = "공격자 카드를 선택하세요"
	} else {
		bmsg = "골드가 부족합니다.\n추가할 카드를 선택하세요."
		return
	}

	if cardS.skill != "" {
		switch cardS.skill {
		case "company":
			playerNow = &player[cardS.pn]
			cardS.bn, cardS.pn = 20, 20
			cardS.skill = ""
		}

	}

}

func pickControl(mp, mc int) {

	if mp == 0 && mc == 0 {
		pickplayer = 0
		jipick = 0
	} else if mp == 1 && mc == 0 {
		pickplayer = 1
		jipick = 0
	} else if mp == 2 && mc == 0 {
		if twoplay {
			return
		}
		pickplayer = 2
		jipick = 0
	} else if mc == 2 {
		go pickedC(mc)
		pickChan <- "" // wait until pickedC() ends and closes pickChan channel
		pickNumber(ji)

	} else if mc == 3 {
		go pickedC(mc)
		pickChan <- ""
		pickNumber(ji)
	}

}

func dealControl(mp, mc int) {
	if !dealerS {
		if mc == 0 || mc == 2 {
			dOffer = mp
			dealP = dOffer
			bmsg = "거래할 카드를 선택하세요"
			dealerS = true

		} else if mc == 1 || mc == 3 {
			dReciever = mp
			dealP = dReciever
			bmsg = "거래할 카드를 선택하세요"
			dealerS = true
		}
	} else {
		if dealP == dOffer {
			deckNum := 4*mp + mc
			dOfferD = deckNum
			dealerS = false
		} else if dealP == dReciever {
			deckNum := 4*mp + mc
			dRecieverD = deckNum
			dealerS = false
		}

	}
}
