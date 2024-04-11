package main

import (
	"fmt"
	"go-gba/gba"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowWidth  = 1200
	windowHeight = 800
)

type emu struct {
	ticks  uint64
	start  time.Time
	screen *ebiten.Image

	gba *gba.GBA
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	romBytes, err := os.ReadFile("roms/tetris.gba")
	if err != nil {
		panic(err)
	}

	e := &emu{
		ticks:  0,
		start:  time.Now(),
		screen: ebiten.NewImage(gba.ScreenWidth, gba.ScreenHeight),

		gba: gba.NewGBA(romBytes),
	}
	if err = ebiten.RunGame(e); err != nil {
		panic(err)
	}
}

func (e *emu) Update() error {
	e.ticks++
	e.updateTitle()

	e.gba.Update()

	var displayBytes []byte
	gbaDisplay := e.gba.GetDisplay()
	for _, row := range gbaDisplay {
		for _, pixel := range row {
			displayBytes = append(displayBytes, colorToBytes(pixel)...)
		}
	}
	e.screen.WritePixels(displayBytes)

	return nil
}

func (e *emu) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	op.GeoM.Scale(4, 4)
	screen.DrawImage(e.screen, op)
}

func (e *emu) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (e *emu) updateTitle() {
	realTPS := float64(e.ticks) / time.Since(e.start).Seconds()
	ebiten.SetWindowTitle(fmt.Sprintf("go-gba | Updates/s: %.2f Ticks: %d TPS: %.2f FPS: %.2f", realTPS, e.ticks, ebiten.ActualTPS(), ebiten.ActualFPS()))
}
