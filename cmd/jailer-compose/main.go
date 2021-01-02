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
	"fmt"
	"os"

	"github.com/NeoBSD/jailer"
)

// These variables get set during link time. See Makefile
var (
	hostOS string
	commit string
	date   string
)

func init() {
	jailer.BuildCommit = commit
	jailer.BuildDate = date
	jailer.BuildOS = hostOS
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
