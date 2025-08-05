////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_instructions.go - Aug-3-2025 by aldebap
//
//	Emulator for Sharp SM83 CPU - generic instructions
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

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

// execute instruction LD_A_ADDR_BC
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
