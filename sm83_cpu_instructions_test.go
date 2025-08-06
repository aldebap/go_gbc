////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_instructions.go - Aug-3-2025 by aldebap
//
//	Test cases for Sharp SM83 CPU - generic instructions
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"
)

// LD_B_n instruction unit tests
func Test_LD_X_n(t *testing.T) {

	var err error

	t.Run(fmt.Sprintf(">>> LD X, n (0x%02x): scenario 1 - load B 8 bits register", LD_B_n), func(t *testing.T) {

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
			0x0003, 0x0000, 0x00, 0x00, 0x7e00, 0x0000, 0x0000)

		//	forced fetch instruction + two cicles to execute the instruction
		cpu.pc++
		cpu.cpu_state = EXECUTION_CYCLE_1

		for i := range 2 {
			err = cpu.executeInstruction_LD_X_n(&cpu.b, "B")
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
