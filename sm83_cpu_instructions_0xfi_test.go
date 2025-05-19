////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_instructions_0xfi_test.go - Apr-24-2025 by aldebap
//
//	Test cases for Sharp SM83 CPU - instructions 0xf0 - 0xff
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"
)

// LD_A_ADDR_nn instruction unit tests
func Test_LD_A_ADDR_nn(t *testing.T) {

	var err error

	t.Run(">>> LD A, (nn): scenario 1 - load acumulator from memory", func(t *testing.T) {

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
			0x04,
			0x00,
			NOP,
			0x9a,
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
			0x0004, 0x0000, 0x00, 0x9a, 0x0000, 0x0000, 0x0000)

		//	eight cicles to execute the test program
		for range 5 {
			cpu.MachineCycle()
		}

		got := cpu.DumpRegisters()

		//	check the invocation result
		if want != got {
			t.Errorf("failed executing instruction LD A, (nn): expected: %s\n\tresult: %s", want, got)
		}
	})
}
