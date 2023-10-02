package config

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DatasetName      string `json:"DATASET_NAME,omitempty"`
	SplitTestSize    string `json:"SPLIT_TEST_SIZE,omitempty"`
	SplitRandomState string `json:"SPLIT_RANDOM_STATE,omitempty"`
}

func GetConfig() Configuration {
	configuration := Configuration{}
	fileName := "config/env.json"
	err := gonfig.GetConf(fileName, &configuration)
	if err != nil {
		return Configuration{}
	}
	return configuration
}
