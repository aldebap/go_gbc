////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_instructions_0x0i_test.go - Apr-24-2025 by aldebap
//
//	Test cases for Sharp SM83 CPU - instructions 0x00 - 0x0f
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"

	go_gbc "github.com/aldebap/go_gbc"
)

// NOP instruction unit tests
func Test_NOP(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> NOP (0x%02x): scenario 1 - do nothing", go_gbc.NOP), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.NOP,
			go_gbc.NOP,
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

		//	two cicles to execute the test program
		for range 2 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction NOP: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD_BC_nn instruction unit tests
func Test_LD_BC_nn(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD BC, nn (0x%02x): scenario 1 - load BC 16 bits register", go_gbc.LD_BC_nn), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_BC_nn,
			0x52,
			0xf0,
			go_gbc.NOP,
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
			0x0004, 0x0000, 0x00, 0x00, 0xf052, 0x0000, 0x0000)

		//	four cicles to execute the test program
		for range 4 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD BC, nn: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD_ADDR_BC_A instruction unit tests
func Test_LD_ADDR_BC_A(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD (BC), A (0x%02x): scenario 1 - write A into (BC)", go_gbc.LD_ADDR_BC_A), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_BC_nn,
			0x00,
			0xC0,
			go_gbc.LD_A_n,
			0x6c,
			go_gbc.LD_ADDR_BC_A,
			go_gbc.LD_A_n,
			0x00,
			go_gbc.LD_A_ADDR_nn,
			0x00,
			0xC0,
			go_gbc.NOP,
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
		ram := go_gbc.NewRAM_memory(8)
		if ram == nil {
			t.Errorf("fail creating new RAM memory")
		}

		//	connect the RAM memory to the CPU
		err = cpu.ConnectMemory(ram, 0xC000)
		if err != nil {
			t.Errorf("fail connecting RAM to CPU: %s", err.Error())
		}

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x000c, 0x0000, 0x00, 0x6c, 0xc000, 0x0000, 0x0000)

		//	eight cicles to execute the test program
		for range 14 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD (BC), A: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// INC_BC instruction unit tests
func Test_INC_BC(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> INC BC (0x%02x): scenario 1 - increment without carry out", go_gbc.INC_BC), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_BC_nn,
			0x07,
			0x00,
			go_gbc.INC_BC,
			go_gbc.NOP,
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
			0x0005, 0x0000, 0x00, 0x00, 0x0008, 0x0000, 0x0000)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction INC BC: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> INC BC (0x%02x): scenario 2 - increment with carry out", go_gbc.INC_BC), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_BC_nn,
			0xff,
			0xff,
			go_gbc.INC_BC,
			go_gbc.NOP,
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
			0x0005, 0x0000, go_gbc.FLAG_Z, 0x00, 0x0000, 0x0000, 0x0000)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction INC BC: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// INC_B instruction unit tests
func Test_INC_B(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> INC B (0x%02x): scenario 1 - increment without carry out", go_gbc.INC_B), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_BC_nn,
			0x07,
			0x2c,
			go_gbc.INC_B,
			go_gbc.NOP,
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
			0x0005, 0x0000, 0x00, 0x00, 0x2d07, 0x0000, 0x0000)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction INC B: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> INC B (0x%02x): scenario 2 - increment with carry out", go_gbc.INC_B), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_BC_nn,
			0x07,
			0xff,
			go_gbc.INC_B,
			go_gbc.NOP,
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
			0x0005, 0x0000, go_gbc.FLAG_Z, 0x00, 0x0007, 0x0000, 0x0000)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction INC B: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// DEC_B instruction unit tests
func Test_DEC_B(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> DEC B (0x%02x): scenario 1 - decrement without carry out", go_gbc.DEC_B), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_BC_nn,
			0x07,
			0x2c,
			go_gbc.DEC_B,
			go_gbc.NOP,
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
			0x0005, 0x0000, 0x00, 0x00, 0x2b07, 0x0000, 0x0000)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction DEC B: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> DEC B (0x%02x): scenario 2 - decrement with carry out", go_gbc.DEC_B), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_BC_nn,
			0x07,
			0x00,
			go_gbc.DEC_B,
			go_gbc.NOP,
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
			0x0005, 0x0000, 0x00, 0x00, 0xff07, 0x0000, 0x0000)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction DEC B: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD_B_n instruction unit tests
func Test_LD_B_n(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD B, n (0x%02x): scenario 1 - load B 8 bits register", go_gbc.LD_B_n), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_B_n,
			0x7e,
			go_gbc.NOP,
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
			0x0003, 0x0000, 0x00, 0x00, 0x7e00, 0x0000, 0x0000)

		//	three cicles to execute the test program
		for range 3 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD B, n: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// RLCA instruction unit tests
func Test_RLCA(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> RLCA (0x%02x): scenario 1 - no circular bit", go_gbc.RLCA), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_A_n,
			0x40,
			go_gbc.RLCA,
			go_gbc.NOP,
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
			t.Errorf("failed executing instruction RLCA: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> RLCA (0x%02x): scenario 2 - circular bit", go_gbc.RLCA), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_A_n,
			0xc0,
			go_gbc.RLCA,
			go_gbc.NOP,
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
			0x0004, 0x0000, 0x00, 0x81, 0x0000, 0x0000, 0x0000)

		//	four cicles to execute the test program
		for range 4 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction RLCA: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD_ADDR_nn_SP instruction unit tests
