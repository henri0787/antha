// antha/AnthaStandardLibrary/Packages/enzymes/_typeIIsassembly_test.go: Part of the Antha language
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

package enzymes

import (
	"fmt"
	"strings"
	"testing"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/search"
	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/enzymes/lookup"
	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/sequences"
)

var pSEVA651 = wtype.DNASequence{
	Nm:      "test Elektrified pSEVA651",
	Seq:     strings.ToUpper("ttaattaaagcggataacaatttcacacaggaggccgcctaggccgcggccgcgcgaattcgagctcggtacccggggatcctctagagtcgacctgcaggcatgcaagcttgcggccgcgtcgtgactgggaaaaccctggcgactagtcttggactcctgttgatagatccagtaatgacctcagaactccatctggatttgttcagaacgctcggttgccgccgggcgttttttattggtgagaatccaggggtccccaataattacgatttaaatttgacataagcctgttcggttcgtaaactgtaatgcaagtagcgtatgcgctcacgcaactggtccagaaccttgaccgaacgcagcggtggtaacggcgcagtggcggttttcatggcttgttatgactgtttttttgtacagcctatgcctcgggcatccaagcagcaagcgcgttacgccgtgggtcgatgtttgatgttatggagcagcaacgatgttacgcagcagcaacgatgttacgcagcagggcagtcgccctaaaacaaagttaggtggctcaagtatgggcatcattcgcacatgtaggctcggccctgaccaagtcaaatccatgcgggctgctcttgatcttttcggtcgtgagttcggagacgtagccacctactcccaacatcagccggactccgattacctcgggaacttgctccgtagtaagacattcatcgcgcttgctgccttcgaccaagaagcggttgttggcgctctcgcggcttacgttctgcccaagtttgagcagccgcgtagtgagatctatatctatgatctcgcagtctccggagagcaccggaggcagggcattgccaccgcgctcatcaatctcctcaagcatgaggccaacgcgcttggtgcttatgtgatctacgtgcaagcagattacggtgacgatcccgcagtggctctctatacaaagttgggcatacgggaagaagtgatgcactttgatatcgacccaagtaccgccacctaacaattcgttcaagccgagatcggcttcccggccgcggagttgttcggtaaattggacaacggtccgcgcgttgtccttttccgctgcataaccctgcttcggggtcattatagcgattttttcggtatatccatcctttttcgcacgatatacaggattttgccaaagggttcgtgtagactttccttggtgtatccaacggcgtcagccgggcaggataggtgaagtaggcccacccgcgagcgggtgttccttcttcactgtcccttattcgcacctggcggtgctcaacgggaatcctgctctgcgaggctggccgtaggccggcctcagcctgccgccttgggccgggtgatgtcgtacttgcccgccgcgaactcggttaccgtccagcccagcgcgaccagctccggcaacgcctcgcgcacccgctggcggcgcttgcgcatggtcgaaccactggcctctgacggccagacatagccgcacaaggtatctatggaagccttgccggttttgccggggtcgatccagccacacagccgctggtgcagcaggcgggcggtttcgctgtccagcgcccgcacctcgtccatgctgatgcgcacatgctggccgccacccatgacggcctgcgcgatcaaggggttcagggccacgtacaggcgcccgtccgcctcgtcgctggcgtactccgacagcagccgaaacccctgccgcttgcggccattctgggcgatgatggataccttccaaaggcgctcgatgcagtcctgtatgtgcttgagcgccccaccactatcgacctctgccccgatttcctttgccagcgcccgatagctacctttgaccacatggcattcagcggtgacggcctcccacttgggttccaggaacagccggagctgccgtccgccttcggtcttgggttccgggccaagcactaggccattaggcccagccatggccaccagcccttgcaggatgcgcagatcatcagcgcccagcggctccgggccgctgaactcgatccgcttgccgtcgccgtagtcatacgtcacgtccagcttgctgcgcttgcgctcgccccgcttgagggcacggaacaggccgggggccagacagtgcgccgggtcgtgccggacgtggctgaggctgtgcttgttcttaggcttcaccacggggcacccccttgctcttgcgctgcctctccagcacggcgggcttgagcaccccgccgtcatgccgcctgaaccaccgatcagcgaacggtgcgccatagttggccttgctcacaccgaagcggacgaagaaccggcgctggtcgtcgtccacaccccattcctcggcctcggcgctggtcatgctcgacaggtaggactgccagcggatgttatcgaccagtaccgagctgccccggctggcctgctgctggtcgcctgcgcccatcatggccgcgcccttgctggcatggtgcaggaacacgatagagcacccggtatcggcggcgatggcctccatgcgaccgatgacctgggccatggggccgctggcgttttcttcctcgatgtggaaccggcgcagcgtgtccagcaccatcaggcggcggccctcggcggcgcgcttgaggccgtcgaaccactccggggccatgatgttgggcaggctgccgatcagcggctggatcagcaggccgtcagccacggcttgccgttcctcggcgctgaggtgcgccccaagggcgtgcaggcggtgatgaatggcggtgggcgggtcttcggcgggcaggtagatcaccgggccggtgggcagttcgcccacctccagcagatccggcccgcctgcaatctgtgcggccagttgcagggccagcatggatttaccggcaccaccgggcgacaccagcgccccgaccgtaccggccaccatgttgggcaaaacgtagtccagcggtggcggcgctgctgcgaacgcctccagaatattgataggcttatgggtagccattgattgcctcctttgcaggcagttggtggttaggcgctggcggggtcactacccccgccctgcgccgctctgagttcttccaggcactcgcgcagcgcctcgtattcgtcgtcggtcagccagaacttgcgctgacgcatccctttggccttcatgcgctcggcatatcgcgcttggcgtacagcgtcagggctggccagcaggtcgccggtctgcttgtccttttggtctttcatatcagtcaccgagaaacttgccggggccgaaaggcttgtcttcgcggaacaaggacaaggtgcagccgtcaaggttaaggctggccatatcagcgactgaaaagcggccagcctcggccttgtttgacgtataaccaaagccaccgggcaaccaatagcccttgtcacttttgatcaggtagaccgaccctgaagcgcttttttcgtattccataaaacccccttctgtgcgtgagtactcatagtataacaggcgtgagtaccaacgcaagcactacatgctgaaatctggcccgcccctgtccatgcctcgctggcggggtgccggtgcccgtgccagctcggcccgcgcaagctggacgctgggcagacccatgaccttgctgacggtgcgctcgatgtaatccgcttcgtggccgggcttgcgctctgccagcgctgggctggcctcggccatggccttgccgatttcctcggcactgcggccccggctggccagcttctgcgcggcgataaagtcgcacttgctgaggtcatcaccgaagcgcttgaccagcccggccatctcgctgcggtactcgtccagcgccgtgcgccggtggcggctaagctgccgctcgggcagttcgaggctggccagcctgcgggccttctcctgctgccgctgggcctgctcgatctgctggccagcctgctgcaccagcgccgggccagcggtggcggtcttgcccttggattcacgcagcagcacccacggctgataaccggcgcgggtggtgtgcttgtccttgcggttggtgaagcccgccaagcggccatagtggcggctgtcggcgctggccgggtcggcgtcgtactcgctggccagcgtccgggcaatctgcccccgaagttcaccgcctgcggcgtcggccaccttgacccatgcctgatagttcttcgggctggtttccactaccagggcaggctcccggccctcggctttcatgtcatccaggtcaaactcgctgaggtcgtccaccagcaccagaccatgccgctcctgctcggcgggcctgatatacacgtcattgccctgggcattcatccgcttgagccatggcgtgttctggagcacttcggcggctgaccattcccggttcatcatctggccggtggtggcgtccctgacgccgatatcgaagcgctcacagcccatggccttgagctgtcggcctatggcctgcaaagtcctgtcgttcttcatcgggccaccaagcgattcccacacattatacgagccggaagcataaagtgtaaagcctagatccgaaggatgagccgggctgaatgatcgaccgagacaggccctgcggggctgcacacgcgcccccacccttcgggtagggggaaaggccgctaaagcggctaaaagcgctccagcgtatttctgcggggtttggtgtggggtttagcgggctttgcccgcctttccccctgccgcgcagcggtggggcggtgtgtagcctagcgcagcgaatagaccagctatccggcctctggccgggcatattgggcaagggcagcagcgccccacaagggcgctgataaccgcgcctagtggattattcttagataatcatggatggatttttccaacaccccgccagcccccgcccctgctgggtttgcaggtttgggggcgtgacagttattgcaggggttcgtgacagttattgcaggggggcgtgacagttattgcaggggttcgtgacagttagggcgcgcccagctgtctagggcggcggatttgtcctactcaggagagcgttcaccgacaaacaacagataaaacgaaaggcccagtctttcgactgagcctttcgttttatttgatgcctGCTCTTCTATGatatatatatataGGTAGAAGAGC"),
	Plasmid: true,
}
var pSEVA651_2 = wtype.DNASequence{
	Nm:      "test Elektrified pSEVA651 sites to be removed",
	Seq:     strings.ToUpper("ttaattaaagcggataacaatttcacacaggaggccgcctaggccgcggccgcgcgaattcgagctcggtacccggggatcctctagagtcgacctgcaggcatgcaagcttgcggccgcgtcgtgactgggaaaaccctggcgactagtcttggactcctgttgatagatccagtaatgacctcagaactccatctggatttgttcagaacgctcggttgccgccgggcgttttttattggtgagaatccaggggtccccaataattacgatttaaatttgacataagcctgttcggttcgtaaactgtaatgcaagtagcgtatgcgctcacgcaactggtccagaaccttgaccgaacgcagcggtggtaacggcgcagtggcggttttcatggcttgttatgactgtttttttgtacagcctatgcctcgggcatccaagcagcaagcgcgttacgccgtgggtcgatgtttgatgttatggagcagcaacgatgttacgcagcagcaacgatgttacgcagcagggcagtcgccctaaaacaaagttaggtggctcaagtatgggcatcattcgcacatgtaggctcggccctgaccaagtcaaatccatgcgggctgctcttgatcttttcggtcgtgagttcggagacgtagccacctactcccaacatcagccggactccgattacctcgggaacttgctccgtagtaagacattcatcgcgcttgctgccttcgaccaagaagcggttgttggcgctctcgcggcttacgttctgcccaagtttgagcagccgcgtagtgagatctatatctatgatctcgcagtctccggagagcaccggaggcagggcattgccaccgcgctcatcaatctcctcaagcatgaggccaacgcgcttggtgcttatgtgatctacgtgcaagcagattacggtgacgatcccgcagtggctctctatacaaagttgggcatacgggaagaagtgatgcactttgatatcgacccaagtaccgccacctaacaattcgttcaagccgagatcggcttcccggccgcggagttgttcggtaaattggacaacggtccgcgcgttgtccttttccgctgcataaccctgcttcggggtcattatagcgattttttcggtatatccatcctttttcgcacgatatacaggattttgccaaagggttcgtgtagactttccttggtgtatccaacggcgtcagccgggcaggataggtgaagtaggcccacccgcgagcgggtgttccttcttcactgtcccttattcgcacctggcggtgctcaacgggaatcctgctctgcgaggctggccgtaggccggcctcagcctgccgccttgggccgggtgatgtcgtacttgcccgccgcgaactcggttaccgtccagcccagcgcgaccagctccggcaacgcctcgcgcacccgctggcggcgcttgcgcatggtcgaaccactggcctctgacggccagacatagccgcacaaggtatctatggaagccttgccggttttgccggggtcgatccagccacacagccgctggtgcagcaggcgggcggtttcgctgtccagcgcccgcacctcgtccatgctgatgcgcacatgctggccgccacccatgacggcctgcgcgatcaaggggttcagggccacgtacaggcgcccgtccgcctcgtcgctggcgtactccgacagcagccgaaacccctgccgcttgcggccattctgggcgatgatggataccttccaaaggcgctcgatgcagtcctgtatgtgcttgagcgccccaccactatcgacctctgccccgatttcctttgccagcgcccgatagctacctttgaccacatggcattcagcggtgacggcctcccacttgggttccaggaacagccggagctgccgtccgccttcggtcttgggttccgggccaagcactaggccattaggcccagccatggccaccagcccttgcaggatgcgcagatcatcagcgcccagcggctccgggccgctgaactcgatccgcttgccgtcgccgtagtcatacgtcacgtccagcttgctgcgcttgcgctcgccccgcttgagggcacggaacaggccgggggccagacagtgcgccgggtcgtgccggacgtggctgaggctgtgcttgttcttaggcttcaccacggggcacccccttgctcttgcgctgcctctccagcacggcgggcttgagcaccccgccgtcatgccgcctgaaccaccgatcagcgaacggtgcgccatagttggccttgctcacaccgaagcggacgaagaaccggcgctggtcgtcgtccacaccccattcctcggcctcggcgctggtcatgctcgacaggtaggactgccagcggatgttatcgaccagtaccgagctgccccggctggcctgctgctggtcgcctgcgcccatcatggccgcgcccttgctggcatggtgcaggaacacgatagagcacccggtatcggcggcgatggcctccatgcgaccgatgacctgggccatggggccgctggcgttttcttcctcgatgtggaaccggcgcagcgtgtccagcaccatcaggcggcggccctcggcggcgcgcttgaggccgtcgaaccactccggggccatgatgttgggcaggctgccgatcagcggctggatcagcaggccgtcagccacggcttgccgttcctcggcgctgaggtgcgccccaagggcgtgcaggcggtgatgaatggcggtgggcgggtcttcggcgggcaggtagatcaccgggccggtgggcagttcgcccacctccagcagatccggcccgcctgcaatctgtgcggccagttgcagggccagcatggatttaccggcaccaccgggcgacaccagcgccccgaccgtaccggccaccatgttgggcaaaacgtagtccagcggtggcggcgctgctgcgaacgcctccagaatattgataggcttatgggtagccattgattgcctcctttgcaggcagttggtggttaggcgctggcggggtcactacccccgccctgcgccgctctgagttcttccaggcactcgcgcagcgcctcgtattcgtcgtcggtcagccagaacttgcgctgacgcatccctttggccttcatgcgctcggcatatcgcgcttggcgtacagcgtcagggctggccagcaggtcgccggtctgcttgtccttttggtctttcatatcagtcaccgagaaacttgccggggccgaaaggcttgtcttcgcggaacaaggacaaggtgcagccgtcaaggttaaggctggccatatcagcgactgaaaagcggccagcctcggccttgtttgacgtataaccaaagccaccgggcaaccaatagcccttgtcacttttgatcaggtagaccgaccctgaagcgcttttttcgtattccataaaacccccttctgtgcgtgagtactcatagtataacaggcgtgagtaccaacgcaagcactacatgctgaaatctggcccgcccctgtccatgcctcgctggcggggtgccggtgcccgtgccagctcggcccgcgcaagctggacgctgggcagacccatgaccttgctgacggtgcgctcgatgtaatccgcttcgtggccgggcttgcgctctgccagcgctgggctggcctcggccatggccttgccgatttcctcggcactgcggccccggctggccagcttctgcgcggcgataaagtcgcacttgctgaggtcatcaccgaagcgcttgaccagcccggccatctcgctgcggtactcgtccagcgccgtgcgccggtggcggctaagctgccgctcgggcagttcgaggctggccagcctgcgggccttctcctgctgccgctgggcctgctcgatctgctggccagcctgctgcaccagcgccgggccagcggtggcggtcttgcccttggattcacgcagcagcacccacggctgataaccggcgcgggtggtgtgcttgtccttgcggttggtgaagcccgccaagcggccatagtggcggctgtcggcgctggccgggtcggcgtcgtactcgctggccagcgtccgggcaatctgcccccgaagttcaccgcctgcggcgtcggccaccttgacccatgcctgatagttcttcgggctggtttccactaccagggcaggctcccggccctcggctttcatgtcatccaggtcaaactcgctgaggtcgtccaccagcaccagaccatgccgctcctgctcggcgggcctgatatacacgtcattgccctgggcattcatccgcttgagccatggcgtgttctggagcacttcggcggctgaccattcccggttcatcatctggccggtggtggcgtccctgacgccgatatcgaagcgctcacagcccatggccttgagctgtcggcctatggcctgcaaagtcctgtcgttcttcatcgggccaccaagcgattcccacacattatacgagccggaagcataaagtgtaaagcctagatccgaaggatgagccgggctgaatgatcgaccgagacaggccctgcggggctgcacacgcgcccccacccttcgggtagggggaaaggccgctaaagcggctaaaagcgctccagcgtatttctgcggggtttggtgtggggtttagcgggctttgcccgcctttccccctgccgcgcagcggtggggcggtgtgtagcctagcgcagcgaatagaccagctatccggcctctggccgggcatattgggcaagggcagcagcgccccacaagggcgctgataaccgcgcctagtggattattcttagataatcatggatggatttttccaacaccccgccagcccccgcccctgctgggtttgcaggtttgggggcgtgacagttattgcaggggttcgtgacagttattgcaggggggcgtgacagttattgcaggggttcgtgacagttagggcgcgcccagctgtctagggcggcggatttgtcctactcaggagagcgttcaccgacaaacaacagataaaacgaaaggcccagtctttcgactgagcctttcgttttatttgatgcctataATGAGAAGAGCatatatGCTCTTCTGGTatat"),
	Plasmid: true,
}
var comet1part = wtype.DNASequence{
	Nm:      "test comet 1 part",
	Seq:     strings.ToUpper("gctcttctatgacggcattgacggaaggcgcaaaattgttcgaaaaagaaatcccatacatcaccgaactggaaggcgatgttgaaggtatgaagttcatcattaagggtgagggcaccggcgatgcaactacgggcaccattaaagcgaagtatatctgcaccaccggtgacgttccggtgccgtggagcacgctggtcaccaccctgacctatggcgcgcagtgtttcgcgaagtacggtccggaactgaaggacttctataagagctgtatgcctgagggctatgttcaggagcgtaccattacctttgagggtgatggtgtctttaagacgcgtgctgaggtgacctttgagaatggttccgtgtacaatcgcgtgaaactgaatggtcaaggttttaagaaagatggtcacgtgctgggcaaaaacctggagtttaactttactccgcattgcctgtgcatttggggcgaccaagcgaaccacggtctgaaaagcgcgttcaagattatgcacgagattacgggtagcaaagaggacttcatcgtggccgaccacacgcagatgaacaccccgatcggtggcggtccggtccatgtcccggagtaccaccacttgaccgtttggacctctttcggtaaagacccggatgatgacgaaacggatcatctgaatattgttgaggttatcaaagccgtcgacctggaaacttaccgttaatgataatgaggtagaagagc"),
	Plasmid: false,
}

