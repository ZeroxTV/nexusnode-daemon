package programs

import (
	"docker.io/go-docker/api/types/container"
	"docker.io/go-docker/api/types/network"
	"nexusnode.de/nexusnode-daemon/dockerclient"
	"context"
	"encoding/json"
	"nexusnode.de/nexusnode-daemon/util"
	"io/ioutil"
	"os"
	"path/filepath"
)

var programLocation = "/srv/nexusnode/programs/"

type Program struct {
	Name          string                   `json:"name"`
	Image         string                   `json:"image"`
	Config        container.Config         `json:"config"`
	HostConfig    container.HostConfig     `json:"hostConfig"`
	NetworkConfig network.NetworkingConfig `json:"networkConfig"`
}

func (p *Program) createContainer(name string) {
	dockerclient.Client.ContainerCreate(context.Background(), &p.Config, &p.HostConfig, &p.NetworkConfig, name)
}

func LoadNamedProgram(name string) (Program, error) {
	file, err := os.Open(filepath.Join(programLocation, filepath.Base(name + ".json")))
	defer file.Close()
	if err != nil {
		util.Log("Could not open program-json for", name, "Error:\n", err)
		return Program{}, err
	}
	return LoadProgramFile(file)
}

func LoadProgramFile(file *os.File) (Program, error) {
	byteValue, _ := ioutil.ReadAll(file)

	var program Program
	json.Unmarshal(byteValue, &program)
	return program, nil
}

func LoadProgramJSON(programJson string) (Program, error) {
	var program Program
	err := json.Unmarshal([]byte(programJson), &program)
	if err != nil {
		util.Log("Failed to create program from supplied json. Error:\n", err)
		return program, err
	}
	return program, nil
}

func (p *Program) SaveProgram() {
	programJson, err := json.Marshal(p)
	if err != nil {
		util.Log("Could not save program", p.Name, "to file. Error:\n", err)
		return
	}
	ioutil.WriteFile(filepath.Join(programLocation, filepath.Base(p.Name + ".json")), programJson, 0644)

}