package gba

import (
	"fmt"
	"os"
)

type registers struct {
	R [16]uint32
}

type cpu struct {
	gba *GBA

	registers

	instPipeline [3]uint32

	steps uint64
}

func newCpu(gba *GBA, rom []byte) *cpu {
	c := &cpu{
		gba: gba,
	}

	c.R[15] = merge4(rom[0:4])

	return c
}

func (cpu *cpu) step() {
	cpu.steps++

	if cpu.steps >= 10 {
		os.Exit(0)
	}

	cpu.instPipeline[0] = cpu.instPipeline[1]
	cpu.instPipeline[1] = cpu.instPipeline[2]

	cpu.stepThumb()
}

func (cpu *cpu) stepThumb() {
	fmt.Printf("[stepThumb] cpu.steps: %v. Instruction pipeline: [0x%08X, 0x%08X, 0x%08X]\n", cpu.steps, cpu.instPipeline[0], cpu.instPipeline[1], cpu.instPipeline[2])

	pc := cpu.R[15]

	fmt.Printf("[stepThumb] PC: 0x%08X\n", pc)

	// fetch
	cpu.instPipeline[2] = uint32(cpu.gba.read16(pc))

	// decode??

	// execute
	cpu.execThumb()
}

func (cpu *cpu) execThumb() {
	fmt.Printf("[execThumb] cpu.steps: %v. Executing instruction: %08X\n", cpu.steps, cpu.instPipeline[0])
}
