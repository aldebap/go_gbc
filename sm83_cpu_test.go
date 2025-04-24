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

	t.Run(">>> INC_A: scenario 1 - increment without carry out", func(t *testing.T) {

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
			0x3c, // INC A
			0x00, // NOP
		})
		if err != nil {
			t.Errorf("fail loading test program: %s", err.Error())
		}

		//	connect the ROM memory to the CPU
		err = cpu.ConnectMemory(rom, 0x0000)
		if err != nil {
			t.Errorf("fail connecting ROM to CPU: %s", err.Error())
		}

		want := fmt.Sprintf("PC: 0x%04x; SP: 0x%04x; AF: 0x%04x; BC: 0x%04x; DE: 0x%04x; HL: 0x%04x",
			0x0002, 0x0000, 0x0100, 0x0000, 0x0000, 0x0000)

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
