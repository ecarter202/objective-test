package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
)

type JSON struct{}

func (j JSON) Load(filePath string) error {
	fmt.Printf("Loading JSON file %s...\n", filePath)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.Wrap(err, "reading json file")
	}

	if err := json.Unmarshal(data, &dataset); err != nil {
		return errors.Wrap(err, "unmarshaling json data")
	}

	return nil
}
