package parser

import (
	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
)

// LoadProject loads the Docker compose project from the current directory
func LoadProject() {
	configDetails := types.ConfigDetails{}
	loader.Load(configDetails)
}

// SaveProject saves the Docker compose project to the current directory
func SaveProject() {

}
