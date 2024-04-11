package gba

import "math/rand"

func randomByte() byte {
	return byte(rand.Intn(256))
}

func merge4(b []byte) uint32 {
	if len(b) != 4 {
		panic("merge4: byte array length is not 4")
	}
	//return uint32(b[3])<<24 | uint32(b[2])<<16 | uint32(b[1])<<8 | uint32(b[0])
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}
