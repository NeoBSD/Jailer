// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"

// 	"github.com/NeoBSD/jailer"
// )

// func main() {
// 	if err := run(); err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// }

// func run() error {
// 	path := "testdata/jailer-compose/jailer-compose.yml"
// 	compose, err := jailer.ReadComposeFile(path)
// 	if err != nil {
// 		return err
// 	}

// 	// fmt.Println(compose)

// 	js, err := json.Marshal(compose)
// 	fmt.Println(string(js))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/NeoBSD/jailer"
)

// These variables get set during link time. See Makefile
var (
	hostOS string
	commit string
	date   string
)

func init() {

	// Set build info
	jailer.BuildCommit = commit
	jailer.BuildDate = date
	jailer.BuildOS = hostOS

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Errorln(err)
		os.Exit(1)
	}
}