var comet2part1 = wtype.DNASequence{
	Nm:      "test comet 2.1 part",
	Seq:     strings.ToUpper("gctcttctatgacggcattgacggaaggcgcaaaattgttcgaaaaagaaatcccatacatcaccgaactggaaggcgatgttgaaggtatgaagttcatcattaagggtgagggcaccggcgatgcaactacgggcaccattaaagcgaagtatatctgcaccaccggtgacgttccggtgccgtggagcacgctggtcaccaccctgacctatggcgcgcagtgtttcgcgaagtacggtccggaactggaagagc"),
	Plasmid: false,
}
var comet2part2 = wtype.DNASequence{
	Nm:      "test comet 2.2 part",
	Seq:     strings.ToUpper("gctcttcgactgaaggacttctataagagctgtatgcctgagggctatgttcaggagcgtaccattacctttgagggtgatggtgtctttaagacgcgtgctgaggtgacctttgagaatggttccgtgtacaatcgcgtgaaactgaatggtcaaggttttaagaaagatggtcacgtgctgggcaaaaacctggagtttaactttactccgcattgcctgtgcatttggggcgaccaagcgaaccacggtctgaaaagcgcgttcaagattatgcacgagattacgggtagcaaagaggacttcatcgtggccgaccacacgcagatgaacaccccgatcggtggcggtccggtccatgtcccggagtaccaccacttgaccgtttggacctctttcggtaaagacccggatgatgacgaaacggatcatctgaatattgttgaggttatcaaagccgtcgacctggaaacttaccgttaatgataatgaggtagaagagc"),
	Plasmid: false,
}

func bbsI() wtype.TypeIIs{
	enzyme, _ := lookup.TypeIIs("bbsI")
	return enzyme
}

func ecoRI() wtype.RestrictionEnzyme{
	enzyme, _ := lookup.RestrictionEnzyme("EcoRI")
	return enzyme
}

type restrictionsitetest struct {
	sequence   wtype.DNASequence
	enzymelist []wtype.RestrictionEnzyme
	sitesfound []RestrictionSites
}

var restrictionsitetests = []restrictionsitetest{
	{sequence: pSEVA651,
	enzymelist:	[]wtype.RestrictionEnzyme{SapI.RestrictionEnzyme, BsaI.RestrictionEnzyme},
	sitesfound:	[]RestrictionSites{
			{
				SapI.RestrictionEnzyme,
				[]sequences.PositionPair{
					sequences.PositionPair{StartPosition:5154,EndPosition: 5160,Reverse: false},
					sequences.PositionPair{StartPosition:5188, EndPosition:5182,Reverse:true},
				},
			},
			{
				BsaI.RestrictionEnzyme,
				[]sequences.PositionPair{	
				},
			},
		},
	},
	{sequence: wtype.DNASequence{Nm:"Forward SapI test",Seq:"GCTCTTCTGGT"},
		enzymelist: []wtype.RestrictionEnzyme{SapI.RestrictionEnzyme},
		sitesfound: []RestrictionSites{
			{
				SapI.RestrictionEnzyme,
				[]sequences.PositionPair{
					sequences.PositionPair{StartPosition:1,EndPosition: 7,Reverse: false},
				},
			},
		},
	},
	{sequence: wtype.DNASequence{Nm:"Reverse SapI test",Seq:"GGTAGAAGAGC"},
		enzymelist: []wtype.RestrictionEnzyme{SapI.RestrictionEnzyme},
		sitesfound: []RestrictionSites{
			{
				SapI.RestrictionEnzyme,
				[]sequences.PositionPair{
					sequences.PositionPair{StartPosition: 11,EndPosition:5,Reverse: true},
				},
			},
		},
	},
	{sequence: wtype.DNASequence{Nm:"Forward and Reverse SapI test",Seq:"GCTCTTCTGGTGGTAGAAGAGC"},
		enzymelist: []wtype.RestrictionEnzyme{SapI.RestrictionEnzyme},
		sitesfound: []RestrictionSites{
			{
				SapI.RestrictionEnzyme,
				[]sequences.PositionPair{
					sequences.PositionPair{StartPosition:1, EndPosition:7,Reverse: false},
					sequences.PositionPair{StartPosition:22,EndPosition:16, Reverse: true},
				},
			},
		},
	},
	{sequence: wtype.DNASequence{Nm:"2 Forward and Reverse SapI test",Seq:"GCTCTTCTGGTGCTCTTCTGGTGGTAGAAGAGC"},
		enzymelist: []wtype.RestrictionEnzyme{SapI.RestrictionEnzyme},
		sitesfound: []RestrictionSites{
			{
				SapI.RestrictionEnzyme,
				[]sequences.PositionPair{
					sequences.PositionPair{StartPosition:1, EndPosition:7, Reverse: false},
					sequences.PositionPair{StartPosition:12, EndPosition:18, Reverse: false},
					sequences.PositionPair{StartPosition:33, EndPosition:27,Reverse: true},
				},
			},
		},
	},
}

