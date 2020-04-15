package main

import (
	"html/template"
	"log"
	"os"
	"sort"

	"github.com/Masterminds/sprig"
)

type function struct {
	Name        string
	Description string
	Origin      string
	URL         string
}

var t = template.Must(template.New("Name").ParseFiles("templates/index.tmpl"))

func makeFunctions(slice []string, origin string) []function {
	var funcs []function
	for _, f := range slice {
		funcs = append(funcs, function{Name: f, Origin: origin})
	}
	return funcs
}

func goBuiltinFuncs() []function {
	// Go doesn't export the function names anywhere so we have to list them ourselves
	funcs := []string{"and", "call", "html", "index", "slice", "js", "len", "not", "or", "print", "printf", "println", "urlquery", "eq", "ge", "gt", "le", "lt", "ne"}
	return makeFunctions(funcs, "Go BuiltIn")
}

func sprigFuncs() []function {
	var funcs []string
	for k := range sprig.FuncMap() {
		funcs = append(funcs, k)
	}
	return makeFunctions(funcs, "Sprig")
}

func helmFuncs() []function {

	funcs := []string{"toToml", "toYaml", "fromYaml", "fromYamlArray", "toJson", "fromJson", "fromJsonArray"}
	return makeFunctions(funcs, "Helm")
}

func main() {
	allFuncs := []function{}
	allFuncs = append(allFuncs, goBuiltinFuncs()...)
	allFuncs = append(allFuncs, sprigFuncs()...)
	allFuncs = append(allFuncs, helmFuncs()...)
	sort.Slice(allFuncs, func(i, j int) bool { return allFuncs[i].Name < allFuncs[j].Name })

	err := t.ExecuteTemplate(os.Stdout, "index.tmpl", allFuncs)
	if err != nil {
		log.Fatalf("Unable to execute template: %v", err)
	}
}
