package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Plans struct {
	Plans []Plan `yaml:plans`
}

type Plan struct {
	Url string `yaml:url`
	Res string `yaml:res`
}

func GetPlan() {
	// yamlを読み込む
	buf, err := ioutil.ReadFile("zgoc.yml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("buf: %+v\n", string(buf))

	var data map[interface{}]interface{}
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("d: %+v\n", data["plans"])
}
