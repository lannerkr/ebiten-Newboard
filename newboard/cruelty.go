package newboard

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
