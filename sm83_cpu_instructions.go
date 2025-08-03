////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_instructions.go - Aug-3-2025 by aldebap
//
//	Emulator for Sharp SM83 CPU - generic instructions
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

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

/*
// execute instruction INC_BC
func (c *SM83_CPU) executeInstruction_INC_BC() error {

	var bc = uint16(c.b)<<8 | uint16(c.c)

	bc++

	c.b = uint8((bc & 0xff00) >> 8)
	c.c = uint8(bc & 0x00ff)

	if bc == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] INC BC: 0x%02x%02x\n", c.b, c.c)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction INC_B
func (c *SM83_CPU) executeInstruction_INC_B() error {

	c.b++

	if c.b == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] INC B: 0x%02x\n", c.b)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction DEC_B
func (c *SM83_CPU) executeInstruction_DEC_B() error {

	c.b--

	if c.b == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] DEC B: 0x%02x\n", c.b)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_B_n
func (c *SM83_CPU) executeInstruction_LD_B_n() error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_msb, err = c.fetchInstructionArgument()
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		c.b = c.n_msb
	}

	if c.trace {
		fmt.Printf("[trace] LD B, n: 0x%02x\n", c.b)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_ADDR_nn_SP
func (c *SM83_CPU) executeInstruction_LD_ADDR_nn_SP() error {
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
		err = c.writeByteIntoMemory(uint16(c.n_msb)<<8|uint16(c.n_lsb), uint8(c.sp&0x00ff))
		c.cpu_state = EXECUTION_CYCLE_4

		return err

	case EXECUTION_CYCLE_4:
		err = c.writeByteIntoMemory(uint16(c.n_msb)<<8|uint16(c.n_lsb)+1, uint8((c.sp&0xff00)>>8))
		c.cpu_state = EXECUTION_CYCLE_5

		return err

	case EXECUTION_CYCLE_5:
	}

	if c.trace {
		fmt.Printf("[trace] LD (nn), SP: 0x%04x\n", c.sp)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction ADD_HL_BC
func (c *SM83_CPU) executeInstruction_ADD_HL_BC() error {

	var result uint16

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		result = uint16(c.l) + uint16(c.c)
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
			result = uint16(c.h) + uint16(c.b)
		} else {
			result = uint16(c.h) + uint16(c.b) + 1
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
		fmt.Printf("[trace] ADD HL, BC: 0x%02x%02x\n", c.h, c.l)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_A_ADDR_BC
func (c *SM83_CPU) executeInstruction_LD_A_ADDR_BC() error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_lsb, err = c.readByteFromMemory(uint16(c.b)<<8 | uint16(c.c))
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		c.a = c.n_lsb
	}

	if c.trace {
		fmt.Printf("[trace] LD A, (BC): 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction DEC_BC
func (c *SM83_CPU) executeInstruction_DEC_BC() error {

	var bc = uint16(c.b)<<8 | uint16(c.c)

	bc--

	c.b = uint8((bc & 0xff00) >> 8)
	c.c = uint8(bc & 0x00ff)

	if bc == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] DEC BC: 0x%02x%02x\n", c.b, c.c)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}
*/
