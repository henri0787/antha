// /anthalib/simulator/liquidhandling/simulator.go: Part of the Antha language
// Copyright (C) 2015 The Antha authors. All rights reserved.
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
//
// For more information relating to the software or licensing issues please
// contact license@antha-lang.org or write to the Antha team c/o
// Synthace Ltd. The London Bioscience Innovation Centre
// 2 Royal College St, London NW1 0NH UK

package liquidhandling

import (
    "io/ioutil"
    "strings"
    "fmt"
	"github.com/antha-lang/antha/microArch/driver"
	"github.com/antha-lang/antha/microArch/driver/liquidhandling"
	"github.com/antha-lang/antha/microArch/simulator"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
)

func summariseWell2Channel(well []string, channels []int) string {
    ret := make([]string, len(well))
    for i := range well {
        ret[i] = fmt.Sprintf("%s->channel%v", well[i], channels[i])
    }
    return strings.Join(ret, ", ")
}

// Simulate a liquid handler Driver
type VirtualLiquidHandler struct {
    properties *liquidhandling.LHProperties 
    initialized bool
    finalized   bool
    //Need to store:
    // LHProperties
    // plate(s) at each layout position
    //   contents of each well of each plate / tip box
    //     liquid type / tip type
    // contents of each loaded tip
    //   liquid type
    // tips on each adaptor
    //   LHAdapter know how many tips it has, and what type they are (assumed equal)
    //   but it can't tell which position they're on
    //   tip properties
    // adaptors on each head
    //   Adapeter properties
    // head location
    //   Head properties
    errors      []*simulator.SimulationError
    max_error   simulator.ErrorSeverity
    worst_error *simulator.SimulationError
    log []string
}

//Create a new VirtualLiquidHandler which mimics an LHDriver
func NewVirtualLiquidHandler(props *liquidhandling.LHProperties) *VirtualLiquidHandler {
    var vlh VirtualLiquidHandler
    vlh.initialized = false
    vlh.finalized   = false

    vlh.properties = props.Dup()
    vlh.log = make([]string, 0)
    vlh.errors = make([]*simulator.SimulationError, 0)
    vlh.worst_error = nil

    vlh.validateProperties()

    return &vlh
}

//Write to the log
func (self *VirtualLiquidHandler) LogLine(line string) {
    self.log = append(self.log, line)
}

//save the log
func (self *VirtualLiquidHandler) SaveLog(filename string) {
    ioutil.WriteFile(filename, []byte(strings.Join(self.log, "\n")), 0644)
}

//Get the list of errors, and the maximum error severity
func (self *VirtualLiquidHandler) GetErrors() ([]*simulator.SimulationError, simulator.ErrorSeverity) {
    return self.errors, self.GetErrorSeverity()
}

//AddError Add an error
func (self *VirtualLiquidHandler) AddError(err *simulator.SimulationError) {
    if err.Severity() > self.GetErrorSeverity() {
        self.worst_error = err
    }
    self.errors = append(self.errors, err)
}

//GetErrorSeverity get the severity of the worst error encountered so far
func (self *VirtualLiquidHandler) GetErrorSeverity() simulator.ErrorSeverity {
    if self.worst_error != nil {
        return self.worst_error.Severity()
    }
    return simulator.SeverityNone
}

//IsError return true iff an error has been raised with SeverityError
func (self *VirtualLiquidHandler) IsError() bool {
    return self.GetErrorSeverity() == simulator.SeverityError
}

//GetWorstError get the worst error encountered so far
func (self *VirtualLiquidHandler) GetWorstError() *simulator.SimulationError {
    return self.worst_error
}

