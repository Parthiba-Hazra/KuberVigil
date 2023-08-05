/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Parthiba-Hazra/kubervigil/cmd"
	_ "github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd"
	_ "github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd/apis"
	_ "github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd/apiversion"
	_ "github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd/configmap"
	_ "github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd/currentAPIversion"
	_ "github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd/daemonsets"
	_ "github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd/deployments"
	_ "github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd/ingress"
	_ "github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd/pods"
	_ "github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd/services"
	_ "github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd/statefulsets"
)

func main() {
	cmd.Execute()
}