func TestRestrictionsitefinder(t *testing.T) {
	for _, test := range restrictionsitetests {
		sitesFound := RestrictionSiteFinder(test.sequence, test.enzymelist)
		for i := 0; i < len(sitesFound); i++ {
			if sitesFound[i].NumberOfSites() != test.sitesfound[i].NumberOfSites() {
				t.Error(
					"For", test.sequence.Name(), "\n",
					"and", test.enzymelist[i].Name(), "\n",
					"expected", test.sitesfound[i].NumberOfSites(), "sites \n",
					"got", sitesFound[i].NumberOfSites(), "sites \n",
				)
			}
			if len(sitesFound[i].ForwardPositions()) != len(test.sitesfound[i].ForwardPositions()){
					t.Error(
						"For", test.sequence.Name(), "\n",
						"and", test.enzymelist[i].Name(), "\n",
						"expected", test.sitesfound[i].ForwardPositions(), "\n",
						"got", sitesFound[i].ForwardPositions(), "\n",
					)
					
			}else{
				for j := range sitesFound[i].ForwardPositions() {
					if sitesFound[i].ForwardPositions()[j] != test.sitesfound[i].ForwardPositions()[j] {
						t.Error(
							"For", test.sequence.Name(), "\n",
							"and", test.enzymelist[i].Name(), "\n",
							"expected forward positions ", test.sitesfound[i].ForwardPositions(), "\n",
							"got", sitesFound[i].ForwardPositions(), "\n",
						)
					}
				}
			}
			if len(sitesFound[i].ReversePositions()) != len(test.sitesfound[i].ReversePositions()){
					t.Error(
						"For", test.sequence.Name(), "\n",
						"and", test.enzymelist[i].Name(), "\n",
						"expected reverse positions ", test.sitesfound[i].ReversePositions(), "\n",
						"got", sitesFound[i].ReversePositions(), "\n",
					)
					
			}else {
				for j := range sitesFound[i].ReversePositions() {
					if sitesFound[i].ReversePositions()[j] != test.sitesfound[i].ReversePositions()[j] {
						t.Error(
							"For", test.sequence.Name(), "\n",
							"and", test.enzymelist[i].Name(), "\n",
							"expected", test.sitesfound[i].ReversePositions(), "\n",
							"got", sitesFound[i].ReversePositions(), "\n",
						)
					}
				}
			}
		}
	}
}

type digesttest struct {
	sequence          wtype.DNASequence
	enzyme            wtype.RestrictionEnzyme
	Finalfragments    []string
	fivePrimeOverhangs []string
	threePrimeUnderhangs []string
}

