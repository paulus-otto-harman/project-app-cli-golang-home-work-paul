package component

import (
	"fmt"
)

type Param struct {
	Key   string
	Value interface{}
}

func P(key string, value interface{}) Param {
	return Param{key, value}
}

func Args(params ...Param) map[string]interface{} {
	var parameters = make(map[string]interface{})
	for _, param := range params {
		parameters[param.Key] = param.Value
	}
	return parameters
}

func ToString(data interface{}, err error) (string, error) {
	// TODO : validasi
	return fmt.Sprintf("%v", data), err
}

func ToInt(data interface{}, err error) (int, error) {
	// TODO : validasi

	if err != nil {
		panic("Sistem mendeteksi error. Aplikasi dihentikan.")
	}

	number, ok := data.(int)
	if !ok {
		return 0, err
	}
	return number, err
}

func Input(params ...map[string]interface{}) (interface{}, error) {
	if len(params) > 0 {
		if value, ok := params[0]["label"]; ok {
			fmt.Printf("%s ", value)
		}

		if value, ok := params[0]["type"]; ok {
			if value == "number" {
				var inputAngka int
				_, err := fmt.Scanln(&inputAngka)
				return inputAngka, err
			}
		}
	}

	// default parameters :: gunakan jika komponen ini tidak memiliki parameter sama sekali
	var inputTeks string
	_, err := fmt.Scanln(&inputTeks)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("%v", inputTeks), err
}
