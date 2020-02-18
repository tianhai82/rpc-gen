package rpc_gen

import (
	"bytes"
	"text/template"
)

func genBody(genConfig GenConfig) (string, error) {
	it := template.New("bodyTemplate")
	it, err := it.Parse(functionTemplate)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBufferString("")
	for _, service := range genConfig.Services {
		service.Path = genConfig.BasePath + service.Path
		err = it.Execute(buf, service)
		if err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}
