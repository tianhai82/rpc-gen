package rpc_gen

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

type GenConfig struct {
	Folder   string
	BasePath string
	Services []Service
}

type Service struct {
	Cache        bool
	Path         string
	FunctionName string
	Input        *Param
	Output       *Param
}

type Param struct {
	Class     interface{}
	ClassName string
	IsArray   bool
}

func CreateTsServiceClients(configs []GenConfig) {
	for _, genConfig := range configs {
		err := createTsServiceClient(genConfig)
		if err != nil {
			fmt.Printf("Error creating ts client for %s: %s", genConfig.Folder, err.Error())
		}
	}
}

func createTsServiceClient(genConfig GenConfig) error {
	err := genModel(&genConfig)
	if err != nil {
		return err
	}
	importString, err := genImport(genConfig)
	if err != nil {
		return err
	}
	cacheString := "\nconst cache = new Map<string, any>();"
	bodyString, err := genBody(genConfig)
	if err != nil {
		return err
	}
	f, err := os.Create(fmt.Sprintf("%s/rpcs.ts", genConfig.Folder))
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString("/* Do not change, this code is generated from Golang structs */\n\n")
	f.WriteString("/* eslint-disable import/prefer-default-export, max-len */\n")
	f.WriteString(importString)
	f.WriteString(cacheString)
	f.WriteString(bodyString)
	return nil
}

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
