package newboard

func preP(pn int) int {
	var prepn int
	if pn == 0 {
		prepn = 2
	} else {
		prepn = pn - 1
	}
	return prepn
}
func nexP(pn int) int {
	var nexpn int
	if pn == 2 {
		nexpn = 0
	} else {
		nexpn = pn + 1
	}
	return nexpn
}
func preB(bn int) int {
	var prebn int
	if bn == 0 {
		prebn = 3
	} else {
		prebn = bn - 1
	}
	return prebn
}
func nexB(bn int) int {
	var nexbn int
	if bn == 3 {
		nexbn = 0
	} else {
		nexbn = bn + 1
	}
	return nexbn

}

func (pl *Players) deckRemainCard() int {
	rcard := 0
	for d := 0; d < 10; d++ {
		if !deckCard[pl.pn][d].cardOn && deckCard[pl.pn][d].bNum != 99 {
			rcard++
		}
	}

	return rcard
}
