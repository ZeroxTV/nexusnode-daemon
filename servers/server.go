package servers

import (
	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"context"
	"nexusnode.de/nexusnode-daemon/util"
	"nexusnode.de/nexusnode-daemon/dockerclient"
	"docker.io/go-docker/api/types/container"
	"docker.io/go-docker/api/types/network"
)

type Server struct {
	id string
	dock *docker.Client
}

//Create a container
func Create(name string) (*Server, error) {
	var config container.Config
	var hostConfig container.HostConfig
	var networkConfig network.NetworkingConfig
	cont, err := dockerclient.Client.ContainerCreate(context.Background(), &config, &hostConfig, &networkConfig, name)
	if err != nil {
		util.Log("Failed to create container. Error:\n", err)
		return nil, err
	}
	return &Server{
		cont.ID,
		dockerclient.Client}, nil
}

//Start the container
func (s *Server) Start() {
	info, err := s.dock.ContainerInspect(context.Background(), s.id)
	state := info.State.Status
	if err != nil {
		util.Log("Failed to start container", s.id, "because of error:\n", err)
		return
	}
	if !(state == "running" || state == "restarting") {
		s.dock.ContainerStart(context.Background(), s.id, types.ContainerStartOptions{})
	} else if state == "paused" {
		err = s.dock.ContainerUnpause(context.Background(), s.id)

		if err != nil {
			util.Log("Failed to start container", s.id, "because of error:\n", err)
			return
		}
	}
}

//Stop the container
func (s *Server) Stop() {
	info, err := s.dock.ContainerInspect(context.Background(), s.id)
	state := info.State.Status
	if err != nil {
		util.Log("Failed to stop container", s.id, "because of error:\n", err)
		return
	}
	if state == "running" || state == "restarting" {
		err = s.dock.ContainerStop(context.Background(), s.id, nil)

		if err != nil {
			util.Log("Failed to stop container", s.id, "because of error:\n", err)
			return
		}
	}
}

//Restart the container
func (s *Server) Restart() {
	info, err := s.dock.ContainerInspect(context.Background(), s.id)
	state := info.State.Status
	if err != nil {
		util.Log("Failed to restart container", s.id, "because of error:\n", err)
		return
	}
	if state != "restarting" {
		err = s.dock.ContainerRestart(context.Background(), s.id, nil)

		if err != nil {
			util.Log("Failed to restart container", s.id, "because of error:\n", err)
			return
		}
	}
}

//Kill/Force stop the container
func (s *Server) Kill() {
	info, err := s.dock.ContainerInspect(context.Background(), s.id)
	state := info.State.Status
	if err != nil {
		util.Log("Failed to kill container", s.id, "because of error:\n", err)
		return
	}
	if state == "running" || state == "restarting" {
		err = s.dock.ContainerKill(context.Background(), s.id, "SIGKILL")

		if err != nil {
			util.Log("Failed to kill container", s.id, "because of error:\n", err)
			return
		}
	}
}

//Reinstall the image, running in the container
//Basically deletes and recreates the container
func (s *Server) Reinstall() {
	//TODO
	//TODO Create Backup before recreating
}