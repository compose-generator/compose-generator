/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"strings"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateExecServiceInitCommands runs all service init commands of the stack in a closed environment
func GenerateExecServiceInitCommands(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
	for _, template := range selectedTemplates.GetAll() {
		if len(template.ServiceInitCmd) > 0 {
			// Retrieve service init commands
			cmds := []string{}
			for _, cmd := range template.ServiceInitCmd {
				cmds = append(cmds, util.ReplaceVarsInString(cmd, project.Vars))
			}
			// Execute service init commands for this template
			spinner := startProcess("Generating configuration for " + template.Label + " ...")
			executeOnToolbox(strings.Join(cmds, " && "))
			stopProcess(spinner)
		}
	}
}
