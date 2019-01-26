package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Threads []Request `yaml:"threads"`
}

type Request struct {
	URL     string
	Headers []interface{}
}

func GetConfig(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// interface{}だと速度が低下するため準備しておいた構造体に値を読み込む
	var config Config
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", config.Threads)
	fmt.Printf("%T\n", config.Threads)
	// リクエスト群の整形
	for _, request := range config.Threads {
		fmt.Print("\n range\n\n")
		fmt.Printf("%+v\n", request.URL)
		fmt.Printf("%T\n", request.URL)
		fmt.Printf("%+v\n", request.Headers)
		fmt.Printf("%T\n", request.Headers)
		fmt.Print("\n headers \n\n")
		for _, header := range request.Headers {
			fmt.Printf("%s\n", header)
			fmt.Printf("%+v\n", header.(map[interface{}]interface{})["X-ACCESS-TOKEN"])
			fmt.Printf("%T\n", header)
			fmt.Printf("%T\n", header.(interface{}))
		}
	}
	return nil
}
