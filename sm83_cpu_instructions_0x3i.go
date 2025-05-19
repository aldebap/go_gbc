////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_instructions_0x3i.go - Apr-23-2025 by aldebap
//
//	Emulator for Sharp SM83 CPU - instructions 0x30 - 0x3f
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

// execute instruction LD_SP_nn
func (c *SM83_CPU) executeInstruction_LD_SP_nn() error {
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
		c.sp = uint16(c.n_msb)<<8 | uint16(c.n_lsb)
	}

	if c.trace {
		fmt.Printf("[trace] LD SP, nn: 0x%04x\n", c.sp)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction INC_A
func (c *SM83_CPU) executeInstruction_INC_A() error {

	c.a++

	if c.a == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] INC A: 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction DEC_A
func (c *SM83_CPU) executeInstruction_DEC_A() error {

	c.a--

	if c.a == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] DEC A: 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_A_n
func (c *SM83_CPU) executeInstruction_LD_A_n() error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_msb, err = c.fetchInstructionArgument()
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		c.a = c.n_msb
	}

	if c.trace {
		fmt.Printf("[trace] LD A, n: 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}
