////////////////////////////////////////////////////////////////////////////////
//	ram_memory.go - Apr-28-2025 by aldebap
//
//	implementation of RAM memory bank
////////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

type RAM_memory struct {
	size  uint16
	array []uint8
}

// create a new RAM memory
func NewRAM_memory(size uint16) *RAM_memory {

	return &RAM_memory{
		size:  size,
		array: make([]uint8, size),
	}
}

// return memory bank size
func (m *RAM_memory) Len() uint16 {
	return m.size
}

// write a byte into RAM memory
func (m *RAM_memory) WriteByte(address uint16, value uint8) error {
	if address < 0 || address >= m.size {
		return fmt.Errorf("address out of bounds")
	}

	m.array[address] = value

	return nil
}

// read a byte from RAM memory
func (m *RAM_memory) ReadByte(address uint16) (uint8, error) {
	if address < 0 || address >= m.size {
		return 0, fmt.Errorf("address out of bounds")
	}

	return m.array[address], nil
}

// write a word into RAM memory
func (m *RAM_memory) WriteWord(address uint16, value uint16) error {
	if address < 0 || address >= m.size-1 {
		return fmt.Errorf("address out of bounds")
	}

	m.array[address] = uint8(value & 0x00ff)
	m.array[address+1] = uint8(value >> 8 & 0x00ff)

	return nil
}

// read a word from RAM memory
func (m *RAM_memory) ReadWord(address uint16) (uint16, error) {
	if address < 0 || address >= m.size-1 {
		return 0, fmt.Errorf("address out of bounds")
	}

	return uint16(m.array[address])<<8 | uint16(m.array[address+1]), nil
}
