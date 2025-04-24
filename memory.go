////////////////////////////////////////////////////////////////////////////////
//	memory.go - Apr-23-2025 by aldebap
//
//	interface for a memory bank
////////////////////////////////////////////////////////////////////////////////

package main

type memory interface {
	Len() uint16

	WriteByte(address uint16, value uint8) error
	ReadByte(address uint16) (uint8, error)

	WriteWord(address uint16, value uint16) error
	ReadWord(address uint16) (uint16, error)
}
