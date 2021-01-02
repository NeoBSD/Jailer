package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/NeoBSD/jailer"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	path := "cmd/jailer-compose/testdata/jailer-compose.yml"
	compose, err := jailer.ReadComposeFile(path)
	if err != nil {
		return err
	}

	// fmt.Println(compose)

	js, err := json.Marshal(compose)
	fmt.Println(string(js))
	if err != nil {
		return err
	}

	return nil
}
