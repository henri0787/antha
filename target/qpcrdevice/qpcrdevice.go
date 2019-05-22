package qpcrdevice

import (
	"errors"
	"fmt"
	"os"

	"github.com/antha-lang/antha/driver"
	framework "github.com/antha-lang/antha/driver/antha_framework_v1"
	quantstudio "github.com/antha-lang/antha/driver/antha_quantstudio_v1"
	"github.com/antha-lang/antha/instructions"
	"github.com/antha-lang/antha/laboratory/effects"
	"github.com/antha-lang/antha/target"
	"github.com/antha-lang/antha/workflow"
	"github.com/golang/protobuf/proto"
)

// Ensure satisfies Device interface
var (
	_ effects.Device = (*QPCRDevice)(nil)
)

func NewQPCRInstances(tgt *target.Target, config workflow.QPCRConfig) error {
	for id := range config.Devices {
		if err := tgt.AddDevice(New(id)); err != nil {
			return err
		}
	}
	return nil
}

// QPCRDevice defines the state of a qpcr device device
type QPCRDevice struct {
	id workflow.DeviceInstanceID
}

// NewQPCRDevice returns a new QPCR Machine
func New(id workflow.DeviceInstanceID) *QPCRDevice {
	return &QPCRDevice{
		id: id,
	}
}

func (a *QPCRDevice) Id() workflow.DeviceInstanceID {
	return a.id
}

// CanCompile implements a Device
func (a *QPCRDevice) CanCompile(req instructions.Request) bool {
	can := instructions.Request{
		Selector: []instructions.NameValue{target.DriverSelectorV1QPCRDevice},
	}
	return can.Contains(req)
}

func (a *QPCRDevice) Connect(*workflow.Workflow) error {
	return nil
}

func (a *QPCRDevice) Close() {}

func (dev *QPCRDevice) transform(inst *instructions.QPCRInstruction) (*target.Run, error) {
	if inst.Definition == "" {
		return nil, errors.New("Blank experiment file for qPCR instruction.")
	}

	message := &quantstudio.TemplatedRequest{
		SessionInstrument: &quantstudio.SessionInstrument{
			Session:    &quantstudio.Session{},
			Instrument: &quantstudio.Instrument{},
		},
		TemplateFile: &quantstudio.ExperimentFile{
			Url: inst.Definition,
		},
		Barcode: &quantstudio.Barcode{
			Barcode: inst.Barcode,
		},
	}

	messageBytes, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}

	jobID := inst.TagAs
	if len(jobID) == 0 {
		jobID = os.Getenv("METADATA_JOB_ID")
	}

	instID := ""
	if len(inst.ComponentIn) > 0 {
		instID = inst.ComponentIn[0].ID
	}

	call := driver.Call{
		Method: "RunFramework",
		Args: &framework.CommandRequest{
			Description: &framework.CommandDescription{
				CommandName: inst.Command,
				DeviceId:    "QPCRDevice",
				CommandId:   instID,
			},
			Data: &framework.CommandData{
				Data: messageBytes,
			},
			Metadata: &framework.CommandMetadata{
				JobId: jobID,
			},
		},
		Reply: &framework.CommandResponse{},
	}

	return &target.Run{
		Device: dev,
		Label:  "Perform qPCR Analysis",
		Calls:  []driver.Call{call},
	}, nil
}

func (dev *QPCRDevice) makePrompt(inst *instructions.QPCRInstruction) instructions.Inst {
	bc := inst.Barcode
	if bc != "" {
		bc = " (" + bc + ")" // deliberate leading space
	}
	return &target.Prompt{
		Message: fmt.Sprintf("Ensure that the experiment file %s is configured, then put the plate%s into the qPCR device. Check that the AnthaHub software is running on the computer connected to the qPCR device. Then, click the button below to schedule the qPCR run.",
			inst.Definition, bc),
	}
}

// Compile implements a qPCR device.
func (dev *QPCRDevice) Compile(labEffects *effects.LaboratoryEffects, dir string, nodes []instructions.Node) (instructions.Insts, error) {
	if len(nodes) > 1 {
		return nil, fmt.Errorf("Currently only permit a single qPCR instruction per workflow. Received %d", len(nodes))
	}

	insts := make(instructions.Insts, 0, 2*len(nodes))
	for _, node := range nodes {
		if cmd, ok := node.(*instructions.Command); !ok {
			return nil, fmt.Errorf("cannot compile %T", node)
		} else if inst, ok := cmd.Inst.(*instructions.QPCRInstruction); !ok {
			return nil, fmt.Errorf("cannot compile %T", cmd.Inst)
		} else if call, err := dev.transform(inst); err != nil {
			return nil, err
		} else {
			insts = append(insts, dev.makePrompt(inst), call)
		}
	}

	insts.SequentialOrder()
	return insts, nil
}
