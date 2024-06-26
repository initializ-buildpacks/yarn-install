package main

import (
	"log"
	"os"

	"github.com/initializ-buildpacks/yarn-install/cmd/setup-symlinks/internal"
)

func main() {
	projectPath, set := os.LookupEnv("NODE_PROJECT_PATH")
	if !set {
		var err error
		projectPath, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}

	err := internal.Run(os.Args[0], projectPath)
	if err != nil {
		log.Fatal(err)
	}
}
