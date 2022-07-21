package newboard

import (
	"embed"
	"image"
	"image/color"
	"path"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

//go:embed assets/*
var Assets embed.FS

var (
	images = map[string]*ebiten.Image{}
)

func LoadImages() error {
	const dir = "assets"

	ents, err := Assets.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, ent := range ents {
		name := ent.Name()
		ext := filepath.Ext(name)
		if ext != ".png" {
			continue
		}

		f, err := Assets.Open(path.Join(dir, name))
		if err != nil {
			return err
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			return err
		}

		key := name[:len(name)-len(ext)]
		images[key] = ebiten.NewImageFromImage(img)
	}
	return nil
}

func fillBack(s *ebiten.Image) error {

	width := ScreenWidth
	height := ScreenHeight
	var opboard [12]*ebiten.DrawImageOptions

	background := images["background"]
	boardImage := images["boardN"]
	baw, bah = background.Size()
	bow, boh = boardImage.Size()

	Nboard := ebiten.NewImage(cardw, cardh)
	opboardn := &ebiten.DrawImageOptions{}
	opboardn.GeoM.Scale(float64(cardw)/float64(bow), float64(cardh)/float64(boh))
	Nboard.DrawImage(boardImage, opboardn)

	opback := &ebiten.DrawImageOptions{}
	opback.GeoM.Translate(0, 0)
	opback.GeoM.Scale(float64(width)/float64(baw), float64(height)/float64(bah))
	s.DrawImage(background, opback)

	for i := 0; i < 12; i++ {
		opboard[i] = &ebiten.DrawImageOptions{}
		opboard[i].GeoM.Translate(bC[i].bx, bC[i].by)
		s.DrawImage(Nboard, opboard[i])
	}

	p1y, p2y, p3y := int(bC[0].by), int(bC[4].by), int(bC[8].by)
	text.Draw(s, "Player 1", arcadeFont, width/120*105, p1y+30, color.Black)
	text.Draw(s, "Player 2", arcadeFont, width/120*105, p2y+30, color.Black)
	text.Draw(s, "Player 3", arcadeFont, width/120*105, p3y+30, color.Black)

	text.Draw(s, "CASTLE", arcadeFont, width/120*103, p1y+60, color.RGBA{0, 0, 0xff, 0xff})
	text.Draw(s, "CASTLE", arcadeFont, width/120*103, p2y+60, color.RGBA{0, 0, 0xff, 0xff})
	text.Draw(s, "CASTLE", arcadeFont, width/120*103, p3y+60, color.RGBA{0, 0, 0xff, 0xff})
	text.Draw(s, "GOLD", arcadeFont, width/120*114, p1y+60, color.Black)
	text.Draw(s, "GOLD", arcadeFont, width/120*114, p2y+60, color.Black)
	text.Draw(s, "GOLD", arcadeFont, width/120*114, p3y+60, color.Black)

	nextButton := images["touch"].SubImage(image.Rect(80, 0, 160, 64)).(*ebiten.Image)
	homeButton := images["touch"].SubImage(image.Rect(240, 0, 320, 64)).(*ebiten.Image)

	ophome := &ebiten.DrawImageOptions{}
	ophome.GeoM.Translate(float64(buttonPos[0].sx), float64(buttonPos[0].sy))
	opnext := &ebiten.DrawImageOptions{}
	opnext.GeoM.Translate(float64(buttonPos[1].sx), float64(buttonPos[1].sy))

	s.DrawImage(nextButton, opnext)
	s.DrawImage(homeButton, ophome)

	return nil
}

func drawTouchedNext(screen *ebiten.Image) {
	nextButton := images["touched"].SubImage(image.Rect(80, 0, 160, 64)).(*ebiten.Image)
	opnext := &ebiten.DrawImageOptions{}
	opnext.GeoM.Translate(float64(buttonPos[1].sx+2), float64(buttonPos[1].sy+2))

	screen.DrawImage(nextButton, opnext)
}

func drawTouchedHome(screen *ebiten.Image) {
	homeButton := images["touched"].SubImage(image.Rect(240, 0, 320, 64)).(*ebiten.Image)
	ophome := &ebiten.DrawImageOptions{}
	ophome.GeoM.Translate(float64(buttonPos[0].sx+2), float64(buttonPos[0].sy+2))

	screen.DrawImage(homeButton, ophome)
}
