package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type AutoGenerated struct {
	Variables []struct {
		Name  string `yaml:"name"`
		Value string `yaml:"value"`
	} `yaml:"variables"`
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Please provide a yaml file as argument")
	}
	fileName := os.Args[1]

	if len(fileName) == 0 {
		log.Fatalln("Please provide a yaml file as argument")
	}

	t := AutoGenerated{}

	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("Error reading YAML file: %s\n", err)
		return
	}

	err = yaml.Unmarshal(yamlFile, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	out := ""
	bash := ""
	for _, v := range t.Variables {
		out = out + v.Name + "=" + v.Value + "\n"
		bash = bash + v.Name + "=" + v.Value + ";"
	}
	log.Println(bash)

	targetPath := filepath.Dir(fileName)
	p := filepath.Join(targetPath, ".env")
	err = ioutil.WriteFile(p, []byte(out), 0644)
	if err != nil {
		log.Printf("Error writing .env: %s\n", err)
		return
	}
}
