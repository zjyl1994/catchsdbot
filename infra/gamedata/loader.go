package gamedata

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed gamedata.yaml
var dataJson []byte

func Load() (*GameData, error) {
	var gamedata GameData
	err := yaml.Unmarshal(dataJson, &gamedata)
	return &gamedata, err
}
