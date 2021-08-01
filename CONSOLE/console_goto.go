package CONSOLE

import (
	"github.com/cassianoperin/6502"
	"fmt"
)

// Console goto command
func Console_Command_Goto(text_slice []string) {

	var (
		current_PC      uint16
		step_count      int    = 0
		loop_count      uint16 = 0
		breakpoint_flag bool
	)

	// Test the command syntax
	if len(text_slice) == 1 || len(text_slice) > 2 {

		// Print goto usage
		fmt.Printf("Usage: goto <PC_ADDRESS>\n\n")

	} else {

		// Check if input is Decimar of Hexadecimal and convert to integer
		mem_arg, error_flag := Console_Hex_or_Dec(text_slice[1])

		if !error_flag {

			for loop_count < CORE.Loop_detection {

				// -------------------------- Start Checks --------------------------- //

				// Check Goto step limits
				if step_count > goto_limit {
					break // Exit for loop
				}

				// Check Breakpoints
				breakpoint_flag = Console_Check_breakpoints(breakpoint_flag)

				// Exit for loop if breakpoint has been found
				if breakpoint_flag {
					break
				}

				// -------------- Finish checks and return to execution -------------- //
				current_PC = CORE.PC

				select {
				case <-CORE.Second_timer: // Show the header and debug each second

					// Execute one instruction
					Console_Step(opcode_map, text_slice[0])

				default: // Just run the CPU

					// Execute one instruction without print
					Console_Step_without_debug(opcode_map, text_slice[0])

				}

				// Increase steps count
				step_count++

				// Check for goto_limit and print debug message prior to quit loop
				if step_count > goto_limit { // Print limit reached message
					fmt.Printf("Goto limit reached (%d)\n\n", goto_limit)
				}

				// Increase the loop counter
				if current_PC == CORE.PC {
					loop_count++
				}

				// Check for loops and print debug message prior to quit loop
				if loop_count >= CORE.Loop_detection {
					fmt.Printf("Loop detected on PC=%04X (%d repetitions)\n", CORE.PC, CORE.Loop_detection)
				}

				// Check if GOTO address was reached
				if CORE.PC == uint16(mem_arg) {
					fmt.Printf("Reached PC=0x%02X\n\n", mem_arg)
					break // Exit for loop
				}

			}

			// Print Header
			Console_PrintHeader()

		} else {
			fmt.Printf("Invalid value %s\n\n", text_slice[1])
		}
	}

}
