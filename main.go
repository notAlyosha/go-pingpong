package main

/*===============================================================*/

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 900
	screenHeight = 600
	step         = 5
)

/*===============================================================*/

type position struct {
	X int
	Y int
}

type direction struct {
	isDown  bool
	isRight bool
}

type cube struct {
	pos  position
	dir  direction
	size int
}

func (c *cube) Draw(s *ebiten.Image) {
	ebitenutil.DrawRect(s, 0, 0, screenWidth, screenHeight, color.RGBA{R: 60, G: 60, B: 60, A: 60})
	ebitenutil.DrawRect(s, float64(c.pos.X), float64(c.pos.Y), float64(c.size), float64(c.size), color.RGBA{R: 130, G: 190, B: 230, A: 255})

}

func (c *cube) ChangeDir() {
	if c.pos.X < 0 {
		c.dir.isRight = true
	}

	if c.pos.Y < 0 {
		c.dir.isDown = true
	}

	if c.pos.X > screenWidth-c.size {
		c.dir.isRight = false
	}

	if c.pos.Y > screenHeight-c.size {
		c.dir.isDown = false
	}
}

func (c *cube) SetPosition() {

	if c.dir.isDown {
		c.pos.Y += step
	} else {
		c.pos.Y -= step
	}

	if c.dir.isRight {
		c.pos.X += step
	} else {
		c.pos.X -= step
	}
}

/*===============================================================*/

type Game struct {
	C cube
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	g.C.ChangeDir()
	g.C.SetPosition()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.C.Draw(screen)

}

func main() {
	c := cube{pos: position{X: 40, Y: 60}, dir: direction{isDown: true, isRight: true}, size: 20}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(&Game{C: c}); err != nil {
		panic(err)
	}
}
