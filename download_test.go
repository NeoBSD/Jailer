package jailer_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/NeoBSD/jailer"
	"github.com/matryer/is"
)

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func TestDownloadFile(t *testing.T) {
	is := is.New(t)
	file, err := ioutil.TempFile(".", "test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	url := "https://raw.githubusercontent.com/NeoBSD/jailer/master/README.md"
	filepath := file.Name()
	err = jailer.DownloadFile(filepath, url)
	is.NoErr(err)
	is.True(fileExists(filepath))
}
