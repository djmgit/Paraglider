package yamlparser

import (
	"paraglider/glider/models"
	"io/ioutil"
)

func ParseYaml(yamlPath string) (*models.Config, error) {
	data, err := ioutil.ReadFile(yamlPath)

	if err != nil {
		return nil, err
	}

	config := models.Config{}

	// Parse the read content of the file as YAML
	err = config.Parse(data)

	if err != nil {
		return nil, err
	}

	return &config, nil
}