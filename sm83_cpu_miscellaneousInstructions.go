////////////////////////////////////////////////////////////////////////////////
//	sm83_cpu_miscellaneousInstructions.go - Apr-23-2025 by aldebap
//
//	Emulator for Sharp SM83 CPU - miscellaneous instructions
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

/*
DAA
NOP  --> NOP
STOP --> STOP
*/

// execute instruction NOP
func (c *SM83_CPU) executeInstruction_NOP() error {

	if c.trace {
		fmt.Printf("[trace] NOP\n")
	}

	//	fecth next instruction in the same cycle
	return c.fetchInstruction()
}

// execute instruction STOP
func (c *SM83_CPU) executeInstruction_STOP() error {

	if c.trace {
		fmt.Printf("[trace] STOP\n")
	}

	//	TODO: add a flag to STOP/HALT CPU
	return nil
}
