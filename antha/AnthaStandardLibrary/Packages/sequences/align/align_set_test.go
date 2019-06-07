package align

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Synthace/antha/antha/AnthaStandardLibrary/Packages/sequences/parse/fasta"
	"github.com/Synthace/antha/antha/anthalib/wtype"
	"github.com/Synthace/antha/laboratory"
	"github.com/Synthace/antha/laboratory/testlab"
)

func TestDNASet(t *testing.T) {

	const fastaText = `
>CCE57619 cdna plasmid:HUSEC2011CHR1:pHUSEC2011-2:1219:1338:-1 gene:HUS2011_pII0002 gene_biotype:protein_coding transcript_biotype:protein_coding description:hypothetical protein
ATGTTTTATGAAGGGAGCAATGCCTCAGCATCAGGTTACGGGGTGACTCACGTAAGGGAC
AGGCAGATGGCAGCTCAGCCACAGGCAGCACTGCAGGAAACTGAATATAAACTGCAGTGA
>CCE57620 cdna plasmid:HUSEC2011CHR1:pHUSEC2011-2:1422:2162:-1 gene:HUS2011_pII0003 gene_biotype:protein_coding transcript_biotype:protein_coding description:site-specific recombinase
ATGAACAATGTCATTCCCCTGCAGAATTCACCAGAACGCGTCTCCCTGTTACCCATTGCG
CCGGGGGTGGATTTTGCAACAGCGCTCTCCCTGAGAAGAATGGCCACTTCCACGGGGGCC
ACACCGGCCTACCTGCTGGCCCCGGAAGTGAGTGCCCTTCTTTTCTATATGCCGGATCAG
CGTCACCATATGCTGTTCGCCACCCTCTGGAATACCGGAATGCGTATTGGCGAAGCCCGG
ATGCTGACACCGGAATCATTTGACCTGGATGGAGTAAGACCGTTTGTGCGGATCCAGTCC
GAAAAAGTGCGTGCGCGACGCGGACGCCCGCCAAAAGATGAAGTGCGCCTGGTTCCGCTG
ACAGATATAAGCTATGTCAGGCAGATGGAAAGCTGGATGATCACCACCCGGCCCCGTCGT
CGTGAACCATTATGGGCCGTGACCGACGAAACCATGCGCAACTGGCTGAAGCAGGCTGTC
AGACGGGCCGAAGCTGACGGGGTACACTTTTCGATTCCGGTAACACCACACACTTTCCGG
CACAGCTATATCATGCACATGCTCTATCACCGCCAGCCCCGGAAAGTCATCCAGGCACTG
GCTGGTCACAGGGATCCACGTTCGATGGAGGTCTATACACGGGTGTTTGCGCTGGATATG
GCTGCCACGCTGGCAGTGCCTTTCACAGGTGACGGACGGGATGCTGCAGAGATCCTGCGT
ACACTGCCTCCCCTGAAGTAA
>CCE57621 cdna plasmid:HUSEC2011CHR1:pHUSEC2011-2:2531:2692:-1 gene:HUS2011_pII0004 gene_biotype:protein_coding transcript_biotype:protein_coding description:plasmid stabilisation system family protein
ATGAGTAATCATAATATCGGTACTCCCCGTCCTGAACTGGGGGAATACACATTCGCACTA
CCCGTTGAACGGCATATGGTTTATTTTCTGCAAACTGATACTGAAATTGTTATTATTCGT
ATATTAAGTCAGCATCAGGATGCCAGCCGTCATTTCAACTGA
>CCE57622 cdna plasmid:HUSEC2011CHR1:pHUSEC2011-2:2894:3076:-1 gene:HUS2011_pII0005 gene_biotype:protein_coding transcript_biotype:protein_coding description:putative transcriptional regulator
ATGGCCAGAACAATGACAGTTGCTCTCGGAGATGAACTCTGGGAGTACATAGAATCTCTC
ATAGAATCAGGTGATTATCGTACACAGAGTGAGGTAATCCGCGAGTCACTTCGTCTCCTT
CGAGAGAAACAGGCAGAGTCACGTCTCCAGGTGTCCTGCGGGATTTACTGGCAGAAGGCT
TGA
>CCE57623 cdna plasmid:HUSEC2011CHR1:pHUSEC2011-2:3145:3420:-1 gene:HUS2011_pII0006 gene_biotype:protein_coding transcript_biotype:protein_coding description:plasmid stabilisation system family protein
ATGGAACTGAAGTGGACCAGTAAGGCGCTTTCTGATTTGTCGCGGTTATATGATTTTCTG
GTGCTGGCCAGTAAACCTGCTGCCGCCAGAACGGTACAGTCCCTGACACAGGCACCGGTC
ATTCTGTTAACTCATCCACGTATGGGAGAACAGTTGTTTCAGTTTGAACCCAGGGAGGTC
AGACGGATTTTTGCTGGCGAGTACGAAATCCGTTACGAAATTAATGGCCAGACTATTTAT
GTATTGCGTCTGTGGCACACACGAGAAAACAGGTAG
>CCE57624 cdna plasmid:HUSEC2011CHR1:pHUSEC2011-2:3420:3698:-1 gene:HUS2011_pII0007 gene_biotype:protein_coding transcript_biotype:protein_coding description:ribbon-helix-helix protein, copG family
ATGAAAAACAATGCCGCACAAGCAACAAAAGTAATTACCGCGCATGTGCCATTACCTATG
GCTGATAAAGTCGACCAGATGGCCGCCAGACTGGAACGTTCCCGGGGCTGGGTTATCAAA
CAGGCGCTTTCTGCATGGCTTGCCCAGGAGGAGGAGCGTAATCGCCTGACGCTGGAAGCC
CTGGACGATGTGACATCCGGACAGGTTATCGACCATCAGGCTGTACAGGCCTGGTCGGAC
AGCCTCAGTACTGACAATCCGTTACCGGTGCCACGCTGA
`

	testlab.WithTestLab(t, "", &testlab.TestElementCallbacks{
		Steps: func(lab *laboratory.Laboratory) error {
			fastaFile, err := lab.FileManager.WriteAll([]byte(fastaText), "ecoli-cdna")
			if err != nil {
				return err
			}
			database, err := fasta.FastaToDNASequences(lab, fastaFile)
			if err != nil {
				return err
			}

			primer := wtype.DNASequence{Seq: "ATGGAACTGAAGTGG"}

			algorithm, found := Algorithms["SWAffine"]
			if !found {
				return errors.New("algorithm not found")
			}

			testLimit := 4

			results, err := DNASet(primer, database, algorithm, testLimit)
			if err != nil {
				return err
			}

			if len(results) != testLimit {
				return fmt.Errorf("Number of results: got %d, want %d\n", len(results), testLimit)
			}

			if results[0].Template.Name() != "CCE57623" {
				return fmt.Errorf("Unexpected top scoring result: got %s, want %s\n", results[0].Template.Name(), "CCE57623")
			}

			if results[0].Score() != 15 {
				return fmt.Errorf("Unexpected top score: got %d, want %d\n", results[0].Score(), 15)
			}
			return nil
		},
	})
}

