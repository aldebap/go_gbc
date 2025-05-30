////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_instructions_0x0i.go - Apr-23-2025 by aldebap
//
//	Emulator for Sharp SM83 CPU - instructions 0x00 - 0x0f
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

// execute instruction NOP
func (c *SM83_CPU) executeInstruction_NOP() error {

	if c.trace {
		fmt.Printf("[trace] NOP\n")
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_BC_nn
func (c *SM83_CPU) executeInstruction_LD_BC_nn() error {
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
		c.b = c.n_msb
		c.c = c.n_lsb
	}

	if c.trace {
		fmt.Printf("[trace] LD BC, nn: 0x%02x%02x\n", c.b, c.c)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_ADDR_BC_A
func (c *SM83_CPU) executeInstruction_LD_ADDR_BC_A() error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		err = c.writeByteIntoMemory(uint16(c.b)<<8|uint16(c.c), c.a)
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
	}

	if c.trace {
		fmt.Printf("[trace] LD (BC), A: 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

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

// execute instruction RLCA
func (c *SM83_CPU) executeInstruction_RLCA() error {

	if c.a&0x80 == 0x80 {
		c.a = c.a<<1 | 0x01
	} else {
		c.a = c.a << 1
	}

	if c.a == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] RLCA: 0x%02x\n", c.a)
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

// execute instruction INC_C
func (c *SM83_CPU) executeInstruction_INC_C() error {

	c.c++

	if c.c == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] INC C: 0x%02x\n", c.c)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction DEC_C
func (c *SM83_CPU) executeInstruction_DEC_C() error {

	c.c--

	if c.c == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] DEC C: 0x%02x\n", c.c)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_C_n
func (c *SM83_CPU) executeInstruction_LD_C_n() error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_msb, err = c.fetchInstructionArgument()
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		c.c = c.n_msb
	}

	if c.trace {
		fmt.Printf("[trace] LD C, n: 0x%02x\n", c.c)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction RLCA
func (c *SM83_CPU) executeInstruction_RRCA() error {

	if c.a&0x01 == 0x01 {
		c.a = c.a>>1 | 0x80
	} else {
		c.a = c.a >> 1
	}

	if c.a == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] RRCA: 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}
