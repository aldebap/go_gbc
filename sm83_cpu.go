////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu.go - Apr-23-2025 by aldebap
//
//	Emulator for Sharp SM83 CPU
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

// SM83 CPU states
const (
	FETCHING_INSTRUCTION = 1
	EXECUTION_CYCLE_1    = 2
	EXECUTION_CYCLE_2    = 3
	EXECUTION_CYCLE_3    = 4
	EXECUTION_CYCLE_4    = 5
	EXECUTION_CYCLE_5    = 6
)

// SM83 CPU flags
const (
	FLAG_Z = uint8(0x80)
	FLAG_N = uint8(0x40)
	FLAG_H = uint8(0x20)
	FLAG_C = uint8(0x10)
)

// opcode constants
const (
	NOP           = uint8(0x00)
	LD_BC_nn      = uint8(0x01)
	LD_ADDR_BC_A  = uint8(0x02)
	INC_BC        = uint8(0x03)
	INC_B         = uint8(0x04)
	DEC_B         = uint8(0x05)
	LD_B_n        = uint8(0x06)
	RLCA          = uint8(0x07)
	LD_ADDR_nn_SP = uint8(0x08)
	ADD_HL_BC     = uint8(0x09)
	LD_A_ADDR_BC  = uint8(0x0a)
	DEC_BC        = uint8(0x0b)
	INC_C         = uint8(0x0c)
	DEC_C         = uint8(0x0d)
	LD_C_n        = uint8(0x0e)
	RRCA          = uint8(0x0f)

	STOP         = uint8(0x10)
	LD_DE_nn     = uint8(0x11)
	LD_ADDR_DE_A = uint8(0x12)
	INC_DE       = uint8(0x13)
	INC_D        = uint8(0x14)
	DEC_D        = uint8(0x15)
	LD_D_n       = uint8(0x16)
	RLA          = uint8(0x17)
	JR_E         = uint8(0x18)
	ADD_HL_DE    = uint8(0x19)
	LD_A_ADDR_DE = uint8(0x1a)
	DEC_DE       = uint8(0x1b)
	INC_E        = uint8(0x1c)
	DEC_E        = uint8(0x1d)
	LD_E_n       = uint8(0x1e)
	RRA          = uint8(0x1f)

	JR_NZ_e           = uint8(0x20)
	LD_HL_nn          = uint8(0x21)
	LD_ADDR_HL_PLUS_A = uint8(0x22)
	INC_HL            = uint8(0x23)
	INC_H             = uint8(0x24)
	DEC_H             = uint8(0x25)
	LD_H_n            = uint8(0x26)
	DAA               = uint8(0x27)
	JR_Z_e            = uint8(0x28)
	ADD_HL_HL         = uint8(0x29)
	LD_A_ADDR_HL_PLUS = uint8(0x2a)
	DEC_HL            = uint8(0x2b)
	INC_L             = uint8(0x2c)
	DEC_L             = uint8(0x2d)
	LD_L_n            = uint8(0x2e)
	CPL               = uint8(0x2f)

	JR_NC_e            = uint8(0x30)
	LD_SP_nn           = uint8(0x31)
	LD_ADDR_HL_MINUS_A = uint8(0x32)
	INC_SP             = uint8(0x33)
	INC_A              = uint8(0x3c)
	DEC_A              = uint8(0x3d)
	LD_A_n             = uint8(0x3e)

	LD_A_ADDR_nn = uint8(0xfa)
)

// SM83 CPU internal registers and connections
type SM83_CPU struct {
	pc    uint16
	ir    uint8
	ie    uint8
	a     uint8
	b     uint8
	c     uint8
	d     uint8
	e     uint8
	h     uint8
	l     uint8
	s     uint8
	p     uint8
	flags uint8

	trace     bool
	cpu_state uint8
	n_lsb     uint8
	n_msb     uint8

	memoryBank        []memory
	memoryBankAddress []uint16
}

// create a new SM83 CPU
func NewSM83_CPU(trace bool) *SM83_CPU {

	return &SM83_CPU{
		pc:    0,
		ir:    0,
		ie:    0,
		a:     0,
		b:     0,
		c:     0,
		d:     0,
		e:     0,
		h:     0,
		l:     0,
		s:     0,
		p:     0,
		flags: 0,

		trace:     trace,
		cpu_state: FETCHING_INSTRUCTION,

		memoryBank:        nil,
		memoryBankAddress: nil,
	}
}

