package composer

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/antha-lang/antha/workflow"
)

type renderer struct {
	varCount uint64
	varMemo  map[workflow.ElementInstanceName]string
}

func (r *renderer) varName(name workflow.ElementInstanceName) string {
	if res, found := r.varMemo[name]; found {
		return res
	}

	res := make([]rune, 0, len(name))
	ensureUpper := true
	for _, r := range []rune(name) {
		switch {
		case 'a' <= r && r <= 'z' && ensureUpper:
			res = append(res, unicode.ToUpper(r))
			ensureUpper = false
		case 'a' <= r && r <= 'z':
			res = append(res, r)
		case 'A' <= r && r <= 'Z':
			res = append(res, r)
			ensureUpper = false
		case strings.ContainsRune(" -_", r):
			ensureUpper = true
		case unicode.IsDigit(r):
			res = append(res, r)
			ensureUpper = true
		}
	}
	resStr := fmt.Sprintf("elem%s%d", string(res), r.varCount)
	r.varCount++
	r.varMemo[name] = resStr
	return resStr
}

type mainRenderer struct {
	renderer
	mainComposer *mainComposer
}

func renderMain(w io.Writer, mc *mainComposer) error {
	mr := &mainRenderer{
		renderer: renderer{
			varMemo: make(map[workflow.ElementInstanceName]string),
		},
		mainComposer: mc,
	}
	funcs := template.FuncMap{
		"elementTypes": mr.elementTypes,
		"varName":      mr.varName,
		"token":        mr.token,
		"id":           func() string { return "" },
		"name":         func() string { return "" },
		"inDir":        func() string { return "" },
		"outDir":       func() string { return "" },
	}
	if t, err := template.New("maintpl").Funcs(funcs).Parse(maintpl); err != nil {
		return err
	} else {
		return t.ExecuteTemplate(w, "main", mr.mainComposer.Workflow)
	}
}

func (mr *mainRenderer) elementTypes() map[workflow.ElementTypeName]*TranspilableElementType {
	return mr.mainComposer.elementTypes
}

func (mr *mainRenderer) token(elem workflow.ElementInstanceName, param workflow.ElementParameterName) (string, error) {
	return mr.mainComposer.token(mr.mainComposer.Workflow, elem, param)
}

type testRenderer struct {
	renderer
	testWorkflow *testWorkflow
}

func renderTest(w io.Writer, twf *testWorkflow) error {
	tr := &testRenderer{
		renderer: renderer{
			varMemo: make(map[workflow.ElementInstanceName]string),
		},
		testWorkflow: twf,
	}
	idStr := fmt.Sprintf("%d", tr.testWorkflow.index)
	name := strings.Title(tr.testWorkflow.workflow.Meta.NameAsGoIdentifier())
	funcs := template.FuncMap{
		"elementTypes": tr.elementTypes,
		"varName":      tr.varName,
		"token":        tr.token,
		"id":           func() string { return idStr },
		"name":         func() string { return name },
		"inDir":        func() string { return tr.testWorkflow.inDir },
		"outDir": func() string {
			return filepath.Join(tr.testWorkflow.OutDir, "simulations", fmt.Sprintf("%s-%s", idStr, name))
		},
	}
	if t, err := template.New("maintpl").Funcs(funcs).Parse(maintpl); err != nil {
		return err
	} else {
		return t.ExecuteTemplate(w, "test", tr.testWorkflow.workflow)
	}
}

func (tr *testRenderer) elementTypes() map[workflow.ElementTypeName]*TranspilableElementType {
	return tr.testWorkflow.elementTypes
}

func (tr *testRenderer) token(elem workflow.ElementInstanceName, param workflow.ElementParameterName) (string, error) {
	return tr.testWorkflow.token(tr.testWorkflow.workflow, elem, param)
}

func renderGoGenerate(w io.Writer) error {
	const content = `//go:generate go-bindata data/...
package main
`
	_, err := w.Write([]byte(content))
	return err
}

var maintpl = `
{{define "header"}}// Code generated by antha composer.
package main
{{end}}

{{define "main-imports"}}import (
	"bytes"
	"os"
	"io/ioutil"

	"github.com/antha-lang/antha/laboratory"
	"github.com/ugorji/go/codec"

{{range elementTypes}}{{if .IsAnthaElement}}	{{printf "%q" .ImportPath}}
{{end}}{{end}})
{{end}}

{{define "test-imports"}}import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/antha-lang/antha/laboratory"
	"github.com/antha-lang/antha/laboratory/testlab"
	"github.com/ugorji/go/codec"

{{range elementTypes}}{{if .IsAnthaElement}}	{{printf "%q" .ImportPath}}
{{end}}{{end}})
{{end}}


{{define "main-main"}}func main() {
	labBuild := laboratory.NewLaboratoryBuilder(ioutil.NopCloser(bytes.NewBuffer(MustAsset("data/workflow.json"))))
	runWorkflow(labBuild)
	if err := labBuild.Decommission(); err != nil {
		os.Exit(1)
	}
}
{{end}}

{{define "test-test"}}func TestWorkflow_{{id}}_{{name}}(t *testing.T) {
	t.Parallel()
	labBuild := testlab.NewTestLabBuilder(t, {{printf "%q" inDir}}, {{printf "%q" outDir}}, ioutil.NopCloser(bytes.NewBuffer(MustAsset("data/workflow{{id}}.json"))))
	runWorkflow{{id}}(labBuild)
	labBuild.Compare()
	if err := labBuild.Decommission(); err != nil {
		t.Fatal(err)
	}
}
{{end}}

{{define "run-workflow"}}func runWorkflow{{id}}(labBuild *laboratory.LaboratoryBuilder) {
	jh := &codec.JsonHandle{}
	if err := labBuild.RegisterJsonExtensions(jh); err != nil {
		return
	}

	// Register line maps for the elements we're using
{{range elementTypes}}{{if .IsAnthaElement}}	{{.Name}}.RegisterLineMap(labBuild)
{{end}}{{end}}
	// Create the elements
{{range $name, $inst := .Elements.Instances}}	{{if $inst.IsUsed}}{{varName $name}} := {{end}}{{$inst.ElementTypeName}}.New(labBuild, {{printf "%q" $name}})
{{end}}
	// Add wiring
{{range .Elements.InstancesConnections}}	labBuild.AddConnection({{varName .Source.ElementInstance}}, {{varName .Target.ElementInstance}}, func() { {{varName .Target.ElementInstance}}.{{token .Target.ElementInstance .Target.ParameterName}}.{{.Target.ParameterName}} = {{varName .Source.ElementInstance}}.{{token .Source.ElementInstance .Source.ParameterName}}.{{.Source.ParameterName}} })
{{end}}
	// Set parameters
{{range $name, $inst := .Elements.Instances}}{{range $param, $value := $inst.Parameters}}	if err := codec.NewDecoderBytes([]byte({{printf "%q" $value}}), jh).Decode(&{{varName $name}}.{{token $name $param}}.{{$param}}); err != nil {
		// This will catch things like invalid json where we error before we hit any of our own code.
		labBuild.RecordError(err, true)
		return
	}
{{end}}{{end}}
	// Run!
	labBuild.RunElements()
	labBuild.Compile()
	labBuild.Export()
}
{{end}}

{{define "main"}}{{template "header" .}}
{{template "main-imports" .}}
{{template "main-main" .}}
{{template "run-workflow" .}}{{end}}

{{define "test"}}{{template "header" .}}
{{template "test-imports" .}}
{{template "test-test" .}}
{{template "run-workflow" .}}{{end}}
`
