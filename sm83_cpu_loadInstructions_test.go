////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_loadInstructions_test.go - Aug-3-2025 by aldebap
//
//	Test cases for Sharp SM83 CPU - load instructions
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"
)

// LD X, Y instruction unit tests
func Test_LD_X_Y(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD X, Y: scenario 1 - load B 8 bits register"), func(t *testing.T) {

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
			LD_B_A,
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
			0x0002, 0x0000, 0x00, 0x7e, 0x7e00, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.a = 0x7e
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_LD_X_Y(&cpu.b, "B", cpu.a, "A")
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD X, Y: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD X, n instruction unit tests
func Test_LD_X_n(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD X, n: scenario 1 - load C 8 bits register"), func(t *testing.T) {

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
			LD_C_n,
			0xe7,
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
			0x0003, 0x0000, 0x00, 0x00, 0x00e7, 0x0000, 0x0000)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_LD_X_n(&cpu.c, "C")
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD X, n: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD_XX_nn instruction unit tests
func Test_LD_XX_nn(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD XX, nn: scenario 1 - load BC 16 bits register"), func(t *testing.T) {

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
			0x52,
			0xf0,
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
			0x0004, 0x0000, 0x00, 0x00, 0xf052, 0x0000, 0x0000)

		//	forced fetch instruction + three cicles to execute the instruction
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 3 {
			err = cpu.executeInstruction_LD_XX_nn(&cpu.b, &cpu.c, "BC")
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD XX, nn: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD_ADDR_HL_X instruction unit tests
func Test_LD_ADDR_HL_X(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD (HL), X: scenario 1 - write A into (HL)"), func(t *testing.T) {

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
			LD_ADDR_HL_A,
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
			0x0002, 0x0000, 0x00, 0x6c, 0x0000, 0x0000, 0xc000)
		wantData := uint8(0x6c)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.a = 0x6c
		cpu.h = 0xc0
		cpu.l = 0x00
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_LD_ADDR_HL_X(cpu.a, "A")
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD (HL), X: expected: %s\n\tresult: %s", want, got)
		}

		gotData, err := ram.ReadByte(0x0000)
		if err != nil {
			t.Errorf("fail reading result from RAM: %s", err.Error())
		}

		if wantData != gotData {
			t.Errorf("failed executing instruction LD (HL), X: expected: %02x\n\tresult: %02x", wantData, gotData)
		}
	})
}

// LD_ADDR_HL_n instruction unit tests
func Test_LD_ADDR_HL_n(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD (HL), n: scenario 1 - write n into (HL)"), func(t *testing.T) {

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
			LD_ADDR_HL_n,
			0x7d,
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
			0x0003, 0x0000, 0x00, 0x00, 0x0000, 0x0000, 0xc000)
		wantData := uint8(0x7d)

		//	forced fetch instruction + three cicles to execute the instruction
		cpu.h = 0xc0
		cpu.l = 0x00
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 3 {
			err = cpu.executeInstruction_LD_ADDR_HL_n()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD (HL), n: expected: %s\n\tresult: %s", want, got)
		}

		gotData, err := ram.ReadByte(0x0000)
		if err != nil {
			t.Errorf("fail reading result from RAM: %s", err.Error())
		}

		if wantData != gotData {
			t.Errorf("failed executing instruction LD (HL), n: expected: %02x\n\tresult: %02x", wantData, gotData)
		}
	})
}

// LD_X_ADDR_HL instruction unit tests
func Test_LD_X_ADDR_HL(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD X, (HL): scenario 1 - load X from (HL)"), func(t *testing.T) {

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
			LD_C_ADDR_HL,
			NOP,
			0xa5,
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
			0x0002, 0x0000, 0x00, 0x00, 0x00a5, 0x0000, 0x0002)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.h = 0x00
		cpu.l = 0x02
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_LD_X_ADDR_HL(&cpu.c, "C")
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD X, (HL): expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD_ADDR_XX_A instruction unit tests
func Test_LD_ADDR_XX_A(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD (XX), A: scenario 1 - write A into (XX)"), func(t *testing.T) {

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
			LD_ADDR_DE_A,
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
			0x0002, 0x0000, 0x00, 0x49, 0x0000, 0xc000, 0x0000)
		wantData := uint8(0x49)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.a = 0x49
		cpu.d = 0xc0
		cpu.e = 0x00
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_LD_ADDR_XX_A(cpu.d, cpu.e, "DE")
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD (XX), A: expected: %s\n\tresult: %s", want, got)
		}

		gotData, err := ram.ReadByte(0x0000)
		if err != nil {
			t.Errorf("fail reading result from RAM: %s", err.Error())
		}

		if wantData != gotData {
			t.Errorf("failed executing instruction LD (XX), A: expected: %02x\n\tresult: %02x", wantData, gotData)
		}
	})
}

