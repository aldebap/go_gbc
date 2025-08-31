////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_instructions_0x2i.go - Apr-23-2025 by aldebap
//
//	Emulator for Sharp SM83 CPU - instructions 0x20 - 0x2f
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

// execute instruction JR_NZ_e
func (c *SM83_CPU) executeInstruction_JR_NZ_e() error {

	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_msb, err = c.fetchInstructionArgument()
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		if c.flags&FLAG_Z == 0 {
			c.pc += uint16(int8(c.n_msb))
		}
	}

	if c.trace {
		fmt.Printf("[trace] JR_NZ_e\n")
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction DAA
func (c *SM83_CPU) executeInstruction_DAA() error {

	// TODO: review implementation of DAA
	/*
		DAA: function() {
			var a=Z80._r.a;
			if((Z80._r.f&0x20)||((Z80._r.a&15)>9))
				Z80._r.a+=6;
			Z80._r.f&=0xEF;
			if((Z80._r.f&0x20)||(a>0x99)) {
				Z80._r.a+=0x60;
				Z80._r.f|=0x10;
			}
			Z80._r.m=1;
		},
	*/

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		if c.flags&FLAG_H != 0 || c.a&0x0F > 0x09 {
			c.a += 0x06
		}
		c.cpu_state = EXECUTION_CYCLE_2

		return nil

	case EXECUTION_CYCLE_2:
		c.flags &= ^FLAG_N
		c.cpu_state = EXECUTION_CYCLE_3

		return nil

	case EXECUTION_CYCLE_3:
		if c.flags&FLAG_H != 0 || c.a > 0x99 {
			c.a += 0x60
			c.flags |= FLAG_C
		}
		c.cpu_state = EXECUTION_CYCLE_4

		return nil

	case EXECUTION_CYCLE_4:
	}

	if c.trace {
		fmt.Printf("[trace] DAA: 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction JR_Z_e
func (c *SM83_CPU) executeInstruction_JR_Z_e() error {

	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_msb, err = c.fetchInstructionArgument()
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		if c.flags&FLAG_Z != 0 {
			c.pc += uint16(int8(c.n_msb))
		}
	}

	if c.trace {
		fmt.Printf("[trace] JR_Z_e\n")
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction CPL
func (c *SM83_CPU) executeInstruction_CPL() error {

	// TODO: review implementation of CPL
	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		if c.flags&FLAG_H != 0 || c.a&0x0F > 0x09 {
			c.a += 0x06
		}
		c.cpu_state = EXECUTION_CYCLE_2

		return nil

	case EXECUTION_CYCLE_2:
		c.flags &= ^FLAG_N
		c.cpu_state = EXECUTION_CYCLE_3

		return nil

	case EXECUTION_CYCLE_3:
		if c.flags&FLAG_H != 0 || c.a > 0x99 {
			c.a += 0x60
			c.flags |= FLAG_C
		}
		c.cpu_state = EXECUTION_CYCLE_4

		return nil

	case EXECUTION_CYCLE_4:
	}

	if c.trace {
		fmt.Printf("[trace] CPL: 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}
