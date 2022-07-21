package main

import (
	"com.lannerkr.NewBoard/newboard"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := newboard.NewGame()

	const scale = 1
	ebiten.SetWindowSize(newboard.ScreenWidth*scale, newboard.ScreenHeight*scale)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	//ebiten.SetMaxTPS(10)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}

}
