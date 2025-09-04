////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_8bitArithmeticInstructions_test.go - Aug-3-2025 by aldebap
//
//	Test cases for Sharp SM83 CPU - 83 CPU - 8 bit arithmetic instructions
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"
)

// ADC_X instruction unit tests
func Test_ADC_X(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> ADC_X: scenario 1 - add B + carry = 0, without carry out"), func(t *testing.T) {

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
			ADC_B,
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
			0x0002, 0x0000, 0x00, 0x0c, 0x0200, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.a = 0x0a
		cpu.b = 0x02
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_ADC_X(cpu.b, "B")
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction ADC X: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> ADC_X: scenario 2 - add B + carry = 1, without carry out"), func(t *testing.T) {

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
			ADC_B,
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
			0x0002, 0x0000, 0x00, 0x0d, 0x0200, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.a = 0x0a
		cpu.b = 0x02
		cpu.flags = FLAG_C
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_ADC_X(cpu.b, "B")
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction ADC X: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> ADC_X: scenario 3 - add B + carry = 0, with carry out / zero"), func(t *testing.T) {

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
			ADC_B,
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
			0x0002, 0x0000, 0xb0, 0x00, 0x0200, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.a = 0xfe
		cpu.b = 0x02
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_ADC_X(cpu.b, "B")
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction ADC X: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> ADC_X: scenario 4 - add B + carry = 0, with carry out / non zero"), func(t *testing.T) {

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
			ADC_B,
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
			0x0002, 0x0000, 0x30, 0x01, 0x0200, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.a = 0xff
		cpu.b = 0x02
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_ADC_X(cpu.b, "B")
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction ADC X: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// ADC_ADDR_HL instruction unit tests
func Test_ADC_ADDR_HL(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> ADC_(HL): scenario 1 - add (HL) + carry = 0, without carry out"), func(t *testing.T) {

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
			ADC_ADDR_HL,
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
			0x01,
			0x02,
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
			0x0002, 0x0000, 0x00, 0x0c, 0x0000, 0x0000, 0xc001)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.a = 0x0a
		cpu.h = 0xc0
		cpu.l = 0x01
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_ADC_ADDR_HL()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction ADC (HL): expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> ADC_(HL): scenario 2 - add (HL) + carry = 1, without carry out"), func(t *testing.T) {

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
			ADC_ADDR_HL,
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
			0x01,
			0x02,
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
			0x0002, 0x0000, 0x00, 0x0d, 0x0000, 0x0000, 0xc001)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.a = 0x0a
		cpu.h = 0xc0
		cpu.l = 0x01
		cpu.flags = FLAG_C
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_ADC_ADDR_HL()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction ADC (HL): expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> ADC_(HL): scenario 3 - add (HL) + carry = 0, with carry out / zero"), func(t *testing.T) {

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
			ADC_ADDR_HL,
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
			0x01,
			0x02,
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
			0x0002, 0x0000, 0xb0, 0x00, 0x0000, 0x0000, 0xc005)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.a = 0xb3
		cpu.h = 0xc0
		cpu.l = 0x05
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_ADC_ADDR_HL()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction ADC (HL): expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> ADC_(HL): scenario 4 - add (HL) + carry = 0, with carry out / non zero"), func(t *testing.T) {

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
			ADC_ADDR_HL,
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
			0x01,
			0x02,
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
			0x0002, 0x0000, 0x30, 0x10, 0x0000, 0x0000, 0xc005)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.a = 0xc3
		cpu.h = 0xc0
		cpu.l = 0x05
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_ADC_ADDR_HL()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction ADC (HL): expected: %s\n\tresult: %s", want, got)
		}
	})
}

// ADC_ADDR_n instruction unit tests
func Test_ADC_ADDR_n(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> ADC_n: scenario 1 - add n + carry = 0, without carry out"), func(t *testing.T) {

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
			ADC_n,
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

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x0003, 0x0000, 0x00, 0x0c, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.a = 0x0a
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_ADC_n()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction ADC n: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> ADC_n: scenario 2 - add n + carry = 1, without carry out"), func(t *testing.T) {

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
			ADC_n,
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

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; Flags: 0x%02x; A: 0x%02x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x0003, 0x0000, 0x00, 0x0d, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.a = 0x0a
		cpu.flags = FLAG_C
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_ADC_n()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction ADC n: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> ADC_n: scenario 3 - add n + carry = 0, with carry out / zero"), func(t *testing.T) {

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
			ADC_n,
			0x4d,
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
			0x0003, 0x0000, 0xb0, 0x00, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.a = 0xb3
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_ADC_n()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction ADC n: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> ADC_n: scenario 4 - add n + carry = 0, with carry out / non zero"), func(t *testing.T) {

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
			ADC_n,
			0x4d,
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
			0x0003, 0x0000, 0x30, 0x10, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.a = 0xc3
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_ADC_n()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction ADC n: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// DEC_X instruction unit tests
func Test_DEC_X(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> DEC X: scenario 1 - decrement without carry out"), func(t *testing.T) {

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
			DEC_B,
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
			0x0002, 0x0000, 0x00, 0x00, 0x2b00, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.b = 0x2c
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_DEC_X(&cpu.b, "B")
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction DEC X: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> DEC X: scenario 2 - decrement with carry out"), func(t *testing.T) {

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
			DEC_B,
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
			0x0002, 0x0000, 0x00, 0x00, 0xff00, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.b = 0x00
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_DEC_X(&cpu.b, "B")
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction DEC B: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// INC X instruction unit tests
func Test_INC_X(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> INC X: scenario 1 - increment without carry out"), func(t *testing.T) {

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
			INC_B,
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
			0x0002, 0x0000, 0x00, 0x00, 0x2d00, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.b = 0x2c
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_INC_X(&cpu.b, "B")
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction INC X: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> INC X: scenario 2 - increment with carry out"), func(t *testing.T) {

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
			INC_B,
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
			0x0002, 0x0000, FLAG_Z, 0x00, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.b = 0xff
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_INC_X(&cpu.b, "B")
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction INC X: expected: %s\n\tresult: %s", want, got)
		}
	})
}
