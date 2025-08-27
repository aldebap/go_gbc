////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_loadInstructions.go - Aug-3-2025 by aldebap
//
//	Emulator for Sharp SM83 CPU - load instructions
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

/*
LD r8,r8    --> LD_X_Y
LD r8,n8    --> LD_X_n
LD r16,n16
LD [HL],r8
LD [HL],n8
LD r8,[HL]  --> LD_X_ADDR_HL
LD [r16],A
LD [n16],A
LDH [n16],A
LDH [C],A
LD A,[r16]
LD A,[n16]
LDH A,[n16]
LDH A,[C]
LD [HLI],A
LD [HLD],A
LD A,[HLI]
LD A,[HLD]
*/

// execute instruction LD_X_Y
func (c *SM83_CPU) executeInstruction_LD_X_Y(r *uint8, reg string, vr uint8, vreg string) error {
	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		*r = vr
	}

	if c.trace {
		fmt.Printf("[trace] LD %s, %s: 0x%02x\n", reg, vreg, vr)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_X_n
func (c *SM83_CPU) executeInstruction_LD_X_n(r *uint8, reg string) error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_msb, err = c.fetchInstructionArgument()
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		*r = c.n_msb
	}

	if c.trace {
		fmt.Printf("[trace] LD %s, n: 0x%02x\n", reg, *r)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

/*
LD r16,n16
LD [HL],r8
LD [HL],n8
*/

// execute instruction LD_X_ADDR_HL
func (c *SM83_CPU) executeInstruction_LD_X_ADDR_HL(r *uint8, reg string) error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_lsb, err = c.readByteFromMemory(uint16(c.h)<<8 | uint16(c.l))
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		*r = c.n_lsb
	}

	if c.trace {
		fmt.Printf("[trace] LD %s, (HL): 0x%02x\n", reg, c.n_lsb)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

/*
LD [r16],A
LD [n16],A
LDH [n16],A
LDH [C],A
LD A,[r16]
LD A,[n16]
LDH A,[n16]
LDH A,[C]
LD [HLI],A
LD [HLD],A
LD A,[HLI]
LD A,[HLD]
*/

// execute instruction LD_ADDR_XX_Y
func (c *SM83_CPU) executeInstruction_LD_ADDR_XX_Y(msr uint8, lsr uint8, reg string, vr uint8, vreg string) error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		err = c.writeByteIntoMemory(uint16(msr)<<8|uint16(lsr), vr)
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
	}

	if c.trace {
		fmt.Printf("[trace] LD (%s), %s: 0x%02x\n", reg, vreg, vr)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_XX_nn
func (c *SM83_CPU) executeInstruction_LD_XX_nn(msr *uint8, lsr *uint8, reg string) error {
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
		*msr = c.n_msb
		*lsr = c.n_lsb
	}

	if c.trace {
		fmt.Printf("[trace] LD %s, nn: 0x%02x%02x\n", reg, *msr, *lsr)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_A_ADDR_XX
func (c *SM83_CPU) executeInstruction_LD_A_ADDR_XX(msr uint8, lsr uint8, reg string) error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_lsb, err = c.readByteFromMemory(uint16(msr)<<8 | uint16(lsr))
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		c.a = c.n_lsb
	}

	if c.trace {
		fmt.Printf("[trace] LD A, (%s): 0x%02x\n", reg, c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}