func (self *VirtualLiquidHandler) validateProperties() {
    
    //check a property
    check_prop := func(l []string, name string) {
        //is empty
        if len(l) == 0 {
            self.AddError(simulator.NewSimulationError(simulator.SeverityWarning, 
                                                       fmt.Sprintf("No %s specified", name),
                                                       nil))
        }
        //all locations defined
        for _,loc := range l {
            if !self.locationIsKnown(loc) {
                self.AddError(simulator.NewSimulationError(simulator.SeverityWarning, 
                        fmt.Sprintf("Undefined location \"%s\" referenced in %s", loc, name),
                                                           nil))
            }
        }
    }

    check_prop(self.properties.Tip_preferences, "tip preferences")
    check_prop(self.properties.Input_preferences, "input preferences")
    check_prop(self.properties.Output_preferences, "output preferences")
    check_prop(self.properties.Tipwaste_preferences, "tipwaste preferences")
    check_prop(self.properties.Wash_preferences, "wash preferences")
    check_prop(self.properties.Waste_preferences, "waste preferences")
}

func (self *VirtualLiquidHandler) locationIsKnown(location string) bool {
    for loc := range self.properties.Layout {
        if loc == location {
            return true
        }
    }
    return false
}

func (self *VirtualLiquidHandler) checkReady(fnName string) {
    if !self.initialized {
        self.AddError(simulator.NewSimulationError(simulator.SeverityWarning,
            fmt.Sprintf("Instruction \"%s\" before Initialize", fnName),
            nil))
    }
    if self.finalized {
        self.AddError(simulator.NewSimulationError(simulator.SeverityWarning,
            fmt.Sprintf("Instruction \"%s\" after Finalize", fnName),
            nil))
    }
}

// ------------------------------------------------------------------------ ExtendedLHDriver

//Move command - used
func (self *VirtualLiquidHandler) Move(deckposition []string, wellcoords []string, reference []int, 
                                       offsetX, offsetY, offsetZ []float64, plate_type []string, 
                                       head int) driver.CommandStatus {
    self.checkReady("Move")
    self.LogLine(fmt.Sprintf(`Move(
    deckposition = %v,
    wellcoords = %v,
    reference = %v,
    offsetX,Y,Z = (%v, %v, %v),
    plate_type = %v,
    head = %v)`, deckposition, wellcoords, reference, offsetX, offsetY, offsetZ, plate_type, head))
    //Asserts:
    //deckposition exists - why is it a list?
    //wellcoords exists within the plate at deckposition
    //reference is in allowable range
    //offsetX, offsetY, offsetZ are within the well (but I guess they needn't be...)
    //plate_type matches the type of plate at deckposition
    //head is valid
    return driver.CommandStatus{true, driver.OK, "MOVE ACK"}
}

//Move raw - not yet implemented in compositerobotinstruction
func (self *VirtualLiquidHandler) MoveRaw(head int, x, y, z float64) driver.CommandStatus {
    self.checkReady("MoveRaw")
    self.LogLine(fmt.Sprintf(`MoveRaw(
    head = %v,
    offsetX,Y,Z = (%v, %v, %v))`, head, x,y,z))
    //Asserts:
    //head exists
    //x,y,x are within the machine
    return driver.CommandStatus{true, driver.OK, "MOVERAW ACK"}
}

//Aspirate - used
func (self *VirtualLiquidHandler) Aspirate(volume []float64, overstroke []bool, head int, multi int, 
                                           platetype []string, what []string, llf []bool) driver.CommandStatus {
    self.checkReady("Aspirate")
    self.LogLine(fmt.Sprintf(`Aspirate(
    volume = %v,
    overstroke = %v,
    head = %v,
    multi = %v,
    platetype = %v,
    what = %v,
    llf = %v)`, volume, overstroke, head, multi, platetype, what, llf))
    //volumes are equal if adapter isn't independent
    //tips are loaded in each adapter location the aspirates
    //volume is smaller that the tips' maximum capacity
    //the tip hasn't been used for a different liquid
    //head exists
    //multi matches number of tips loaded
    //platetype matches the plate at the location we moved to
    //what matches the expected liquid class
    //llf is the right size, cannot vary unless independent
    return driver.CommandStatus{true, driver.OK, "ASPIRATE ACK"}
}

