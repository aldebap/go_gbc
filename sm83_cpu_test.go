////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_test.go - Apr-24-2025 by aldebap
//
//	Test cases for Sharp SM83 CPU
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"
)

// INC_A instruction unit tests
func Test_INC_A(t *testing.T) {

	var err error

	t.Run(">>> INC A: scenario 1 - increment without carry out", func(t *testing.T) {

		//	create a new SM83 CPU
		cpu := NewSM83_CPU()
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			INC_A,
			NOOP,
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
			0x0002, 0x0000, 0x00, 0x01, 0x0000, 0x0000, 0x0000)

		//	two cicles to execute the test program
		cpu.MachineCycle()
		cpu.MachineCycle()

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
		cpu := NewSM83_CPU()
		if cpu == nil {
			t.Errorf("fail creating new SM83 CPU")
		}

		//	create a new ROM memory and load it with the test program
		rom := &ROM_memory{}
		if rom == nil {
			t.Errorf("fail creating new ROM memory")
		}
		err = rom.Load([]uint8{
			INC_A,
			DEC_A,
			NOOP,
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
			0x0003, 0x0000, 0x01, 0x00, 0x0000, 0x0000, 0x0000)

		//	three cicles to execute the test program
		cpu.MachineCycle()
		cpu.MachineCycle()
		cpu.MachineCycle()

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
		cpu := NewSM83_CPU()
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
			NOOP,
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
		cpu.MachineCycle()
		cpu.MachineCycle()
		cpu.MachineCycle()

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD A, n: expected: %s\n\tresult: %s", want, got)
		}
	})
}
