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
	if YesNoQuestion("Do you want to expose ports of your service?", false) {
		Pel()
		// Create list if not exists
		if service.Ports == nil {
			service.Ports = []spec.ServicePortConfig{}
		}
		// Question loop
		for another := true; another; another = YesNoQuestion("Expose another port?", true) {
			// Ask for inner and outer port
			portInner := TextQuestionWithValidator("Which port do you want to expose? (inner port)", util.PortValidator)
			portOuter := TextQuestionWithValidator("To which destination port on the host machine?", util.PortValidator)
			portInnerInt, err := strconv.ParseUint(portInner, 10, 32)
			if err != nil {
				Error("Port could not be converted to uint32", err, false)
				return
			}
			portOuterInt, err := strconv.ParseUint(portOuter, 10, 32)
			if err != nil {
				Error("Port could not be converted to uint32", err, false)
				return
			}

			// Add port to service
			service.Ports = append(service.Ports, spec.ServicePortConfig{
				Mode:      "ingress",
				Protocol:  "tcp",
				Target:    uint32(portInnerInt),
				Published: uint32(portOuterInt),
			})
		}
		Pel()
	}
}
