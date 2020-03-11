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

	config, err := yamlparser.ParseYaml(yamlconfig)

	if err != nil {
		fmt.Println("Could not parse config file.")
		return
	}

	frontendAddr := strings.Split(config.Frontend.Bind, ":")
	frontendHost := frontendAddr[0]
	frontendPort, err := strconv.Atoi(frontendAddr[1])
	frontendPrivateIP := config.Frontend.PrivateIP

	backendTargets := make([]models.TargetBackendHolder, 1, 1)

	for _, backend := range config.Frontend.Backends {

		backendAddr := strings.Split(backend, ":")
		backendHost := backendAddr[0]
		backendPort, _ := strconv.Atoi(backendAddr[1])

		backendTargets = append(backendTargets, models.TargetBackendHolder{
			BackendIP: backendHost,
			BackendPort: backendPort,
			LbIP: frontendHost,
			LbPort: frontendPort,
			LbPrivateIP: frontendPrivateIP,

		})
	}

	if lbStartStop == "start" {
		err = addBackendTargets(&backendTargets)

		if err != nil {
			fmt.Println("Unable to add backends")
		}
	} else {
		err = removeBackendTargets(&backendTargets)

		if err != nil {
			fmt.Println("Unabel to remove backends")
		}
	}
}

func addBackendTargets(backendTargets *[]models.TargetBackendHolder) error {

	for _, backendTarget := range *backendTargets {
		err := glidercore.CreateTargetForLb(backendTarget)

		if err != nil {
			return err
		}
	}

	return nil
}

func removeBackendTargets(backendTargets *[]models.TargetBackendHolder) error {

	for _, backendTarget := range *backendTargets {
		err := glidercore.RemoveTargetForLb(backendTarget)

		if err != nil {
			return err
		}
	}

	return nil
}

