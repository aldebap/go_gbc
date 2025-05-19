////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_instructions_0x0i_test.go - Apr-24-2025 by aldebap
//
//	Test cases for Sharp SM83 CPU - instructions 0x00 - 0x0f
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"
)

// LD_ADDR_nn_SP instruction unit tests
func Test_LD_ADDR_nn_SP(t *testing.T) {

	var err error

	t.Run(">>> LD (nn), SP: scenario 1 - write SP into (nn)", func(t *testing.T) {

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
			LD_SP_nn,
			0xb1,
			0x72,
			LD_ADDR_nn_SP,
			0x00,
			0xC0,
			LD_A_ADDR_nn,
			0x01,
			0xC0,
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
			0x000a, 0x72b1, 0x00, 0x72, 0x0000, 0x0000, 0x0000)

		//	thirteen cicles to execute the test program
		for range 13 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD (nn), SP: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// ADD_HL_BC instruction unit tests
func Test_ADD_HL_BC(t *testing.T) {

	var err error

	t.Run(">>> ADD HL, BC: scenario 1 - increment without carry out in L + C", func(t *testing.T) {

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
			LD_BC_nn,
			0xb1,
			0x72,
			ADD_HL_BC,
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
			0x0005, 0x0000, FLAG_H, 0x00, 0x72b1, 0x0000, 0x72b1)

		//	six cicles to execute the test program
		for range 6 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction ADD HL, BC: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(">>> ADD HL, BC: scenario 2 - increment with carry out in L + C", func(t *testing.T) {

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
			LD_BC_nn,
			0xf1,
			0x00,
			ADD_HL_BC,
			ADD_HL_BC,
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
			0x0006, 0x0000, 0x00, 0x00, 0x00f1, 0x0000, 0x01e2)

		//	eight cicles to execute the test program
		for range 8 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction ADD HL, BC: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(">>> ADD HL, BC: scenario 3 - increment with carry out in H + B", func(t *testing.T) {

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
			LD_BC_nn,
			0x00,
			0xf1,
			ADD_HL_BC,
			ADD_HL_BC,
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
			0x0006, 0x0000, FLAG_C, 0x00, 0xf100, 0x0000, 0xe200)

		//	eight cicles to execute the test program
		for range 8 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction ADD HL, BC: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD_A_ADDR_BC instruction unit tests
func Test_LD_A_ADDR_BC(t *testing.T) {

	var err error

	t.Run(">>> LD A, (BC): scenario 1 - load acumulator from memory", func(t *testing.T) {

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
			LD_BC_nn,
			0x05,
			0x00,
			LD_A_ADDR_BC,
			NOP,
			0x75,
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
			0x0005, 0x0000, 0x00, 0x75, 0x0005, 0x0000, 0x0000)

		//	six cicles to execute the test program
		for range 6 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD A, (nn): expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD_SP_nn instruction unit tests
func Test_LD_SP_nn(t *testing.T) {

	var err error

	t.Run(">>> LD SP, nn: scenario 1 - load SP 16 bits register", func(t *testing.T) {

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
			LD_SP_nn,
			0x0c,
			0x61,
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
			0x0004, 0x610c, 0x00, 0x00, 0x0000, 0x0000, 0x0000)

		//	four cicles to execute the test program
		for range 4 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD SP, nn: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// INC_A instruction unit tests
func Test_INC_A(t *testing.T) {

	var err error

	t.Run(">>> INC A: scenario 1 - increment without carry out", func(t *testing.T) {

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
			LD_A_n,
			0x7e,
			INC_A,
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
			0x0004, 0x0000, 0x00, 0x7f, 0x0000, 0x0000, 0x0000)

		//	four cicles to execute the test program
		for range 4 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction INC_A: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(">>> INC A: scenario 2 - increment with carry out", func(t *testing.T) {

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
			LD_A_n,
			0xff,
			INC_A,
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
			0x0004, 0x0000, FLAG_Z, 0x00, 0x0000, 0x0000, 0x0000)

		//	four cicles to execute the test program
		for range 4 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction INC_A: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// DEC_A instruction unit tests
func Test_DEC_A(t *testing.T) {

	var err error

	t.Run(">>> DEC A: scenario 1 - decrement without carry out", func(t *testing.T) {

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
			LD_A_n,
			0x7e,
			DEC_A,
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
			0x0004, 0x0000, 0x00, 0x7d, 0x0000, 0x0000, 0x0000)

		//	four cicles to execute the test program
		for range 4 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction DEC_A: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(">>> DEC A: scenario 2 - decrement with carry out", func(t *testing.T) {

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
			LD_A_n,
			0x00,
			DEC_A,
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
			0x0004, 0x0000, 0x00, 0xff, 0x0000, 0x0000, 0x0000)

		//	four cicles to execute the test program
		for range 4 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction DEC_A: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD_A_n instruction unit tests
func Test_LD_A_n(t *testing.T) {

	var err error

	t.Run(">>> LD A, n: scenario 1 - load acumulator", func(t *testing.T) {

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
			LD_A_n,
			0x7e,
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
			0x0003, 0x0000, 0x00, 0x7e, 0x0000, 0x0000, 0x0000)

		//	three cicles to execute the test program
		for range 3 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD A, n: expected: %s\n\tresult: %s", want, got)
		}
	})
}