var digesttests = []digesttest{
	digesttest{
		sequence:          pSEVA651,
		enzyme:            SapI.RestrictionEnzyme,
		Finalfragments:    []string{strings.ToUpper("GGTAGAAGAGCttaattaaagcggataacaatttcacacaggaggccgcctaggccgcggccgcgcgaattcgagctcggtacccggggatcctctagagtcgacctgcaggcatgcaagcttgcggccgcgtcgtgactgggaaaaccctggcgactagtcttggactcctgttgatagatccagtaatgacctcagaactccatctggatttgttcagaacgctcggttgccgccgggcgttttttattggtgagaatccaggggtccccaataattacgatttaaatttgacataagcctgttcggttcgtaaactgtaatgcaagtagcgtatgcgctcacgcaactggtccagaaccttgaccgaacgcagcggtggtaacggcgcagtggcggttttcatggcttgttatgactgtttttttgtacagcctatgcctcgggcatccaagcagcaagcgcgttacgccgtgggtcgatgtttgatgttatggagcagcaacgatgttacgcagcagcaacgatgttacgcagcagggcagtcgccctaaaacaaagttaggtggctcaagtatgggcatcattcgcacatgtaggctcggccctgaccaagtcaaatccatgcgggctgctcttgatcttttcggtcgtgagttcggagacgtagccacctactcccaacatcagccggactccgattacctcgggaacttgctccgtagtaagacattcatcgcgcttgctgccttcgaccaagaagcggttgttggcgctctcgcggcttacgttctgcccaagtttgagcagccgcgtagtgagatctatatctatgatctcgcagtctccggagagcaccggaggcagggcattgccaccgcgctcatcaatctcctcaagcatgaggccaacgcgcttggtgcttatgtgatctacgtgcaagcagattacggtgacgatcccgcagtggctctctatacaaagttgggcatacgggaagaagtgatgcactttgatatcgacccaagtaccgccacctaacaattcgttcaagccgagatcggcttcccggccgcggagttgttcggtaaattggacaacggtccgcgcgttgtccttttccgctgcataaccctgcttcggggtcattatagcgattttttcggtatatccatcctttttcgcacgatatacaggattttgccaaagggttcgtgtagactttccttggtgtatccaacggcgtcagccgggcaggataggtgaagtaggcccacccgcgagcgggtgttccttcttcactgtcccttattcgcacctggcggtgctcaacgggaatcctgctctgcgaggctggccgtaggccggcctcagcctgccgccttgggccgggtgatgtcgtacttgcccgccgcgaactcggttaccgtccagcccagcgcgaccagctccggcaacgcctcgcgcacccgctggcggcgcttgcgcatggtcgaaccactggcctctgacggccagacatagccgcacaaggtatctatggaagccttgccggttttgccggggtcgatccagccacacagccgctggtgcagcaggcgggcggtttcgctgtccagcgcccgcacctcgtccatgctgatgcgcacatgctggccgccacccatgacggcctgcgcgatcaaggggttcagggccacgtacaggcgcccgtccgcctcgtcgctggcgtactccgacagcagccgaaacccctgccgcttgcggccattctgggcgatgatggataccttccaaaggcgctcgatgcagtcctgtatgtgcttgagcgccccaccactatcgacctctgccccgatttcctttgccagcgcccgatagctacctttgaccacatggcattcagcggtgacggcctcccacttgggttccaggaacagccggagctgccgtccgccttcggtcttgggttccgggccaagcactaggccattaggcccagccatggccaccagcccttgcaggatgcgcagatcatcagcgcccagcggctccgggccgctgaactcgatccgcttgccgtcgccgtagtcatacgtcacgtccagcttgctgcgcttgcgctcgccccgcttgagggcacggaacaggccgggggccagacagtgcgccgggtcgtgccggacgtggctgaggctgtgcttgttcttaggcttcaccacggggcacccccttgctcttgcgctgcctctccagcacggcgggcttgagcaccccgccgtcatgccgcctgaaccaccgatcagcgaacggtgcgccatagttggccttgctcacaccgaagcggacgaagaaccggcgctggtcgtcgtccacaccccattcctcggcctcggcgctggtcatgctcgacaggtaggactgccagcggatgttatcgaccagtaccgagctgccccggctggcctgctgctggtcgcctgcgcccatcatggccgcgcccttgctggcatggtgcaggaacacgatagagcacccggtatcggcggcgatggcctccatgcgaccgatgacctgggccatggggccgctggcgttttcttcctcgatgtggaaccggcgcagcgtgtccagcaccatcaggcggcggccctcggcggcgcgcttgaggccgtcgaaccactccggggccatgatgttgggcaggctgccgatcagcggctggatcagcaggccgtcagccacggcttgccgttcctcggcgctgaggtgcgccccaagggcgtgcaggcggtgatgaatggcggtgggcgggtcttcggcgggcaggtagatcaccgggccggtgggcagttcgcccacctccagcagatccggcccgcctgcaatctgtgcggccagttgcagggccagcatggatttaccggcaccaccgggcgacaccagcgccccgaccgtaccggccaccatgttgggcaaaacgtagtccagcggtggcggcgctgctgcgaacgcctccagaatattgataggcttatgggtagccattgattgcctcctttgcaggcagttggtggttaggcgctggcggggtcactacccccgccctgcgccgctctgagttcttccaggcactcgcgcagcgcctcgtattcgtcgtcggtcagccagaacttgcgctgacgcatccctttggccttcatgcgctcggcatatcgcgcttggcgtacagcgtcagggctggccagcaggtcgccggtctgcttgtccttttggtctttcatatcagtcaccgagaaacttgccggggccgaaaggcttgtcttcgcggaacaaggacaaggtgcagccgtcaaggttaaggctggccatatcagcgactgaaaagcggccagcctcggccttgtttgacgtataaccaaagccaccgggcaaccaatagcccttgtcacttttgatcaggtagaccgaccctgaagcgcttttttcgtattccataaaacccccttctgtgcgtgagtactcatagtataacaggcgtgagtaccaacgcaagcactacatgctgaaatctggcccgcccctgtccatgcctcgctggcggggtgccggtgcccgtgccagctcggcccgcgcaagctggacgctgggcagacccatgaccttgctgacggtgcgctcgatgtaatccgcttcgtggccgggcttgcgctctgccagcgctgggctggcctcggccatggccttgccgatttcctcggcactgcggccccggctggccagcttctgcgcggcgataaagtcgcacttgctgaggtcatcaccgaagcgcttgaccagcccggccatctcgctgcggtactcgtccagcgccgtgcgccggtggcggctaagctgccgctcgggcagttcgaggctggccagcctgcgggccttctcctgctgccgctgggcctgctcgatctgctggccagcctgctgcaccagcgccgggccagcggtggcggtcttgcccttggattcacgcagcagcacccacggctgataaccggcgcgggtggtgtgcttgtccttgcggttggtgaagcccgccaagcggccatagtggcggctgtcggcgctggccgggtcggcgtcgtactcgctggccagcgtccgggcaatctgcccccgaagttcaccgcctgcggcgtcggccaccttgacccatgcctgatagttcttcgggctggtttccactaccagggcaggctcccggccctcggctttcatgtcatccaggtcaaactcgctgaggtcgtccaccagcaccagaccatgccgctcctgctcggcgggcctgatatacacgtcattgccctgggcattcatccgcttgagccatggcgtgttctggagcacttcggcggctgaccattcccggttcatcatctggccggtggtggcgtccctgacgccgatatcgaagcgctcacagcccatggccttgagctgtcggcctatggcctgcaaagtcctgtcgttcttcatcgggccaccaagcgattcccacacattatacgagccggaagcataaagtgtaaagcctagatccgaaggatgagccgggctgaatgatcgaccgagacaggccctgcggggctgcacacgcgcccccacccttcgggtagggggaaaggccgctaaagcggctaaaagcgctccagcgtatttctgcggggtttggtgtggggtttagcgggctttgcccgcctttccccctgccgcgcagcggtggggcggtgtgtagcctagcgcagcgaatagaccagctatccggcctctggccgggcatattgggcaagggcagcagcgccccacaagggcgctgataaccgcgcctagtggattattcttagataatcatggatggatttttccaacaccccgccagcccccgcccctgctgggtttgcaggtttgggggcgtgacagttattgcaggggttcgtgacagttattgcaggggggcgtgacagttattgcaggggttcgtgacagttagggcgcgcccagctgtctagggcggcggatttgtcctactcaggagagcgttcaccgacaaacaacagataaaacgaaaggcccagtctttcgactgagcctttcgttttatttgatgcctGCTCTTCT"), strings.ToUpper("ATGatatatatatata")},
		fivePrimeOverhangs: []string{"GGT", "ATG"},
		threePrimeUnderhangs: []string{"CAT", "ACC"},
	},
	// BsaI is a non cutter
	digesttest{
		sequence:          pSEVA651,
		enzyme:            BsaI.RestrictionEnzyme,
		Finalfragments:    []string{strings.ToUpper("ttaattaaagcggataacaatttcacacaggaggccgcctaggccgcggccgcgcgaattcgagctcggtacccggggatcctctagagtcgacctgcaggcatgcaagcttgcggccgcgtcgtgactgggaaaaccctggcgactagtcttggactcctgttgatagatccagtaatgacctcagaactccatctggatttgttcagaacgctcggttgccgccgggcgttttttattggtgagaatccaggggtccccaataattacgatttaaatttgacataagcctgttcggttcgtaaactgtaatgcaagtagcgtatgcgctcacgcaactggtccagaaccttgaccgaacgcagcggtggtaacggcgcagtggcggttttcatggcttgttatgactgtttttttgtacagcctatgcctcgggcatccaagcagcaagcgcgttacgccgtgggtcgatgtttgatgttatggagcagcaacgatgttacgcagcagcaacgatgttacgcagcagggcagtcgccctaaaacaaagttaggtggctcaagtatgggcatcattcgcacatgtaggctcggccctgaccaagtcaaatccatgcgggctgctcttgatcttttcggtcgtgagttcggagacgtagccacctactcccaacatcagccggactccgattacctcgggaacttgctccgtagtaagacattcatcgcgcttgctgccttcgaccaagaagcggttgttggcgctctcgcggcttacgttctgcccaagtttgagcagccgcgtagtgagatctatatctatgatctcgcagtctccggagagcaccggaggcagggcattgccaccgcgctcatcaatctcctcaagcatgaggccaacgcgcttggtgcttatgtgatctacgtgcaagcagattacggtgacgatcccgcagtggctctctatacaaagttgggcatacgggaagaagtgatgcactttgatatcgacccaagtaccgccacctaacaattcgttcaagccgagatcggcttcccggccgcggagttgttcggtaaattggacaacggtccgcgcgttgtccttttccgctgcataaccctgcttcggggtcattatagcgattttttcggtatatccatcctttttcgcacgatatacaggattttgccaaagggttcgtgtagactttccttggtgtatccaacggcgtcagccgggcaggataggtgaagtaggcccacccgcgagcgggtgttccttcttcactgtcccttattcgcacctggcggtgctcaacgggaatcctgctctgcgaggctggccgtaggccggcctcagcctgccgccttgggccgggtgatgtcgtacttgcccgccgcgaactcggttaccgtccagcccagcgcgaccagctccggcaacgcctcgcgcacccgctggcggcgcttgcgcatggtcgaaccactggcctctgacggccagacatagccgcacaaggtatctatggaagccttgccggttttgccggggtcgatccagccacacagccgctggtgcagcaggcgggcggtttcgctgtccagcgcccgcacctcgtccatgctgatgcgcacatgctggccgccacccatgacggcctgcgcgatcaaggggttcagggccacgtacaggcgcccgtccgcctcgtcgctggcgtactccgacagcagccgaaacccctgccgcttgcggccattctgggcgatgatggataccttccaaaggcgctcgatgcagtcctgtatgtgcttgagcgccccaccactatcgacctctgccccgatttcctttgccagcgcccgatagctacctttgaccacatggcattcagcggtgacggcctcccacttgggttccaggaacagccggagctgccgtccgccttcggtcttgggttccgggccaagcactaggccattaggcccagccatggccaccagcccttgcaggatgcgcagatcatcagcgcccagcggctccgggccgctgaactcgatccgcttgccgtcgccgtagtcatacgtcacgtccagcttgctgcgcttgcgctcgccccgcttgagggcacggaacaggccgggggccagacagtgcgccgggtcgtgccggacgtggctgaggctgtgcttgttcttaggcttcaccacggggcacccccttgctcttgcgctgcctctccagcacggcgggcttgagcaccccgccgtcatgccgcctgaaccaccgatcagcgaacggtgcgccatagttggccttgctcacaccgaagcggacgaagaaccggcgctggtcgtcgtccacaccccattcctcggcctcggcgctggtcatgctcgacaggtaggactgccagcggatgttatcgaccagtaccgagctgccccggctggcctgctgctggtcgcctgcgcccatcatggccgcgcccttgctggcatggtgcaggaacacgatagagcacccggtatcggcggcgatggcctccatgcgaccgatgacctgggccatggggccgctggcgttttcttcctcgatgtggaaccggcgcagcgtgtccagcaccatcaggcggcggccctcggcggcgcgcttgaggccgtcgaaccactccggggccatgatgttgggcaggctgccgatcagcggctggatcagcaggccgtcagccacggcttgccgttcctcggcgctgaggtgcgccccaagggcgtgcaggcggtgatgaatggcggtgggcgggtcttcggcgggcaggtagatcaccgggccggtgggcagttcgcccacctccagcagatccggcccgcctgcaatctgtgcggccagttgcagggccagcatggatttaccggcaccaccgggcgacaccagcgccccgaccgtaccggccaccatgttgggcaaaacgtagtccagcggtggcggcgctgctgcgaacgcctccagaatattgataggcttatgggtagccattgattgcctcctttgcaggcagttggtggttaggcgctggcggggtcactacccccgccctgcgccgctctgagttcttccaggcactcgcgcagcgcctcgtattcgtcgtcggtcagccagaacttgcgctgacgcatccctttggccttcatgcgctcggcatatcgcgcttggcgtacagcgtcagggctggccagcaggtcgccggtctgcttgtccttttggtctttcatatcagtcaccgagaaacttgccggggccgaaaggcttgtcttcgcggaacaaggacaaggtgcagccgtcaaggttaaggctggccatatcagcgactgaaaagcggccagcctcggccttgtttgacgtataaccaaagccaccgggcaaccaatagcccttgtcacttttgatcaggtagaccgaccctgaagcgcttttttcgtattccataaaacccccttctgtgcgtgagtactcatagtataacaggcgtgagtaccaacgcaagcactacatgctgaaatctggcccgcccctgtccatgcctcgctggcggggtgccggtgcccgtgccagctcggcccgcgcaagctggacgctgggcagacccatgaccttgctgacggtgcgctcgatgtaatccgcttcgtggccgggcttgcgctctgccagcgctgggctggcctcggccatggccttgccgatttcctcggcactgcggccccggctggccagcttctgcgcggcgataaagtcgcacttgctgaggtcatcaccgaagcgcttgaccagcccggccatctcgctgcggtactcgtccagcgccgtgcgccggtggcggctaagctgccgctcgggcagttcgaggctggccagcctgcgggccttctcctgctgccgctgggcctgctcgatctgctggccagcctgctgcaccagcgccgggccagcggtggcggtcttgcccttggattcacgcagcagcacccacggctgataaccggcgcgggtggtgtgcttgtccttgcggttggtgaagcccgccaagcggccatagtggcggctgtcggcgctggccgggtcggcgtcgtactcgctggccagcgtccgggcaatctgcccccgaagttcaccgcctgcggcgtcggccaccttgacccatgcctgatagttcttcgggctggtttccactaccagggcaggctcccggccctcggctttcatgtcatccaggtcaaactcgctgaggtcgtccaccagcaccagaccatgccgctcctgctcggcgggcctgatatacacgtcattgccctgggcattcatccgcttgagccatggcgtgttctggagcacttcggcggctgaccattcccggttcatcatctggccggtggtggcgtccctgacgccgatatcgaagcgctcacagcccatggccttgagctgtcggcctatggcctgcaaagtcctgtcgttcttcatcgggccaccaagcgattcccacacattatacgagccggaagcataaagtgtaaagcctagatccgaaggatgagccgggctgaatgatcgaccgagacaggccctgcggggctgcacacgcgcccccacccttcgggtagggggaaaggccgctaaagcggctaaaagcgctccagcgtatttctgcggggtttggtgtggggtttagcgggctttgcccgcctttccccctgccgcgcagcggtggggcggtgtgtagcctagcgcagcgaatagaccagctatccggcctctggccgggcatattgggcaagggcagcagcgccccacaagggcgctgataaccgcgcctagtggattattcttagataatcatggatggatttttccaacaccccgccagcccccgcccctgctgggtttgcaggtttgggggcgtgacagttattgcaggggttcgtgacagttattgcaggggggcgtgacagttattgcaggggttcgtgacagttagggcgcgcccagctgtctagggcggcggatttgtcctactcaggagagcgttcaccgacaaacaacagataaaacgaaaggcccagtctttcgactgagcctttcgttttatttgatgcctGCTCTTCTATGatatatatatataGGTAGAAGAGC")},
		fivePrimeOverhangs: []string{""},
		threePrimeUnderhangs: []string{""},
	},
	digesttest{
		sequence:          pSEVA651,
		enzyme:            ecoRI(),
		Finalfragments:    []string{strings.ToUpper("cggtacccggggatcctctagagtcgacctgcaggcatgcaagcttgcggccgcgtcgtgactgggaaaaccctggcgactagtcttggactcctgttgatagatccagtaatgacctcagaactccatctggatttgttcagaacgctcggttgccgccgggcgttttttattggtgagaatccaggggtccccaataattacgatttaaatttgacataagcctgttcggttcgtaaactgtaatgcaagtagcgtatgcgctcacgcaactggtccagaaccttgaccgaacgcagcggtggtaacggcgcagtggcggttttcatggcttgttatgactgtttttttgtacagcctatgcctcgggcatccaagcagcaagcgcgttacgccgtgggtcgatgtttgatgttatggagcagcaacgatgttacgcagcagcaacgatgttacgcagcagggcagtcgccctaaaacaaagttaggtggctcaagtatgggcatcattcgcacatgtaggctcggccctgaccaagtcaaatccatgcgggctgctcttgatcttttcggtcgtgagttcggagacgtagccacctactcccaacatcagccggactccgattacctcgggaacttgctccgtagtaagacattcatcgcgcttgctgccttcgaccaagaagcggttgttggcgctctcgcggcttacgttctgcccaagtttgagcagccgcgtagtgagatctatatctatgatctcgcagtctccggagagcaccggaggcagggcattgccaccgcgctcatcaatctcctcaagcatgaggccaacgcgcttggtgcttatgtgatctacgtgcaagcagattacggtgacgatcccgcagtggctctctatacaaagttgggcatacgggaagaagtgatgcactttgatatcgacccaagtaccgccacctaacaattcgttcaagccgagatcggcttcccggccgcggagttgttcggtaaattggacaacggtccgcgcgttgtccttttccgctgcataaccctgcttcggggtcattatagcgattttttcggtatatccatcctttttcgcacgatatacaggattttgccaaagggttcgtgtagactttccttggtgtatccaacggcgtcagccgggcaggataggtgaagtaggcccacccgcgagcgggtgttccttcttcactgtcccttattcgcacctggcggtgctcaacgggaatcctgctctgcgaggctggccgtaggccggcctcagcctgccgccttgggccgggtgatgtcgtacttgcccgccgcgaactcggttaccgtccagcccagcgcgaccagctccggcaacgcctcgcgcacccgctggcggcgcttgcgcatggtcgaaccactggcctctgacggccagacatagccgcacaaggtatctatggaagccttgccggttttgccggggtcgatccagccacacagccgctggtgcagcaggcgggcggtttcgctgtccagcgcccgcacctcgtccatgctgatgcgcacatgctggccgccacccatgacggcctgcgcgatcaaggggttcagggccacgtacaggcgcccgtccgcctcgtcgctggcgtactccgacagcagccgaaacccctgccgcttgcggccattctgggcgatgatggataccttccaaaggcgctcgatgcagtcctgtatgtgcttgagcgccccaccactatcgacctctgccccgatttcctttgccagcgcccgatagctacctttgaccacatggcattcagcggtgacggcctcccacttgggttccaggaacagccggagctgccgtccgccttcggtcttgggttccgggccaagcactaggccattaggcccagccatggccaccagcccttgcaggatgcgcagatcatcagcgcccagcggctccgggccgctgaactcgatccgcttgccgtcgccgtagtcatacgtcacgtccagcttgctgcgcttgcgctcgccccgcttgagggcacggaacaggccgggggccagacagtgcgccgggtcgtgccggacgtggctgaggctgtgcttgttcttaggcttcaccacggggcacccccttgctcttgcgctgcctctccagcacggcgggcttgagcaccccgccgtcatgccgcctgaaccaccgatcagcgaacggtgcgccatagttggccttgctcacaccgaagcggacgaagaaccggcgctggtcgtcgtccacaccccattcctcggcctcggcgctggtcatgctcgacaggtaggactgccagcggatgttatcgaccagtaccgagctgccccggctggcctgctgctggtcgcctgcgcccatcatggccgcgcccttgctggcatggtgcaggaacacgatagagcacccggtatcggcggcgatggcctccatgcgaccgatgacctgggccatggggccgctggcgttttcttcctcgatgtggaaccggcgcagcgtgtccagcaccatcaggcggcggccctcggcggcgcgcttgaggccgtcgaaccactccggggccatgatgttgggcaggctgccgatcagcggctggatcagcaggccgtcagccacggcttgccgttcctcggcgctgaggtgcgccccaagggcgtgcaggcggtgatgaatggcggtgggcgggtcttcggcgggcaggtagatcaccgggccggtgggcagttcgcccacctccagcagatccggcccgcctgcaatctgtgcggccagttgcagggccagcatggatttaccggcaccaccgggcgacaccagcgccccgaccgtaccggccaccatgttgggcaaaacgtagtccagcggtggcggcgctgctgcgaacgcctccagaatattgataggcttatgggtagccattgattgcctcctttgcaggcagttggtggttaggcgctggcggggtcactacccccgccctgcgccgctctgagttcttccaggcactcgcgcagcgcctcgtattcgtcgtcggtcagccagaacttgcgctgacgcatccctttggccttcatgcgctcggcatatcgcgcttggcgtacagcgtcagggctggccagcaggtcgccggtctgcttgtccttttggtctttcatatcagtcaccgagaaacttgccggggccgaaaggcttgtcttcgcggaacaaggacaaggtgcagccgtcaaggttaaggctggccatatcagcgactgaaaagcggccagcctcggccttgtttgacgtataaccaaagccaccgggcaaccaatagcccttgtcacttttgatcaggtagaccgaccctgaagcgcttttttcgtattccataaaacccccttctgtgcgtgagtactcatagtataacaggcgtgagtaccaacgcaagcactacatgctgaaatctggcccgcccctgtccatgcctcgctggcggggtgccggtgcccgtgccagctcggcccgcgcaagctggacgctgggcagacccatgaccttgctgacggtgcgctcgatgtaatccgcttcgtggccgggcttgcgctctgccagcgctgggctggcctcggccatggccttgccgatttcctcggcactgcggccccggctggccagcttctgcgcggcgataaagtcgcacttgctgaggtcatcaccgaagcgcttgaccagcccggccatctcgctgcggtactcgtccagcgccgtgcgccggtggcggctaagctgccgctcgggcagttcgaggctggccagcctgcgggccttctcctgctgccgctgggcctgctcgatctgctggccagcctgctgcaccagcgccgggccagcggtggcggtcttgcccttggattcacgcagcagcacccacggctgataaccggcgcgggtggtgtgcttgtccttgcggttggtgaagcccgccaagcggccatagtggcggctgtcggcgctggccgggtcggcgtcgtactcgctggccagcgtccgggcaatctgcccccgaagttcaccgcctgcggcgtcggccaccttgacccatgcctgatagttcttcgggctggtttccactaccagggcaggctcccggccctcggctttcatgtcatccaggtcaaactcgctgaggtcgtccaccagcaccagaccatgccgctcctgctcggcgggcctgatatacacgtcattgccctgggcattcatccgcttgagccatggcgtgttctggagcacttcggcggctgaccattcccggttcatcatctggccggtggtggcgtccctgacgccgatatcgaagcgctcacagcccatggccttgagctgtcggcctatggcctgcaaagtcctgtcgttcttcatcgggccaccaagcgattcccacacattatacgagccggaagcataaagtgtaaagcctagatccgaaggatgagccgggctgaatgatcgaccgagacaggccctgcggggctgcacacgcgcccccacccttcgggtagggggaaaggccgctaaagcggctaaaagcgctccagcgtatttctgcggggtttggtgtggggtttagcgggctttgcccgcctttccccctgccgcgcagcggtggggcggtgtgtagcctagcgcagcgaatagaccagctatccggcctctggccgggcatattgggcaagggcagcagcgccccacaagggcgctgataaccgcgcctagtggattattcttagataatcatggatggatttttccaacaccccgccagcccccgcccctgctgggtttgcaggtttgggggcgtgacagttattgcaggggttcgtgacagttattgcaggggggcgtgacagttattgcaggggttcgtgacagttagggcgcgcccagctgtctagggcggcggatttgtcctactcaggagagcgttcaccgacaaacaacagataaaacgaaaggcccagtctttcgactgagcctttcgttttatttgatgcctGCTCTTCTGGTAGAAGAGCttaattaaagcggataacaatttcacacaggaggccgcctaggccgcggccgcgcg"), strings.ToUpper("aattcgagct")},
		fivePrimeOverhangs: []string{"AGCT", "AATT"},
		threePrimeUnderhangs: []string{"AATT", "AGCT"},
	},
	digesttest{
		sequence:          wtype.DNASequence{Nm:"Forward SapI test",Seq:"GCTCTTCTGGTAAA"},
		enzyme:            SapI.RestrictionEnzyme,
		Finalfragments:    []string{"GCTCTTCT", "GGTAAA"},
		fivePrimeOverhangs: []string{"blunt","GGT"},
		threePrimeUnderhangs: []string{"ACC","blunt"},
	},
	digesttest{
		sequence:          wtype.DNASequence{Nm:"Forward SapI test plasmid",Seq:"AAAAAATGGTAAAGCTCTTCCCCCCCC", Plasmid: true},
		enzyme:            SapI.RestrictionEnzyme,
		Finalfragments:    []string{"CCCCCCAAAAAATGGTAAAGCTCTTCC"},
		fivePrimeOverhangs: []string{"CCC"},
		threePrimeUnderhangs: []string{"GGG"},
	},
	digesttest{
		sequence:          wtype.DNASequence{Nm:"Reverse SapI test",Seq:"AAAGGTAGAAGAGC"},
		enzyme:            SapI.RestrictionEnzyme,
		Finalfragments:    []string{"AAA","GGTAGAAGAGC"},
		fivePrimeOverhangs: []string{"blunt","GGT"},
		threePrimeUnderhangs: []string{"ACC","blunt"},
	},
	digesttest{
		sequence:          wtype.DNASequence{Nm:"2 Forward and Reverse SapI test",Seq:"GCTCTTCTGGTGCTCTTCTGGTGGTAGAAGAGC"},
		enzyme:            SapI.RestrictionEnzyme,
		Finalfragments:    []string{"GCTCTTCT","GGTGCTCTTCT", "GGT","GGTAGAAGAGC"},
		fivePrimeOverhangs: []string{"blunt","GGT","GGT","GGT"},
		threePrimeUnderhangs: []string{"ACC","ACC","ACC","blunt"},
	},
}


