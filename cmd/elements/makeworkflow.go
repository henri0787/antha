package main

import (
	"flag"
	"regexp"

	"github.com/antha-lang/antha/logger"
	"github.com/antha-lang/antha/workflow"
)

func makeWorkflow(l *logger.Logger, args []string) error {
	flagSet := flag.NewFlagSet(flag.CommandLine.Name()+" makeWorkflow", flag.ContinueOnError)
	flagSet.Usage = workflow.NewFlagUsage(flagSet, "Modify workflow adding selected elements")

	var regexStr, inDir, toFile string
	flagSet.StringVar(&regexStr, "regex", "", "Regular expression to match against element type path (optional)")
	flagSet.StringVar(&inDir, "indir", "", "Directory from which to read files (optional)")
	flagSet.StringVar(&toFile, "to", "", "File to write to (default: will write to stdout)")

	if err := flagSet.Parse(args); err != nil {
		return err
	} else if wfPaths, err := workflow.GatherPaths(flagSet, inDir); err != nil {
		return err
	} else if rs, err := workflow.ReadersFromPaths(wfPaths); err != nil {
		return err
	} else if acc, err := workflow.WorkflowFromReaders(rs...); err != nil {
		return err
	} else if regex, err := regexp.Compile(regexStr); err != nil {
		return err
	} else {
		err := findElements(l, wfPaths, func(r *workflow.Repository, et *workflow.ElementType) error {
			if !regex.MatchString(string(et.ElementPath)) {
				return nil
			}
			etCopy := *et
			wf := &workflow.Workflow{
				SchemaVersion: workflow.CurrentSchemaVersion,
				Elements: workflow.Elements{
					Types: workflow.ElementTypes{&etCopy},
					Instances: workflow.ElementInstances{
						workflow.ElementInstanceName(etCopy.Name()): &workflow.ElementInstance{
							ElementTypeName: etCopy.Name(),
						},
					},
				},
			}
			return acc.Merge(wf)
		})
		if err != nil {
			return err
		}
		return acc.WriteToFile(toFile, true)
	}
}
