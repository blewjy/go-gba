package gba

import (
	"fmt"
	"os"
)

type registers struct {
	R    [16]uint32
	RFiq [16]uint32
	RSvc [16]uint32
	RAbt [16]uint32
	RIrq [16]uint32
	RUnd [16]uint32
	CPSR cpsr      // Current Program Status Register
	SPSR [5]uint32 // fiq, svc, abt, irq, und
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

	c.R[15] = 0x08

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

	fmt.Printf("\n[cpu.steps: %v][stepArm] PC: 0x%08X\n", cpu.steps, cpu.R[15])

	// fetch
	cpu.instPipeline[0] = cpu.instPipeline[1]
	cpu.instPipeline[1] = cpu.instPipeline[2]
	cpu.instPipeline[2] = cpu.gba.read32(cpu.R[15])

	fmt.Printf("[cpu.steps: %v][stepArm] Instruction pipeline: [0x%08X, 0x%08X, 0x%08X]\n", cpu.steps, cpu.instPipeline[0], cpu.instPipeline[1], cpu.instPipeline[2])

	// decode??

	// execute
	cpu.execArm()

	// incr PC
	cpu.R[15] = cpu.R[15] + 4

}

func (cpu *cpu) execArm() {
	inst := cpu.instPipeline[0]
	fmt.Printf("[cpu.steps: %v][execArm] Executing instruction: 0x%08X (0b%032b)\n", cpu.steps, inst, inst)

	cond := cpu.getCond(inst)
	fmt.Printf("[cpu.steps: %v][execArm] Cond: 0b%04b\n", cpu.steps, cond)

	if !cpu.testCond(cond) {
		fmt.Printf("[cpu.steps: %v][execArm] Condition check failed, go next\n", cpu.steps)
		return
	}

	switch {
	case cpu.checkZero(inst):
		fmt.Printf("[cpu.steps: %v][execArm] Zero instruction, continue first\n", cpu.steps)
	case cpu.checkArmBranch(inst):
		cpu.execArmBranch(inst)
	default:
		fmt.Printf("[cpu.steps: %v][execArm] Instruction not implemented!\n", cpu.steps)
		os.Exit(1)
	}
}

type Condition uint8

const (
	EQ   Condition = 0b0000
	NE             = 0b0001
	CSHS           = 0b0010
	CCLO           = 0b0011
	MI             = 0b0100
	PL             = 0b0101
	VS             = 0b0110
	VC             = 0b0111
	HI             = 0b1000
	LS             = 0b1001
	GE             = 0b1010
	LT             = 0b1011
	GT             = 0b1100
	LE             = 0b1101
	AL             = 0b1110
	NV             = 0b1111
)

func (cpu *cpu) getCond(inst uint32) Condition {
	cond := Condition(uint8(inst >> 28))
	assert(cond < 16)
	return cond
}

func (cpu *cpu) testCond(cond Condition) bool {
	switch cond {
	case EQ:
		return cpu.CPSR.Z() == 1
	case NE:
		return cpu.CPSR.Z() == 0
	case CSHS:
		return cpu.CPSR.C() == 1
	case CCLO:
		return cpu.CPSR.C() == 0
	case MI:
		return cpu.CPSR.N() == 1
	case PL:
		return cpu.CPSR.N() == 0
	case VS:
		return cpu.CPSR.V() == 1
	case VC:
		return cpu.CPSR.V() == 0
	case HI:
		return cpu.CPSR.C() == 1 && cpu.CPSR.Z() == 0
	case LS:
		return cpu.CPSR.C() == 0 || cpu.CPSR.Z() == 1
	case GE:
		return cpu.CPSR.N() == cpu.CPSR.V()
	case LT:
		return cpu.CPSR.N() != cpu.CPSR.V()
	case GT:
		return cpu.CPSR.Z() == 0 && cpu.CPSR.N() == cpu.CPSR.V()
	case LE:
		return cpu.CPSR.Z() == 1 || cpu.CPSR.N() != cpu.CPSR.V()
	case AL:
		return true
	case NV:
		return false
	default:
		panic("unhandled condition")
	}
}

func (cpu *cpu) checkZero(inst uint32) bool {
	return inst == 0
}

func (cpu *cpu) checkArmBranch(inst uint32) bool {
	return inst&0b0000_1110_0000_0000_0000_0000_0000_0000 == 0b0000_1010_0000_0000_0000_0000_0000_0000
}

func (cpu *cpu) execArmBranch(inst uint32) {
	fmt.Printf("[cpu.steps: %v][execArmBranch] inst: %032b\n", cpu.steps, inst)

	L := getBit32(inst, 24)
	offset := int32(getBitRange32(inst, 23, 0))

	fmt.Printf("[cpu.steps: %v][execArmBranch] L: %01b, offset := %024b (%d)\n", cpu.steps, L, offset, offset)

	cpu.R[15] = addSigned32(cpu.R[15]+8, offset<<2)
}
