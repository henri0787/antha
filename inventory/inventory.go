package inventory

import (
	"github.com/Synthace/antha/inventory/components"
	"github.com/Synthace/antha/inventory/plates"
	"github.com/Synthace/antha/inventory/tipboxes"
	"github.com/Synthace/antha/inventory/tipwastes"
	"github.com/Synthace/antha/laboratory/effects/id"
	"github.com/Synthace/antha/workflow"
)

type Inventory struct {
	Components *components.Inventory
	TipWastes  *tipwastes.Inventory
	TipBoxes   *tipboxes.Inventory
	Plates     *plates.Inventory
}

func NewInventory(idGen *id.IDGenerator) *Inventory {
	return &Inventory{
		Components: components.NewInventory(idGen),
		TipWastes:  tipwastes.NewInventory(idGen),
		TipBoxes:   tipboxes.NewInventory(idGen),
		Plates:     plates.NewInventory(idGen),
	}
}

func (inv *Inventory) LoadForWorkflow(wf *workflow.Workflow) {
	// In the cloud we should always be supplied with a real inventory
	// of PlateTypes. Thus this is a convenience for working locally
	// from the command line.
	if wf == nil || len(wf.Inventory.PlateTypes) == 0 {
		inv.Plates.LoadLibrary()
	} else {
		inv.Plates.SetPlateTypes(wf.Inventory.PlateTypes)
	}

	// Similarly, the long-term intention is that these should come in
	// with the workflow too...
	inv.TipBoxes.LoadLibrary()
}
