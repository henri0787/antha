// liquidhandling/lhinterfaces.go: Part of the Antha language
// Copyright (C) 2014 the Antha authors. All rights reserved.
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
// contact license@antha-lang.Org or write to the Antha team c/o
// Synthace Ltd. The London Bioscience Innovation Centre
// 2 Royal College St, London NW1 0NH UK

// defines types for dealing with liquid handling requests
package wtype

import "math"

//BBox is a simple LHObject representing a bounding box, 
//useful for checking if there's stuff in the way 
type BBox struct {
    position    Coordinates
    size        Coordinates
}

func NewBBox(pos, size Coordinates) *BBox {
    if size.X < 0. {
        pos.X = pos.X + size.X
        size.X = -size.X
    }
    if size.Y < 0. {
        pos.Y = pos.Y + size.Y
        size.Y = -size.Y
    }
    if size.Z < 0. {
        pos.Z = pos.Z + size.Z
        size.Z = -size.Z
    }
    r := BBox{pos, size}
    return &r
}

func NewBBox6f(pos_x, pos_y, pos_z, size_x, size_y, size_z float64) *BBox {
    return NewBBox(Coordinates{ pos_x,  pos_y,  pos_z}, 
                   Coordinates{size_x, size_y, size_z})
}

func NewXBox4f(pos_y, pos_z, size_y, size_z float64) *BBox {
    return NewBBox(Coordinates{-math.MaxFloat64 / 2.,  pos_y,  pos_z}, 
                   Coordinates{math.MaxFloat64, size_y, size_z})
}

func NewYBox4f(pos_x, pos_z, size_x, size_z float64) *BBox {
    return NewBBox(Coordinates{ pos_x,  -math.MaxFloat64 / 2.,  pos_z}, 
                   Coordinates{size_x, math.MaxFloat64, size_z})
}

func NewZBox4f(pos_x, pos_y, size_x, size_y float64) *BBox {
    return NewBBox(Coordinates{ pos_x,  pos_y,  -math.MaxFloat64 / 2.}, 
                   Coordinates{size_x, size_y, math.MaxFloat64})
}

func (self BBox) GetPosition() Coordinates {
    return self.position
}
func (self BBox) ZMax() float64 {
    return self.position.Z + self.size.Z
}

func (self BBox) GetSize() Coordinates {
    return self.size
}

func (self *BBox) SetPosition(c Coordinates) {
    self.position = c
}

func (self *BBox) SetSize(c Coordinates) {
    self.size = c
}

func (self BBox) Contains(rhs Coordinates) bool {
    return (rhs.X >= self.position.X && rhs.X < self.position.X + self.size.X &&
            rhs.Y >= self.position.Y && rhs.Y < self.position.Y + self.size.Y &&
            rhs.Z >= self.position.Z && rhs.Z < self.position.Z + self.size.Z)
}

//Intersects just checks for bounding box intersection
func (self BBox) Intersects(rhs BBox) bool {
    //test a single dimension. 
    //(a,b) are the start and end of the first position
    //(c,d) are the start and end of the second pos
    // assert(a > b  and  d > c)
    f := func(a,b,c,d float64) bool {
        return !(c >= b || d <= a)
    }

    s := self.position.Add(self.size)
    r :=  rhs.GetPosition().Add(rhs.GetSize())
    return (f(self.position.X, s.X, rhs.GetPosition().X, r.X) &&
            f(self.position.Y, s.Y, rhs.GetPosition().Y, r.Y) &&
            f(self.position.Z, s.Z, rhs.GetPosition().Z, r.Z))
}
