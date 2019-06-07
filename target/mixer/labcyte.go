package mixer

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Synthace/antha/antha/anthalib/wtype"
	"github.com/Synthace/antha/instructions"
	"github.com/Synthace/antha/inventory"
	"github.com/Synthace/antha/laboratory/effects"
	"github.com/Synthace/antha/logger"
	"github.com/Synthace/antha/target"
	"github.com/Synthace/antha/workflow"
)

var (
	_ effects.Device = (*LabcyteInstance)(nil)
)

type LabcyteInstance struct {
	MaxPlates            float64
	MaxWells             float64
	ResidualVolumeWeight float64

	global *GlobalMixerConfig
	*workflow.LabcyteInstanceConfig

	*BaseMixer
}

func NewLabcyteInstances(logger *logger.Logger, tgt *target.Target, inv *inventory.Inventory, global *GlobalMixerConfig, config workflow.LabcyteConfig) error {
	defaultsWF := config.Defaults
	if defaultsWF == nil {
		defaultsWF = &workflow.LabcyteInstanceConfig{}
	}

	var (
		defaultMaxPlates            = 4.5
		defaultMaxWells             = 278.0
		defaultResidualVolumeWeight = 1.0
	)

	defaults := &LabcyteInstance{
		MaxPlates:             floatValue(defaultsWF.MaxPlates, &defaultMaxPlates),
		MaxWells:              floatValue(defaultsWF.MaxWells, &defaultMaxWells),
		ResidualVolumeWeight:  floatValue(defaultsWF.ResidualVolumeWeight, &defaultResidualVolumeWeight),
		LabcyteInstanceConfig: defaultsWF,
	}
	if err := defaults.Validate(inv); err != nil {
		return err
	}

	for id, instWF := range config.Devices {
		instance := &LabcyteInstance{
			MaxPlates:             floatValue(instWF.MaxPlates, &defaults.MaxPlates),
			MaxWells:              floatValue(instWF.MaxWells, &defaults.MaxWells),
			ResidualVolumeWeight:  floatValue(instWF.MaxPlates, &defaults.ResidualVolumeWeight),
			global:                global,
			LabcyteInstanceConfig: instWF,
			BaseMixer:             NewBaseMixer(logger, id, instWF.ParsedConnection, target.LabcyteSubType),
		}
		if err := instance.Validate(inv); err != nil {
			return err
		} else if err := tgt.AddDevice(instance); err != nil {
			return err
		}
	}

	return nil
}

func (inst *LabcyteInstance) Validate(inv *inventory.Inventory) error {
	switch {
	case inst.MaxPlates <= 0:
		return errors.New("Validation error: MaxPlates must be > 0")
	case inst.MaxWells <= 0:
		return errors.New("Validation error: MaxWells must be > 0")
	case inst.ResidualVolumeWeight < 0:
		return errors.New("Validation error: ResidualVolumeWeight must be >= 0")
	}

	// TODO: add extra validation here!
	for _, ptns := range [][]wtype.PlateTypeName{inst.InputPlateTypes, inst.OutputPlateTypes} {
		for _, ptn := range ptns {
			if _, err := inv.Plates.NewPlateType(ptn); err != nil {
				return err
			}
		}
	}

	return nil
}

func (inst *LabcyteInstance) Connect(wf *workflow.Workflow) error {
	if inst.properties == nil {
		if data, err := json.Marshal(inst.Model); err != nil {
			return err
		} else if err := inst.connect(wf, data); err != nil {
			return err
		} else if err := inst.properties.ApplyUserPreferences(inst.LayoutPreferences); err != nil {
			inst.Close()
			return err
		}
	}
	return nil
}

func (inst *LabcyteInstance) Compile(labEffects *effects.LaboratoryEffects, dir string, nodes []instructions.Node) (instructions.Insts, error) {
	instrs, err := checkInstructions(nodes)
	if err != nil {
		return nil, err
	}

	mix, err := mixOpts{
		Device:     inst,
		LabEffects: labEffects,
		Base:       inst.BaseMixer,
		Global:     inst.global,
		Instrs:     instrs,
		InputWeights: map[string]float64{
			"MAX_N_PLATES":           inst.MaxPlates,
			"MAX_N_WELLS":            inst.MaxWells,
			"RESIDUAL_VOLUME_WEIGHT": inst.ResidualVolumeWeight,
		},
		InputPlateTypes:  inst.InputPlateTypes,
		OutputPlateTypes: inst.OutputPlateTypes,

		OutDir:      dir,
		ContentName: fmt.Sprintf("%v-%v.txt", inst.Id(), instrs[0].BlockID),
	}.mix()
	if err != nil {
		return nil, err
	}

	return []instructions.Inst{mix}, nil
}
