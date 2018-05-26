package main

import (
	"docker.io/go-docker"
	"context"
	"docker.io/go-docker/api/types"
	"fmt"
)

func main() {
	cli, err := docker.NewEnvClient()

	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All:true})

	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}
