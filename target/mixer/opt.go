package mixer

var (
	defaultMaxPlates            = 4.5
	defaultMaxWells             = 278.0
	defaultResidualVolumeWeight = 1.0
	DefaultOpt                  = Opt{
		MaxPlates:            &defaultMaxPlates,
		MaxWells:             &defaultMaxWells,
		ResidualVolumeWeight: &defaultResidualVolumeWeight,
		InputPlateType:       []string{"pcrplate_skirted"},
		OutputPlateType:      []string{"pcrplate_skirted"},
	}
)

type Opt struct {
	MaxPlates            *float64
	MaxWells             *float64
	ResidualVolumeWeight *float64
	InputPlateType       []string
	OutputPlateType      []string
	PlanningVersion      *int
}

// Merge two configs together and return the result. Values in the argument
// override those in the receiver.
func (a Opt) Merge(x *Opt) Opt {
	if x == nil {
		return a
	}
	if x.MaxPlates != nil {
		a.MaxPlates = x.MaxPlates
	}
	if x.MaxWells != nil {
		a.MaxWells = x.MaxWells
	}
	if x.ResidualVolumeWeight != nil {
		a.ResidualVolumeWeight = x.ResidualVolumeWeight
	}
	if len(x.InputPlateType) != 0 {
		a.InputPlateType = x.InputPlateType
	}
	if len(x.OutputPlateType) != 0 {
		a.OutputPlateType = x.OutputPlateType
	}

	return a
}
