package newboard

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type boardStr struct {
	bx float64
	by float64
}

type controlStr struct {
	sx int
	sy int
	ex int
	ey int
}

type mousePos struct {
	mp, mc int
}

type Cardstr struct {
	name  string
	price int
	img   *ebiten.Image // 카드 이름
	char  string        // 속성
	cImg  *ebiten.Image // 능력 그림
	capa  string        // 능력
	hp    int
	dp    int
	skill string
}

type deckCardstr struct {
	card    Cardstr
	cardImg *ebiten.Image
	pNum    int // player number //player card로부터 deckcard 선택을 위해 사용
	bNum    int // card가 놓인 borad number
	deckNum int // player deck 의 넘버 [0 ~ 9] //player card로부터 deckcard 선택을 위해 사용

	cardOn bool     // 카드가 보드에서 사용중인지 확인
	used   bool     //공격 실행 여부 확인, 공격 사용 시 true로 변경
	debuf  []string //카드가 받는 debuf 표시
}

type cardBoardStr struct {
	card *deckCardstr
}

type Players struct {
	pHP    int
	pMoney int
	pn     int
	bn     int // card selecting을 위해 선택된 board number
	next   *Players
	turn   int
}

type buffStr struct {
	apl   *Players
	aturn int
	apc   *deckCardstr
	tpl   *Players
	tturn int
	tpc   *deckCardstr
	buff  string
}

type cardSelectStr struct {
	sel    bool
	skill  string
	pn, bn int
}

var (
	baw, bah int //= 320, 200	background image size
	bow, boh int //= 160, 240	board card image size

	cardw, cardh int     = ScreenWidth / 6, ScreenHeight * 4 / 16
	sepw, seph   float64 = float64(cardw) / 6, float64(cardh) / 4
	cardop       [3][4]*ebiten.DrawImageOptions

	bC [12]boardStr = [12]boardStr{
		{sepw, seph}, {2*sepw + float64(cardw), seph}, {3*sepw + 2*float64(cardw), seph}, {4*sepw + 3*float64(cardw), seph},
		{sepw, 2*seph + float64(cardh)}, {2*sepw + float64(cardw), 2*seph + float64(cardh)}, {3*sepw + 2*float64(cardw), 2*seph + float64(cardh)}, {4*sepw + 3*float64(cardw), 2*seph + float64(cardh)},
		{sepw, 3*seph + 2*float64(cardh)}, {2*sepw + float64(cardw), 3*seph + 2*float64(cardh)}, {3*sepw + 2*float64(cardw), 3*seph + 2*float64(cardh)}, {4*sepw + 3*float64(cardw), 3*seph + 2*float64(cardh)},
	}

	cardPos [12]controlStr = [12]controlStr{
		{int(bC[0].bx), int(bC[0].by), int(bC[0].bx) + cardw, int(bC[0].by) + cardh}, {int(bC[1].bx), int(bC[1].by), int(bC[1].bx) + cardw, int(bC[1].by) + cardh}, {int(bC[2].bx), int(bC[2].by), int(bC[2].bx) + cardw, int(bC[2].by) + cardh}, {int(bC[3].bx), int(bC[3].by), int(bC[3].bx) + cardw, int(bC[3].by) + cardh},
		{int(bC[4].bx), int(bC[4].by), int(bC[4].bx) + cardw, int(bC[4].by) + cardh}, {int(bC[5].bx), int(bC[5].by), int(bC[5].bx) + cardw, int(bC[5].by) + cardh}, {int(bC[6].bx), int(bC[6].by), int(bC[6].bx) + cardw, int(bC[6].by) + cardh}, {int(bC[7].bx), int(bC[7].by), int(bC[7].bx) + cardw, int(bC[7].by) + cardh},
		{int(bC[8].bx), int(bC[8].by), int(bC[8].bx) + cardw, int(bC[8].by) + cardh}, {int(bC[9].bx), int(bC[9].by), int(bC[9].bx) + cardw, int(bC[9].by) + cardh}, {int(bC[10].bx), int(bC[10].by), int(bC[10].bx) + cardw, int(bC[10].by) + cardh}, {int(bC[11].bx), int(bC[11].by), int(bC[11].bx) + cardw, int(bC[11].by) + cardh},
	}

	buttonPos [2]controlStr = [2]controlStr{
		{5*int(sepw) + 4*cardw, ScreenHeight - 2*int(seph), 5*int(sepw) + 4*cardw + 80, ScreenHeight - 2*int(seph) + 64},
		{6*int(sepw) + 4*cardw + 80, ScreenHeight - 2*int(seph), 6*int(sepw) + 4*cardw + 160, ScreenHeight - 2*int(seph) + 64},
	}

	allCard     [cardTotal]Cardstr
	NallCard    [cardTotal]Cardstr
	deckCard    [3][10]deckCardstr
	playingCard [3][10]deckCardstr
	cardBoard   [3][4]cardBoardStr

	player    [3]Players
	playerNow *Players

	playerBuff []buffStr

	mpos mousePos

	cardS = cardSelectStr{false, "", 20, 20}

	turnChan   = make(chan string)
	attackChan = make(chan string)

	strLocal map[string]string

	// plimg, plused [3]*ebiten.Image
)

