////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_instructions_0xfi.go - Apr-23-2025 by aldebap
//
//	Emulator for Sharp SM83 CPU - instructions 0xf0 - 0xff
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

// execute instruction LD_A_ADDR_nn
func (c *SM83_CPU) executeInstruction_LD_A_ADDR_nn() error {
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
		c.n_lsb, err = c.readByteFromMemory(uint16(c.n_msb)<<8 | uint16(c.n_lsb))
		c.cpu_state = EXECUTION_CYCLE_4

		return err

	case EXECUTION_CYCLE_4:
		c.a = c.n_lsb
	}

	if c.trace {
		fmt.Printf("[trace] LD A, (nn): 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}