// connect a new memory bank to the CPU
func (c *SM83_CPU) ConnectMemory(memoryBank memory, intialAddress uint16) error {
	if c.memoryBank == nil {
		c.memoryBank = make([]memory, 0)
		c.memoryBankAddress = make([]uint16, 0)
	}

	//	check for bank overlapping
	for i := 0; i < len(c.memoryBank); i++ {
		//	TODO: check for overlapping
	}

	c.memoryBank = append(c.memoryBank, memoryBank)
	c.memoryBankAddress = append(c.memoryBankAddress, intialAddress)

	return nil
}

// run one machine cycle
func (c *SM83_CPU) MachineCycle() error {
	var err error

	switch c.cpu_state {
	case FETCHING_INSTRUCTION:
		err = c.fetchInstruction()
		if err != nil {
			if c.trace {
				fmt.Printf("[error] %s\n", err.Error())
			}
			return err
		}

	case EXECUTION_CYCLE_1, EXECUTION_CYCLE_2, EXECUTION_CYCLE_3, EXECUTION_CYCLE_4, EXECUTION_CYCLE_5:
		err = c.executeInstruction()
		if err != nil {
			if c.trace {
				fmt.Printf("[error] %s\n", err.Error())
			}
			return err
		}
	}

	return nil
}

// fetch one instruction from memory
func (c *SM83_CPU) fetchInstruction() error {
	var err error

	for i := 0; i < len(c.memoryBankAddress); i++ {
		if c.pc >= c.memoryBankAddress[i] && c.pc < c.memoryBankAddress[i]+c.memoryBank[i].Len() {
			c.ir, err = c.memoryBank[i].ReadByte(c.pc)
			if err != nil {
				return err
			}
			break
		}
	}

	c.pc++
	c.cpu_state = EXECUTION_CYCLE_1

	return nil
}

// fetch one instruction argument from memory
func (c *SM83_CPU) fetchInstructionArgument() (uint8, error) {
	var err error
	var aux uint8

	for i := 0; i < len(c.memoryBankAddress); i++ {
		if c.pc >= c.memoryBankAddress[i] && c.pc < c.memoryBankAddress[i]+c.memoryBank[i].Len() {
			aux, err = c.memoryBank[i].ReadByte(c.pc)
			if err != nil {
				return 0, err
			}
			break
		}
	}

	c.pc++

	return aux, nil
}

// write a byte into memory address
func (c *SM83_CPU) writeByteIntoMemory(address uint16, value uint8) error {

	for i := 0; i < len(c.memoryBankAddress); i++ {
		if address >= c.memoryBankAddress[i] && address < c.memoryBankAddress[i]+c.memoryBank[i].Len() {
			return c.memoryBank[i].WriteByte(address-c.memoryBankAddress[i], value)
		}
	}

	return fmt.Errorf("no memory bank connected to address: %04x", address)
}

// read a byte from memory address
func (c *SM83_CPU) readByteFromMemory(address uint16) (uint8, error) {

	for i := 0; i < len(c.memoryBankAddress); i++ {
		if address >= c.memoryBankAddress[i] && address < c.memoryBankAddress[i]+c.memoryBank[i].Len() {
			return c.memoryBank[i].ReadByte(address - c.memoryBankAddress[i])
		}
	}

	return 0, fmt.Errorf("no memory bank connected to address: %04x", address)
}

