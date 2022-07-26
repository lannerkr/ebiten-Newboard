package newboard

func (pca *deckCardstr) attpc(pct *deckCardstr) {
	pct.card.hp -= pca.card.dp
	pca.used = true

}
func (pca *deckCardstr) attpl(pct *deckCardstr) {
	player[pct.pNum].pHP -= pca.card.dp
	pca.used = true

}

func (pca *deckCardstr) flyingFunc(pct *deckCardstr) {
	for _, buff := range pct.debuf {
		if buff == "lneck" {
			return
		}
	}
	pca.attpl(pct)
}