func TestDigest(t *testing.T) {
	for _, test := range digesttests {
		fragments, fivePrimeStickyEnds, threePrimeStickeyEnds := Digest(test.sequence, test.enzyme)
		if len(fragments) != len(test.Finalfragments){
			t.Error(
				"For", test.sequence.Name(), "\n",
				"and", test.enzyme.Name(), "\n",
				"expected fragments: \n", strings.Join(test.Finalfragments,"\n"), "\n",
				"got \n", strings.Join(fragments,"\n"), "\n",
			)
					
		}else {
			for i := 0; i < len(fragments); i++ {
				if fragments[i] != test.Finalfragments[i] {
					t.Error(
						"For", test.sequence.Name(), "\n",
						"and", test.enzyme.Name(), "\n",
						"expected fragment: ", test.Finalfragments[i], "\n",
						"got", fragments[i], "\n",
					)
				}
			}
		}
		if len(fivePrimeStickyEnds) != len(test.fivePrimeOverhangs){
			t.Error(
				"For", test.sequence.Name(), "\n",
				"and", test.enzyme.Name(), "\n",
				"expected 5 prime sticky ends:", test.fivePrimeOverhangs, "\n",
				"got", fivePrimeStickyEnds, "\n",
			)
					
		}else {
			var fail bool
			for i := 0; i < len(fivePrimeStickyEnds); i++ {
				if fivePrimeStickyEnds[i] != test.fivePrimeOverhangs[i] {
						fail = true
						break
				}
			}
			if fail {
				t.Error(
						"For", test.sequence.Name(), "\n",
						"and", test.enzyme.Name(), "\n",
						"expected 5 prime sticky end", test.fivePrimeOverhangs, "\n",
						"got", fivePrimeStickyEnds, "\n",
					)
			}
		}
		if len(threePrimeStickeyEnds) != len(test.threePrimeUnderhangs){
			t.Error(
				"For", test.sequence.Name(), "\n",
				"and", test.enzyme.Name(), "\n",
				"expected 3 prime sticky ends:", test.threePrimeUnderhangs, "\n",
				"got", threePrimeStickeyEnds, "\n",
			)
					
		}else {
			var fail bool
			
			for i := 0; i < len(threePrimeStickeyEnds); i++ {
				if threePrimeStickeyEnds[i] != test.threePrimeUnderhangs[i] {
					fail = true
					break
				}
			}
			if fail {
				t.Error(
					"For", test.sequence.Name(), "\n",
					"and", test.enzyme.Name(), "\n",
					"expected 3 prime sticky ends:", test.threePrimeUnderhangs, "\n",
					"got", threePrimeStickeyEnds, "\n",
				)
			}
		}
	}
}

// simple reverse complement check to test testing methodology initially
type testsequence struct {
	sequence string
	revcomp  string
}

var seqs = []testsequence{
	{"ATG", "CAT"},
	{"YNB", "VNR"},
}

func TestRevComp(t *testing.T) {
	for _, sequence := range seqs {
		r := wtype.RevComp(sequence.sequence)
		if r != sequence.revcomp {
			t.Error(
				"For", sequence.sequence, "/n",
				"expected", sequence.revcomp, "\n",
				"got", r, "\n",
			)
		}
	}

}

type assemblytest struct {
	parts                  []wtype.DNASequence
	vector                 wtype.DNASequence
	enzyme                 wtype.TypeIIs
	desiredassemblyproduct wtype.DNASequence
}

