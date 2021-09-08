package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"strings"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateExecDemoAppInitCommands runs all demo app init commands of the stack in a closed environment
func GenerateExecDemoAppInitCommands(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
	for _, template := range selectedTemplates.GetAll() {
		if len(template.DemoAppInitCmd) > 0 {
			// Retrieve service init commands
			cmds := []string{}
			for _, cmd := range template.DemoAppInitCmd {
				cmds = append(cmds, util.ReplaceVarsInString(cmd, project.Vars))
			}
			// Execute demo app init commands for this template
			p("Generating demo app for " + template.Label + " ... ")
			executeOnLinux(strings.Join(cmds, " && "))
			done()
		}
	}
}
