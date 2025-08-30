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
LD r8,r8    --> LD_X_Y       (https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7#LD_r8,r8)
LD r8,n8    --> LD_X_n       (https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7#LD_r8,n8)
LD r16,n16  --> LD_XX_nn     (https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7#LD_r16,n16)
LD [HL],r8  --> LD_ADDR_HL_X (https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7#LD__HL_,r8)
LD [HL],n8  --> LD_ADDR_HL_n (https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7#LD__HL_,n8)
LD r8,[HL]  --> LD_X_ADDR_HL (https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7#LD_r8,_HL_)
LD [r16],A  --> LD_ADDR_XX_A (https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7#LD__r16_,A)
LD [n16],A  --> LD_ADDR_nn_A (https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7#LD__n16_,A)
LDH [n16],A --> LDH_ADDR_n_A (https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7#LDH__n16_,A)
LDH [C],A   --> LDH_ADDR_C_A (https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7#LDH__C_,A)
LD A,[r16]  --> LD_A_ADDR_XX (https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7#LD_A,_r16_)
LD A,[n16]
LDH A,[n16]
LDH A,[C]
LD [HLI],A
LD [HLD],A
LD A,[HLI]
LD A,[HLD]
*/

// execute instruction LD_X_Y
func (c *SM83_CPU) executeInstruction_LD_X_Y(x *uint8, x_reg string, y uint8, y_reg string) error {

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		*x = y
	}

	if c.trace {
		fmt.Printf("[trace] LD %s, %s: 0x%02x\n", x_reg, y_reg, y)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_X_n
func (c *SM83_CPU) executeInstruction_LD_X_n(x *uint8, x_reg string) error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_msb, err = c.fetchInstructionArgument()
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		*x = c.n_msb
	}

	if c.trace {
		fmt.Printf("[trace] LD %s, n: 0x%02x\n", x_reg, *x)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_XX_nn
func (c *SM83_CPU) executeInstruction_LD_XX_nn(x_msr *uint8, x_lsr *uint8, x_reg string) error {
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
		*x_msr = c.n_msb
		*x_lsr = c.n_lsb
	}

	if c.trace {
		fmt.Printf("[trace] LD %s, nn: 0x%02x%02x\n", x_reg, *x_msr, *x_lsr)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_ADDR_HL_X
func (c *SM83_CPU) executeInstruction_LD_ADDR_HL_X(x uint8, x_reg string) error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		err = c.writeByteIntoMemory(uint16(c.h)<<8|uint16(c.l), x)
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
	}

	if c.trace {
		fmt.Printf("[trace] LD (HL), %s: 0x%02x\n", x_reg, x)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_ADDR_HL_n
func (c *SM83_CPU) executeInstruction_LD_ADDR_HL_n() error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_lsb, err = c.fetchInstructionArgument()
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		err = c.writeByteIntoMemory(uint16(c.h)<<8|uint16(c.l), c.n_lsb)
		c.cpu_state = EXECUTION_CYCLE_3

		return err

	case EXECUTION_CYCLE_3:
	}

	if c.trace {
		fmt.Printf("[trace] LD (HL), n: 0x%02x\n", c.n_lsb)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_X_ADDR_HL
func (c *SM83_CPU) executeInstruction_LD_X_ADDR_HL(x *uint8, x_reg string) error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_lsb, err = c.readByteFromMemory(uint16(c.h)<<8 | uint16(c.l))
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		*x = c.n_lsb
	}

	if c.trace {
		fmt.Printf("[trace] LD %s, (HL): 0x%02x\n", x_reg, c.n_lsb)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_ADDR_XX_A
func (c *SM83_CPU) executeInstruction_LD_ADDR_XX_A(x_msr uint8, x_lsr uint8, x_reg string) error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		err = c.writeByteIntoMemory(uint16(x_msr)<<8|uint16(x_lsr), c.a)
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
	}

	if c.trace {
		fmt.Printf("[trace] LD (%s), A: 0x%02x\n", x_reg, c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_ADDR_nn_A
func (c *SM83_CPU) executeInstruction_LD_ADDR_nn_A() error {
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
		err = c.writeByteIntoMemory(uint16(c.n_msb)<<8|uint16(c.n_lsb), c.a)
		c.cpu_state = EXECUTION_CYCLE_4

		return err

	case EXECUTION_CYCLE_4:
	}

	if c.trace {
		fmt.Printf("[trace] LD (0x%02x%02x), A: 0x%02x\n", c.n_msb, c.n_lsb, c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LDH_ADDR_n_A
func (c *SM83_CPU) executeInstruction_LDH_ADDR_n_A() error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_lsb, err = c.fetchInstructionArgument()
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		c.n_msb = 0xff
		err = c.writeByteIntoMemory(uint16(c.n_msb)<<8|uint16(c.n_lsb), c.a)
		c.cpu_state = EXECUTION_CYCLE_3

		return err

	case EXECUTION_CYCLE_3:
	}

	if c.trace {
		fmt.Printf("[trace] LDH (0x%02x%02x), A: 0x%02x\n", c.n_msb, c.n_lsb, c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LDH_ADDR_C_A
func (c *SM83_CPU) executeInstruction_LDH_ADDR_C_A() error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_lsb = c.c
		c.n_msb = 0xff
		err = c.writeByteIntoMemory(uint16(c.n_msb)<<8|uint16(c.n_lsb), c.a)
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
	}

	if c.trace {
		fmt.Printf("[trace] LDH (C), A: 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_A_ADDR_XX
func (c *SM83_CPU) executeInstruction_LD_A_ADDR_XX(x_msr uint8, x_lsr uint8, x_reg string) error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_lsb, err = c.readByteFromMemory(uint16(x_msr)<<8 | uint16(x_lsr))
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		c.a = c.n_lsb
	}

	if c.trace {
		fmt.Printf("[trace] LD A, (%s): 0x%02x\n", x_reg, c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

/*
LD A,[n16]
LDH A,[n16]
LDH A,[C]
LD [HLI],A
LD [HLD],A
LD A,[HLI]
LD A,[HLD]
*/
