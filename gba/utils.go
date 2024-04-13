package gba

import "math/rand"

func randomByte() byte {
	return byte(rand.Intn(256))
}

// merge4 Little endian
func merge4(b []byte) uint32 {
	if len(b) != 4 {
		panic("merge4: byte array length is not 4")
	}
	return uint32(b[3])<<24 | uint32(b[2])<<16 | uint32(b[1])<<8 | uint32(b[0])
	//return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

func getBit32(b, n uint32) uint32 {
	assert(n < 32)
	bit := (b >> n) & 1
	assert(bit <= 1)
	return bit
}

// getBitRange32 is inclusive, n1 must be greater than n2
func getBitRange32(b, n1, n2 uint32) uint32 {
	assert(n1 < 32 && n2 < 32 && n2 < n1)
	return (b >> n2) & (1<<n1 - 1)
}

func addSigned32(b uint32, a int32) uint32 {
	if a > 0 {
		return b + uint32(a)
	} else {
		return b - uint32(-a)
	}
}

func assert(condition bool) {
	if !condition {
		panic("assertion failed")
	}
}