var assemblytests = []assemblytest{
	assemblytest{
		parts: []wtype.DNASequence{comet1part},
		vector: 	pSEVA651_2,
		enzyme:	SapI,
		desiredassemblyproduct: 	wtype.DNASequence{
			Nm: "simulated assembly sequence", 
			Seq: strings.ToUpper("GGTATATTTAATTAAAGCGGATAACAATTTCACACAGGAGGCCGCCTAGGCCGCGGCCGCGCGAATTCGAGCTCGGTACCCGGGGATCCTCTAGAGTCGACCTGCAGGCATGCAAGCTTGCGGCCGCGTCGTGACTGGGAAAACCCTGGCGACTAGTCTTGGACTCCTGTTGATAGATCCAGTAATGACCTCAGAACTCCATCTGGATTTGTTCAGAACGCTCGGTTGCCGCCGGGCGTTTTTTATTGGTGAGAATCCAGGGGTCCCCAATAATTACGATTTAAATTTGACATAAGCCTGTTCGGTTCGTAAACTGTAATGCAAGTAGCGTATGCGCTCACGCAACTGGTCCAGAACCTTGACCGAACGCAGCGGTGGTAACGGCGCAGTGGCGGTTTTCATGGCTTGTTATGACTGTTTTTTTGTACAGCCTATGCCTCGGGCATCCAAGCAGCAAGCGCGTTACGCCGTGGGTCGATGTTTGATGTTATGGAGCAGCAACGATGTTACGCAGCAGCAACGATGTTACGCAGCAGGGCAGTCGCCCTAAAACAAAGTTAGGTGGCTCAAGTATGGGCATCATTCGCACATGTAGGCTCGGCCCTGACCAAGTCAAATCCATGCGGGCTGCTCTTGATCTTTTCGGTCGTGAGTTCGGAGACGTAGCCACCTACTCCCAACATCAGCCGGACTCCGATTACCTCGGGAACTTGCTCCGTAGTAAGACATTCATCGCGCTTGCTGCCTTCGACCAAGAAGCGGTTGTTGGCGCTCTCGCGGCTTACGTTCTGCCCAAGTTTGAGCAGCCGCGTAGTGAGATCTATATCTATGATCTCGCAGTCTCCGGAGAGCACCGGAGGCAGGGCATTGCCACCGCGCTCATCAATCTCCTCAAGCATGAGGCCAACGCGCTTGGTGCTTATGTGATCTACGTGCAAGCAGATTACGGTGACGATCCCGCAGTGGCTCTCTATACAAAGTTGGGCATACGGGAAGAAGTGATGCACTTTGATATCGACCCAAGTACCGCCACCTAACAATTCGTTCAAGCCGAGATCGGCTTCCCGGCCGCGGAGTTGTTCGGTAAATTGGACAACGGTCCGCGCGTTGTCCTTTTCCGCTGCATAACCCTGCTTCGGGGTCATTATAGCGATTTTTTCGGTATATCCATCCTTTTTCGCACGATATACAGGATTTTGCCAAAGGGTTCGTGTAGACTTTCCTTGGTGTATCCAACGGCGTCAGCCGGGCAGGATAGGTGAAGTAGGCCCACCCGCGAGCGGGTGTTCCTTCTTCACTGTCCCTTATTCGCACCTGGCGGTGCTCAACGGGAATCCTGCTCTGCGAGGCTGGCCGTAGGCCGGCCTCAGCCTGCCGCCTTGGGCCGGGTGATGTCGTACTTGCCCGCCGCGAACTCGGTTACCGTCCAGCCCAGCGCGACCAGCTCCGGCAACGCCTCGCGCACCCGCTGGCGGCGCTTGCGCATGGTCGAACCACTGGCCTCTGACGGCCAGACATAGCCGCACAAGGTATCTATGGAAGCCTTGCCGGTTTTGCCGGGGTCGATCCAGCCACACAGCCGCTGGTGCAGCAGGCGGGCGGTTTCGCTGTCCAGCGCCCGCACCTCGTCCATGCTGATGCGCACATGCTGGCCGCCACCCATGACGGCCTGCGCGATCAAGGGGTTCAGGGCCACGTACAGGCGCCCGTCCGCCTCGTCGCTGGCGTACTCCGACAGCAGCCGAAACCCCTGCCGCTTGCGGCCATTCTGGGCGATGATGGATACCTTCCAAAGGCGCTCGATGCAGTCCTGTATGTGCTTGAGCGCCCCACCACTATCGACCTCTGCCCCGATTTCCTTTGCCAGCGCCCGATAGCTACCTTTGACCACATGGCATTCAGCGGTGACGGCCTCCCACTTGGGTTCCAGGAACAGCCGGAGCTGCCGTCCGCCTTCGGTCTTGGGTTCCGGGCCAAGCACTAGGCCATTAGGCCCAGCCATGGCCACCAGCCCTTGCAGGATGCGCAGATCATCAGCGCCCAGCGGCTCCGGGCCGCTGAACTCGATCCGCTTGCCGTCGCCGTAGTCATACGTCACGTCCAGCTTGCTGCGCTTGCGCTCGCCCCGCTTGAGGGCACGGAACAGGCCGGGGGCCAGACAGTGCGCCGGGTCGTGCCGGACGTGGCTGAGGCTGTGCTTGTTCTTAGGCTTCACCACGGGGCACCCCCTTGCTCTTGCGCTGCCTCTCCAGCACGGCGGGCTTGAGCACCCCGCCGTCATGCCGCCTGAACCACCGATCAGCGAACGGTGCGCCATAGTTGGCCTTGCTCACACCGAAGCGGACGAAGAACCGGCGCTGGTCGTCGTCCACACCCCATTCCTCGGCCTCGGCGCTGGTCATGCTCGACAGGTAGGACTGCCAGCGGATGTTATCGACCAGTACCGAGCTGCCCCGGCTGGCCTGCTGCTGGTCGCCTGCGCCCATCATGGCCGCGCCCTTGCTGGCATGGTGCAGGAACACGATAGAGCACCCGGTATCGGCGGCGATGGCCTCCATGCGACCGATGACCTGGGCCATGGGGCCGCTGGCGTTTTCTTCCTCGATGTGGAACCGGCGCAGCGTGTCCAGCACCATCAGGCGGCGGCCCTCGGCGGCGCGCTTGAGGCCGTCGAACCACTCCGGGGCCATGATGTTGGGCAGGCTGCCGATCAGCGGCTGGATCAGCAGGCCGTCAGCCACGGCTTGCCGTTCCTCGGCGCTGAGGTGCGCCCCAAGGGCGTGCAGGCGGTGATGAATGGCGGTGGGCGGGTCTTCGGCGGGCAGGTAGATCACCGGGCCGGTGGGCAGTTCGCCCACCTCCAGCAGATCCGGCCCGCCTGCAATCTGTGCGGCCAGTTGCAGGGCCAGCATGGATTTACCGGCACCACCGGGCGACACCAGCGCCCCGACCGTACCGGCCACCATGTTGGGCAAAACGTAGTCCAGCGGTGGCGGCGCTGCTGCGAACGCCTCCAGAATATTGATAGGCTTATGGGTAGCCATTGATTGCCTCCTTTGCAGGCAGTTGGTGGTTAGGCGCTGGCGGGGTCACTACCCCCGCCCTGCGCCGCTCTGAGTTCTTCCAGGCACTCGCGCAGCGCCTCGTATTCGTCGTCGGTCAGCCAGAACTTGCGCTGACGCATCCCTTTGGCCTTCATGCGCTCGGCATATCGCGCTTGGCGTACAGCGTCAGGGCTGGCCAGCAGGTCGCCGGTCTGCTTGTCCTTTTGGTCTTTCATATCAGTCACCGAGAAACTTGCCGGGGCCGAAAGGCTTGTCTTCGCGGAACAAGGACAAGGTGCAGCCGTCAAGGTTAAGGCTGGCCATATCAGCGACTGAAAAGCGGCCAGCCTCGGCCTTGTTTGACGTATAACCAAAGCCACCGGGCAACCAATAGCCCTTGTCACTTTTGATCAGGTAGACCGACCCTGAAGCGCTTTTTTCGTATTCCATAAAACCCCCTTCTGTGCGTGAGTACTCATAGTATAACAGGCGTGAGTACCAACGCAAGCACTACATGCTGAAATCTGGCCCGCCCCTGTCCATGCCTCGCTGGCGGGGTGCCGGTGCCCGTGCCAGCTCGGCCCGCGCAAGCTGGACGCTGGGCAGACCCATGACCTTGCTGACGGTGCGCTCGATGTAATCCGCTTCGTGGCCGGGCTTGCGCTCTGCCAGCGCTGGGCTGGCCTCGGCCATGGCCTTGCCGATTTCCTCGGCACTGCGGCCCCGGCTGGCCAGCTTCTGCGCGGCGATAAAGTCGCACTTGCTGAGGTCATCACCGAAGCGCTTGACCAGCCCGGCCATCTCGCTGCGGTACTCGTCCAGCGCCGTGCGCCGGTGGCGGCTAAGCTGCCGCTCGGGCAGTTCGAGGCTGGCCAGCCTGCGGGCCTTCTCCTGCTGCCGCTGGGCCTGCTCGATCTGCTGGCCAGCCTGCTGCACCAGCGCCGGGCCAGCGGTGGCGGTCTTGCCCTTGGATTCACGCAGCAGCACCCACGGCTGATAACCGGCGCGGGTGGTGTGCTTGTCCTTGCGGTTGGTGAAGCCCGCCAAGCGGCCATAGTGGCGGCTGTCGGCGCTGGCCGGGTCGGCGTCGTACTCGCTGGCCAGCGTCCGGGCAATCTGCCCCCGAAGTTCACCGCCTGCGGCGTCGGCCACCTTGACCCATGCCTGATAGTTCTTCGGGCTGGTTTCCACTACCAGGGCAGGCTCCCGGCCCTCGGCTTTCATGTCATCCAGGTCAAACTCGCTGAGGTCGTCCACCAGCACCAGACCATGCCGCTCCTGCTCGGCGGGCCTGATATACACGTCATTGCCCTGGGCATTCATCCGCTTGAGCCATGGCGTGTTCTGGAGCACTTCGGCGGCTGACCATTCCCGGTTCATCATCTGGCCGGTGGTGGCGTCCCTGACGCCGATATCGAAGCGCTCACAGCCCATGGCCTTGAGCTGTCGGCCTATGGCCTGCAAAGTCCTGTCGTTCTTCATCGGGCCACCAAGCGATTCCCACACATTATACGAGCCGGAAGCATAAAGTGTAAAGCCTAGATCCGAAGGATGAGCCGGGCTGAATGATCGACCGAGACAGGCCCTGCGGGGCTGCACACGCGCCCCCACCCTTCGGGTAGGGGGAAAGGCCGCTAAAGCGGCTAAAAGCGCTCCAGCGTATTTCTGCGGGGTTTGGTGTGGGGTTTAGCGGGCTTTGCCCGCCTTTCCCCCTGCCGCGCAGCGGTGGGGCGGTGTGTAGCCTAGCGCAGCGAATAGACCAGCTATCCGGCCTCTGGCCGGGCATATTGGGCAAGGGCAGCAGCGCCCCACAAGGGCGCTGATAACCGCGCCTAGTGGATTATTCTTAGATAATCATGGATGGATTTTTCCAACACCCCGCCAGCCCCCGCCCCTGCTGGGTTTGCAGGTTTGGGGGCGTGACAGTTATTGCAGGGGTTCGTGACAGTTATTGCAGGGGGGCGTGACAGTTATTGCAGGGGTTCGTGACAGTTAGGGCGCGCCCAGCTGTCTAGGGCGGCGGATTTGTCCTACTCAGGAGAGCGTTCACCGACAAACAACAGATAAAACGAAAGGCCCAGTCTTTCGACTGAGCCTTTCGTTTTATTTGATGCCTATAATGACGGCATTGACGGAAGGCGCAAAATTGTTCGAAAAAGAAATCCCATACATCACCGAACTGGAAGGCGATGTTGAAGGTATGAAGTTCATCATTAAGGGTGAGGGCACCGGCGATGCAACTACGGGCACCATTAAAGCGAAGTATATCTGCACCACCGGTGACGTTCCGGTGCCGTGGAGCACGCTGGTCACCACCCTGACCTATGGCGCGCAGTGTTTCGCGAAGTACGGTCCGGAACTGAAGGACTTCTATAAGAGCTGTATGCCTGAGGGCTATGTTCAGGAGCGTACCATTACCTTTGAGGGTGATGGTGTCTTTAAGACGCGTGCTGAGGTGACCTTTGAGAATGGTTCCGTGTACAATCGCGTGAAACTGAATGGTCAAGGTTTTAAGAAAGATGGTCACGTGCTGGGCAAAAACCTGGAGTTTAACTTTACTCCGCATTGCCTGTGCATTTGGGGCGACCAAGCGAACCACGGTCTGAAAAGCGCGTTCAAGATTATGCACGAGATTACGGGTAGCAAAGAGGACTTCATCGTGGCCGACCACACGCAGATGAACACCCCGATCGGTGGCGGTCCGGTCCATGTCCCGGAGTACCACCACTTGACCGTTTGGACCTCTTTCGGTAAAGACCCGGATGATGACGAAACGGATCATCTGAATATTGTTGAGGTTATCAAAGCCGTCGACCTGGAAACTTACCGTTAATGATAATGA"),
			Plasmid: true,
		},
	},
	assemblytest{
		parts: []wtype.DNASequence{comet2part1, comet2part2},
		vector:	pSEVA651_2,
		enzyme:	SapI,
		desiredassemblyproduct: 	wtype.DNASequence{
			Nm: "simulated assembly sequence", 
			Seq: strings.ToUpper("GGTATATTTAATTAAAGCGGATAACAATTTCACACAGGAGGCCGCCTAGGCCGCGGCCGCGCGAATTCGAGCTCGGTACCCGGGGATCCTCTAGAGTCGACCTGCAGGCATGCAAGCTTGCGGCCGCGTCGTGACTGGGAAAACCCTGGCGACTAGTCTTGGACTCCTGTTGATAGATCCAGTAATGACCTCAGAACTCCATCTGGATTTGTTCAGAACGCTCGGTTGCCGCCGGGCGTTTTTTATTGGTGAGAATCCAGGGGTCCCCAATAATTACGATTTAAATTTGACATAAGCCTGTTCGGTTCGTAAACTGTAATGCAAGTAGCGTATGCGCTCACGCAACTGGTCCAGAACCTTGACCGAACGCAGCGGTGGTAACGGCGCAGTGGCGGTTTTCATGGCTTGTTATGACTGTTTTTTTGTACAGCCTATGCCTCGGGCATCCAAGCAGCAAGCGCGTTACGCCGTGGGTCGATGTTTGATGTTATGGAGCAGCAACGATGTTACGCAGCAGCAACGATGTTACGCAGCAGGGCAGTCGCCCTAAAACAAAGTTAGGTGGCTCAAGTATGGGCATCATTCGCACATGTAGGCTCGGCCCTGACCAAGTCAAATCCATGCGGGCTGCTCTTGATCTTTTCGGTCGTGAGTTCGGAGACGTAGCCACCTACTCCCAACATCAGCCGGACTCCGATTACCTCGGGAACTTGCTCCGTAGTAAGACATTCATCGCGCTTGCTGCCTTCGACCAAGAAGCGGTTGTTGGCGCTCTCGCGGCTTACGTTCTGCCCAAGTTTGAGCAGCCGCGTAGTGAGATCTATATCTATGATCTCGCAGTCTCCGGAGAGCACCGGAGGCAGGGCATTGCCACCGCGCTCATCAATCTCCTCAAGCATGAGGCCAACGCGCTTGGTGCTTATGTGATCTACGTGCAAGCAGATTACGGTGACGATCCCGCAGTGGCTCTCTATACAAAGTTGGGCATACGGGAAGAAGTGATGCACTTTGATATCGACCCAAGTACCGCCACCTAACAATTCGTTCAAGCCGAGATCGGCTTCCCGGCCGCGGAGTTGTTCGGTAAATTGGACAACGGTCCGCGCGTTGTCCTTTTCCGCTGCATAACCCTGCTTCGGGGTCATTATAGCGATTTTTTCGGTATATCCATCCTTTTTCGCACGATATACAGGATTTTGCCAAAGGGTTCGTGTAGACTTTCCTTGGTGTATCCAACGGCGTCAGCCGGGCAGGATAGGTGAAGTAGGCCCACCCGCGAGCGGGTGTTCCTTCTTCACTGTCCCTTATTCGCACCTGGCGGTGCTCAACGGGAATCCTGCTCTGCGAGGCTGGCCGTAGGCCGGCCTCAGCCTGCCGCCTTGGGCCGGGTGATGTCGTACTTGCCCGCCGCGAACTCGGTTACCGTCCAGCCCAGCGCGACCAGCTCCGGCAACGCCTCGCGCACCCGCTGGCGGCGCTTGCGCATGGTCGAACCACTGGCCTCTGACGGCCAGACATAGCCGCACAAGGTATCTATGGAAGCCTTGCCGGTTTTGCCGGGGTCGATCCAGCCACACAGCCGCTGGTGCAGCAGGCGGGCGGTTTCGCTGTCCAGCGCCCGCACCTCGTCCATGCTGATGCGCACATGCTGGCCGCCACCCATGACGGCCTGCGCGATCAAGGGGTTCAGGGCCACGTACAGGCGCCCGTCCGCCTCGTCGCTGGCGTACTCCGACAGCAGCCGAAACCCCTGCCGCTTGCGGCCATTCTGGGCGATGATGGATACCTTCCAAAGGCGCTCGATGCAGTCCTGTATGTGCTTGAGCGCCCCACCACTATCGACCTCTGCCCCGATTTCCTTTGCCAGCGCCCGATAGCTACCTTTGACCACATGGCATTCAGCGGTGACGGCCTCCCACTTGGGTTCCAGGAACAGCCGGAGCTGCCGTCCGCCTTCGGTCTTGGGTTCCGGGCCAAGCACTAGGCCATTAGGCCCAGCCATGGCCACCAGCCCTTGCAGGATGCGCAGATCATCAGCGCCCAGCGGCTCCGGGCCGCTGAACTCGATCCGCTTGCCGTCGCCGTAGTCATACGTCACGTCCAGCTTGCTGCGCTTGCGCTCGCCCCGCTTGAGGGCACGGAACAGGCCGGGGGCCAGACAGTGCGCCGGGTCGTGCCGGACGTGGCTGAGGCTGTGCTTGTTCTTAGGCTTCACCACGGGGCACCCCCTTGCTCTTGCGCTGCCTCTCCAGCACGGCGGGCTTGAGCACCCCGCCGTCATGCCGCCTGAACCACCGATCAGCGAACGGTGCGCCATAGTTGGCCTTGCTCACACCGAAGCGGACGAAGAACCGGCGCTGGTCGTCGTCCACACCCCATTCCTCGGCCTCGGCGCTGGTCATGCTCGACAGGTAGGACTGCCAGCGGATGTTATCGACCAGTACCGAGCTGCCCCGGCTGGCCTGCTGCTGGTCGCCTGCGCCCATCATGGCCGCGCCCTTGCTGGCATGGTGCAGGAACACGATAGAGCACCCGGTATCGGCGGCGATGGCCTCCATGCGACCGATGACCTGGGCCATGGGGCCGCTGGCGTTTTCTTCCTCGATGTGGAACCGGCGCAGCGTGTCCAGCACCATCAGGCGGCGGCCCTCGGCGGCGCGCTTGAGGCCGTCGAACCACTCCGGGGCCATGATGTTGGGCAGGCTGCCGATCAGCGGCTGGATCAGCAGGCCGTCAGCCACGGCTTGCCGTTCCTCGGCGCTGAGGTGCGCCCCAAGGGCGTGCAGGCGGTGATGAATGGCGGTGGGCGGGTCTTCGGCGGGCAGGTAGATCACCGGGCCGGTGGGCAGTTCGCCCACCTCCAGCAGATCCGGCCCGCCTGCAATCTGTGCGGCCAGTTGCAGGGCCAGCATGGATTTACCGGCACCACCGGGCGACACCAGCGCCCCGACCGTACCGGCCACCATGTTGGGCAAAACGTAGTCCAGCGGTGGCGGCGCTGCTGCGAACGCCTCCAGAATATTGATAGGCTTATGGGTAGCCATTGATTGCCTCCTTTGCAGGCAGTTGGTGGTTAGGCGCTGGCGGGGTCACTACCCCCGCCCTGCGCCGCTCTGAGTTCTTCCAGGCACTCGCGCAGCGCCTCGTATTCGTCGTCGGTCAGCCAGAACTTGCGCTGACGCATCCCTTTGGCCTTCATGCGCTCGGCATATCGCGCTTGGCGTACAGCGTCAGGGCTGGCCAGCAGGTCGCCGGTCTGCTTGTCCTTTTGGTCTTTCATATCAGTCACCGAGAAACTTGCCGGGGCCGAAAGGCTTGTCTTCGCGGAACAAGGACAAGGTGCAGCCGTCAAGGTTAAGGCTGGCCATATCAGCGACTGAAAAGCGGCCAGCCTCGGCCTTGTTTGACGTATAACCAAAGCCACCGGGCAACCAATAGCCCTTGTCACTTTTGATCAGGTAGACCGACCCTGAAGCGCTTTTTTCGTATTCCATAAAACCCCCTTCTGTGCGTGAGTACTCATAGTATAACAGGCGTGAGTACCAACGCAAGCACTACATGCTGAAATCTGGCCCGCCCCTGTCCATGCCTCGCTGGCGGGGTGCCGGTGCCCGTGCCAGCTCGGCCCGCGCAAGCTGGACGCTGGGCAGACCCATGACCTTGCTGACGGTGCGCTCGATGTAATCCGCTTCGTGGCCGGGCTTGCGCTCTGCCAGCGCTGGGCTGGCCTCGGCCATGGCCTTGCCGATTTCCTCGGCACTGCGGCCCCGGCTGGCCAGCTTCTGCGCGGCGATAAAGTCGCACTTGCTGAGGTCATCACCGAAGCGCTTGACCAGCCCGGCCATCTCGCTGCGGTACTCGTCCAGCGCCGTGCGCCGGTGGCGGCTAAGCTGCCGCTCGGGCAGTTCGAGGCTGGCCAGCCTGCGGGCCTTCTCCTGCTGCCGCTGGGCCTGCTCGATCTGCTGGCCAGCCTGCTGCACCAGCGCCGGGCCAGCGGTGGCGGTCTTGCCCTTGGATTCACGCAGCAGCACCCACGGCTGATAACCGGCGCGGGTGGTGTGCTTGTCCTTGCGGTTGGTGAAGCCCGCCAAGCGGCCATAGTGGCGGCTGTCGGCGCTGGCCGGGTCGGCGTCGTACTCGCTGGCCAGCGTCCGGGCAATCTGCCCCCGAAGTTCACCGCCTGCGGCGTCGGCCACCTTGACCCATGCCTGATAGTTCTTCGGGCTGGTTTCCACTACCAGGGCAGGCTCCCGGCCCTCGGCTTTCATGTCATCCAGGTCAAACTCGCTGAGGTCGTCCACCAGCACCAGACCATGCCGCTCCTGCTCGGCGGGCCTGATATACACGTCATTGCCCTGGGCATTCATCCGCTTGAGCCATGGCGTGTTCTGGAGCACTTCGGCGGCTGACCATTCCCGGTTCATCATCTGGCCGGTGGTGGCGTCCCTGACGCCGATATCGAAGCGCTCACAGCCCATGGCCTTGAGCTGTCGGCCTATGGCCTGCAAAGTCCTGTCGTTCTTCATCGGGCCACCAAGCGATTCCCACACATTATACGAGCCGGAAGCATAAAGTGTAAAGCCTAGATCCGAAGGATGAGCCGGGCTGAATGATCGACCGAGACAGGCCCTGCGGGGCTGCACACGCGCCCCCACCCTTCGGGTAGGGGGAAAGGCCGCTAAAGCGGCTAAAAGCGCTCCAGCGTATTTCTGCGGGGTTTGGTGTGGGGTTTAGCGGGCTTTGCCCGCCTTTCCCCCTGCCGCGCAGCGGTGGGGCGGTGTGTAGCCTAGCGCAGCGAATAGACCAGCTATCCGGCCTCTGGCCGGGCATATTGGGCAAGGGCAGCAGCGCCCCACAAGGGCGCTGATAACCGCGCCTAGTGGATTATTCTTAGATAATCATGGATGGATTTTTCCAACACCCCGCCAGCCCCCGCCCCTGCTGGGTTTGCAGGTTTGGGGGCGTGACAGTTATTGCAGGGGTTCGTGACAGTTATTGCAGGGGGGCGTGACAGTTATTGCAGGGGTTCGTGACAGTTAGGGCGCGCCCAGCTGTCTAGGGCGGCGGATTTGTCCTACTCAGGAGAGCGTTCACCGACAAACAACAGATAAAACGAAAGGCCCAGTCTTTCGACTGAGCCTTTCGTTTTATTTGATGCCTATAATGACGGCATTGACGGAAGGCGCAAAATTGTTCGAAAAAGAAATCCCATACATCACCGAACTGGAAGGCGATGTTGAAGGTATGAAGTTCATCATTAAGGGTGAGGGCACCGGCGATGCAACTACGGGCACCATTAAAGCGAAGTATATCTGCACCACCGGTGACGTTCCGGTGCCGTGGAGCACGCTGGTCACCACCCTGACCTATGGCGCGCAGTGTTTCGCGAAGTACGGTCCGGAACTGAAGGACTTCTATAAGAGCTGTATGCCTGAGGGCTATGTTCAGGAGCGTACCATTACCTTTGAGGGTGATGGTGTCTTTAAGACGCGTGCTGAGGTGACCTTTGAGAATGGTTCCGTGTACAATCGCGTGAAACTGAATGGTCAAGGTTTTAAGAAAGATGGTCACGTGCTGGGCAAAAACCTGGAGTTTAACTTTACTCCGCATTGCCTGTGCATTTGGGGCGACCAAGCGAACCACGGTCTGAAAAGCGCGTTCAAGATTATGCACGAGATTACGGGTAGCAAAGAGGACTTCATCGTGGCCGACCACACGCAGATGAACACCCCGATCGGTGGCGGTCCGGTCCATGTCCCGGAGTACCACCACTTGACCGTTTGGACCTCTTTCGGTAAAGACCCGGATGATGACGAAACGGATCATCTGAATATTGTTGAGGTTATCAAAGCCGTCGACCTGGAAACTTACCGTTAATGATAATGA"),
			Plasmid: true,
		},
	},
}

