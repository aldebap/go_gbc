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
ADC A,r8  	--> ADC_X       (https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7#ADC_A,r8)
ADC A,[HL]  --> ADC_ADDR_HL (https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7#ADC_A,_HL_)
ADC A,n8    --> ADC_n       (https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7#ADC_A,n8)
ADD A,r8
ADD A,[HL]
ADD A,n8
CP A,r8
CP A,[HL]
CP A,n8
DEC r8     --> DEC_X
DEC [HL]
INC r8     --> INC_X
INC [HL]   -->
SBC A,r8   -->
SBC A,[HL] -->
SBC A,n8
SUB A,r8
SUB A,[HL]
SUB A,n8
*/

// execute instruction ADC_X
func (c *SM83_CPU) executeInstruction_ADC_X(r uint8, reg string) error {

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		var aux16 = uint16(c.a) + uint16(r)

		if (c.flags & FLAG_C) != 0 {
			aux16++
		}

		c.flags = 0x00

		if aux16&0x00ff == 0 {
			c.flags |= FLAG_Z
		}
		if (c.a&0x0f)+(r&0x0f)+(c.flags&FLAG_C) > 0x0f {
			c.flags |= FLAG_H
		}
		if aux16&0xff00 != 0 {
			c.flags |= FLAG_C
		}

		c.a = uint8(aux16 & 0x00ff)
	}

	if c.trace {
		fmt.Printf("[trace] ADC %s: 0x%02x\n", reg, r)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction ADC_ADDR_HL
func (c *SM83_CPU) executeInstruction_ADC_ADDR_HL() error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_lsb, err = c.readByteFromMemory(uint16(c.h)<<8 | uint16(c.l))
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		var aux16 = uint16(c.a) + uint16(c.n_lsb)

		if (c.flags & FLAG_C) != 0 {
			aux16++
		}

		c.flags = 0x00

		if aux16&0x00ff == 0 {
			c.flags |= FLAG_Z
		}
		if (c.a&0x0f)+(c.n_lsb&0x0f)+(c.flags&FLAG_C) > 0x0f {
			c.flags |= FLAG_H
		}
		if aux16&0xff00 != 0 {
			c.flags |= FLAG_C
		}

		c.a = uint8(aux16 & 0x00ff)
	}

	if c.trace {
		fmt.Printf("[trace] ADC (HL): 0x%02x\n", c.n_lsb)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction ADC_n
func (c *SM83_CPU) executeInstruction_ADC_n() error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_lsb, err = c.fetchInstructionArgument()
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		var aux16 = uint16(c.a) + uint16(c.n_lsb)

		if (c.flags & FLAG_C) != 0 {
			aux16++
		}

		c.flags = 0x00

		if aux16&0x00ff == 0 {
			c.flags |= FLAG_Z
		}
		if (c.a&0x0f)+(c.n_lsb&0x0f)+(c.flags&FLAG_C) > 0x0f {
			c.flags |= FLAG_H
		}
		if aux16&0xff00 != 0 {
			c.flags |= FLAG_C
		}

		c.a = uint8(aux16 & 0x00ff)
	}

	if c.trace {
		fmt.Printf("[trace] ADC n: 0x%02x\n", c.n_lsb)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

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