// LD_ADDR_nn_A instruction unit tests
func Test_LD_ADDR_nn_A(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD (nn), A: scenario 1 - write A into (nn)"), func(t *testing.T) {

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
			LD_ADDR_nn_A,
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
			0x0004, 0x0000, 0x00, 0x32, 0x0000, 0x0000, 0x0000)
		wantData := uint8(0x32)

		//	forced fetch instruction + efour cicles to execute the instruction
		cpu.a = 0x32
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 4 {
			err = cpu.executeInstruction_LD_ADDR_nn_A()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD (nn), A: expected: %s\n\tresult: %s", want, got)
		}

		gotData, err := ram.ReadByte(0x0000)
		if err != nil {
			t.Errorf("fail reading result from RAM: %s", err.Error())
		}

		if wantData != gotData {
			t.Errorf("failed executing instruction LD (nn), A: expected: %02x\n\tresult: %02x", wantData, gotData)
		}
	})
}

// LDH_ADDR_n_A instruction unit tests
func Test_LDH_ADDR_n_A(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LDH (n), A: scenario 1 - write A into (0xFF n)"), func(t *testing.T) {

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
			LDH_ADDR_n_A,
			0x02,
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
		err = cpu.ConnectMemory(ram, 0xff00)
		if err != nil {
			t.Errorf("fail connecting RAM to CPU: %s", err.Error())
		}

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x0003, 0x0000, 0x00, 0x46, 0x0000, 0x0000, 0x0000)
		wantData := uint8(0x46)

		//	forced fetch instruction + three cicles to execute the instruction
		cpu.a = 0x46
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 3 {
			err = cpu.executeInstruction_LDH_ADDR_n_A()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LDH (n), A: expected: %s\n\tresult: %s", want, got)
		}

		gotData, err := ram.ReadByte(0x0002)
		if err != nil {
			t.Errorf("fail reading result from RAM: %s", err.Error())
		}

		if wantData != gotData {
			t.Errorf("failed executing instruction LDH (n), A: expected: %02x\n\tresult: %02x", wantData, gotData)
		}
	})
}

// LDH_ADDR_C_A instruction unit tests
func Test_LDH_ADDR_C_A(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LDH (C), A: scenario 1 - write A into (0xFF C)"), func(t *testing.T) {

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
			LDH_ADDR_C_A,
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
		err = cpu.ConnectMemory(ram, 0xff00)
		if err != nil {
			t.Errorf("fail connecting RAM to CPU: %s", err.Error())
		}

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x0002, 0x0000, 0x00, 0x21, 0x0004, 0x0000, 0x0000)
		wantData := uint8(0x21)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.a = 0x21
		cpu.c = 0x04
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_LDH_ADDR_C_A()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LDH (C), A: expected: %s\n\tresult: %s", want, got)
		}

		gotData, err := ram.ReadByte(0x0004)
		if err != nil {
			t.Errorf("fail reading result from RAM: %s", err.Error())
		}

		if wantData != gotData {
			t.Errorf("failed executing instruction LDH (C), A: expected: %02x\n\tresult: %02x", wantData, gotData)
		}
	})
}

