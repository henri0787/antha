package simulaterequestpb

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/Synthace/antha-runner/protobuf"
	inventorypb "github.com/Synthace/microservice/cmd/inventory/protobuf"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wtype/liquidtype"
	"github.com/antha-lang/antha/laboratory/effects"
	"github.com/antha-lang/antha/logger"
	"github.com/antha-lang/antha/workflow"
	"github.com/antha-lang/antha/workflow/migrate"
	"github.com/golang/protobuf/proto"
)

type SimulateRequestProtobufProvider struct {
	pb               *protobuf.SimulateRequest
	fm               *effects.FileManager
	repoMap          workflow.ElementTypesByRepository
	gilsonDeviceName string
	logger           *logger.Logger
}

func NewProvider(
	inputReader io.Reader,
	fm *effects.FileManager,
	repoMap workflow.ElementTypesByRepository,
	gilsonDeviceName string,
	logger *logger.Logger,
) (*SimulateRequestProtobufProvider, error) {
	bytes, err := ioutil.ReadAll(inputReader)
	if err != nil {
		return nil, err
	}

	pb := &protobuf.SimulateRequest{}
	if err := proto.Unmarshal(bytes, pb); err != nil {
		return nil, err
	}

	return &SimulateRequestProtobufProvider{
		pb:               pb,
		fm:               fm,
		repoMap:          repoMap,
		gilsonDeviceName: gilsonDeviceName,
		logger:           logger,
	}, nil
}

func (p *SimulateRequestProtobufProvider) GetWorkflowID() (workflow.BasicId, error) {
	id, err := workflow.RandomBasicId("")
	if err != nil {
		return "", err
	}
	return id, nil
}

func (p *SimulateRequestProtobufProvider) GetMeta() (workflow.Meta, error) {
	// No-op for this provider type, it doesn't model metadata
	return workflow.Meta{}, nil
}

func (p *SimulateRequestProtobufProvider) GetRepositories() (workflow.Repositories, error) {
	// No-op for this provider type, it doesn't model repositories
	return workflow.Repositories{}, nil
}

func (p *SimulateRequestProtobufProvider) getElementTypes() (workflow.ElementTypes, error) {
	seen := make(map[string]struct{}, len(p.pb.Processes))
	types := make(workflow.ElementTypes, 0, len(p.pb.Processes))
	for _, v := range p.pb.Processes {
		if _, found := seen[v.Component]; found {
			continue
		}

		seen[v.Component] = struct{}{}
		et, err := migrate.UniqueElementType(p.repoMap, workflow.ElementTypeName(v.Component))
		if err != nil {
			return nil, err
		}
		types = append(types, et)
	}

	return types, nil
}

func (p *SimulateRequestProtobufProvider) getElementConnections() (workflow.ElementInstancesConnections, error) {
	connections := make(workflow.ElementInstancesConnections, 0, len(p.pb.Connections))
	for _, c := range p.pb.Connections {
		connections = append(connections, workflow.ElementConnection{
			Source: workflow.ElementSocket{
				ElementInstance: workflow.ElementInstanceName(c.Source.Process),
				ParameterName:   workflow.ElementParameterName(c.Source.Port),
			},
			Target: workflow.ElementSocket{
				ElementInstance: workflow.ElementInstanceName(c.Target.Process),
				ParameterName:   workflow.ElementParameterName(c.Target.Port),
			},
		})
	}
	return connections, nil
}

func (p *SimulateRequestProtobufProvider) GetElements() (workflow.Elements, error) {
	types, err := p.getElementTypes()
	if err != nil {
		return workflow.Elements{}, err
	}

	connections, err := p.getElementConnections()
	if err != nil {
		return workflow.Elements{}, err
	}

	return workflow.Elements{
		Types:                types,
		InstancesConnections: connections,
	}, nil
}

func (p *SimulateRequestProtobufProvider) GetInventory() (workflow.Inventory, error) {
	// TODO: implement
	return workflow.Inventory{}, nil
}

