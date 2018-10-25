package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

func GetPlan(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	// fmt.Printf("buf: %+v\n", string(buf))

	var data map[interface{}]interface{}
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", data["plans"])
	fmt.Printf("%T\n", data["plans"])

	for _, val := range data["plans"].([]interface{}) {
		value := val.(map[interface{}]interface{})
		fmt.Printf("%+v\n", value["url"])
		fmt.Printf("%T\n", value["url"])
	}
	return nil
}
