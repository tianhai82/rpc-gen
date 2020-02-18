package rpc_gen

import (
	"fmt"
	"reflect"

	"github.com/tianhai82/typescriptify-golang-structs/typescriptify"
)

func genModel(config *GenConfig) error {
	clsMap := map[interface{}]bool{}
	for i, service := range config.Services {
		if service.Input != nil {
			cls := service.Input.Class
			switch cls.(type) {
			case int:
				config.Services[i].Input.ClassName = "number"
			case int64:
				config.Services[i].Input.ClassName = "number"
			case float64:
				config.Services[i].Input.ClassName = "number"
			case bool:
				config.Services[i].Input.ClassName = "boolean"
			case string:
				config.Services[i].Input.ClassName = "string"
			default:
				clsMap[service.Input.Class] = true
				config.Services[i].Input.ClassName = reflect.TypeOf(service.Input.Class).Name()
			}
		}
	}
	for i, service := range config.Services {
		if service.Output != nil {
			cls := service.Output.Class
			switch cls.(type) {
			case int:
				config.Services[i].Output.ClassName = "number"
			case int64:
				config.Services[i].Output.ClassName = "number"
			case float64:
				config.Services[i].Output.ClassName = "number"
			case bool:
				config.Services[i].Output.ClassName = "boolean"
			case string:
				config.Services[i].Output.ClassName = "string"
			default:
				clsMap[service.Output.Class] = true
				config.Services[i].Output.ClassName = reflect.TypeOf(service.Output.Class).Name()
			}
		}
	}

	scriptify := typescriptify.New()
	for key := range clsMap {
		scriptify.Add(key)
	}
	return scriptify.ConvertToFile(fmt.Sprintf("%s/models.ts", config.Folder))
}