//Dispense - used
func (self *VirtualLiquidHandler) Dispense(volume []float64, blowout []bool, head int, multi int, 
                                           platetype []string, what []string, llf []bool) driver.CommandStatus {
    self.checkReady("Dispense")
    self.LogLine(fmt.Sprintf(`Dispense(
    volume = %v,
    blowout = %v,
    head = %v,
    multi = %v,
    platetype = %v,
    what = %v,
    llf = %v)`, volume, blowout, head, multi, platetype, what, llf))
    //Volumes are equal if adapter isn't indepentent
    //Volumes are at most equal to the volume in the tip
    //blowout is the right length
    //head exists
    //multi is valid
    //platetype matches the type of plate that we're next to
    //what matches the liquid class that was aspirated
    //llf is the right size and follows independence constraint
    return driver.CommandStatus{true, driver.OK, "DISPENSE ACK"}
}

//LoadTips - used
func (self *VirtualLiquidHandler) LoadTips(channels []int, head, multi int, 
                                           platetype, position, well []string) driver.CommandStatus {
    self.checkReady("LoadTips")
    self.LogLine(fmt.Sprintf(`LoadTips(
    channels = %v,
    head = %v,
    multi = %v,
    platetype = %v,
    position = %v,
    well = %v)`, channels, head, multi, platetype, position, well))
    
    
    //test length of channel, platetype, position, and well
    wrong_length := make([]string, 0)
    if len(channels) != multi {
        wrong_length = append(wrong_length, "channels")
    }
    if len(platetype) != multi {
        wrong_length = append(wrong_length, "platetype")
    }
    if len(position) != multi {
        wrong_length = append(wrong_length, "position")
    }
    if len(well) != multi {
        wrong_length = append(wrong_length, "well")
    }
    if len(wrong_length) > 0 {
        self.AddError(simulator.NewSimulationError(simulator.SeverityError,
            fmt.Sprintf("LoadTips: %s should be of length multi=%v",
                strings.Join(wrong_length, ", "),
                multi), nil))
        return driver.CommandStatus{true, driver.OK, "LOADTIPS ACK"}
    }

    //check head exists
    if head < 0 || head >= len(self.properties.Heads) {
        self.AddError(simulator.NewSimulationError(simulator.SeverityError,
            fmt.Sprintf("LoadTips: request for invalid Head %v", head), 
            nil))
        return driver.CommandStatus{true, driver.OK, "LOADTIPS ACK"}
    }

    //check that the adaptor is currently empty
    adaptor := self.properties.Heads[head].Adaptor
    if adaptor.Ntipsloaded != 0 {
        //because not having a ternary operator is so cool.
        stip := "tip"
        if adaptor.Ntipsloaded > 1 {
            stip = "tips"
        }
        self.AddError(simulator.NewSimulationError(simulator.SeverityError,
            fmt.Sprintf("LoadTips: Cannot load tips while adaptor already contains %v %s", 
                adaptor.Ntipsloaded, stip), 
            nil))
        return driver.CommandStatus{true, driver.OK, "LOADTIPS ACK"}
    }

    //check that channel values are valid and there are no duplicates
    encountered := map[int]bool{}
    for _, channel := range channels {
        if channel < 0 || channel >= adaptor.Params.Multi {
            self.AddError(simulator.NewSimulationError(simulator.SeverityError, 
                fmt.Sprintf("Cannot load tip to channel %v of %v-channel adaptor", 
                            channel, adaptor.Params.Multi), nil))
            return driver.CommandStatus{true, driver.OK, "LOADTIPS ACK"}
        }
        if encountered[channel] {
            self.AddError(simulator.NewSimulationError(simulator.SeverityError,
                fmt.Sprintf("LoadTips: Channel%v appears more than once", channel),
                nil))
            return driver.CommandStatus{true, driver.OK, "LOADTIPS ACK"}
        } else {
            encountered[channel] = true
        }
    }


    //check that position are all equal and point to a tipbox
    tipbox_pos := position[0]
    positions_equal := true
    for _,pos := range position {
        positions_equal = positions_equal && (pos == tipbox_pos)
    }
    if !positions_equal {
        self.AddError(simulator.NewSimulationError(simulator.SeverityError, 
            "LoadTips: Cannot load tips from multiple locations", 
            nil))
    }
    tipbox, ok := self.properties.PlateLookup[self.properties.PosLookup[tipbox_pos]].(*wtype.LHTipbox)
    if !ok {
        self.AddError(simulator.NewSimulationError(simulator.SeverityError,
            fmt.Sprintf("LoadTips: location \"%s\" doesn't contain a tipbox", tipbox_pos),
            nil))
        return driver.CommandStatus{true, driver.OK, "LOADTIPS ACK"}
    }
    

    //check that platetype are all the same, and match the type of the tipbox
    ptype := platetype[0]
    ptype_equal := true
    for _,p := range platetype {
        ptype_equal = ptype_equal && (ptype == p)
    }
    if !ptype_equal {
        self.AddError(simulator.NewSimulationError(simulator.SeverityWarning,
            "LoadTips: platetype should all be the same",
            nil))
    }
    if ptype != tipbox.Type {
        self.AddError(simulator.NewSimulationError(simulator.SeverityWarning,
            fmt.Sprintf("LoadTips: Requested plate type \"%s\" but plate at %s is of type \"%s\"", ptype, tipbox_pos, tipbox.Type),
            nil))
    }

    //check that well is valid
    wellcoords := make([]wtype.WellCoords, multi)
    for i := range well {
        wellcoords[i] = wtype.MakeWellCoords(well[i])
        
        //check range
        if wellcoords[i].IsZero() {
            self.AddError(simulator.NewSimulationError(simulator.SeverityError,
                fmt.Sprintf("LoadTips: Couldn't parse well \"%s\"", well[i]),
                nil))
            continue
        }
        if !tipbox.ContainsCoords(&wellcoords[i]) {
            self.AddError(simulator.NewSimulationError(simulator.SeverityError,
                fmt.Sprintf("LoadTips: Request for well %s, but tipbox size is [%vx%v]", well[i], tipbox.Ncols, tipbox.Nrows),
                nil))
            continue
        }
    }

    //Check that the head origin is sane
    head_origin_ := make([]wtype.WellCoords, multi)
    for i,wc := range wellcoords {
        head_origin_[i] = wc
        if adaptor.Params.Orientation == wtype.LHVChannel {
            head_origin_[i].Y -= channels[i]
        } else if adaptor.Params.Orientation == wtype.LHHChannel {
            head_origin_[i].X -= channels[i]
        }
    }
    head_origin := head_origin_[0]
    origins_equal := true
    for _,o := range head_origin_ {
        origins_equal = origins_equal && head_origin.Equals(o)
    }
    if !origins_equal {
        self.AddError(simulator.NewSimulationError(simulator.SeverityError,
            fmt.Sprintf("LoadTips: Cannot load %s, tip spacing doesn't match channel spacing",
                summariseWell2Channel(well, channels)),
            nil))
        return driver.CommandStatus{true, driver.OK, "LOADTIPS ACK"}
    }

    //check for tip collisions on other channels
    if !adaptor.Params.Independent {
        collisions := []string{}
        for i := 0; i < adaptor.Params.Multi; i++ {
            pos := head_origin
            if adaptor.Params.Orientation == wtype.LHVChannel {
                pos.Y += i
            } else if adaptor.Params.Orientation == wtype.LHHChannel {
                pos.X += i
            }
            
            //if there's a tip at this location, and we don't intend to pick it up, and the adaptor is not independent
            if tipbox.HasTipAt(&pos) && !encountered[i] {
               collisions = append(collisions, pos.FormatA1()) 
            }
        }
        if len(collisions) > 0 {
            stip := "tip"
            if len(collisions) > 1 {
                stip = "tips"
            }
            self.AddError(simulator.NewSimulationError(simulator.SeverityError,
                fmt.Sprintf("LoadTips: Cannot load %s due to %s at %s (Head%v is not independent)",
                    summariseWell2Channel(well, channels), stip, strings.Join(collisions, ","), head),
                nil))
        }
    }


    //if we found an error, bail on the operation
    if self.IsError() {
        return driver.CommandStatus{true, driver.OK, "LOADTIPS ACK"}
    }
    
    //move the tips from the plate and add them to the adaptor
    var tip *wtype.LHTip
    for i := range well {
        
        tip = tipbox.Tips[wellcoords[i].X][wellcoords[i].Y]
        if tip == nil {
            self.AddError(simulator.NewSimulationError(simulator.SeverityError,
                fmt.Sprintf("LoadTips: Cannot load %s->channel%v as %s is empty", well[i], channels[i], well[i]),
                nil))
            continue
        }
        tipbox.Tips[wellcoords[i].X][wellcoords[i].Y] = nil
        //TODO: add the tip to the particular location in the adaptor
    }
    adaptor.LoadTips(multi, tip)

    return driver.CommandStatus{true, driver.OK, "LOADTIPS ACK"}
}

