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

// execute instruction RLA
func (c *SM83_CPU) executeInstruction_RLA() error {

	c.a = c.a << 1

	if c.a == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] RLA: 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction JR_e
func (c *SM83_CPU) executeInstruction_JR_e() error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_msb, err = c.fetchInstructionArgument()
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		c.pc += uint16(int8(c.n_msb))
	}

	if c.trace {
		fmt.Printf("[trace] JR_e: 0x%02x\n", c.n_msb)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction RRA
func (c *SM83_CPU) executeInstruction_RRA() error {

	c.a = c.a >> 1

	// TODO: review implementation of RRA
	if c.a == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] RRA: 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}