func TestJoinXnumberofparts(t *testing.T) {
	for _, assembly := range assemblytests {
		fragments, plasmidProducts,err := FindAllAssemblyProducts(assembly.vector, assembly.parts, assembly.enzyme)
		if present, _ := search.InSequences(plasmidProducts, assembly.desiredassemblyproduct);!present{
			t.Error(
				"For", assembly, "\n",
				"expected", assembly.desiredassemblyproduct, "\n",
				"got plasmids: ", plasmidProducts, "\n",
				"got fragments: ", fragments, "\n",
			)
		}
		if err != nil {
			t.Error(err.Error())
		}
	}
}

type positionTest struct {
	sequence          wtype.DNASequence
	enzyme            wtype.TypeIIs
	positionPair 	   sequences.PositionPair
	cutPosition 		int
	UpStreamOverhang string
	DownStreamOverhang string
}

var positionTests = []positionTest{
	positionTest{
		sequence:          wtype.DNASequence{Nm:"Forward SapI test AAA",Seq:"GCTCTTCTGGTAAA"},
		enzyme:            SapI,
		positionPair: sequences.PositionPair{
			StartPosition: 1,
			EndPosition: 7,
			Reverse: false,
		},
		cutPosition: 8,
		UpStreamOverhang: "ACC",
		DownStreamOverhang: "GGT",
	},
	positionTest{
		sequence:          wtype.DNASequence{Nm:"Forward SapI test Plasmid AAA",Seq:"GGTAAAGCTCTTCT",Plasmid:true},
		enzyme:            SapI,
		positionPair: sequences.PositionPair{
			StartPosition: 7,
			EndPosition: 13,
			Reverse: false,
		},
		cutPosition: 14, 
		UpStreamOverhang: "ACC",
		DownStreamOverhang: "GGT",
	},
	positionTest{
		sequence:          wtype.DNASequence{Nm:"Forward SapI test",Seq:"GCTCTTCTGGT",Plasmid:true},
		enzyme:            SapI,
		positionPair: sequences.PositionPair{
			StartPosition: 1,
			EndPosition: 7,
			Reverse: false,
		},
		cutPosition: 8,
		UpStreamOverhang: "ACC",
		DownStreamOverhang: "GGT",
	},
	positionTest{
		sequence:          wtype.DNASequence{Nm:"Rev SapI test",Seq:"AAAGGTAGAAGAGC"},
		enzyme:            SapI,
		positionPair: sequences.PositionPair{
			StartPosition: 14,
			EndPosition: 8,
			Reverse: true,
		},
		cutPosition: 3,
		UpStreamOverhang: "ACC",
		DownStreamOverhang: "GGT",
	},
	positionTest{
		sequence:          wtype.DNASequence{Nm:"Rev SapI test",Seq:"GGTAGAAGAGC"},
		enzyme:            SapI,
		positionPair: sequences.PositionPair{
			StartPosition: 11,
			EndPosition: 5,
			Reverse: true,
		},
		cutPosition: 0,
		UpStreamOverhang: "ACC",
		DownStreamOverhang: "GGT",
	},
}

