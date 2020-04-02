package glide

import (
	"paraglider/glider/yamlparser"
	"paraglider/glider/glidercore"
	"paraglider/glider/models"
	"strings"
	"strconv"
	"fmt"
)

func Glide(yamlconfig, lbStartStop string) {

	// parse the default or provided yaml
	configPointer, err := yamlparser.ParseYaml(yamlconfig)
	config := *configPointer

	if err != nil {
		fmt.Println("Could not parse config file.")
		return
	}

	// Split the frontend bind address into host and port
	frontendAddr := strings.Split(config.Frontend.Bind, ":")
	frontendHost := frontendAddr[0]
	frontendPort, err := strconv.Atoi(frontendAddr[1])

	// Private IP is the ip via which backends will recognize the LB server
	// Private IP and bind IP can be the same as well
	frontendPrivateIP := config.Frontend.PrivateIP

	backendTargets := make([]models.TargetBackendHolder, 0, 0)

	for _, backend := range config.Frontend.Backends {

		backendAddr := strings.Split(backend, ":")
		backendHost := backendAddr[0]
		backendPort, _ := strconv.Atoi(backendAddr[1])

		// Create a new target and append it to the list
		backendTargets = append(backendTargets, models.TargetBackendHolder{
			BackendIP: backendHost,
			BackendPort: int(backendPort),
			LbIP: frontendHost,
			LbPort: frontendPort,
			LbPrivateIP: frontendPrivateIP,

		})
	}

	if lbStartStop == "start" {

		// Add the backend targets
		err = addBackendTargets(&backendTargets)

		if err != nil {
			fmt.Printf("%v\n", err)
			fmt.Println("Unable to add backends")
		}
	} else {

		// Optionally, remove the backend targets
		err = removeBackendTargets(&backendTargets)

		if err != nil {
			fmt.Println("Unabel to remove backends")
		}
	}
}

func addBackendTargets(backendTargets *[]models.TargetBackendHolder) error {

	roundRobinTurn := len(*backendTargets)
	for _, backendTarget := range *backendTargets {

		// Loop over the backend targets and create a target using iptable rules for each one
		// of them
		err := glidercore.CreateTargetForLb(backendTarget, roundRobinTurn)

		if err != nil {
			return err
		}

		roundRobinTurn -= 1
	}

	return nil
}

func removeBackendTargets(backendTargets *[]models.TargetBackendHolder) error {

	roundRobinTurn := len(*backendTargets)
	for _, backendTarget := range *backendTargets {

		// Loop over the backend targets and remove each one of them from iptables= rules.
		err := glidercore.RemoveTargetForLb(backendTarget, roundRobinTurn)

		if err != nil {
			return err
		}

		roundRobinTurn -= 1
	}

	return nil
}

