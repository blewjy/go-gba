package gba

type cpsr uint32

func (c cpsr) N() uint32 {
	return uint32(c) >> 31
}

func (c cpsr) Z() uint32 {
	return uint32(c) >> 30
}

func (c cpsr) C() uint32 {
	return uint32(c) >> 29
}

func (c cpsr) V() uint32 {
	return uint32(c) >> 28
}
