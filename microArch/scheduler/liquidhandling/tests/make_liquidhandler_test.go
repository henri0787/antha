package tests

// this version of the liquid handler factory is JUST for testing
// so has no public calls to return liquid handlers

import (
	"fmt"

	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
	"github.com/antha-lang/antha/laboratory"
	"github.com/antha-lang/antha/laboratory/effects/id"
	"github.com/antha-lang/antha/microArch/driver/liquidhandling"
	"github.com/antha-lang/antha/workflow"
)

func setUpTipsFor(lab *laboratory.Laboratory, lhp *liquidhandling.LHProperties) {
	err := lab.Inventory.TipBoxes.ForEach(func(tb wtype.LHTipbox) error {
		if tb.Mnfr == lhp.Mnfr || lhp.Mnfr == "MotherNature" {
			// Don't return filter tips: the iteration order through the
			// inventory is non deterministic, so if we return filter
			// tips then we risk the tests failing due to the planner
			// choosing to use the filter tips.
			// This conditional is copied from instruction-plugins
			if !tb.Tiptype.Filtered && tb.Tiptype.Type != "LVGilson200" {
				lhp.Tips = append(lhp.Tips, tb.Tips[0][0])
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

const (
	HVMinRate = 0.225
	HVMaxRate = 37.5
	LVMinRate = 0.0225
	LVMaxRate = 3.75
)

func getIndependentConfig(idGen *id.IDGenerator) *wtype.LHChannelParameter {
	minvol := wunit.NewVolume(0.5, "ul")
	maxvol := wunit.NewVolume(250, "ul")
	minspd := wunit.NewFlowRate(LVMinRate, "ml/min")
	maxspd := wunit.NewFlowRate(HVMaxRate, "ml/min")

	return wtype.NewLHChannelParameter(idGen, "config", "GilsonPipetmax", minvol, maxvol, minspd, maxspd, 8, true, wtype.LHVChannel, 0)
}

func getHVConfig(idGen *id.IDGenerator) *wtype.LHChannelParameter {
	minvol := wunit.NewVolume(10, "ul")
	maxvol := wunit.NewVolume(250, "ul")
	minspd := wunit.NewFlowRate(HVMinRate, "ml/min")
	maxspd := wunit.NewFlowRate(HVMaxRate, "ml/min")

	return wtype.NewLHChannelParameter(idGen, "HVconfig", "GilsonPipetmax", minvol, maxvol, minspd, maxspd, 8, false, wtype.LHVChannel, 0)
}

func getLVConfig(idGen *id.IDGenerator) *wtype.LHChannelParameter {
	newminvol := wunit.NewVolume(0.5, "ul")
	newmaxvol := wunit.NewVolume(20, "ul")
	newminspd := wunit.NewFlowRate(LVMinRate, "ml/min")
	newmaxspd := wunit.NewFlowRate(LVMaxRate, "ml/min")

	return wtype.NewLHChannelParameter(idGen, "LVconfig", "GilsonPipetmax", newminvol, newmaxvol, newminspd, newmaxspd, 8, false, wtype.LHVChannel, 1)
}

func makeLayout(lab *laboratory.Laboratory) *liquidhandling.LHProperties {
	// gilson pipetmax

	layout := make(map[string]*wtype.LHPosition)
	i := 0
	x0 := 3.886
	y0 := 3.513
	z0 := -82.035
	xi := 149.86
	yi := 95.25
	xp := x0 // nolint
	yp := y0
	zp := z0
	for y := 0; y < 3; y++ {
		xp = x0
		for x := 0; x < 3; x++ {
			pos := wtype.NewLHPosition(fmt.Sprintf("position_%d", i+1), wtype.Coordinates3D{X: xp, Y: yp, Z: zp}, wtype.SBSFootprint)
			layout[pos.Name] = pos
			i += 1
			xp += xi
		}
		yp += yi
	}
	lhp := liquidhandling.NewLHProperties(lab.IDGenerator, "Pipetmax", "Gilson", liquidhandling.LLLiquidHandler, liquidhandling.DisposableTips, layout)
	// get tips permissible from the factory
	setUpTipsFor(lab, lhp)

	lhp.Preferences = &workflow.LayoutOpt{
		Tipboxes:  workflow.Addresses{"position_9", "position_6", "position_3", "position_5", "position_2"},
		Inputs:    workflow.Addresses{"position_4", "position_5", "position_6", "position_9", "position_8", "position_3"},
		Outputs:   workflow.Addresses{"position_7", "position_8", "position_9", "position_6", "position_5", "position_3"},
		Washes:    workflow.Addresses{"position_8"},
		Tipwastes: workflow.Addresses{"position_1"},
		Wastes:    workflow.Addresses{"position_9"},
	}

	return lhp
}

func makeIndependentLH(lab *laboratory.Laboratory) *liquidhandling.LHProperties {
	lhp := makeLayout(lab)
	// set manufacturer so the right tipChooser is selected
	lhp.Mnfr = "Hamilton"

	// test independent liquidhandler has only one head to avoid multi-head instruction issues
	config := getIndependentConfig(lab.IDGenerator)
	adaptor := wtype.NewLHAdaptor(lab.IDGenerator, "DummyAdaptor", "Gilson", config)
	head := wtype.NewLHHead(lab.IDGenerator, "Head", "Gilson", config)
	head.Adaptor = adaptor

	ha := wtype.NewLHHeadAssembly(nil)
	ha.AddPosition(wtype.Coordinates3D{X: 0, Y: 0, Z: 0})
	if err := ha.LoadHead(head); err != nil {
		panic(err)
	}
	lhp.Heads = append(lhp.Heads, head)
	lhp.Adaptors = append(lhp.Adaptors, adaptor)
	lhp.HeadAssemblies = append(lhp.HeadAssemblies, ha)

	return lhp
}

func makeGilson(lab *laboratory.Laboratory) *liquidhandling.LHProperties {
	lhp := makeLayout(lab)

	hvconfig := getHVConfig(lab.IDGenerator)
	hvadaptor := wtype.NewLHAdaptor(lab.IDGenerator, "DummyAdaptor", "Gilson", hvconfig)
	hvhead := wtype.NewLHHead(lab.IDGenerator, "HVHead", "Gilson", hvconfig)
	hvhead.Adaptor = hvadaptor

	lvconfig := getLVConfig(lab.IDGenerator)
	lvadaptor := wtype.NewLHAdaptor(lab.IDGenerator, "DummyAdaptor", "Gilson", lvconfig)
	lvhead := wtype.NewLHHead(lab.IDGenerator, "LVHead", "Gilson", lvconfig)
	lvhead.Adaptor = lvadaptor

	ha := wtype.NewLHHeadAssembly(nil)
	ha.AddPosition(wtype.Coordinates3D{X: 0, Y: -18.08, Z: 0})
	ha.AddPosition(wtype.Coordinates3D{X: 0, Y: 0, Z: 0})
	if err := ha.LoadHead(hvhead); err != nil {
		panic(err)
	}
	if err := ha.LoadHead(lvhead); err != nil {
		panic(err)
	}
	lhp.Heads = append(lhp.Heads, hvhead, lvhead)
	lhp.Adaptors = append(lhp.Adaptors, hvadaptor, lvadaptor)
	lhp.HeadAssemblies = append(lhp.HeadAssemblies, ha)

	return lhp
}