package newboard

import (
	"math/rand"
	"time"
)

func cardimport() {

	nilcard := images["nil"]

	blowfish := images["복어"]
	mosquito := images["모기"]
	eggv := images["알"]
	raptor := images["랩터"]
	chicken := images["닭"]
	crow := images["까마귀"]
	seaAnemone := images["말미잘"]
	wolf := images["늑대"]
	rattlesnake := images["방울뱀"]
	poison := images["독"]
	vampire := images["흡혈"]
	eggC := images["알"]
	revenge := images["보복"]
	flying := images["비행"]
	tentacle := images["긴촉수"]
	threat := images["위협"]
	fire := images["불"]
	run := images["달리기"]
	dealershop := images["거래소"]
	dealer := images["거래상"]
	bomber := images["폭탄수레"]
	selfdestruct := images["자폭"]
	stupid := images["바보게"]
	mistaken := images["착각"]
	satellite := images["인공위성"]
	spacesupply := images["우주보급"]

	shell := images["껍데기"]
	wit := images["꾀"]
	rush := images["돌진"]
	stag := images["사슴벌레"]
	ship := images["양"]
	incision := images["자르기"]

	//bigmon := images["거대괴물"]
	bigbug := images["거대벌레"]
	fmachine := images["격투머신"]
	bug := images["구더기"]
	nessie := images["네시"]
	runover := images["메치기"]
	submersion := images["물에숨어지내다"]
	mistery := images["미스테리"]
	//defence := images["방어"]
	grow := images["성장"]
	dark := images["암흑"]
	orc := images["오크"]
	dragon := images["용"]
	//turtle := images["육지거북"]
	charge := images["차지"]
	elephant := images["코끼리"]
	wield := images["휘두르기"]

	allCard[0] = Cardstr{"복 어", 2, blowfish, "____", poison, "독: 상대를 공격하면\n다음 자신의 턴에 상\n대를 죽인다", 2, 1, "poison"}
	allCard[1] = Cardstr{"모 기", 1, mosquito, "____", vampire, "흡혈: 상대를 죽이면\n체력을 3 더한다", 2, 1, "vampire"}
	allCard[2] = Cardstr{"   알", 3, eggv, "____", eggC, "부화: 2턴동안 살아\n남는다면, 덱의 카드\n로 변화한다", 3, 0, "eggC"}
	allCard[3] = Cardstr{"랩 터", 2, raptor, "____", revenge, "보복: 자신을 공격시\n자신을 공격한 대상\n의 체력을 1만큼 공격", 2, 2, "revenge"}
	allCard[4] = Cardstr{"   딝", 2, chicken, "____", flying, "비행: 앞에 상대가 있\n어도 상대의 성을 바\n로 공격한다", 2, 1, "flying"}
	allCard[5] = Cardstr{"까마귀", 3, crow, "____", flying, "비행: 앞에 상대가 있\n어도 상대의 성을 바\n로 공격한다", 2, 2, "flying"}
	allCard[6] = Cardstr{"말미잘", 2, seaAnemone, "____", tentacle, "긴촉수: 턴이 끝날때\n마다, 상대에게 1만\n큼 공격(2번죽을때까지)", 3, 1, "tentacle"}
	allCard[7] = Cardstr{"늑 대", 3, wolf, "____", threat, "위협: 자신의 공격을\n받은 상대의 공격력\n을 1 깍는다(중첩불가)", 3, 2, "threat"}
	allCard[8] = Cardstr{"방울뱀", 3, rattlesnake, "____", threat, "위협: 자신의 공격을\n받은 상대의 공격력\n을 1 깍는다(중첩불가)", 2, 3, "threat"}
	allCard[9] = Cardstr{"지옥견", 4, nilcard, "____", fire, "불: 카드를 공격하면\n카드와 성 두개를 모\n두 공격한다", 4, 3, "fire"}
	allCard[10] = Cardstr{"여 우", 2, nilcard, "____", wit, "꾀: 죽을 상황에 왼쪽\n의 칸이 비어 있으면\n그곳으로 도망친다", 2, 2, "wit"}
	allCard[11] = Cardstr{"메갈로돈", 4, nilcard, "____", nilcard, "잔혹함: 이 카드를 소\n환하기 위해서는 $2\n이상의 카드를 지불해야 함", 7, 4, "cruelty"}
	allCard[12] = Cardstr{"늑 대", 3, nilcard, "____", nilcard, "전략: 양쪽상대 모두\n를 공격할 수 있습니다", 3, 2, "strategy"}
	allCard[13] = Cardstr{"팬 더", 4, nilcard, "____", nilcard, "귀여움: 같은 라인에\n있는 상대의 공격력\n을 1 떨어트립니다", 4, 3, "cute"}
	allCard[14] = Cardstr{"돌고래", 3, nilcard, "____", nilcard, "물:피공격시 다음턴\n에 1회복/초음파:공격\n시 상대는 1턴간 얼음", 3, 2, "dolphin"}
	allCard[15] = Cardstr{"물거미", 3, nilcard, "____", nilcard, "공기방탄: 상대 공격\n력에 상관없이, 2의 피\n해만 받는다", 3, 1, "airproof"}
	allCard[16] = Cardstr{"까 치", 3, nilcard, "____", nilcard, "동료: 죽으면 그자리\n에 원하는 카드를 즉\n시 놓는다", 3, 2, "company"}
	allCard[17] = Cardstr{"다람쥐", 2, nilcard, "____", nilcard, "나무타기: 적을 죽이\n면, 옆의 카드를 한번\n더 공격한다", 2, 1, "climb"}
	allCard[18] = Cardstr{"골 렘", 5, nilcard, "____", nilcard, "거대: 공격시, 오른쪽\n의 카드도 함께 공격", 9, 1, "huge"}
	allCard[19] = Cardstr{"생 쥐", 0, nilcard, "____", nilcard, "     ", 3, 1, ""}
	allCard[20] = Cardstr{"태양석", 3, nilcard, "____", nilcard, "태양신: 매턴 마다 공\n격력이 2 증가한다", 5, 0, "helios"}
	allCard[21] = Cardstr{"폭탄수레", 3, bomber, "____", selfdestruct, "자폭: 상대를 공격 후\n자신도 사망한다", 2, 5, "selfdestruct"}
	allCard[22] = Cardstr{"백 마", 2, nilcard, "____", run, "달리기:자신의턴이\n끝날때,모든카드를\n왼쪽으로1칸움직인다", 4, 1, "run"}
	allCard[23] = Cardstr{"바보게", 3, stupid, "____", mistaken, "착각: 상대를 죽이면,\n상대를 내 패에 넣고\n그 자리에 간다", 2, 2, "mistaken"}
	allCard[24] = Cardstr{"인공위성", 3, satellite, "____", spacesupply, "우주보급: 매 자신의\n턴마다 양 옆 카드의\nHP를 1 올린다", 3, 1, "spacesupply"}
	allCard[25] = Cardstr{"원숭이", 2, nilcard, "____", nilcard, "나무타기: 적을 죽이\n면, 옆의 카드를 한번\n더 공격한다", 3, 1, "climb"}
	allCard[26] = Cardstr{"낙 타", 4, nilcard, "____", nilcard, "버티기: 판에 내려진\n후, 2턴동안 체력 1로 살\n아남습니다", 4, 2, "survive"}
	allCard[27] = Cardstr{"   꿩", 3, nilcard, "____", flying, "비행: 앞에 상대가 있\n어도 상대의 성을 바\n로 공격한다", 4, 2, "flying"}
	allCard[28] = Cardstr{"토 끼", 1, nilcard, "____", nilcard, "S코인: 이 동물이 살\n아있을 경우 매 턴마\n다 1코인이 더 들어온다", 1, 1, "sCoin"}
	allCard[29] = Cardstr{"변형체", 3, nilcard, "____", nilcard, "변형: 상대의 방어특\n성을 무시합니다", 4, 3, "trans"}
	allCard[30] = Cardstr{"붕 어", 0, nilcard, "____", nilcard, "물: 공격을 받으면 다\n음 자신의 턴에 1 회복\n한다", 2, 1, "water"}
	allCard[31] = Cardstr{"코뿔소", 3, nilcard, "____", rush, "돌진: 공격으로 상대\n를 죽이지 못한다면\n부가피해 1을 받는다", 4, 3, "rush"}
	allCard[32] = Cardstr{"사슴벌레", 3, stag, "____", incision, "자르기: 죽인 카드를\n잘라 3턴간 다신 판에\n올라오지 못하게한다", 3, 2, "incision"}
	allCard[33] = Cardstr{"말코손바닥사슴", 4, nilcard, "____", nilcard, "뿔: 정면과 양옆도 공\n격할 수 있다", 4, 3, "horn"}
	allCard[34] = Cardstr{"흡혈박쥐", 3, nilcard, "____", vampire, "흡혈: 상대를 죽이면\n체력을 3 더한다", 3, 2, "vampire"}
	allCard[35] = Cardstr{"   양", 3, ship, "____", shell, "껍데기: 한번에 죽지\n않으면, 해당 공격을 무\n효화 한다.(1번만 발동)", 3, 1, "shell"}
	allCard[36] = Cardstr{"거대괴물", 4, nilcard, "____", nilcard, "   ", 5, 3, " "}
	allCard[37] = Cardstr{"타 조", 3, nilcard, "____", nilcard, "신속: 첫 배치턴에 상\n대를 공격 할 수 있다", 2, 3, "advance"}
	allCard[38] = Cardstr{"토 템", 3, nilcard, "____", nilcard, "리더쉽:옆카드 공격\n력 1업. 긴목: 피공시\n비행특성 무효화", 5, 0, "totem"}
	allCard[39] = Cardstr{"기 린", 4, nilcard, "____", nilcard, "긴목: 피공시 비행특\n성 무효화", 5, 2, "lneck"}
	allCard[40] = Cardstr{"꽃사슴", 2, nilcard, "____", nilcard, "점프: 턴이끝나면 왼\n쪽으로 이동. 왼쪽이\n 안되면 오른쪽으로 이동", 2, 2, "jump"}
	allCard[41] = Cardstr{"산 양", 2, nilcard, "____", nilcard, "점프: 턴이끝나면 왼\n쪽으로 이동. 왼쪽이\n 안되면 오른쪽으로 이동", 2, 3, "jump"}
	allCard[42] = Cardstr{"좀 비", 3, nilcard, "____", nilcard, "독: 상대를 공격하면\n다음 자신의 턴에 상\n대를 죽인다", 3, 2, "poison"}
	allCard[43] = Cardstr{"펭 귄", 3, nilcard, "____", nilcard, "다이빙: 한번에 죽지\n않으면, 다이빙하여\n공격을 받지않는다", 3, 2, "diving"}
	allCard[44] = Cardstr{"고릴라", 4, nilcard, "____", nilcard, "전략: 양쪽상대 모두\n를 공격할 수 있습니다", 4, 4, "strategy"}
	allCard[45] = Cardstr{"오크", 5, orc, "____", wield, "휘두르기: 상대를\n때리면 바로 오른쪽\n으로 날린다", 3, 5, "wield"}
	allCard[46] = Cardstr{"용", 7, dragon, "____", charge, "준비: 공격을 한후,\n1턴간 준비가 필요하다", 8, 6, "charge"}
	allCard[47] = Cardstr{"네시", 4, nessie, "____", submersion, "물에숨어지내다: 1턴\n은 공격을 받지만, 다\n음턴은 공격을 받지 않는다", 4, 2, "submersion"}
	allCard[48] = Cardstr{"  ", 4, dark, "____", mistery, "미스테리: 50%의 확률\n로 공격이 맞지 않는다", 3, 2, "mistery"}
	allCard[49] = Cardstr{"격투머신", 4, fmachine, "____", runover, "메치기: 상대를 공\n격하면 다른 상대의\n칸으로 옮긴다", 4, 2, "runover"}
	allCard[50] = Cardstr{"구더기", 2, bug, "____", grow, "성장: 3턴이 지나면 카드를 뒤집는다", 1, 1, "grow"}
	allCard[51] = Cardstr{"달팽이", 3, nilcard, "____", shell, "껍데기: 한번에 죽지\n않으면, 해당 공격을 무\n효화 한다.(1번만 발동)", 4, 1, "shell"}
	allCard[52] = Cardstr{"코끼리", 4, elephant, "____", nilcard, "힘: 공격력이나 체력\n이 1인 상대의 공격은\n통하지 않는다", 4, 3, "power"}

	allCard[53] = Cardstr{"CV 맨", 4, nilcard, "____", nilcard, "흉내: 해당 라인에 있\n는 모든 카드의 특정\n을 가진다", 3, 3, "copy"}
	allCard[54] = Cardstr{"앵무새", 3, nilcard, "____", nilcard, "흉내: 해당 라인에 있\n는 모든 카드의 특정\n을 가진다", 2, 2, "copy"}

	bigbugCard = Cardstr{"거대벌레", 2, bigbug, "____", nilcard, "   ", 4, 6, " "}

	NallCard[1] = Cardstr{"거래소", 3, dealershop, "____", dealer, "거래상: 이 카드가 놓여\n질때, 플레이어간 카\n드를 교환할수 있음", 4, 1, "dealer"}
	NallCard[2] = Cardstr{"CV 맨", 4, nilcard, "____", nilcard, "흉내: 해당 라인에 있\n는 모든 카드의 특정\n을 가진다", 3, 3, "copy"}
	NallCard[3] = Cardstr{"앵무새", 3, nilcard, "____", nilcard, "흉내: 해당 라인에 있\n는 모든 카드의 특정\n을 가진다", 2, 2, "copy"}

}

func shuffleCard() {
	var a [cardTotal]int
	for i := 0; i < int(cardTotal); i++ {
		a[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	for c := 0; c < 30; {
		for p := 0; p < 3; p++ {
			for d := 0; d < 10; d++ {
				cn := a[c]
				deckCard[p][d] = deckCardstr{allCard[cn], theCardimg(allCard[cn]), p, 20, d, false, false, nil}
				c++
			}
		}
	}

	//deckCard[0][1] = deckCardstr{allCard[24], theCardimg(allCard[24]), 0, 20, 1, false, false, nil}
	//deckCard[1][1] = deckCardstr{allCard[53], theCardimg(allCard[53]), 1, 20, 1, false, false, nil}

}