func playerinit() {
	p := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			cardop[i][j] = &ebiten.DrawImageOptions{}
			cardop[i][j].GeoM.Translate(bC[p].bx, bC[p].by)
			p++
		}
	}

	player = [3]Players{{20, 5, 0, 20, &player[1], 1}, {20, 5, 1, 20, &player[2], 0}, {20, 5, 2, 20, &player[0], 0}}
	looseplayer = [3]bool{false, false, false}
	playerBuff = nil
	playerNow = &player[0]
	cardS.sel = false
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			cardBoard[i][j].card = nil
		}
		for j := 0; j < 10; j++ {
			playingCard[i][j].cardOn = false
		}
	}
	if twoplay {
		for j := 0; j < 4; j++ {
			playingCard[2][j].putCard(&deckCard[2][j], j)
		}
	}
	//fmt.Println(player[0].pn, player[1].pn, player[2].pn)
	//fmt.Println(playerNow.pn)
}

func init() {

	// shuffle cards
	// deckCard[0][7] = deckCardstr{allCard[0], theCardimg(allCard[0]), 0, 20, 7, false, false, ""}
	// deckCard[1][3] = deckCardstr{allCard[8], theCardimg(allCard[8]), 1, 20, 3, false, false, ""}
	// deckCard[2][5] = deckCardstr{allCard[21], theCardimg(allCard[21]), 2, 20, 5, false, false, ""}

	// playing card choose

	// playingCard[0][0].putCard(&deckCard[0][7], 0)
	// playingCard[0][1].putCard(&deckCard[0][5], 1)

	// playingCard[1][0].putCard(&deckCard[1][4], 0)
	// playingCard[1][2].putCard(&deckCard[1][2], 2)

	// playingCard[2][1].putCard(&deckCard[2][1], 1)
	// playingCard[2][3].putCard(&deckCard[2][9], 3)

	//playingCard[0][0].offCard()

}

func (pc *deckCardstr) putCard(dc *deckCardstr, bn int) {
	dc.cardOn = true
	dc.bNum = bn
	*pc = *dc
	pc.boardC().card = pc

}

func (pc *deckCardstr) offCard() {
	pc.cardOn = false
	pc.boardC().card = nil
	pn := pc.pNum
	dn := pc.deckNum
	deckCard[pn][dn].cardOn = false
	deckCard[pn][dn].bNum = 20

	if elimode {
		deckCard[pn][dn].bNum = 99
	}

	pc.removePlayerBuff(pc, "")
}

func (pc *deckCardstr) boardC() *cardBoardStr {
	pn, bn := pc.pNum, pc.bNum
	cardB := &cardBoard[pn][bn]
	return cardB

}