//UnloadTips - used
func (self *VirtualLiquidHandler) UnloadTips(channels []int, head, multi int, 
                                             platetype, position, well []string) driver.CommandStatus {
    self.checkReady("UnloadTips")
    self.LogLine(fmt.Sprintf(`UnloadTips(
    channels = %v,
    head = %v,
    multi = %v,
    platetype = %v,
    position = %v,
    well = %v)`, channels, head, multi, platetype, position, well))
    //Tips are loaded in channels
    //independence constraints are met
    //head exists
    //multi is correct
    //platetype matches the plate we're over
    //platetype is tip-waste
    //position and well are correct
    return driver.CommandStatus{true, driver.OK, "UNLOADTIPS ACK"}
}

//SetPipetteSpeed - used
func (self *VirtualLiquidHandler) SetPipetteSpeed(head, channel int, rate float64) driver.CommandStatus {
    self.checkReady("SetPipetteSpeed")
    self.LogLine(fmt.Sprintf(`SetPipetteSpeed(
    head = %v,
    channel = %v,
    rate = %v)`, head, channel, rate))
    //head exists
    //channel exists
    //speed is within allowable range
    return driver.CommandStatus{true, driver.OK, "SETPIPETTESPEED ACK"}
}

//SetDriveSpeed - used
func (self *VirtualLiquidHandler) SetDriveSpeed(drive string, rate float64) driver.CommandStatus {
    self.checkReady("SetDriveSpeed")
    self.LogLine(fmt.Sprintf(`SetDriveSpeed(
    drive = %v,
    rate = %v)`, drive, rate))
    //drive string?
    //rate is within allowable range (what is this?)
    return driver.CommandStatus{true, driver.OK, "SETDRIVESPEED ACK"}
}

