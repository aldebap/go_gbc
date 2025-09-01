////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_8bitArithmeticInstructions.go - Aug-3-2025 by aldebap
//
//	Emulator for Sharp SM83 CPU - 8 bit arithmetic instructions
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

/*
ADC A,r8
ADC A,[HL]
ADC A,n8
ADD A,r8
ADD A,[HL]
ADD A,n8
CP A,r8
CP A,[HL]
CP A,n8
DEC r8     --> DEC_X
DEC [HL]
INC r8     --> INC_X
INC [HL]
SBC A,r8
SBC A,[HL]
SBC A,n8
SUB A,r8
SUB A,[HL]
SUB A,n8
*/

// execute instruction DEC_X
func (c *SM83_CPU) executeInstruction_DEC_X(r *uint8, reg string) error {

	*r--

	if *r == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] DEC %s: 0x%02x\n", reg, *r)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction INC_X
func (c *SM83_CPU) executeInstruction_INC_X(r *uint8, reg string) error {

	*r++

	if *r == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] INC %s: 0x%02x\n", reg, *r)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}
