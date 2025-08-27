////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_instructions_0x0i.go - Apr-23-2025 by aldebap
//
//	Emulator for Sharp SM83 CPU - instructions 0x00 - 0x0f
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

// execute instruction RLCA
func (c *SM83_CPU) executeInstruction_RLCA() error {

	if c.a&0x80 == 0x80 {
		c.a = c.a<<1 | 0x01
	} else {
		c.a = c.a << 1
	}

	if c.a == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] RLCA: 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_ADDR_nn_SP
func (c *SM83_CPU) executeInstruction_LD_ADDR_nn_SP() error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_lsb, err = c.fetchInstructionArgument()
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		c.n_msb, err = c.fetchInstructionArgument()
		c.cpu_state = EXECUTION_CYCLE_3

		return err

	case EXECUTION_CYCLE_3:
		err = c.writeByteIntoMemory(uint16(c.n_msb)<<8|uint16(c.n_lsb), c.p)
		c.cpu_state = EXECUTION_CYCLE_4

		return err

	case EXECUTION_CYCLE_4:
		err = c.writeByteIntoMemory(uint16(c.n_msb)<<8|uint16(c.n_lsb)+1, c.s)
		c.cpu_state = EXECUTION_CYCLE_5

		return err

	case EXECUTION_CYCLE_5:
	}

	if c.trace {
		fmt.Printf("[trace] LD (nn), SP: 0x%02x%02x\n", c.s, c.p)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction RLCA
func (c *SM83_CPU) executeInstruction_RRCA() error {

	if c.a&0x01 == 0x01 {
		c.a = c.a>>1 | 0x80
	} else {
		c.a = c.a >> 1
	}

	if c.a == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] RRCA: 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}
