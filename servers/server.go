package servers

import (
	"docker.io/go-docker/api/types"
	"context"
	"nexusnode.de/nexusnode-daemon/util"
	"nexusnode.de/nexusnode-daemon/dockerclient"
	"nexusnode.de/nexusnode-daemon/programs"
)

type Server struct {
	id string
	owner string
	program *programs.Program
}

//Create a container
func Create(name string, owner string, program programs.Program) (*Server, error) {
	var config = program.Config
	var hostConfig = program.HostConfig
	var networkConfig = program.NetworkConfig
	cont, err := dockerclient.Client.ContainerCreate(context.Background(), &config, &hostConfig, &networkConfig, name)
	if err != nil {
		util.Log("Failed to create container. Error:\n", err)
		return nil, err
	}
	return &Server{
		cont.ID,
	owner,
	&program}, nil
}

//Start the container
func (s *Server) Start() {
	info, err := dockerclient.Client.ContainerInspect(context.Background(), s.id)
	state := info.State.Status
	if err != nil {
		util.Log("Failed to start container", s.id, "because of error:\n", err)
		return
	}
	if !(state == "running" || state == "restarting") {
		dockerclient.Client.ContainerStart(context.Background(), s.id, types.ContainerStartOptions{})
	} else if state == "paused" {
		err = dockerclient.Client.ContainerUnpause(context.Background(), s.id)

		if err != nil {
			util.Log("Failed to start container", s.id, "because of error:\n", err)
			return
		}
	}
}

//Stop the container
func (s *Server) Stop() {
	info, err := dockerclient.Client.ContainerInspect(context.Background(), s.id)
	state := info.State.Status
	if err != nil {
		util.Log("Failed to stop container", s.id, "because of error:\n", err)
		return
	}
	if state == "running" || state == "restarting" {
		err = dockerclient.Client.ContainerStop(context.Background(), s.id, nil)

		if err != nil {
			util.Log("Failed to stop container", s.id, "because of error:\n", err)
			return
		}
	}
}

//Restart the container
func (s *Server) Restart() {
	info, err := dockerclient.Client.ContainerInspect(context.Background(), s.id)
	state := info.State.Status
	if err != nil {
		util.Log("Failed to restart container", s.id, "because of error:\n", err)
		return
	}
	if state != "restarting" {
		err = dockerclient.Client.ContainerRestart(context.Background(), s.id, nil)

		if err != nil {
			util.Log("Failed to restart container", s.id, "because of error:\n", err)
			return
		}
	}
}

//Kill/Force stop the container
func (s *Server) Kill() {
	info, err := dockerclient.Client.ContainerInspect(context.Background(), s.id)
	state := info.State.Status
	if err != nil {
		util.Log("Failed to kill container", s.id, "because of error:\n", err)
		return
	}
	if state == "running" || state == "restarting" {
		err = dockerclient.Client.ContainerKill(context.Background(), s.id, "SIGKILL")

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