//Stop - unused
func (self *VirtualLiquidHandler) Stop() driver.CommandStatus {
    panic("unimplemented")
}

//Go - unused
func (self *VirtualLiquidHandler) Go() driver.CommandStatus {
    self.checkReady("Go")
    panic("unimplemented")
}

//Initialize - used
func (self *VirtualLiquidHandler) Initialize() driver.CommandStatus {
    self.LogLine("Initialize()")
    if self.initialized {
        self.AddError(simulator.NewSimulationError(simulator.SeverityWarning,
            fmt.Sprintf("Second call to Initialize"),
            nil))
    } else {
        self.initialized = true
    }
    return driver.CommandStatus{true, driver.OK, "INITIALIZE ACK"}
}

//Finalize - used
func (self *VirtualLiquidHandler) Finalize() driver.CommandStatus {
    self.LogLine("Finalize()")
    //check that this is called last, no more calls
    if self.finalized {
        self.AddError(simulator.NewSimulationError(simulator.SeverityWarning,
            fmt.Sprintf("Second call to Finalize"),
            nil))
    } else {
        self.finalized = true
    }
    return driver.CommandStatus{true, driver.OK, "FINALIZE ACK"}
}

//Wait - used
func (self *VirtualLiquidHandler) Wait(time float64) driver.CommandStatus {
    self.checkReady("Wait")
    self.LogLine(fmt.Sprintf(`Wait(time = %v)`, time))
    //time is positive
    //maybe a warning if it's super-long
    return driver.CommandStatus{true, driver.OK, "WAIT ACK"}
}

