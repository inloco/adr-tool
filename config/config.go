package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type osFS struct {
	stat func(string) (os.FileInfo, error)
}

type fileSystem interface {
	Stat(name string) (os.FileInfo, error)
}

func (fs osFS) Stat(name string) (os.FileInfo, error) {
	return fs.stat(name)
}

var fs fileSystem = osFS{stat: os.Stat}

const (
	maxDepth       = 100
	configFileName = ".adr-prefix"

	start = "."
)

func findConfigFile(start string) (string, error) {
	filepath := path.Dir(start)
	for i := 0; i < maxDepth; i++ {
		_, err := fs.Stat(path.Join(filepath, configFileName))
		if err != nil {
			filepath = path.Join(filepath, "..")
		} else {
			return path.Join(filepath, configFileName), nil
		}
	}
	return "", fmt.Errorf("Config file not found")
}

func LoadConfigFile(confPath string) (string, error) {
	path, err := findConfigFile(start)
	if err != nil {
		return "", err
	}

	confBytes, err := ioutil.ReadFile(path)
	return string(confBytes), err
}
