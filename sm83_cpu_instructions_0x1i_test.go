////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_instructions_0x1i_test.go - Aug-7-2025 by aldebap
//
//	Test cases for Sharp SM83 CPU - instructions 0x10 - 0x1f
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"
)

// STOP instruction unit tests
func Test_STOP(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> STOP (0x%02x): scenario 1 - stop CPU", STOP), func(t *testing.T) {

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
			STOP,
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
			0x0001, 0x0000, 0x00, 0x00, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_STOP()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction STOP: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// RLA instruction unit tests
func Test_RLA(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> RLA (0x%02x): scenario 1 - no overflow", RLA), func(t *testing.T) {

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
			RLA,
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
			err = cpu.executeInstruction_RLA()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction RLA: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> RLA (0x%02x): scenario 2 - overflow", RLA), func(t *testing.T) {

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
			RLA,
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
		cpu.a = 0xc0
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_RLA()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction RLA: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// JR e instruction unit tests
func Test_JR_e(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> JR_e (0x%02x): scenario 1 - forward jump", JR_e), func(t *testing.T) {

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
			JR_e,
			0x0a,
			NOP,
			NOP,
			NOP,
			NOP,
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
			0x000d, 0x0000, 0x00, 0x00, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_JR_e()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction JR_e: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> JR_e (0x%02x): scenario 2 - backward jump", JR_e), func(t *testing.T) {

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
			JR_e,
			0xfc,
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
			0x0001, 0x0000, 0x00, 0x00, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.pc = 0x0003
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_JR_e()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction JR_e: expected: %s\n\tresult: %s", want, got)
		}
	})
}

// RRA instruction unit tests
func Test_RRA(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> RRA (0x%02x): scenario 1 - no underflow", RRA), func(t *testing.T) {

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
			RRA,
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
			0x0002, 0x0000, 0x00, 0x40, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.a = 0x80
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_RRA()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction RRA: expected: %s\n\tresult: %s", want, got)
		}
	})

	t.Run(fmt.Sprintf(">>> RRA (0x%02x): scenario 2 - underflow", RRA), func(t *testing.T) {

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
			RRA,
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
			0x0002, 0x0000, 0x80, 0x00, 0x0000, 0x0000, 0x0000)

		//	forced fetch instruction + one cicle to execute the instruction
		cpu.a = 0x01
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 1 {
			err = cpu.executeInstruction_RRA()
			if err != nil {
				t.Errorf("fail on cycle %d: %s", i, err.Error())
			}
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction RRA: expected: %s\n\tresult: %s", want, got)
		}
	})
}
