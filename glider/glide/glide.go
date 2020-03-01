package glide

import (
	"paraglider/glider/yamlparser"
	"paraglider/glider/glidercore"
	"paraglider/glider/models"
	"strings"
	"strconv"
)

func glide(yamlconfig string) {

	config, err := yamlparser.ParseYaml(yamlconfig)

	if err != nil {
		return
	}

	frontendAddr := strings.Split(config.Frontend.Bind, ":")
	frontendHost := frontendAddr[0]
	frontendPort, err := strconv.Atoi(frontendAddr[1])
	frontendPrivateIP := config.LbPrivateIP

	backendTargets := make([]models.TargetBackendHolder, 1, 1)

	for index, backend := range config.Frontend.Backends {

		backendAddr := strings.Split(backend, ":")
		backendHost := backendAddr[0]
		backendPort, err := strconv.Atoi(backendAddr[1])

		append(backendTargets, models.TargetBackendHolder{
			BackendIP: backendHost,
			BackendPort: backendPort,
			LbIP: frontendHost,
			LbPort: frontendPort,
			LbPrivateIP: frontendPrivateIP,

		})
	}
}