// execute instruction
func (c *SM83_CPU) executeInstruction() error {
	switch c.ir {
	case NOP:
		return c.executeInstruction_NOP()

	case LD_BC_nn:
		return c.executeInstruction_LD_XX_nn(&c.b, &c.c, "BC")

	case LD_ADDR_BC_A:
		return c.executeInstruction_LD_ADDR_XX_Y(c.b, c.c, "BC", c.a, "A")

	case INC_BC:
		return c.executeInstruction_INC_XX(&c.b, &c.c, "BC")

	case INC_B:
		return c.executeInstruction_INC_X(&c.b, "B")

	case DEC_B:
		return c.executeInstruction_DEC_X(&c.b, "B")

	case LD_B_n:
		return c.executeInstruction_LD_X_n(&c.b, "B")

	case RLCA:
		return c.executeInstruction_RLCA()

	case LD_ADDR_nn_SP:
		return c.executeInstruction_LD_ADDR_nn_SP()

	case ADD_HL_BC:
		return c.executeInstruction_ADD_HL_BC()

	case LD_A_ADDR_BC:
		return c.executeInstruction_LD_A_ADDR_BC()

	case DEC_BC:
		return c.executeInstruction_DEC_XX(&c.b, &c.c, "BC")

	case INC_C:
		return c.executeInstruction_INC_X(&c.c, "C")

	case DEC_C:
		return c.executeInstruction_DEC_X(&c.c, "C")

	case LD_C_n:
		return c.executeInstruction_LD_X_n(&c.c, "C")

	case RRCA:
		return c.executeInstruction_RRCA()

	case STOP:
		return c.executeInstruction_STOP()

	case LD_DE_nn:
		return c.executeInstruction_LD_XX_nn(&c.d, &c.e, "DE")

	case LD_ADDR_DE_A:
		return c.executeInstruction_LD_ADDR_XX_Y(c.d, c.e, "DE", c.a, "A")

	case INC_DE:
		return c.executeInstruction_INC_XX(&c.d, &c.e, "DE")

	case INC_D:
		return c.executeInstruction_INC_X(&c.d, "D")

	case DEC_D:
		return c.executeInstruction_DEC_X(&c.d, "D")

	case LD_D_n:
		return c.executeInstruction_LD_X_n(&c.d, "D")

	case RLA:
		return c.executeInstruction_RLA()

	case JR_E:
		return c.executeInstruction_JR_E()

	case ADD_HL_DE:
		return c.executeInstruction_ADD_HL_DE()

	case LD_A_ADDR_DE:
		return c.executeInstruction_LD_A_ADDR_DE()

	case DEC_DE:
		return c.executeInstruction_DEC_XX(&c.d, &c.e, "DE")

	case INC_E:
		return c.executeInstruction_INC_X(&c.e, "E")

	case DEC_E:
		return c.executeInstruction_DEC_X(&c.e, "E")

	case LD_E_n:
		return c.executeInstruction_LD_X_n(&c.e, "E")

	case RRA:
		return c.executeInstruction_RRA()

	case JR_NZ_e:
		return c.executeInstruction_JR_NZ_e()

	case LD_HL_nn:
		return c.executeInstruction_LD_XX_nn(&c.h, &c.l, "HL")

	case LD_ADDR_HL_PLUS_A:
		return c.executeInstruction_LD_ADDR_HL_PLUS_A()

	case INC_HL:
		return c.executeInstruction_INC_XX(&c.h, &c.l, "HL")

	case INC_H:
		return c.executeInstruction_INC_X(&c.h, "H")

	case DEC_H:
		return c.executeInstruction_DEC_X(&c.h, "H")

	case LD_H_n:
		return c.executeInstruction_LD_X_n(&c.h, "H")

	case DAA:
		return c.executeInstruction_DAA()

	case JR_Z_e:
		return nil // TODO: implement JR_Z_e

	case ADD_HL_HL:
		return nil // TODO: implement ADD_HL_HL

	case LD_A_ADDR_HL_PLUS:
		return nil // TODO: implement LD_A_ADDR_HL_PLUS

	case DEC_HL:
		return c.executeInstruction_DEC_XX(&c.h, &c.l, "HL")

	case INC_L:
		return c.executeInstruction_INC_X(&c.l, "L")

	case DEC_L:
		return c.executeInstruction_DEC_X(&c.l, "L")

	case LD_L_n:
		return c.executeInstruction_LD_X_n(&c.l, "L")

	case CPL:
		return nil // TODO: implement CPL

	case JR_NC_e:
		return c.executeInstruction_JR_NC_e()

	case LD_SP_nn:
		return c.executeInstruction_LD_XX_nn(&c.s, &c.p, "SP")

	case LD_ADDR_HL_MINUS_A:
		return nil // TODO: implement LD_ADDR_HL_MINUS_A

	case INC_SP:
		return c.executeInstruction_INC_XX(&c.s, &c.p, "SP")

	case INC_A:
		return c.executeInstruction_INC_X(&c.a, "A")

	case DEC_A:
		return c.executeInstruction_DEC_X(&c.a, "A")

	case LD_A_n:
		return c.executeInstruction_LD_X_n(&c.a, "A")

	case LD_A_ADDR_nn:
		return c.executeInstruction_LD_A_ADDR_nn()
	}

	return nil
}

// dump CPU registers
func (c *SM83_CPU) DumpRegisters() string {
	return fmt.Sprintf("PC: 0x%04x; SP: 0x%02x%02x; Flags: 0x%02x; A: 0x%02x; BC: 0x%02x%02x; DE: 0x%02x%02x; HL: 0x%02x%02x",
		c.pc, c.s, c.p, c.flags, c.a, c.b, c.c, c.d, c.e, c.h, c.l)
}
