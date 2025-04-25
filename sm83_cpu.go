////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu.go - Apr-23-2025 by aldebap
//
//	Emulator for Sharp SM83 CPU
////////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

// SM83 CPU states
const (
	FETCH_INSTRUCTION     = 1
	EXECUTE_INSTRUCTION_1 = 2
	EXECUTE_INSTRUCTION_2 = 3
)

// SM83 CPU flags
const (
	FLAG_Z = uint8(0x01)
	FLAG_N = uint8(0x02)
	FLAG_H = uint8(0x04)
)

// opcode constants
const (
	NOOP   = uint8(0x00)
	INC_A  = uint8(0x3c)
	DEC_A  = uint8(0x3d)
	LD_A_n = uint8(0x3e)
)

// SM83 CPU internal registers and connections
type SM83_CPU struct {
	pc  uint16
	sp  uint16
	ir  uint8
	ie  uint8
	aux uint8

	a     uint8
	flags uint8
	bc    uint16
	de    uint16
	hl    uint16

	cpu_state uint8

	memoryBank        []memory
	memoryBankAddress []uint16
}

// create a new SM83 CPU
func NewSM83_CPU() *SM83_CPU {

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

		cpu_state: FETCH_INSTRUCTION,

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
	case FETCH_INSTRUCTION:
		err = c.fetchInstruction()
		if err != nil {
			return err
		}

	case EXECUTE_INSTRUCTION_1, EXECUTE_INSTRUCTION_2:
		err = c.executeInstruction()
		if err != nil {
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
	c.cpu_state = EXECUTE_INSTRUCTION_1

	return nil
}

// fetch one instruction argument from memory
func (c *SM83_CPU) fetchInstructionArgument(address uint16) (uint8, error) {
	for i := 0; i < len(c.memoryBankAddress); i++ {
		if address >= c.memoryBankAddress[i] && address < c.memoryBankAddress[i]+c.memoryBank[i].Len() {
			return c.memoryBank[i].ReadByte(address)
		}
	}

	return 0, fmt.Errorf("memory address not available")
}

// execute instruction
func (c *SM83_CPU) executeInstruction() error {
	switch c.ir {
	case NOOP:
		return c.fetchInstruction()

	case INC_A:
		return c.executeInstruction_INC_A()

	case DEC_A:
		return c.executeInstruction_DEC_A()

	case LD_A_n:
		return c.executeInstruction_LD_A_n()
	}

	return nil
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
	//fmt.Printf("[debug] INC A: 0x%02x\n", c.a)

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
	//fmt.Printf("[debug] DEC A: 0x%02x\n", c.a)

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction LD_A_n
func (c *SM83_CPU) executeInstruction_LD_A_n() error {
	var err error

	switch c.cpu_state {
	case EXECUTE_INSTRUCTION_1:
		c.aux, err = c.fetchInstructionArgument(c.pc)
		if err != nil {
			return err
		}

		c.pc++
		c.cpu_state = EXECUTE_INSTRUCTION_2

		return nil

	case EXECUTE_INSTRUCTION_2:
		c.a = c.aux
	}
	fmt.Printf("[debug] LD A, n: 0x%02x\n", c.a)

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// dump CPU registers
func (c *SM83_CPU) DumpRegisters() string {
	return fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
		c.pc, c.sp, c.flags, c.a, c.bc, c.de, c.hl)
}
