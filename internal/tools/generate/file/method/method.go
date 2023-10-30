package method

import (
	"bytes"
	"path/filepath"
	"text/template"

	"github.com/nomad-software/goat/internal/tools/generate/cli"
	"github.com/nomad-software/goat/internal/tools/generate/env"
	"github.com/nomad-software/goat/internal/tools/generate/file/common"
	"github.com/nomad-software/goat/log"
)

const (
	methTmpl = `{{.Comments}}
func (ele {{.Receiver}}) {{.Name}}({{.Parameters}}){{.ReturnTypes}} {
	{{if .ReturnTypes}}return {{end}}{{.PkgName}}.{{.Name}}(ele.GetID(), {{.Arguments}})
}`
)

type Output struct {
	PkgName    string
	ImportPath string
	Methods    []string
}

func Create(e *env.Env, opt *cli.Options, funcs []common.Function) *Output {
	output := &Output{
		PkgName:    e.GoPackage,
		ImportPath: filepath.Join(env.RepoName, env.CommonDir, opt.PkgName),
		Methods:    make([]string, 0),
	}

	tmpl, err := template.New("method").Parse(methTmpl)
	if err != nil {
		log.Error(err)
		return nil
	}

	for _, fn := range funcs {
		buf := new(bytes.Buffer)
		tmpl.Execute(buf, fn)

		output.Methods = append(output.Methods, buf.String())
	}

	return output
}
