package main

import (
	"log"
	"os"
	"path"
)

func getPwd() string {
	err := os.Chdir(path.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln(err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	return pwd
}
