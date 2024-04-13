package main

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsGo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

func main() {
	pgs.Init(pgs.DebugEnv("DEBUG_PGR")).
		RegisterModule(Redactor()).
		RegisterPostProcessor(pgsGo.GoFmt()).
		Render()
}
