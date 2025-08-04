////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_instructions_0x3i.go - Apr-23-2025 by aldebap
//
//	Emulator for Sharp SM83 CPU - instructions 0x30 - 0x3f
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

// execute instruction JR_NC_e
func (c *SM83_CPU) executeInstruction_JR_NC_e() error {

	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_msb, err = c.fetchInstructionArgument()
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		if c.flags&FLAG_C == 0 {
			c.pc += uint16(int8(c.n_msb))
		}
	}

	if c.trace {
		fmt.Printf("[trace] JR_NC_e\n")
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}
