package coms

import (
	"fmt"
)

func Input(params map[string]interface{}) (interface{}, error) {
	fmt.Printf("%s ", params["label"])

	var inputAngka int
	if params["type"] == "number" {
		_, err := fmt.Scanln(&inputAngka)
		return inputAngka, err
	}

	var inputTeks string
	_, err := fmt.Scanln(&inputTeks)
	return inputTeks, err
}
