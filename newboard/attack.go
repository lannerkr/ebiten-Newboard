package newboard

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func (pca *deckCardstr) attackCard(pct *deckCardstr) {
	for _, buff := range pct.debuf {
		if buff == "shell" {
			pct.remBuff(buff)
			return
		} else if buff == "survive" && (pct.card.hp-pca.card.dp) < 1 && pca.card.skill != "trans" {
			pct.card.hp += 1 - (pct.card.hp - pca.card.dp)
		} else if buff == "submersion" || (buff == "power" && (pca.card.hp <= 1 || pca.card.dp <= 1)) {
			return
		} else if buff == "mistery" {
			rand.Seed(time.Now().UnixNano())
			misteryCheck := rand.Float32() < 0.5
			if misteryCheck {
				//fmt.Println(misteryCheck)
				continue
			} else {
				//fmt.Println(misteryCheck)
				return
			}
		}

	}

	pca.skillDo(pct)

	if pca.card.skill == "flying" {
		for _, buff := range pct.debuf {
			if buff == "lneck" {
				goto LNECK
			}
		}
		return
	}
LNECK:

	if pct.card.skill == "airproof" && pca.card.dp != 2 {
		if pca.card.dp > 2 {
			pct.card.hp += pca.card.dp - 2
		} else {
			pct.card.hp -= 2 - pca.card.dp
		}
	} else if pct.card.skill == "diving" && pct.card.hp > pca.card.dp {
		goto DIVE
	}

	pct.card.hp = pct.card.hp - pca.card.dp
	if pct.card.hp <= 0 {
		pct.offCard()
	}
DIVE:
	pca.used = true

	// 공격 시 시작되는 스킬 function pca.skillDo(pct) 중, 공격이 끝난 후 적용되는 스킬의 go func() 에게 전달하기 위해 attackChan을 open
	// close 되지 못한 attackChan이 있을 경우, 먼저 attackChan을 close
	go func() {
		select {
		case <-attackChan:
		default:
		}
		attackChan <- ""
	}()

	// attackChan이 전달 될 수 있도록 100 msec 대기 후,
	// attackChan이 open 상태이면 close 후 아니면 그냥 function을 종료
	time.Sleep(100 * time.Millisecond)
	select {
	case <-attackChan:
	default:
	}

}

func (pca *deckCardstr) attackCastle(pt *Players) {
	pca.skillDoCastle(pt)

	pt.pHP = pt.pHP - pca.card.dp
	pca.used = true
}

func (pca *deckCardstr) addBuff(pct *deckCardstr) {
	buff := pca.card.skill
	if buff == "dolphin" && pca != pct {
		buff = "sonic"
	} else if buff == "dolphin" && pca == pct {
		buff = "water"
	} else if buff == "totem" {
		if pca != pct {
			buff = "leader"
		} else {
			buff = "lneck"
		}
	}

	pct.debuf = append(pct.debuf, buff)
}

func (pc *deckCardstr) remBuff(s string) {

	switch s {
	case "leader":
		pc.card.dp--
	case "cute":
		pc.card.dp++
	}

	if len(pc.debuf) == 1 {
		pc.debuf = nil
	} else {
	DBUFF:
		for i, buff := range pc.debuf {
			if buff == s {
				pc.debuf[i] = pc.debuf[len(pc.debuf)-1]
				pc.debuf = pc.debuf[:len(pc.debuf)-1]
				goto DBUFF
			}
		}
	}
}

func (dc *deckCardstr) checkcbuf(d string) bool {
	str := fmt.Sprintf("%v", dc.debuf)
	return strings.Contains(str, d)
}

func (pca *deckCardstr) addPlayerBuff(pct *deckCardstr) {
	pla := &player[pca.pNum]
	plt := &player[pct.pNum]
	buff := pca.card.skill
	if buff == "dolphin" && pca != pct {
		buff = "sonic"
	} else if buff == "dolphin" && pca == pct {
		buff = "water"
	} else if buff == "totem" {
		buff = "leader"
	}
	pBuff := buffStr{pla, pla.turn, pca, plt, plt.turn, pct, buff}
	playerBuff = append(playerBuff, pBuff)
	fmt.Println("addPlayerBuff", pca.pNum, pca.card.skill)
}

// (skill은 "") 그리고 (apc는 buff apc 이거나 tcp는 buff tpc) 인 경우,
// (skill은 buff 스킬) 그리고 (apc는 buff apc 이고 tcp는 buff tpc) 인 경우,
// 해당 buff 슬롯에 마지막 buff슬롯을 카피한 후 마지만 슬롯의 buff를 삭제
// playerBuff를 처음부터 다시 검사
func (apc *deckCardstr) removePlayerBuff(tpc *deckCardstr, skill string) {
START:

	for bi, buffs := range playerBuff {
		if skill == "" && (buffs.apc == apc || buffs.tpc == apc) {
			if buffs.buff == "incision" || ((buffs.buff == "poison" || buffs.buff == "tentacle" || buffs.buff == "runover") && buffs.apc == apc) {
				continue
			} else if buffs.apc != buffs.tpc {
				buffs.tpc.remBuff(buffs.buff)
			}
			if buffs.buff == "runover" {
				tpc.pNum = buffs.tpl.pn
				deckCard[tpc.pNum][tpc.deckNum].cardOn = false
			}
			playerBuff[bi] = playerBuff[len(playerBuff)-1]
			playerBuff = playerBuff[:len(playerBuff)-1]
			goto START
		}
		if skill == buffs.buff && (buffs.apc == apc && buffs.tpc == tpc) {
			playerBuff[bi] = playerBuff[len(playerBuff)-1]
			playerBuff = playerBuff[:len(playerBuff)-1]
			goto START
		}
	}
}