// LD_A_ADDR_XX instruction unit tests
func Test_LD_A_ADDR_XX(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD A, (XX): scenario 1 - load A from (BC)"), func(t *testing.T) {

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
			LD_A_ADDR_BC,
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

		//	create a new cartrige ROM memory bank
		cartridgeRom := &ROM_memory{}
		if cartridgeRom == nil {
			t.Errorf("fail creating new cartridge ROM memory")
		}
		err = cartridgeRom.Load([]uint8{
			0x08,
			0x09,
			0x1a,
			0x2b,
			0x3c,
			0x4d,
			0x5e,
			0x6f,
		})
		if err != nil {
			t.Errorf("fail loading test program: %s", err.Error())
		}

		//	connect the RAM memory to the CPU
		err = cpu.ConnectMemory(cartridgeRom, 0x2000)
		if err != nil {
			t.Errorf("fail connecting cartridge ROM to CPU: %s", err.Error())
		}

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x0002, 0x0000, 0x00, 0x4d, 0x2005, 0x0000, 0x0000)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.b = 0x20
		cpu.c = 0x05
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_LD_A_ADDR_XX(cpu.b, cpu.c, "BC")
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD A, (XX): expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD_A_ADDR_nn instruction unit tests
func Test_LD_A_ADDR_nn(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD A, (nn): scenario 1 - load A from (nn)"), func(t *testing.T) {

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
			LD_A_ADDR_nn,
			0x03,
			0x20,
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

		//	create a new cartrige ROM memory bank
		cartridgeRom := &ROM_memory{}
		if cartridgeRom == nil {
			t.Errorf("fail creating new cartridge ROM memory")
		}
		err = cartridgeRom.Load([]uint8{
			0x08,
			0x09,
			0x1a,
			0x2b,
			0x3c,
			0x4d,
			0x5e,
			0x6f,
		})
		if err != nil {
			t.Errorf("fail loading test program: %s", err.Error())
		}

		//	connect the RAM memory to the CPU
		err = cpu.ConnectMemory(cartridgeRom, 0x2000)
		if err != nil {
			t.Errorf("fail connecting cartridge ROM to CPU: %s", err.Error())
		}

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x0004, 0x0000, 0x00, 0x2b, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + four cicles to execute the instruction
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 4 {
			err = cpu.executeInstruction_LD_A_ADDR_nn()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD A, (nn): expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LDH_A_ADDR_n instruction unit tests
func Test_LDH_A_ADDR_n(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LDH A, (n): scenario 1 - load A from (0xFF n)"), func(t *testing.T) {

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
			LDH_A_ADDR_n,
			0x04,
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

		//	create a new cartrige ROM memory bank
		cartridgeRom := &ROM_memory{}
		if cartridgeRom == nil {
			t.Errorf("fail creating new cartridge ROM memory")
		}
		err = cartridgeRom.Load([]uint8{
			0x08,
			0x09,
			0x1a,
			0x2b,
			0x3c,
			0x4d,
			0x5e,
			0x6f,
		})
		if err != nil {
			t.Errorf("fail loading test program: %s", err.Error())
		}

		//	connect the RAM memory to the CPU
		err = cpu.ConnectMemory(cartridgeRom, 0xff00)
		if err != nil {
			t.Errorf("fail connecting cartridge ROM to CPU: %s", err.Error())
		}

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x0003, 0x0000, 0x00, 0x3c, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + three cicles to execute the instruction
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 3 {
			err = cpu.executeInstruction_LDH_A_ADDR_n()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LDH A, (n): expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LDH_A_ADDR_C instruction unit tests
func Test_LDH_A_ADDR_C(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LDH A, (C): scenario 1 - load A from (0xFF C)"), func(t *testing.T) {

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
			LDH_A_ADDR_C,
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

		//	create a new cartrige ROM memory bank
		cartridgeRom := &ROM_memory{}
		if cartridgeRom == nil {
			t.Errorf("fail creating new cartridge ROM memory")
		}
		err = cartridgeRom.Load([]uint8{
			0x08,
			0x09,
			0x1a,
			0x2b,
			0x3c,
			0x4d,
			0x5e,
			0x6f,
		})
		if err != nil {
			t.Errorf("fail loading test program: %s", err.Error())
		}

		//	connect the RAM memory to the CPU
		err = cpu.ConnectMemory(cartridgeRom, 0xff00)
		if err != nil {
			t.Errorf("fail connecting cartridge ROM to CPU: %s", err.Error())
		}

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x0002, 0x0000, 0x00, 0x1a, 0x0002, 0x0000, 0x0000)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.c = 0x02
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_LDH_A_ADDR_C()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LDH A, (C): expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD_ADDR_HLI_A instruction unit tests
func Test_LD_ADDR_HLI_A(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD (HL+), X: scenario 1 - write A into (HL+)"), func(t *testing.T) {

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
			LD_ADDR_HL_A,
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
			0x0002, 0x0000, 0x00, 0xa1, 0x0000, 0x0000, 0xc001)
		wantData := uint8(0xa1)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.a = 0xa1
		cpu.h = 0xc0
		cpu.l = 0x00
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_LD_ADDR_HLI_A()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD (HL+), A: expected: %s\n\tresult: %s", want, got)
		}

		gotData, err := ram.ReadByte(0x0000)
		if err != nil {
			t.Errorf("fail reading result from RAM: %s", err.Error())
		}

		if wantData != gotData {
			t.Errorf("failed executing instruction LD (HL+), A: expected: %02x\n\tresult: %02x", wantData, gotData)
		}
	})
}

