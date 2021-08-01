package CORE

import "fmt"

// SEC  Set Carry Flag
//
//      1 -> C                           N Z C I D V
//                                       - - 1 - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       SEC           38    1     2

func opc_SEC(bytes uint16, opc_cycles byte) {

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		P[0] = 1

		// Print Opcode Debug Message
		opc_SEC_DebugMsg(bytes)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_SEC_DebugMsg(bytes uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tSEC  Set Carry Flag.\tP[0]=1\n", opc_string)
		fmt.Println(dbg_show_message)
	}
}