func TestDNASetBenchmark(t *testing.T) {

	// All ecoli CDNA
	// URI ftp://ftp.ensemblgenomes.org/pub/bacteria/release-42/fasta/bacteria_91_collection/escherichia_coli/cdna/Escherichia_coli.HUSEC2011CHR1.cdna.all.fa.gz
	// ... a few seconds

	testlab.WithTestLab(t, "testdata", &testlab.TestElementCallbacks{
		Steps: func(lab *laboratory.Laboratory) error {

			fastaFile := wtype.NewFile("Escherichia_coli.HUSEC2011CHR1.cdna.all.fa").AsInput()
			database, err := fasta.FastaToDNASequences(lab, fastaFile)
			if err != nil {
				return err
			}
			lab.Logger.Log("readSequences", len(database))

			primer := wtype.DNASequence{Seq: "ATGGAACTGAAGTGG"}

			algorithm, found := Algorithms["SWAffine"]
			if !found {
				return errors.New("algorithm not found")
			}

			testLimit := 4

			results, err := DNASet(primer, database, algorithm, testLimit)
			if err != nil {
				return err
			}

			for _, result := range results {
				lab.Logger.Log("found", result.Template.Name(), "score", result.Score())
			}

			if len(results) != testLimit {
				return fmt.Errorf("Number of results: got %d, want %d\n", len(results), testLimit)
			}
			return nil
		},
	})

}
