package main

import (
	"nexusnode.de/nexusnode-daemon/util"
	"nexusnode.de/nexusnode-daemon/dockerclient"
	"nexusnode.de/nexusnode-daemon/servers"
)

func main() {
	util.Log("###################################################")
	util.Log("#             Nexusnode Panel Daemon              #")
	util.Log("#                   Version 1.0                   #")
	util.Log("###################################################\n")
	util.Log("Checking for updates...")

	checkUpdates()

	util.Log("Connecting to docker installation...")

	dockerclient.ConnectToDocker()

	util.Log("Successfully connected to a running docker installation")
	util.Log("Checking out all containers...")

	dockerclient.PrintAllContainers()
	util.Log("")
	servers.LoadAllServers()
}

func checkUpdates() {
	//TODO Checking for daemon, image and configuration updates
	util.Log("No updates found")
}