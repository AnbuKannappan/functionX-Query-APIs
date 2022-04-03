package utils

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
)

func YamlParser(output string) interface{} {

	var data interface{}
	if err := yaml.Unmarshal([]byte(output), &data); err != nil {
		panic(err)
	}

	data = convert(data)

	if outputJson, err := json.Marshal(data); err != nil {
		panic(err)
	} else {
		var f interface{}
		if err := json.Unmarshal([]byte(outputJson), &f); err != nil {
			panic(err)
		}
		return f
	}
}

func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}
