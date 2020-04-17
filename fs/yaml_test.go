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

func TestReadYaml(t *testing.T) {
	yamlF := "test.yaml"
	story:= Story{}
	if err := ReadYaml(yamlF, &story); err != nil {
		fmt.Println(err.Error())
	}
	// assert equality
	assert.Equal(t, 4, len(story.GoALCmds), "GoALCmds should be equal 4")
}

func TestWriteYaml(t *testing.T) {
	yamlF := "test2.yaml"

	var s Story = Story{
		Name: "New Story",
		Description: "New Story",
		LogFile: "New Story",
		GoALCmds: []string{"cmd1", "cmd2", "cmd3"},
	}

	if err := WriteYaml(yamlF, s); err != nil {
		fmt.Println(err.Error())
	}
	// assert equality
	//assert.Equal(t, 4, len(story.GoALCmds), "GoALCmds should be equal 4")
}

func TestExistsPath(t *testing.T) {
	yamlF := "test.yaml"
	var (
		ok bool
		err error
	)
	if ok,err = ExistsPath(yamlF); err != nil {
		fmt.Println(err.Error())
	}
	// assert equality
	assert.Equal(t, true, ok, "File test.yaml should exist")
	yamlF = "test1.yaml"
	if ok,err = ExistsPath(yamlF); err != nil {
		fmt.Println(err.Error())
	}
	assert.Equal(t, false, ok, "File test.yaml should not exist")
}