func TestCorrectPosition(t *testing.T) {
	for _, test := range positionTests {
		cutPosition := correctTypeIIsCutPosition(test.enzyme, test.positionPair)
		if cutPosition != test.cutPosition{
			t.Error(
				"For", test.enzyme.Name(), test.positionPair, "\n",
				"expected start", test.cutPosition, "\n",
				"got", cutPosition, "\n",
			)
		}
	}
}


func TestMakeOverhangs(t *testing.T) {
	for _, test := range positionTests {
		
		allPositionsFound := sequences.FindAll(&test.sequence,&wtype.DNASequence{Seq: test.enzyme.RecognitionSequence})
		
		upstream, downstream, err := makeOverhangs(test.enzyme, test.positionPair,test.sequence)
		if test.UpStreamOverhang != upstream.Sequence() || test.DownStreamOverhang != downstream.Sequence() {
			t.Error(
				"For", test.enzyme.Name(), test.positionPair, test.sequence.Name(), "\n",
				"expected overhangs", test.UpStreamOverhang, test.DownStreamOverhang, "\n",
				"got", upstream, downstream, "\n",
				"All positions found: ", allPositionsFound, "\n",
			)
		}
		
		if err != nil {
			t.Error(
				"For", test.enzyme.Name(), test.positionPair, test.sequence.Name(), "\n",
				"expected overhangs", test.UpStreamOverhang, test.DownStreamOverhang, "\n",
				"got error: ", err.Error(), "\n",
				"All positions found: ", allPositionsFound, "\n",
			)
		}
	}
}

type fragmentTest struct {
	enzyme wtype.TypeIIs
	upstream sequences.PositionPair
	downstream sequences.PositionPair
	sequence wtype.DNASequence
	fragment wtype.DNASequence
}

var fragmentTests =[]fragmentTest{
	fragmentTest{
		sequence:          wtype.DNASequence{Nm:"Forward SapI fragment test",Seq:"GCTCTTCTGGTGCTCTTCTGGT",Plasmid:true},
		enzyme:            SapI,
		upstream: sequences.PositionPair{
			StartPosition: 1,
			EndPosition: 7,
			Reverse: false,
		},
		downstream: sequences.PositionPair{
			StartPosition: 12,
			EndPosition: 18,
			Reverse: false,
		},
		fragment:          wtype.DNASequence{Nm:"fragment",Seq:"GGTGCTCTTCT",Plasmid:false},
	},
	fragmentTest{
		sequence:          wtype.DNASequence{Nm:"Forward SapI fragment test downstream null ",Seq:"GCTCTTCTGGTGCTCTTCTGGT",Plasmid:true},
		enzyme:            SapI,
		upstream: sequences.PositionPair{
			StartPosition: 1,
			EndPosition: 7,
			Reverse: false,
		},
		downstream: sequences.PositionPair{
			
		},
		fragment:          wtype.DNASequence{Nm:"fragment",Seq:"GGTGCTCTTCTGGT",Plasmid:false},
	},
	fragmentTest{
		sequence:          wtype.DNASequence{Nm:"Forward SapI fragment upstream null ",Seq:"GCTCTTCTGGTGCTCTTCTGGT",Plasmid:true},
		enzyme:            SapI,
		upstream: sequences.PositionPair{
			
		},
		downstream: sequences.PositionPair{
			StartPosition: 12,
			EndPosition: 18,
			Reverse: false,
		},
		fragment:          wtype.DNASequence{Nm:"fragment",Seq:"GCTCTTCTGGTGCTCTTCT",Plasmid:false},
	},
	fragmentTest{
		sequence:          wtype.DNASequence{Nm:"Reverse SapI fragment test",Seq:"GGTAGAAGAGCGGTAGAAGAGC",Plasmid:true},
		enzyme:            SapI,
		upstream: sequences.PositionPair{
			StartPosition: 11,
			EndPosition: 5,
			Reverse: true,
		},
		downstream: sequences.PositionPair{
			StartPosition: 22,
			EndPosition: 16,
			Reverse: true,
		},
		fragment:          wtype.DNASequence{Nm:"fragment",Seq:"GGTAGAAGAGC",Plasmid:false},
	},
	fragmentTest{
		sequence:          wtype.DNASequence{Nm:"Reverse SapI fragment test upstream null",Seq:"GGTAGAAGAGCGGTAGAAGAGC",Plasmid:true},
		enzyme:            SapI,
		upstream: sequences.PositionPair{
			
		},
		downstream: sequences.PositionPair{
			StartPosition: 22,
			EndPosition: 16,
			Reverse: true,
		},
		fragment:          wtype.DNASequence{Nm:"fragment",Seq:"GGTAGAAGAGC",Plasmid:false},
	},
	fragmentTest{
		sequence:          wtype.DNASequence{Nm:"Reverse SapI fragment test downstream null",Seq:"GGTAGAAGAGCGGTAGAAGAGC",Plasmid:true},
		enzyme:            SapI,
		upstream: sequences.PositionPair{
			StartPosition: 22,
			EndPosition: 16,
			Reverse: true,
		},
		downstream: sequences.PositionPair{
			
		},
		fragment:          wtype.DNASequence{Nm:"fragment",Seq:"GGTAGAAGAGC",Plasmid:false},
	},
	fragmentTest{
		sequence:          wtype.DNASequence{Nm:"Forward SapI fragment test wrap around",Seq:"TGGTGCTCTTCTGGTAAAAAAGCTCTTC",Plasmid:true},
		enzyme:            SapI,
		upstream: sequences.PositionPair{
			StartPosition: 5,
			EndPosition: 11,
			Reverse: false,
		},
		downstream: sequences.PositionPair{
			StartPosition: 22,
			EndPosition: 28,
			Reverse: false,
		},
		fragment:          wtype.DNASequence{Nm:"fragment",Seq:"GGTAAAAAAGCTCTTCT",Plasmid:false},
	},
	fragmentTest{
		sequence:          wtype.DNASequence{Nm:"Reverse SapI fragment test wrap around",Seq:"AGAAGAGCGGTAGAAGAGCGGT",Plasmid:true},
		enzyme:            SapI,
		upstream: sequences.PositionPair{
			StartPosition: 8,
			EndPosition: 2,
			Reverse: true,
		},
		downstream: sequences.PositionPair{
			StartPosition: 19,
			EndPosition: 13,
			Reverse: true,
		},
		fragment:          wtype.DNASequence{Nm:"fragment",Seq:"GGTAGAAGAGC",Plasmid:false},
	},
}

func TestMakeFragment(t *testing.T) {
	for _, test := range fragmentTests {
		allPositionsFound := sequences.FindAll(&test.sequence,&wtype.DNASequence{Seq: test.enzyme.RecognitionSequence})

		fragment, err := makeFragment(test.enzyme, test.upstream,test.downstream, test.sequence)
		if fragment.Sequence()!=  test.fragment.Sequence(){
			t.Error(
				"For", fmt.Sprintf("%s",test.sequence.Name()), "\n",
				"expected fragment", test.fragment.Sequence(), "\n",
				"got", fragment.Sequence(), "\n",
				"All positions found: ", allPositionsFound, "\n",
			)
		}
		
		if err != nil {
			t.Error(
				"For", fmt.Sprintf("%s",test.sequence.Name()), "\n",
				"got error: ", err.Error(), "\n",
				"All positions found: ", allPositionsFound, "\n",
			)
		}
	}
}

