package rpc_gen

import (
	"bytes"
	"reflect"
	"text/template"
)

func genImport(genConfig GenConfig, needCaching bool) (string, error) {
	clsMap := map[string]bool{}
	for _, service := range genConfig.Services {
		if service.Input != nil {
			cls := service.Input.Class
			switch cls.(type) {
			case int:
			case int64:
			case float64:
			case float32:
			case bool:
			case string:
				continue
			default:
				clsMap[reflect.TypeOf(service.Input.Class).Name()] = true
			}
		}
	}
	for _, service := range genConfig.Services {
		if service.Output != nil {
			cls := service.Output.Class
			switch cls.(type) {
			case int:
			case int64:
			case float64:
			case float32:
			case bool:
			case string:
				continue
			default:
				clsMap[reflect.TypeOf(service.Output.Class).Name()] = true
			}
		}
	}
	classes := make([]string, 0, len(clsMap))
	for key := range clsMap {
		classes = append(classes, key)
	}
	it := template.New("importTemplate")
	it, err := it.Parse(importTemplate)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBufferString("")
	err = it.Execute(buf, map[string]interface{}{
		"NeedCaching": needCaching,
		"Classes":     classes,
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
