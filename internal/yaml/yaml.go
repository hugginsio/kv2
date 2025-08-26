// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package yaml

import (
	"os"

	"sigs.k8s.io/yaml"
)

func LoadYaml(path string, config any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	return nil
}

func LoadYamlBytes(contents []byte, config any) error {
	err := yaml.Unmarshal(contents, &config)
	if err != nil {
		return err
	}

	return nil
}
