package fs

import (
	"gopkg.in/yaml.v2"
	"os"
)

// Read single module from yaml file and return the object as interface
func ReadYaml(filepath string, inter interface{}) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	// Unmarshall command
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(inter)
	if err != nil {
		return err
	}
	return nil
}

// Returns whether the given file or directory exists
func ExistsPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