// LD_ADDR_HLD_A instruction unit tests
func Test_LD_ADDR_HLD_A(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD (HL-), X: scenario 1 - write A into (HL-)"), func(t *testing.T) {

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
			LD_ADDR_HL_A,
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
			0x0002, 0x0000, 0x00, 0xb2, 0x0000, 0x0000, 0xc006)
		wantData := uint8(0xb2)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.a = 0xb2
		cpu.h = 0xc0
		cpu.l = 0x07
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_LD_ADDR_HLD_A()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD (HL-), A: expected: %s\n\tresult: %s", want, got)
		}

		gotData, err := ram.ReadByte(0x0007)
		if err != nil {
			t.Errorf("fail reading result from RAM: %s", err.Error())
		}

		if wantData != gotData {
			t.Errorf("failed executing instruction LD (HL-), A: expected: %02x\n\tresult: %02x", wantData, gotData)
		}
	})
}

// LD_A_ADDR_HLI instruction unit tests
func Test_LD_A_ADDR_HLI(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD A, (HL+): scenario 1 - load A from (HL+)"), func(t *testing.T) {

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
			LD_A_ADDR_HLI,
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

		//	create a new cartrige ROM memory bank
		cartridgeRom := &ROM_memory{}
		if cartridgeRom == nil {
			t.Errorf("fail creating new cartridge ROM memory")
		}
		err = cartridgeRom.Load([]uint8{
			0x08,
			0x09,
			0x1a,
			0x2b,
			0x3c,
			0x4d,
			0x5e,
			0x6f,
		})
		if err != nil {
			t.Errorf("fail loading test program: %s", err.Error())
		}

		//	connect the RAM memory to the CPU
		err = cpu.ConnectMemory(cartridgeRom, 0xc000)
		if err != nil {
			t.Errorf("fail connecting cartridge ROM to CPU: %s", err.Error())
		}

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x0002, 0x0000, 0x00, 0x3c, 0x0000, 0x0000, 0xc005)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.h = 0xc0
		cpu.l = 0x04
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_LD_A_ADDR_HLI()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD A, (HL+): expected: %s\n\tresult: %s", want, got)
		}
	})
}

// LD_A_ADDR_HLD instruction unit tests
func Test_LD_A_ADDR_HLD(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD A, (HL-): scenario 1 - load A from (HL-)"), func(t *testing.T) {

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
			LD_A_ADDR_HLD,
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

		//	create a new cartrige ROM memory bank
		cartridgeRom := &ROM_memory{}
		if cartridgeRom == nil {
			t.Errorf("fail creating new cartridge ROM memory")
		}
		err = cartridgeRom.Load([]uint8{
			0x08,
			0x09,
			0x1a,
			0x2b,
			0x3c,
			0x4d,
			0x5e,
			0x6f,
		})
		if err != nil {
			t.Errorf("fail loading test program: %s", err.Error())
		}

		//	connect the RAM memory to the CPU
		err = cpu.ConnectMemory(cartridgeRom, 0xc000)
		if err != nil {
			t.Errorf("fail connecting cartridge ROM to CPU: %s", err.Error())
		}

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x0002, 0x0000, 0x00, 0x4d, 0x0000, 0x0000, 0xc004)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.h = 0xc0
		cpu.l = 0x05
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_LD_A_ADDR_HLD()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD A, (HL-): expected: %s\n\tresult: %s", want, got)
		}
	})
}
