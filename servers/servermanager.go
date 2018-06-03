package servers

import (
	"docker.io/go-docker/api/types"
	"nexusnode.de/nexusnode-daemon/util"
	"nexusnode.de/nexusnode-daemon/dockerclient"
	"context"
	"nexusnode.de/nexusnode-daemon/programs"
)

var servers map[string]*Server

func LoadAllServers() {
	containers, err := dockerclient.Client.ContainerList(context.Background(), types.ContainerListOptions{All:true})

	if err != nil {
		util.Log("Failed to list all containers. Error:\n", err)
		return
	}
	for _, container := range containers {
		_, ok := container.Labels["nexusnode"]
		if !ok {
			util.Log("Container", container.ID[:10], "is not a nexusnode server. Skipping.")
			continue
		}
		owner, aok := container.Labels["owner"]
		programJson, bok := container.Labels["program"]
		program, err := programs.LoadProgramJSON(programJson)

		if err != nil || !aok || !bok {
			util.Log("Failed to load server", container.ID, "because of missing or incorrect metadata (labels)")
			continue
		}

		server := Server{
			container.ID,
			owner,
			&program,
		}
		servers[container.ID] = &server
		util.Log("Successfully loaded server", container.ID)
	}
}