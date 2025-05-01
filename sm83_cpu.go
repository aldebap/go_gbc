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

	JR_NC_e  = uint8(0x30)
	LD_SP_nn = uint8(0x31)
	INC_A    = uint8(0x3c)
	DEC_A    = uint8(0x3d)
	LD_A_n   = uint8(0x3e)

	LD_A_ADDR_nn = uint8(0xfa)
)

// SM83 CPU internal registers and connections
type SM83_CPU struct {
	pc uint16
	sp uint16
	ir uint8
	ie uint8

	a     uint8
	flags uint8
	bc    uint16
	de    uint16
	hl    uint16

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
		sp:    0,
		ir:    0,
		ie:    0,
		a:     0,
		flags: 0,
		bc:    0,
		de:    0,
		hl:    0,

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
		return c.executeInstruction_LD_BC_nn()

	case LD_ADDR_BC_A:
		return c.executeInstruction_LD_ADDR_BC_A()

	case INC_BC:
		return c.executeInstruction_INC_BC()

	case INC_B:
		return c.executeInstruction_INC_B()

	case DEC_B:
		return c.executeInstruction_DEC_B()

	case LD_B_n:
		return c.executeInstruction_LD_B_n()

	case RLCA:
		return c.executeInstruction_RLCA()

	case LD_ADDR_nn_SP:
		return c.executeInstruction_LD_ADDR_nn_SP()

	case ADD_HL_BC:
		return c.executeInstruction_ADD_HL_BC()

	case LD_A_ADDR_BC:
		return c.executeInstruction_LD_A_ADDR_BC()

	case LD_SP_nn:
		return c.executeInstruction_LD_SP_nn()

	case INC_A:
		return c.executeInstruction_INC_A()

	case DEC_A:
		return c.executeInstruction_DEC_A()

	case LD_A_n:
		return c.executeInstruction_LD_A_n()

	case LD_A_ADDR_nn:
		return c.executeInstruction_LD_A_ADDR_nn()
	}

	return nil
}

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
		c.bc = uint16(c.n_msb)<<8 | uint16(c.n_lsb)
	}

	if c.trace {
		fmt.Printf("[trace] LD BC, nn: 0x%04x\n", c.bc)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_ADDR_BC_A
func (c *SM83_CPU) executeInstruction_LD_ADDR_BC_A() error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		err = c.writeByteIntoMemory(c.bc, c.a)
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

	c.bc++

	if c.bc == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] INC BC: 0x%04x\n", c.bc)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction INC_B
func (c *SM83_CPU) executeInstruction_INC_B() error {

	msb := uint8((c.bc & 0xff00) >> 8)
	lsb := uint8(c.bc & 0x00ff)

	msb++
	c.bc = uint16(msb)<<8 | uint16(lsb)

	if msb == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] INC B: 0x%02x\n", msb)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction DEC_B
func (c *SM83_CPU) executeInstruction_DEC_B() error {

	msb := uint8((c.bc & 0xff00) >> 8)
	lsb := uint8(c.bc & 0x00ff)

	msb--
	c.bc = uint16(msb)<<8 | uint16(lsb)

	if msb == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] DEC B: 0x%02x\n", msb)
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
		c.bc = uint16(c.n_msb)<<8 | uint16(c.bc&0x00ff)
	}

	if c.trace {
		fmt.Printf("[trace] LD B, n: 0x%02x\n", c.n_msb)
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
		result = uint16(c.hl&0x00ff) + uint16(c.bc&0x00ff)
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
			result = uint16(c.hl&0xff00)>>8 + uint16(c.bc&0xff00)>>8
		} else {
			result = uint16(c.hl&0xff00)>>8 + uint16(c.bc&0xff00)>>8 + 1
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

	c.hl = uint16(c.n_msb)<<8 | uint16(c.n_lsb)

	if c.trace {
		fmt.Printf("[trace] ADD HL, BC: 0x%04x\n", c.hl)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_A_ADDR_BC
func (c *SM83_CPU) executeInstruction_LD_A_ADDR_BC() error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_lsb, err = c.readByteFromMemory(c.bc)
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

// execute instruction LD_SP_nn
func (c *SM83_CPU) executeInstruction_LD_SP_nn() error {
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
		c.sp = uint16(c.n_msb)<<8 | uint16(c.n_lsb)
	}

	if c.trace {
		fmt.Printf("[trace] LD SP, nn: 0x%04x\n", c.sp)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction INC_A
func (c *SM83_CPU) executeInstruction_INC_A() error {

	c.a++

	if c.a == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] INC A: 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction DEC_A
func (c *SM83_CPU) executeInstruction_DEC_A() error {

	c.a--

	if c.a == 0x00 {
		c.flags |= FLAG_Z
	} else {
		c.flags &= ^FLAG_Z
	}
	c.flags &= ^FLAG_N

	if c.trace {
		fmt.Printf("[trace] DEC A: 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_A_n
func (c *SM83_CPU) executeInstruction_LD_A_n() error {
	var err error

	switch c.cpu_state {
	case EXECUTION_CYCLE_1:
		c.n_msb, err = c.fetchInstructionArgument()
		c.cpu_state = EXECUTION_CYCLE_2

		return err

	case EXECUTION_CYCLE_2:
		c.a = c.n_msb
	}

	if c.trace {
		fmt.Printf("[trace] LD A, n: 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_A_ADDR_nn
func (c *SM83_CPU) executeInstruction_LD_A_ADDR_nn() error {
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
		c.n_lsb, err = c.readByteFromMemory(uint16(c.n_msb)<<8 | uint16(c.n_lsb))
		c.cpu_state = EXECUTION_CYCLE_4

		return err

	case EXECUTION_CYCLE_4:
		c.a = c.n_lsb
	}

	if c.trace {
		fmt.Printf("[trace] LD A, (nn): 0x%02x\n", c.a)
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// dump CPU registers
func (c *SM83_CPU) DumpRegisters() string {
	return fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
		c.pc, c.sp, c.flags, c.a, c.bc, c.de, c.hl)
}