//Mix - used
func (self *VirtualLiquidHandler) Mix(head int, volume []float64, platetype []string, cycles []int, 
                                      multi int, what []string, blowout []bool) driver.CommandStatus {
    self.checkReady("Mix")
    self.LogLine(fmt.Sprintf(`Mix(
    head = %v,
    volume = %v,
    platetype = %v,
    cycles = %v,
    multi = %v,
    what = %v,
    blowout = %v)`, head, volume, platetype, cycles, multi, what, blowout))
    //head exists
    //volume is lte volume in wells
    //platetype matches the plate we're over
    //muli is correct
    //what matches expected liquidclass
    //volume, platetype, what, blowout match independence constraint
    return driver.CommandStatus{true, driver.OK, "MIX ACK"}
}

//ResetPistons - used
func (self *VirtualLiquidHandler) ResetPistons(head, channel int) driver.CommandStatus {
    self.checkReady("ResetPistons")
    self.LogLine("ResetPistons()")
    //head exists
    //channel exists
    //what does this do again? probably need to make sure it gets called appropriately
    return driver.CommandStatus{true, driver.OK, "RESETPISTONS ACK"}
}

//AddPlateTo - used
func (self *VirtualLiquidHandler) AddPlateTo(position string, plate interface{}, name string) driver.CommandStatus {
    self.checkReady("AddPlateTo")
    self.LogLine(fmt.Sprintf(`AddPlateTo(
    position = %v,
    plate = %v,
    name = %v)`, position, plate, name))
    //check that the requested position exists
    if !self.locationIsKnown(position) {
        self.AddError(simulator.NewSimulationError(simulator.SeverityWarning,
            fmt.Sprintf("Adding plate \"%s\" to unknown location \"%s\"", name, position),
            nil))
    }
    //check that the requested position is empty
    if self.properties.PosLookup[position] != "" {
        self.AddError(simulator.NewSimulationError(simulator.SeverityError,
            fmt.Sprintf("Adding plate \"%s\" to \"%s\" which is already occupied by plate \"%s\"", 
                name, position, self.properties.PlateLookup[self.properties.PosLookup[position]].(wtype.Named).GetName()),
            nil))
    } else {
        //position can accept a plate of this type
        switch plate := plate.(type) {
        case *wtype.LHPlate:
            plate.PlateName = name
            self.properties.AddPlate(position, plate)
        case *wtype.LHTipbox:
            plate.Boxname = name
            self.properties.AddTipBoxTo(position, plate)
        case *wtype.LHTipwaste:
            plate.Type = name
            self.properties.AddTipWasteTo(position, plate)

        default:
            self.AddError(simulator.NewSimulationError(simulator.SeverityWarning,
                fmt.Sprintf("unknown plate of type %T while adding \"%s\" to location \"%s\"", plate, name, position),
                nil))
        }
    }
    return driver.CommandStatus{true, driver.OK, "ADDPLATETO ACK"}
}

