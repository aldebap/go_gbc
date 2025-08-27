////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_16bitArithmeticInstructions.go - Aug-26-2025 by aldebap
//
//	Emulator for Sharp SM83 CPU - 16 bit arithmetic instructions
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

/*
ADD HL,r16 --> ADD_HL_XX
DEC r16    --> DEC_XX
INC r16    --> INC_XX
*/

// execute instruction ADD_HL_XX
func (c *SM83_CPU) executeInstruction_ADD_HL_XX(msr uint8, lsr uint8, reg string) error {

	var result uint16

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		result = uint16(c.l) + uint16(lsr)
		c.n_lsb = uint8(result & 0x00ff)

		if result&0x0010 != 0x0000 {
			c.flags |= FLAG_H
		} else {
			c.flags &= ^FLAG_H
		}
		if result&0x0100 != 0x0000 {
			c.flags |= FLAG_C
		} else {
			c.flags &= ^FLAG_C
		}
		c.flags &= ^FLAG_N

		c.cpu_state = EXECUTION_CYCLE_2

		return nil

	case EXECUTION_CYCLE_2:
		if c.flags&FLAG_C == 0x00 {
			result = uint16(c.h) + uint16(msr)
		} else {
			result = uint16(c.h) + uint16(msr) + 1
		}
		c.n_msb = uint8(result & 0x00ff)

		if result&0x0010 != 0x0000 {
			c.flags |= FLAG_H
		} else {
			c.flags &= ^FLAG_H
		}
		if result&0x0100 != 0x0000 {
			c.flags |= FLAG_C
		} else {
			c.flags &= ^FLAG_C
		}
		c.flags &= ^FLAG_N
	}

	c.h = c.n_msb
	c.l = c.n_lsb

	if c.trace {
		fmt.Printf("[trace] ADD HL, %s: 0x%02x%02x\n", reg, c.h, c.l)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction DEC_XX
func (c *SM83_CPU) executeInstruction_DEC_XX(msr *uint8, lsr *uint8, reg string) error {

	var reg16 = uint16(*msr)<<8 | uint16(*lsr)

	reg16--

	*msr = uint8((reg16 & 0xff00) >> 8)
	*lsr = uint8(reg16 & 0x00ff)

	if reg16 == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] DEC %s: 0x%02x%02x\n", reg, *msr, *lsr)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction INC_XX
func (c *SM83_CPU) executeInstruction_INC_XX(msr *uint8, lsr *uint8, reg string) error {

	var reg16 = uint16(*msr)<<8 | uint16(*lsr)

	reg16++

	*msr = uint8((reg16 & 0xff00) >> 8)
	*lsr = uint8(reg16 & 0x00ff)

	if reg16 == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] INC %s: 0x%02x%02x\n", reg, *msr, *lsr)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}
