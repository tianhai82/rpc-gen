package rpc_gen

import (
	"fmt"
	"os"
	"reflect"

	"github.com/tianhai82/typescriptify-golang-structs/typescriptify"
)

func genModel(config *GenConfig) error {
	clsMap := map[reflect.Type]bool{}
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
				clsMap[reflect.TypeOf(service.Input.Class)] = true
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
				clsMap[reflect.TypeOf(service.Output.Class)] = true
				config.Services[i].Output.ClassName = reflect.TypeOf(service.Output.Class).Name()
			}
		}
	}

	scriptify := typescriptify.New()
	scriptify.BackupDir = ""
	for key := range clsMap {
		scriptify.AddType(key)
	}
	err := os.MkdirAll(config.Folder, 0700)
	if err != nil {
		return err
	}
	return scriptify.ConvertToFile(fmt.Sprintf("%s/models.ts", config.Folder))
}