//RemoveAllPlates - used
func (self *VirtualLiquidHandler) RemoveAllPlates() driver.CommandStatus {
    self.checkReady("RemoveAllPlates")
    self.LogLine("RemoveAllPlates()")
    //remove plates, no checks required.
    return driver.CommandStatus{true, driver.OK, "REMOVEALLPLATES ACK"}
}

//RemovePlateAt - unused
func (self *VirtualLiquidHandler) RemovePlateAt(position string) driver.CommandStatus {
    self.LogLine(fmt.Sprintf("RemovePlateAt(position = %v)", position))
    //plate exists at position
    return driver.CommandStatus{true, driver.OK, "REMOVEPLATEAT ACK"}
}

//SetPositionState - unused
func (self *VirtualLiquidHandler) SetPositionState(position string, state driver.PositionState) driver.CommandStatus {
    panic("unimplemented")
}

//GetCapabilites - used
func (self *VirtualLiquidHandler) GetCapabilities() (liquidhandling.LHProperties, driver.CommandStatus) {
    self.checkReady("GetCapabilities")
    self.LogLine("GetCapabilities()")
    //no checks required
    return *self.properties, driver.CommandStatus{true, driver.OK, ""} 
}

//GetCurrentPosition - unused
func (self *VirtualLiquidHandler) GetCurrentPosition(head int) (string, driver.CommandStatus) {
    panic("unimplemented")
}

//GetPositionState - unused
func (self *VirtualLiquidHandler) GetPositionState(position string) (string, driver.CommandStatus) {
    panic("unimplemented")
}

//GetHeadState - unused
func (self *VirtualLiquidHandler) GetHeadState(head int) (string, driver.CommandStatus) {
    panic("unimplemented")
}

//GetStatus - unused
func (self *VirtualLiquidHandler) GetStatus() (driver.Status, driver.CommandStatus) {
    panic("unimplemented")
}

//UpdateMetaData - used
func (self *VirtualLiquidHandler) UpdateMetaData(props *liquidhandling.LHProperties) driver.CommandStatus {
    self.checkReady("UpdateMetaData")
    self.LogLine("UpdateMetaData(props *LHProperties)")
    //check that the props and self.props are the same...
    return driver.CommandStatus{true, driver.OK, "UPDATEMETADATA ACK"}
}

//UnloadHead - unused
func (self *VirtualLiquidHandler) UnloadHead(param int) driver.CommandStatus {
    panic("unimplemented")
}

//LoadHead - unused
func (self *VirtualLiquidHandler) LoadHead(param int) driver.CommandStatus {
    panic("unimplemented")
}

//Lights On - not implemented in compositerobotinstruction
func (self *VirtualLiquidHandler) LightsOn() driver.CommandStatus {
    panic("unimplemented")
}

//Lights Off - notimplemented in compositerobotinstruction
func (self *VirtualLiquidHandler) LightsOff() driver.CommandStatus {
    panic("unimplemented")
}

//LoadAdaptor - notimplemented in CRI
func (self *VirtualLiquidHandler) LoadAdaptor(param int) driver.CommandStatus {
    panic("unimplemented")
}

//UnloadAdaptor - notimplemented in CRI
func (self *VirtualLiquidHandler) UnloadAdaptor(param int) driver.CommandStatus {
    panic("unimplemented")
}

//Open - notimplemented in CRI
func (self *VirtualLiquidHandler) Open() driver.CommandStatus {
    panic("unimplemented")
}

//Close - notimplement in CRI
func (self *VirtualLiquidHandler) Close() driver.CommandStatus {
    panic("unimplemented")
}

//Message - unused
func (self *VirtualLiquidHandler) Message(level int, title, text string, showcancel bool) driver.CommandStatus {
    panic("unimplemented")
}

//GetOutputFile - used, but not in instruction stream
func (self *VirtualLiquidHandler) GetOutputFile() (string, driver.CommandStatus) {
    //Probably won't get called on the simulator just yet...
    panic("unimplemented")
}



