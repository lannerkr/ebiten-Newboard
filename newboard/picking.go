package newboard

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func pickingCard(screen *ebiten.Image, ji *int) {
	var plimg, plused [3]*ebiten.Image
	var upop [3]*ebiten.DrawImageOptions
	//var pickplayer int = 11
	plimg[0] = images["picking_p0"]
	plimg[1] = images["picking_p1"]
	plimg[2] = images["picking_p2"]
	plused[0] = images["picked_p0"]
	plused[1] = images["picked_p1"]
	plused[2] = images["picked_p2"]

	for i := 0; i < 3; i++ {
		if i == 2 && twoplay {
			break
		}
		p := i * 4
		upop[i] = &ebiten.DrawImageOptions{}
		upop[i].GeoM.Scale(float64(cardw)/160, float64(cardh)/180)
		upop[i].GeoM.Translate(bC[p].bx, bC[p].by)
		screen.DrawImage(plimg[i], upop[i])
	}

	// click player number
	//pickplayer = 0
	//fmt.Println("card pick player : ", player)

	if pickplayer != 11 {
		screen.DrawImage(plused[pickplayer], upop[pickplayer])

		screen.DrawImage(pickCard[pickplayer][*ji].first.cardImg, cardop[pickplayer][2])
		screen.DrawImage(pickCard[pickplayer][*ji].second.cardImg, cardop[pickplayer][3])

		// click card

		// if picked {
		// 	fmt.Println("picked 001]")
		// 	deckCard[pickplayer][*ji] = pickCard[pickplayer][*ji].first
		// 	fmt.Println("picked 002]")
		// 	if twoplay && pickplayer == 1 {
		// 		deckCard[0][5+*ji] = pickCard[pickplayer][*ji].second
		// 	} else {
		// 		deckCard[nexP(pickplayer)][5+*ji] = pickCard[pickplayer][*ji].second
		// 	}
		// 	fmt.Println("picked 003]")
		// 	picked = false
		// 	<-pickChan
		// 	fmt.Println("picked 004]")
		// }

	}
}

func pickNumber(ji *int) {
	//fmt.Println("pickNumber before : ", *ji)
	//	<-pickNChan
	if *ji < 4 {
		*ji++
	} else {
		pickplayer++
		*ji = 0
		if twoplay && pickplayer == 2 {
			pickBools = false
			// select {
			// case <-pickChan:
			// default:
			// }
		} else if pickplayer == 3 {
			pickBools = false
			// select {
			// case <-pickChan:
			// default:
			// }
		}
	}

	//fmt.Println("pickNumber after : ", *ji)
}

func pickedC(mc int) {
	if mc == 2 {
		//fmt.Println("picked deck : ", *ji)
		deckCard[pickplayer][*ji] = pickCard[pickplayer][*ji].first
		if twoplay && pickplayer == 1 {
			deckCard[0][5+*ji] = pickCard[pickplayer][*ji].second
			deckCard[0][5+*ji].pNum = 0
			deckCard[0][5+*ji].deckNum = 5 + *ji
		} else {
			deckCard[nexP(pickplayer)][5+*ji] = pickCard[pickplayer][*ji].second
			deckCard[nexP(pickplayer)][5+*ji].pNum = nexP(pickplayer)
			deckCard[nexP(pickplayer)][5+*ji].deckNum = 5 + *ji
		}
	} else if mc == 3 {
		//fmt.Println("picked deck : ", *ji)
		deckCard[pickplayer][*ji] = pickCard[pickplayer][*ji].second
		if twoplay && pickplayer == 1 {
			deckCard[0][5+*ji] = pickCard[pickplayer][*ji].first
			deckCard[0][5+*ji].pNum = 0
			deckCard[0][5+*ji].deckNum = 5 + *ji
		} else {
			deckCard[nexP(pickplayer)][5+*ji] = pickCard[pickplayer][*ji].first
			deckCard[nexP(pickplayer)][5+*ji].pNum = nexP(pickplayer)
			deckCard[nexP(pickplayer)][5+*ji].deckNum = 5 + *ji
		}
	}

	<-pickChan
}
