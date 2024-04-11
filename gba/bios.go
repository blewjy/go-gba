package gba

import (
	_ "embed"
	"fmt"
)

//go:embed bios.gba
var biosData []byte

type bios struct{}

func newBios() *bios {

	for i := 0; i < 10; i++ {
		fmt.Printf("[bios] %d: 0x%08X\n", i, merge4(biosData[i*4:i*4+4]))
	}

	return &bios{}
}

func (b *bios) read32(addr uint32) uint32 {
	return merge4(biosData[addr : addr+4])
}
