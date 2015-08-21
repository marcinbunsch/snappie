package main

import (
  "path"
  "github.com/kylelemons/go-gypsy/yaml"
)

func LoadConfig() *yaml.File {

  location := path.Join(GetHomeLocation(), ".snappie.yml")

  config, err := yaml.ReadFile(location)

  if err != nil {
    return nil
  }

  return config

}

