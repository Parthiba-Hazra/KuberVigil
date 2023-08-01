/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Parthiba-Hazra/kubervigil/cmd"
	_ "github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd"
	_ "github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd/apiversion"
)

func main() {
	cmd.Execute()
}
