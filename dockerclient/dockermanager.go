package dockerclient

import (
	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"fmt"
	"context"
	"nexusnode.de/nexusnode-daemon/util"
)

var Client *docker.Client

func ConnectToDocker() {
	client, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}
	Client = client
}

func PrintAllContainers() {
	containers, err := Client.ContainerList(context.Background(), types.ContainerListOptions{All:true})

	if err != nil {
		panic(err)
	}
	util.Log("")
	util.Log(fmt.Sprintf("| %-30v | %v         | %-20v | %-30v |", "NAME", "ID", "STATE", "IMAGE"))
	for _, container := range containers {
		util.Log(fmt.Sprintf("| %-30v | %v | %-20v | %-30v |", container.Names[0][1:], container.ID[:10], container.State, container.Image))
	}
}