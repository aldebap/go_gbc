////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_instructions_0x1i.go - Apr-23-2025 by aldebap
//
//	Emulator for Sharp SM83 CPU - instructions 0x10 - 0x1f
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

// execute instruction STOP
func (c *SM83_CPU) executeInstruction_STOP() error {

	if c.trace {
		fmt.Printf("[trace] STOP\n")
	}

	//	TODO: add a flag to STOP/HALT CPU
	return nil
}

// execute instruction LD_DE_nn
func (c *SM83_CPU) executeInstruction_LD_DE_nn() error {
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
		c.d = c.n_msb
		c.e = c.n_lsb
	}

	if c.trace {
		fmt.Printf("[trace] LD DE, nn: 0x%02x%02x\n", c.d, c.e)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_ADDR_DE_A
func (c *SM83_CPU) executeInstruction_LD_ADDR_DE_A() error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		err = c.writeByteIntoMemory(uint16(c.d)<<8|uint16(c.e), c.a)
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
	}

	if c.trace {
		fmt.Printf("[trace] LD (DE), A: 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction INC_DE
func (c *SM83_CPU) executeInstruction_INC_DE() error {

	var de = uint16(c.d)<<8 | uint16(c.e)

	de++

	c.d = uint8((de & 0xff00) >> 8)
	c.e = uint8(de & 0x00ff)

	if de == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] INC DE: 0x%02x%02x\n", c.d, c.e)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}
