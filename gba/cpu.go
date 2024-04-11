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

	return c
}

func (cpu *cpu) step() {
	cpu.steps++

	if cpu.steps >= 10 {
		os.Exit(0)
	}

	cpu.stepArm()
}

func (cpu *cpu) stepArm() {
	fmt.Printf("[stepArm] cpu.steps: %v\n", cpu.steps)
	fmt.Printf("[stepArm] Instruction pipeline: [0x%08X, 0x%08X, 0x%08X]\n", cpu.instPipeline[0], cpu.instPipeline[1], cpu.instPipeline[2])

	pc := cpu.R[15]

	fmt.Printf("[stepArm] PC: 0x%08X\n", pc)

	// fetch
	cpu.instPipeline[0] = cpu.instPipeline[1]
	cpu.instPipeline[1] = cpu.instPipeline[2]
	cpu.instPipeline[2] = cpu.gba.read32(pc)

	// decode??

	// execute
	cpu.execArm()

	// incr PC
	cpu.R[15] = pc + 4
}

func (cpu *cpu) execArm() {
	fmt.Printf("[execArm] Executing instruction: 0x%08X\n", cpu.instPipeline[0])
}
