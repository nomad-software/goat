package main

import (
	"os"
	"path/filepath"

	"github.com/nomad-software/goat/internal/tools/generate/cli"
	"github.com/nomad-software/goat/internal/tools/generate/env"
	"github.com/nomad-software/goat/internal/tools/generate/file/common"
	"github.com/nomad-software/goat/internal/tools/generate/file/method"
	"github.com/nomad-software/goat/internal/tools/generate/file/output"
	"github.com/nomad-software/goat/log"
)

func main() {
	env := env.Parse()
	opt := cli.Parse()

	if err := opt.Valid(); err != nil {
		log.Panic(err, "options are not valid")
	}

	// fmt.Printf("args: %v\n", os.Args)
	// fmt.Printf("env: %#v\n", env)
	// fmt.Printf("opt: %#v\n", opt)

	path := filepath.Join(env.ProjectDir, env.CommonDir, opt.PkgName, opt.PkgName+".go")
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Panic(err, "cannot read file")
	}

	funcs := common.Parse(opt, bytes)
	if len(funcs) == 0 {
		log.Info("no functions to parse")
		return
	}

	out := method.Create(env, opt, funcs)
	bytes = output.Create(out)

	path = filepath.Join(env.PkgDir, env.GoPackage+"_"+opt.PkgName+".go")

	// fmt.Printf("path: %v\n", path)
	// fmt.Printf("content: %v\n", string(bytes))

	err = os.WriteFile(path, bytes, 0666)
	if err != nil {
		log.Panic(err, "cannot read file")
	}
}