func translatePlates(plates []*inventorypb.Plate) []*wtype.Plate {
	result := make([]*wtype.Plate, len(plates))
	for i, plate := range plates {
		result[i] = &wtype.Plate{
			ID:         plate.Id,
			Type:       wtype.PlateTypeName(plate.Type),
			PlateName:  plate.Name,
			Loc:        plate.Location,
			WlsX:       int(plate.WlsX),
			WlsY:       int(plate.WlsY),
			Wellcoords: nil, // FIXME! Expecting []*wtype.LHWell, plate.WellCoords is []string
			// TODO: Do we need to do anything with plate.Contents and plate.Barcodes??
		}
	}
	return result
}

func (p *SimulateRequestProtobufProvider) getGlobalMixerConfig() (workflow.GlobalMixerConfig, error) {
	config := workflow.GlobalMixerConfig{}
	mc := p.pb.GetMixerConfig()
	if mc != nil {
		if mc.LiquidHandlingPolicyXlsxJmpFile != nil {
			policyMap, err := liquidtype.PolicyMakerFromBytes(mc.LiquidHandlingPolicyXlsxJmpFile, wtype.PolicyName(liquidtype.BASEPolicy))
			if err != nil {
				return workflow.GlobalMixerConfig{}, err
			}
			lhpr := wtype.NewLHPolicyRuleSet()
			lhpr, err = wtype.AddUniversalRules(lhpr, policyMap)
			if err != nil {
				return workflow.GlobalMixerConfig{}, err
			}
			config.CustomPolicyRuleSet = lhpr
		}
		config.IgnorePhysicalSimulation = mc.IgnorePhysicalSimulation
		config.InputPlates = translatePlates(mc.GetInputPlateVals().GetPlates())
		config.UseDriverTipTracking = mc.UseDriverTipTracking
	}
	return config, nil
}

func (p *SimulateRequestProtobufProvider) getGilsonPipetMaxInstanceConfig() (*workflow.GilsonPipetMaxInstanceConfig, error) {
	config := workflow.GilsonPipetMaxInstanceConfig{}
	mc := p.pb.GetMixerConfig()
	if mc != nil {
		config.InputPlateTypes = migrate.UpdatePlateTypes(mc.InputPlateTypes)
		config.MaxPlates = &mc.MaxPlates
		config.MaxWells = &mc.MaxWells
		config.OutputPlateTypes = migrate.UpdatePlateTypes(mc.OutputPlateTypes)
		config.ResidualVolumeWeight = &mc.ResidualVolumeWeight
		config.TipTypes = mc.TipTypes
		config.LayoutPreferences = &workflow.LayoutOpt{
			Inputs:    mc.DriverSpecificInputPreferences,
			Outputs:   mc.DriverSpecificOutputPreferences,
			Tipboxes:  mc.DriverSpecificTipPreferences,
			Tipwastes: mc.DriverSpecificTipWastePreferences,
			Washes:    mc.DriverSpecificWashPreferences,
		}
	}
	return &config, nil
}

func (p *SimulateRequestProtobufProvider) getGilsonPipetMaxConfig() (workflow.GilsonPipetMaxConfig, error) {
	if p.gilsonDeviceName == "" {
		return workflow.GilsonPipetMaxConfig{}, nil
	}

	devices := map[workflow.DeviceInstanceID]*workflow.GilsonPipetMaxInstanceConfig{}
	devID := workflow.DeviceInstanceID(p.gilsonDeviceName)

	if _, found := devices[devID]; found {
		p.logger.Log("warning", fmt.Sprintf("Gilson device %s already exists, and will have configuration replaced with migrated configuration.", p.gilsonDeviceName))
	}

	devConfig, err := p.getGilsonPipetMaxInstanceConfig()
	if err != nil {
		return workflow.GilsonPipetMaxConfig{}, err
	}

	devices[devID] = devConfig

	return workflow.GilsonPipetMaxConfig{
		Devices: devices,
	}, nil
}

func (p *SimulateRequestProtobufProvider) GetConfig() (workflow.Config, error) {
	gmc, err := p.getGlobalMixerConfig()
	if err != nil {
		return workflow.Config{}, err
	}

	gpmc, err := p.getGilsonPipetMaxConfig()
	if err != nil {
		return workflow.Config{}, err
	}

	return workflow.Config{
		GlobalMixer:    gmc,
		GilsonPipetMax: gpmc,
	}, nil
}

func (p *SimulateRequestProtobufProvider) GetTesting() (workflow.Testing, error) {
	// TODO: implement
	return workflow.Testing{}, nil
}
