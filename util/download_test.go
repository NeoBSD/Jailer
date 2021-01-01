package util_test

// import (
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"testing"

// 	"github.com/NeoBSD/jailer/jail"
// )

// // fileExists checks if a file exists and is not a directory before we
// // try using it to prevent further errors.
// func fileExists(filename string) bool {
// 	info, err := os.Stat(filename)
// 	if os.IsNotExist(err) {
// 		return false
// 	}
// 	return !info.IsDir()
// }

// func TestDownloadFile(t *testing.T) {

// 	file, err := ioutil.TempFile(".", "test")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer os.Remove(file.Name())

// 	url := "https://raw.githubusercontent.com/NeoBSD/jailer/master/README.md"
// 	filepath := file.Name()
// 	err = jail.DownloadFile(filepath, url)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if !fileExists(filepath) {
// 		t.Errorf("Expected file at: %s", filepath)
// 	}

// }
