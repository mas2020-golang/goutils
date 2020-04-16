package fs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Story struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"descpription"`
	LogFile     string   `yaml:"logfile"`
	GoALCmds    []string `yaml:"goalcmds,flow"`
}

func TestReadCommand(t *testing.T) {
	yamlF := "test.yaml"
	story:= Story{}
	if err := ReadYaml(yamlF, &story); err != nil {
		fmt.Println(err.Error())
	}
	// assert equality
	assert.Equal(t, 4, len(story.GoALCmds), "GoALCmds should be equal 4")
}

