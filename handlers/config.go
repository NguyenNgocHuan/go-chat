package handlers

import (
	"fmt"
	"go-chat/models"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
}

func ReadFile(cfg *models.Config, path string) {
	f, err := os.Open(path)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}

}

func processError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
