package target

import "github.com/antha-lang/antha/ast"

type Device interface {
	Can(req ast.Request) bool // Can device handle this request
	MoveCost(from Device) int // A non-negative cost to move to this device

	// Produce a single-entry, single-exit DAG of instructions where insts[0]
	// is the entry point and insts[len(insts)-1] is the exit point
	Compile(cmds []ast.Command) (insts []Inst, err error)
}
