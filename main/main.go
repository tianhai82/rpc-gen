// +build generate
package main

import (
	gen "github.com/tianhai82/rpc-gen"
)

type Huat struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

var config = gen.GenConfig{
	Folder:   "./web/src/apis/amount",
	BasePath: "/web/src/apis/amount",
	Services: []gen.Service{
		gen.Service{
			Cache:        true,
			Path:         "/cash",
			FunctionName: "cash",
			Input: &gen.Param{
				Class:     5.5,
				ClassName: "",
				IsArray:   true,
			},
			Output: &gen.Param{
				Class:     Huat{},
				ClassName: "",
				IsArray:   false,
			},
		},
	},
}
var config2 = gen.GenConfig{
	Folder:   "./web/src/apis/age",
	BasePath: "/web/src/apis/age",
	Services: []gen.Service{
		gen.Service{
			Cache:        true,
			Path:         "/year",
			FunctionName: "year",
			Input: &gen.Param{
				Class:     Huat{},
				ClassName: "",
				IsArray:   true,
			},
			Output: &gen.Param{
				Class:     "",
				ClassName: "",
				IsArray:   true,
			},
		},
	},
}

func main() {
	gen.CreateTsServiceClients([]gen.GenConfig{config, config2})
}
