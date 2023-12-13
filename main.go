package main

import (
	"log"
	"os"

	"github.com/moovlaze/rest-api-server/apiserver"
	"gopkg.in/yaml.v3"
)

// var (
// 	configPath string
// )

// func init() {
// 	flag.StringVar(&configPath, "config-path", "configs/apiserver.yaml", "path to config file")
// }

func main() {
	// flag.Parse()


	openConfigFile, err := os.Open("configs/apiserver.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer openConfigFile.Close()

	yamlDecoder := yaml.NewDecoder(openConfigFile)

	config := apiserver.NewConfig()
	yamlDecoder.Decode(config)

	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}