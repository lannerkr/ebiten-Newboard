package newboard

func mousePosition(mx, my int) (mpp, mcc int) {

	if my > cardPos[0].sy && my < cardPos[0].ey {
		mpp = 0
	} else if my > cardPos[4].sy && my < cardPos[4].ey {
		mpp = 1
	} else if my > cardPos[8].sy && my < cardPos[8].ey {
		mpp = 2
		if mx > buttonPos[0].sx && my > buttonPos[0].sy && my < buttonPos[0].ey {
			mpp = 5
		}
	} else {
		mpp = 8
	}

	if mpp == 5 && mx > buttonPos[0].sx && mx < buttonPos[0].ex {
		mcc = 5
	} else if mpp == 5 && mx > buttonPos[1].sx && mx < buttonPos[1].ex {
		mcc = 6
	} else if mx > cardPos[0].sx && mx < cardPos[0].ex {
		mcc = 0
	} else if mx > cardPos[1].sx && mx < cardPos[1].ex {
		mcc = 1
	} else if mx > cardPos[2].sx && mx < cardPos[2].ex {
		mcc = 2
	} else if mx > cardPos[3].sx && mx < cardPos[3].ex {
		mcc = 3
	} else {
		mcc = 8
	}

	return mpp, mcc
}

func mousePositionNG(mx, my int) (mpp int) {
	width := ScreenWidth
	height := ScreenHeight

	if mx > buttonPos[0].sx && mx < buttonPos[0].ex && my > buttonPos[0].sy && my < buttonPos[0].ey {
		mpp = 5
	}

	if my > height/2-200 && my < height/2 {
		if mx > (width/2-320) && mx < (width/2) {
			mpp = 1
		} else if mx > (width/2) && mx < (width/2+320) {
			mpp = 2
		}
	} else if my > (height/2+20) && my < (height/2+120) && mx > (width/2-320) && mx < (width/2) {
		mpp = 3
	} else if my > (height/2+20) && my < (height/2+120) && mx > (width/2) && mx < (width/2+320) {
		mpp = 4
	} else if my > (height/2+20) && my < (height/2+120) && mx > (width/2+320) && mx < (width/2+640) {
		mpp = 6
	}

	return mpp

}
