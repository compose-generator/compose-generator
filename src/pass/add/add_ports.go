/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"strconv"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddPorts asks the user if he/she wants to add ports to the configuration
func AddPorts(service *spec.ServiceConfig, _ *model.CGProject) {
	if yesNoQuestion("Do you want to expose ports of your service?", false) {
		pel()
		// Create list if not exists
		if service.Ports == nil {
			service.Ports = []spec.ServicePortConfig{}
		}
		// Question loop
		for another := true; another; another = yesNoQuestion("Expose another port?", true) {
			// Ask for inner and outer port
			portInner := textQuestionWithValidator("Which port do you want to expose? (inner port)", util.PortValidator)
			portOuter := textQuestionWithValidator("To which destination port on the host machine?", util.PortValidator)
			portInnerInt, err := strconv.ParseUint(portInner, 10, 32)
			if err != nil {
				errorLogger.Println("Inner port could not be converted to uint32: " + err.Error())
				logError("Inner port could not be converted to uint32", false)
				return
			}
			portOuterInt, err := strconv.ParseUint(portOuter, 10, 32)
			if err != nil {
				errorLogger.Println("Outer port could not be converted to uint32: " + err.Error())
				logError("Outer port could not be converted to uint32", false)
				return
			}

			// Add port to service
			service.Ports = append(service.Ports, spec.ServicePortConfig{
				Mode:      "ingress",
				Protocol:  "tcp",
				Target:    uint32(portInnerInt),
				Published: uint32(portOuterInt),
			})
			infoLogger.Println("Adding port mapping " + portOuter + ":" + portInner + " to new service")
		}
		pel()
	}
}
