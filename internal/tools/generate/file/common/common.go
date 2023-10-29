package common

import (
	"regexp"
	"slices"
	"strings"

	"github.com/nomad-software/goat/internal/tools/generate/cli"
)

const (
	FuncPattern = `(?ms)^(//.*?)func (.*?)\((.*?)\)(.*?) {.*?}$`
)

var (
	FuncRegex = regexp.MustCompile(FuncPattern)
)

type Function struct {
	PkgName     string
	Comments    string
	Receiver    string
	Name        string
	Parameters  string
	ReturnTypes string
	Arguments   string
}

func Parse(opt *cli.Options, file []byte) []Function {
	result := make([]Function, 0)
	funcs := make([]string, 0)

	if opt.Functions != "" {
		funcs = strings.Split(opt.Functions, ",")
	}

	matches := FuncRegex.FindAllStringSubmatch(string(file), -1)

match:
	for _, match := range matches {
		fn := Function{
			PkgName:  opt.PkgName,
			Receiver: opt.Receiver,
		}
		for i, val := range match {
			if i == 1 {
				fn.Comments = strings.Trim(val, "\n")
			}
			if i == 2 {
				if len(funcs) > 0 && !slices.Contains(funcs, val) {
					continue match
				}
				fn.Name = val
			}
			if i == 3 {
				fn.Parameters = parseParameters(val)
				fn.Arguments = parseArguments(val)
			}
			if i == 4 {
				fn.ReturnTypes = val
			}
		}
		result = append(result, fn)
	}

	return result
}

func parseParameters(val string) string {
	result := strings.Split(val, ", ")
	return strings.Join(result[1:], ", ")
}

func parseArguments(val string) string {
	result := make([]string, 0)
	params := strings.Split(val, ", ")

	for _, val := range params[1:] {
		split := strings.Split(val, " ")
		result = append(result, split[0])
	}
	return strings.Join(result, ", ")
}
