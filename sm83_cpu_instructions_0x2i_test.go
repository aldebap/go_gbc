////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_test_instructions_0x2i.go - Aug-1-2025 by aldebap
//
//	Test cases for Sharp SM83 CPU - instructions 0x20 - 0x2f
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"
)

// STOP instruction unit tests
func Test_JR_NZ_e(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> JR_NZ_e (0x%02x): scenario 1 - no jump (Z is 1)", JR_NZ_e), func(t *testing.T) {

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
			LD_B_n,
			0xFF,
			INC_B,
			JR_NZ_e,
			0x05,
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
			0x0006, 0x0000, FLAG_Z, 0x00, 0x0000, 0x0000, 0x0000)

		//	two cicles to execute the test program
		for range 6 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction JR_NZ_e: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> JR_NZ_e (0x%02x): scenario 2 - jump forward", JR_NZ_e), func(t *testing.T) {

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
			LD_B_n,
			0xFE,
			INC_B,
			JR_NZ_e,
			0x05,
			NOP,
			NOP,
			NOP,
			NOP,
			NOP,
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
			0x000B, 0x0000, 0x00, 0x00, 0xFF00, 0x0000, 0x0000)

		//	two cicles to execute the test program
		for range 6 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction JR_NZ_e: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> JR_NZ_e (0x%02x): scenario 3 - jump backwards", JR_NZ_e), func(t *testing.T) {

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
			LD_B_n,
			0xFE,
			INC_B,
			JR_NZ_e,
			0xFB, // 0xFB = -5
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
			0x0001, 0x0000, 0x00, 0x00, 0xFF00, 0x0000, 0x0000)

		//	two cicles to execute the test program
		for range 6 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction JR_NZ_e: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD_HL_nn instruction unit tests
func Test_LD_HL_nn(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD HL, nn (0x%02x): scenario 1 - load HL 16 bits register", LD_HL_nn), func(t *testing.T) {

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
			LD_HL_nn,
			0x25,
			0x84,
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
			0x0004, 0x0000, 0x00, 0x00, 0x0000, 0x0000, 0x8425)

		//	four cicles to execute the test program
		for range 4 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD HL, nn: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// Test_LD_ADDR_HL_PLUS_A instruction unit tests
func Test_LD_ADDR_HL_PLUS_A(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD (HL+), A (0x%02x): scenario 1 - write A into (HL)", LD_ADDR_HL_PLUS_A), func(t *testing.T) {

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
			//	TODO: need to create this test scenario
			LD_DE_nn,
			0x00,
			0xC0,
			LD_A_n,
			0xa8,
			LD_ADDR_DE_A,
			LD_A_n,
			0x00,
			LD_A_ADDR_nn,
			0x00,
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
			0x000c, 0x0000, 0x00, 0xa8, 0x0000, 0xc000, 0x0000)

		//	eight cicles to execute the test program
		for range 14 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD (DE), A: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// INC_HL instruction unit tests
func Test_INC_HL(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> INC HL (0x%02x): scenario 1 - increment without carry out", INC_HL), func(t *testing.T) {

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
			LD_HL_nn,
			0x05,
			0x21,
			INC_HL,
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
			0x0005, 0x0000, 0x00, 0x00, 0x0000, 0x0000, 0x2106)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction INC HL: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> INC HL (0x%02x): scenario 2 - increment with carry out", INC_HL), func(t *testing.T) {

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
			LD_HL_nn,
			0xff,
			0xff,
			INC_HL,
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
			0x0005, 0x0000, FLAG_Z, 0x00, 0x0000, 0x0000, 0x0000)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction INC HL: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// INC_H instruction unit tests
func Test_INC_H(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> INC H (0x%02x): scenario 1 - increment without carry out", INC_H), func(t *testing.T) {

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
			LD_HL_nn,
			0xf1,
			0x40,
			INC_H,
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
			0x0005, 0x0000, 0x00, 0x00, 0x0000, 0x0000, 0x41f1)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction INC H: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> INC H (0x%02x): scenario 2 - increment with carry out", INC_H), func(t *testing.T) {

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
			LD_HL_nn,
			0x44,
			0xff,
			INC_H,
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
			0x0005, 0x0000, FLAG_Z, 0x00, 0x0000, 0x0000, 0x0044)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction INC H: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// DEC_H instruction unit tests
func Test_DEC_H(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> DEC H (0x%02x): scenario 1 - decrement without carry out", DEC_H), func(t *testing.T) {

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
			LD_HL_nn,
			0xf1,
			0x40,
			DEC_H,
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
			0x0005, 0x0000, 0x00, 0x00, 0x0000, 0x0000, 0x3ff1)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction DEC H: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> DEC H (0x%02x): scenario 2 - decrement with carry out", DEC_H), func(t *testing.T) {

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
			LD_HL_nn,
			0xf1,
			0x00,
			DEC_H,
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
			0x0005, 0x0000, 0x00, 0x00, 0x0000, 0x0000, 0xfff1)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction DEC H: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD_H_n instruction unit tests
func Test_LD_H_n(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD H, n (0x%02x): scenario 1 - load H 8 bits register", LD_H_n), func(t *testing.T) {

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
			LD_H_n,
			0x74,
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
			0x0003, 0x0000, 0x00, 0x00, 0x0000, 0x0000, 0x7400)

		//	three cicles to execute the test program
		for range 3 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD H, n: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// DAA instruction unit tests
func Test_DAA(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> DAA (0x%02x): scenario 1 - no overflow", DAA), func(t *testing.T) {

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
			//	TODO: need to create this test scenario
			LD_A_n,
			0x40,
			DAA,
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
			0x0004, 0x0000, 0x00, 0x80, 0x0000, 0x0000, 0x0000)

		//	four cicles to execute the test program
		for range 4 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction DAA: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> DAA (0x%02x): scenario 2 - overflow", DAA), func(t *testing.T) {

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
			//	TODO: need to create this test scenario
			LD_A_n,
			0xc0,
			DAA,
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
			0x0004, 0x0000, 0x00, 0x80, 0x0000, 0x0000, 0x0000)

		//	four cicles to execute the test program
		for range 4 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction DAA: expected: %s\n\tresult: %s", want, got)
		}
	})
}
