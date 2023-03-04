package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var flags flag.FlagSet

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f)
		}
		return nil
	})
}

// generateFile generates a _ascii.pb.go file containing gRPC service definitions.
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	var methods []string
	for _, service := range file.Services {
		for _, method := range service.Methods {
			methods = append(methods, method.GoName)
		}
	}
	filename := "SERIVCE.md"

	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("# Service Catalog")
	g.P()
	g.P("## " + file.GeneratedFilenamePrefix)
	for _, method := range methods {
		g.P("- " + method)
	}
	g.P()

	return g
}
