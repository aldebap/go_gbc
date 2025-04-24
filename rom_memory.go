////////////////////////////////////////////////////////////////////////////////
//	rom_memory.go - Apr-23-2025 by aldebap
//
//	implementation of ROM memory bank
////////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

type ROM_memory struct {
	size  uint16
	array []uint8
}

// return memory bank size
func (m *ROM_memory) Len() uint16 {
	return m.size
}

// load a ROM memory from a byte array
func (m *ROM_memory) Load(value []uint8) error {

	m.size = uint16(len(value))
	m.array = make([]uint8, m.size)

	for i := uint16(0); i < m.size; i++ {
		m.array[i] = value[i]
	}

	return nil
}

// write a byte into ROM memory
func (m *ROM_memory) WriteByte(address uint16, value uint8) error {

	return fmt.Errorf("cannot write to ROM memory")
}

// read a byte from ROM memory
func (m *ROM_memory) ReadByte(address uint16) (uint8, error) {
	if address < 0 || address >= m.size {
		return 0, fmt.Errorf("address out of bounds")
	}

	return m.array[address], nil
}

// write a word into ROM memory
func (m *ROM_memory) WriteWord(address uint16, value uint16) error {

	return fmt.Errorf("cannot write to ROM memory")
}

// read a word from ROM memory
func (m *ROM_memory) ReadWord(address uint16) (uint16, error) {
	if address < 0 || address >= m.size-1 {
		return 0, fmt.Errorf("address out of bounds")
	}

	return uint16(m.array[address])<<8 | uint16(m.array[address+1]), nil
}
