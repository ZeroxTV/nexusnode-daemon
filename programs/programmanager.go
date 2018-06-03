package programs

import (
	"io/ioutil"
	"nexusnode.de/nexusnode-daemon/util"
	"os"
	"path/filepath"
)

var programCache map[string]Program

func DownloadAll() {
	//TODO Redownload all programs from panel server
}

func CheckUpdates() {
	//TODO Load updates from panel server
}

func LoadAllPrograms() {
	files, err := ioutil.ReadDir(programLocation)
	if err != nil {
		util.Log("Could not list all programs from file. Error:", err)
		return
	}
	for _, fileinfo := range files {
		file, err := os.Open(filepath.Join(programLocation, fileinfo.Name()))
		if err != nil {
			util.Log("Could not open file", fileinfo.Name() + ". Error:", err)
			file.Close()
			continue
		}
		program, err := LoadProgramFile(file)
		if err == nil {
			AddToCache(program)
		}
	}
}

func AddToCache(program Program) {
	programCache[program.Name] = program
}

func GetProgram(name string) (Program, error) {
	program, ok := programCache[name]
	if ok {
		return program, nil
	} else {
		program, err := LoadNamedProgram(name)
		if err != nil {
			util.Log("Unknown program", name + ".", "Error:", err)
			return program, err
		}
		AddToCache(program)
		return program, nil
	}
}