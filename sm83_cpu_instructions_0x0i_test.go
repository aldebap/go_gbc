////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_instructions_0x0i_test.go - Aug-7-2025 by aldebap
//
//	Test cases for Sharp SM83 CPU - instructions 0x00 - 0x0f
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"
)

// NOP instruction unit tests
func Test_NOP(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> NOP (0x%02x): scenario 1 - do nothing", NOP), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			NOP,
			NOP,
		})
		if err != nil {
			t.Errorf("fail loading test program: %s", err.Error())
		}

		//	connect the ROM memory to the CPU
		err = cpu.ConnectMemory(rom, 0x0000)
		if err != nil {
			t.Errorf("fail connecting ROM to CPU: %s", err.Error())
		}

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x0002, 0x0000, 0x00, 0x00, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_NOP()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction NOP: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// RLCA instruction unit tests
func Test_RLCA(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> RLCA (0x%02x): scenario 1 - no circular bit", RLCA), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			RLCA,
			NOP,
		})
		if err != nil {
			t.Errorf("fail loading test program: %s", err.Error())
		}

		//	connect the ROM memory to the CPU
		err = cpu.ConnectMemory(rom, 0x0000)
		if err != nil {
			t.Errorf("fail connecting ROM to CPU: %s", err.Error())
		}

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x0002, 0x0000, 0x00, 0x80, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.a = 0x40
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_RLCA()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction RLCA: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> RLCA (0x%02x): scenario 2 - circular bit", RLCA), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			RLCA,
			NOP,
		})
		if err != nil {
			t.Errorf("fail loading test program: %s", err.Error())
		}

		//	connect the ROM memory to the CPU
		err = cpu.ConnectMemory(rom, 0x0000)
		if err != nil {
			t.Errorf("fail connecting ROM to CPU: %s", err.Error())
		}

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x0002, 0x0000, 0x00, 0x81, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.a = 0xc0
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_RLCA()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction RLCA: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD (nn), SP instruction unit tests
func Test_LD_ADDR_nn_SP(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD (nn), SP (0x%02x): scenario 1 - write SP into (nn)", LD_ADDR_nn_SP), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			LD_ADDR_nn_SP,
			0x00,
			0xc0,
			NOP,
		})
		if err != nil {
			t.Errorf("fail loading test program: %s", err.Error())
		}

		//	connect the ROM memory to the CPU
		err = cpu.ConnectMemory(rom, 0x0000)
		if err != nil {
			t.Errorf("fail connecting ROM to CPU: %s", err.Error())
		}

		//	create a new RAM memory bank
		ram := NewRAM_memory(8)
		if ram == nil {
			t.Errorf("fail creating new RAM memory")
		}

		//	connect the RAM memory to the CPU
		err = cpu.ConnectMemory(ram, 0xC000)
		if err != nil {
			t.Errorf("fail connecting RAM to CPU: %s", err.Error())
		}

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x0004, 0xc742, 0x00, 0x00, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + five cicles to execute the instruction
		cpu.s = 0xc7
		cpu.p = 0x42
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 5 {
			err = cpu.executeInstruction_LD_ADDR_nn_SP()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD (nn), SP: expected: %s\n\tresult: %s", want, got)
		}

		p, err := ram.ReadByte(0)
		if err != nil {
			t.Errorf("fail reading first byte from RAM: %s", err.Error())
		}

		s, err := ram.ReadByte(1)
		if err != nil {
			t.Errorf("fail reading second byte from RAM: %s", err.Error())
		}

		if s != cpu.s || p != cpu.p {
			t.Errorf("failed executing instruction LD (nn), SP: RAM expected: %02x%02x\n\tresult: %02x%02x", cpu.s, cpu.p, s, p)
		}
	})
}

// RRCA instruction unit tests
func Test_RRCA(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> RRCA (0x%02x): scenario 1 - no circular bit", RRCA), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			RRCA,
			NOP,
		})
		if err != nil {
			t.Errorf("fail loading test program: %s", err.Error())
		}

		//	connect the ROM memory to the CPU
		err = cpu.ConnectMemory(rom, 0x0000)
		if err != nil {
			t.Errorf("fail connecting ROM to CPU: %s", err.Error())
		}

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x0002, 0x0000, 0x00, 0x20, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.a = 0x40
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_RRCA()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction RLCA: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> RRCA (0x%02x): scenario 2 - circular bit", RRCA), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			RRCA,
			NOP,
		})
		if err != nil {
			t.Errorf("fail loading test program: %s", err.Error())
		}

		//	connect the ROM memory to the CPU
		err = cpu.ConnectMemory(rom, 0x0000)
		if err != nil {
			t.Errorf("fail connecting ROM to CPU: %s", err.Error())
		}

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x0002, 0x0000, 0x00, 0x88, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.a = 0x11
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_RRCA()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction RRCA: expected: %s\n\tresult: %s", want, got)
		}
	})
}
