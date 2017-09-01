package main

import (
	"fmt"
	"github.com/paulchiu/xinput-toggle/xinput"
	"os"
)

const Usage = `Usage: xinput-toggle [device-id]

xinput-toggle is a simple command for toggling the enabled/disabled state
of a given device id. For list of devices, run xinput
`

func main() {
	// Check args
	if len(os.Args) < 2 {
		fmt.Println(Usage)
		os.Exit(0)
	}

	// Get device id
	deviceId := os.Args[1]

	// Parse props to find out if the device is enabled
	deviceProperties, getDevicePropertiesError := xinput.GetPropertiesForDeviceId(deviceId)
	if getDevicePropertiesError != nil {
		panic(getDevicePropertiesError)
	}

	deviceIsEnabled := xinput.IsDeviceEnabled(deviceId, deviceProperties)

	// Determine state to toggle to
	toggleToState := xinput.StateEnabled
	if deviceIsEnabled {
		toggleToState = xinput.StateDisabled
	}

	// Update device state
	setStateError := xinput.SetStateForDeviceId(deviceId, toggleToState)
	if setStateError != nil {
		panic(setStateError)
	}

	// Output device state
	fmt.Printf("Device id %s %sd", deviceId, toggleToState)
}