func Test_LD_ADDR_nn_SP(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD (nn), SP (0x%02x): scenario 1 - write SP into (nn)", go_gbc.LD_ADDR_nn_SP), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_SP_nn,
			0x42,
			0xC7,
			go_gbc.LD_ADDR_nn_SP,
			0x00,
			0xc0,
			go_gbc.LD_A_ADDR_nn,
			0x00,
			0xc0,
			go_gbc.NOP,
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
		ram := go_gbc.NewRAM_memory(8)
		if ram == nil {
			t.Errorf("fail creating new RAM memory")
		}

		//	connect the RAM memory to the CPU
		err = cpu.ConnectMemory(ram, 0xC000)
		if err != nil {
			t.Errorf("fail connecting RAM to CPU: %s", err.Error())
		}

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x000a, 0xc742, 0x00, 0x42, 0x0000, 0x0000, 0x0000)

		//	eight cicles to execute the test program
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

	err = nil

	t.Run(fmt.Sprintf(">>> ADD HL, BC (0x%02x): scenario 1 - adding BC to HL without carry", go_gbc.ADD_HL_BC), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.NOP,
		})
		if err != nil {
			t.Errorf("fail loading test program: %s", err.Error())
		}

		//	connect the ROM memory to the CPU
		err = cpu.ConnectMemory(rom, 0x0000)
		if err != nil {
			t.Errorf("fail connecting ROM to CPU: %s", err.Error())
		}
		//	TODO: need to create this test scenario
	})
}

// LD_A_ADDR_BC instruction unit tests
func Test_LD_A_ADDR_BC(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD A, (BC) (0x%02x): scenario 1 - load acumulator from memory", go_gbc.LD_A_ADDR_BC), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_BC_nn,
			0x05,
			0x00,
			go_gbc.LD_A_ADDR_BC,
			go_gbc.NOP,
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

// DEC_BC instruction unit tests
func Test_DEC_BC(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> DEC BC (0x%02x): scenario 1 - decrement without carry out", go_gbc.DEC_BC), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_BC_nn,
			0x07,
			0x00,
			go_gbc.DEC_BC,
			go_gbc.NOP,
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
			0x0005, 0x0000, 0x00, 0x00, 0x0006, 0x0000, 0x0000)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction DEC BC: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> DEC BC (0x%02x): scenario 2 - decrement with carry out", go_gbc.DEC_BC), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_BC_nn,
			0x00,
			0x00,
			go_gbc.DEC_BC,
			go_gbc.NOP,
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
			0x0005, 0x0000, 0x00, 0x00, 0xffff, 0x0000, 0x0000)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction DEC BC: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// INC_C instruction unit tests
func Test_INC_C(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> INC C (0x%02x): scenario 1 - increment without carry out", go_gbc.INC_C), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_BC_nn,
			0x07,
			0x2c,
			go_gbc.INC_C,
			go_gbc.NOP,
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
			0x0005, 0x0000, 0x00, 0x00, 0x2c08, 0x0000, 0x0000)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction INC C: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> INC C (0x%02x): scenario 2 - increment with carry out", go_gbc.INC_C), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_BC_nn,
			0xff,
			0x07,
			go_gbc.INC_C,
			go_gbc.NOP,
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
			0x0005, 0x0000, go_gbc.FLAG_Z, 0x00, 0x0700, 0x0000, 0x0000)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction INC C: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// DEC_C instruction unit tests
func Test_DEC_C(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> DEC C (0x%02x): scenario 1 - decrement without carry out", go_gbc.DEC_C), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_BC_nn,
			0x07,
			0x2c,
			go_gbc.DEC_C,
			go_gbc.NOP,
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
			0x0005, 0x0000, 0x00, 0x00, 0x2c06, 0x0000, 0x0000)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction DEC C: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> DEC C (0x%02x): scenario 2 - decrement with carry out", go_gbc.DEC_C), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_BC_nn,
			0x00,
			0x07,
			go_gbc.DEC_C,
			go_gbc.NOP,
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
			0x0005, 0x0000, 0x00, 0x00, 0x07ff, 0x0000, 0x0000)

		//	five cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction DEC C: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD_C_n instruction unit tests
func Test_LD_C_n(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD C, n (0x%02x): scenario 1 - load C 8 bits register", go_gbc.LD_C_n), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_C_n,
			0xe7,
			go_gbc.NOP,
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
			0x0003, 0x0000, 0x00, 0x00, 0x00e7, 0x0000, 0x0000)

		//	three cicles to execute the test program
		for range 3 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD C, n: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// RRCA instruction unit tests
func Test_RRCA(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> RRCA (0x%02x): scenario 1 - no circular bit", go_gbc.RRCA), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_A_n,
			0x40,
			go_gbc.RRCA,
			go_gbc.NOP,
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
			0x0004, 0x0000, 0x00, 0x20, 0x0000, 0x0000, 0x0000)

		//	four cicles to execute the test program
		for range 4 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction RLCA: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> RRCA (0x%02x): scenario 2 - circular bit", go_gbc.RRCA), func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := go_gbc.NewSM83_CPU(trace)
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &go_gbc.ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			go_gbc.LD_A_n,
			0x11,
			go_gbc.RRCA,
			go_gbc.NOP,
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
			0x0004, 0x0000, 0x00, 0x88, 0x0000, 0x0000, 0x0000)

		//	four cicles to execute the test program
		for range 4 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction RRCA: expected: %s\n\tresult: %s", want, got)
		}
	})
}
