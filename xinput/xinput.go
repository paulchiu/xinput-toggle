package xinput

import (
	"bytes"
	"os/exec"
	"regexp"
	"strings"
)

const (
	ArgumentListProps   = "list-props"
	CommandXinput       = "xinput"
	Eol                 = "\n"
	PropertyEnabled     = "1"
	RegexpDeviceEnabled = `Device Enabled.*:\s+(\d)`
	StateDisabled       = "disable"
	StateEnabled        = "enable"
)

func GetPropertiesForDeviceId(deviceId string) ([]string, error) {
	// Prepare output buffer
	var out bytes.Buffer

	// Execute list props
	listPropsCommand := exec.Command(CommandXinput, ArgumentListProps, deviceId)
	listPropsCommand.Stdout = &out
	listPropsError := listPropsCommand.Run()

	// Log fatal error if occurred
	if listPropsError != nil {
		return nil, listPropsError
	}

	// Parse and return props
	props := strings.Split(out.String(), Eol)
	return props, nil
}

func IsDeviceEnabled(deviceProperties []string) bool {
	// Set default device is enabled state
	deviceIsEnabled := false
	deviceIsEnabledRegExp := regexp.MustCompile(RegexpDeviceEnabled)

	// Attempt to find enabled property state
	for _, property := range deviceProperties {
		match := deviceIsEnabledRegExp.FindStringSubmatch(property)

		if len(match) > 0 {
			deviceIsEnabled = match[1] == PropertyEnabled
			break
		}
	}

	// Return default enabled state
	return deviceIsEnabled
}

func SetStateForDeviceId(deviceId string, state string) error {
	// Prepare output buffer
	var out bytes.Buffer

	// Execute toggle command
	stateChangeCommand := exec.Command(CommandXinput, state, deviceId)
	stateChangeCommand.Stdout = &out
	stateChangeCommandError := stateChangeCommand.Run()
	return stateChangeCommandError
}
