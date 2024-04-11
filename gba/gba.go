package gba

import (
	"image/color"
)

const (
	ScreenWidth  = 240
	ScreenHeight = 160
)

type GBA struct {
	display [][]color.RGBA

	cpu  *cpu
	bios *bios
	cart *cart
}

func NewGBA(rom []byte) *GBA {
	gba := &GBA{}
	gba.cpu = newCpu(gba, rom)
	gba.bios = newBios()
	gba.cart = newCart(rom)

	for i := 0; i < ScreenWidth; i++ {
		gba.display = append(gba.display, []color.RGBA{})
		for j := 0; j < ScreenHeight; j++ {
			gba.display[i] = append(gba.display[i], color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff})
		}
	}

	return gba
}

// Update must be called at 60Hz
func (gba *GBA) Update() {
	for i := 0; i < ScreenWidth; i++ {
		for j := 0; j < ScreenHeight; j++ {
			gba.display[i][j] = color.RGBA{R: randomByte(), G: randomByte(), B: randomByte(), A: 0xff}
		}
	}

	gba.cpu.step()
}

func (gba *GBA) GetDisplay() [][]color.RGBA {
	return gba.display
}

func (gba *GBA) read32(addr uint32) uint32 {
	if addr <= 0x00003FFF { // BIOS - System ROM
		return gba.bios.read32(addr)
	} else if addr <= 0x01FFFFFF {
		panic("0x00004000-0x01FFFFFF: not used")
	} else if addr <= 0x0203FFFF {
		// WRAM - On-board Work RAM
	} else if addr <= 0x02FFFFFF {
		panic("0x02040000-0x02FFFFFF: not used")
	} else if addr <= 0x03007FFF {
		// WRAM - On-chip Work RAM
	} else if addr <= 0x03FFFFFF {
		panic("0x03008000-0x03FFFFFF: not used")
	} else if addr <= 0x040003FE {
		// I/O Registers
	} else if addr <= 0x04FFFFFF {
		panic("0x04000400-0x04FFFFF: not used")
	} else if addr <= 0x050003FF {
		// BG/OBJ Palette RAM
	} else if addr <= 0x05FFFFFF {
		panic("0x05000400-0x05FFFFFF: not used")
	} else if addr <= 0x06017FFF {
		// VRAM - Video RAM
	} else if addr <= 0x06FFFFFF {
		panic("0x06018000-0x06FFFFFF: not used")
	} else if addr <= 0x070003FF {
		// OAM - OBJ Attributes
	} else if addr <= 0x07FFFFFF {
		panic("0x07000400-0x07FFFFFF: not used")
	} else if addr <= 0x09FFFFFF {
		// Game Pak ROM/FlashROM (max 32MB) - Wait State 0
	} else if addr <= 0x0BFFFFFF {
		// Game Pak ROM/FlashROM (max 32MB) - Wait State 1
	} else if addr <= 0x0DFFFFFF {
		// Game Pak ROM/FlashROM (max 32MB) - Wait State 2
	} else if addr <= 0x0E00FFFF {
		// Game Pak SRAM
	} else if addr <= 0x0FFFFFFF {
		panic("0x0E010000-0x0FFFFFFF: not used")
	} else if addr <= 0xFFFFFFFF {
		panic("0x10000000-0xFFFFFFFF: not used (upper 4bits of address bus unused)")
	}

	panic("shit liao lor")
}
