////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu.go - Apr-23-2025 by aldebap
//
//	Emulator for Sharp SM83 CPU
////////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

// SM83 CPU states
const (
	FETCH_INSTRUCTION   = 1
	EXECUTE_INSTRUCTION = 2
)

// opcode constants
const (
	NOOP  = 0x00
	INC_A = 0x3c
)

// SM83 CPU internal registers and connections
type SM83_CPU struct {
	pc uint16
	sp uint16
	ir uint8
	ie uint8

	af uint16
	bc uint16
	de uint16
	hl uint16

	cpu_state uint8

	memoryBank        []memory
	memoryBankAddress []uint16
}

// create a new SM83 CPU
func NewSM83_CPU() *SM83_CPU {

	return &SM83_CPU{
		pc: 0,
		sp: 0,
		ir: 0,
		ie: 0,
		af: 0,
		bc: 0,
		de: 0,
		hl: 0,

		cpu_state: FETCH_INSTRUCTION,

		memoryBank:        nil,
		memoryBankAddress: nil,
	}
}

// connect a new memory bank to the CPU
func (c *SM83_CPU) ConnectMemory(memoryBank memory, intialAddress uint16) error {
	if c.memoryBank == nil {
		c.memoryBank = make([]memory, 0)
		c.memoryBankAddress = make([]uint16, 0)
	}

	//	check for bank overlapping
	for i := 0; i < len(c.memoryBank); i++ {
		//	TODO: check for overlapping
	}

	c.memoryBank = append(c.memoryBank, memoryBank)
	c.memoryBankAddress = append(c.memoryBankAddress, intialAddress)

	return nil
}

// run one machine cycle
func (c *SM83_CPU) MachineCycle() error {
	var err error

	switch c.cpu_state {
	case FETCH_INSTRUCTION:
		err = c.fetchInstruction()
		if err != nil {
			return err
		}

	case EXECUTE_INSTRUCTION:
		err = c.executeInstruction()
		if err != nil {
			return err
		}
	}

	return nil
}

// fetch one instruction from memory
func (c *SM83_CPU) fetchInstruction() error {
	var err error

	for i := 0; i < len(c.memoryBankAddress); i++ {
		if c.pc >= c.memoryBankAddress[i] && c.pc < c.memoryBankAddress[i]+c.memoryBank[i].Len() {
			c.ir, err = c.memoryBank[i].ReadByte(c.pc)
			if err != nil {
				return err
			}
			break
		}
	}

	c.pc++
	c.cpu_state = EXECUTE_INSTRUCTION

	return nil
}

// execute instruction
func (c *SM83_CPU) executeInstruction() error {
	switch c.ir {
	case INC_A:
		if c.af&uint16(0xff00) == uint16(0xff00) {
			c.af = (c.af & uint16(0x00ff))
		} else {
			c.af = ((c.af & uint16(0xff00)) + uint16(0x0100)) | (c.af & uint16(0x00ff))
		}
		//		fmt.Printf("[debug] INC A: 0x%04x\n", c.af)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// dump CPU registers
func (c *SM83_CPU) DumpRegisters() string {
	return fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; AF: 0x%04x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
		c.pc, c.sp, c.af, c.bc, c.de, c.hl)
}
