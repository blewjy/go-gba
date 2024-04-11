package gba

import "fmt"

type cart struct {
	rom []byte
}

func newCart(rom []byte) *cart {

	fmt.Printf("ROM Entry Point: 0x%08X\n", merge4(rom[0:4]))
	fmt.Printf("Game Title: %s\n", rom[0xA0:0xA0+12])

	return &cart{
		rom: rom,
	}
}
