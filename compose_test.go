package jailer_test

import (
	"testing"

	"github.com/NeoBSD/jailer"
	"github.com/matryer/is"
)

func TestReadComposeFileFail(t *testing.T) {
	is := is.New(t)
	_, err := jailer.ReadComposeFile("unknown")
	is.True(err != nil)
}

func TestReadComposeFile(t *testing.T) {
	is := is.New(t)
	path := "testdata/jailer-compose/jailer-compose.yml"
	compose, err := jailer.ReadComposeFile(path)
	is.NoErr(err)
	is.Equal(compose.Version, "0.1")
	is.Equal(len(compose.Services), 2)
	is.Equal(compose.Services[0].Label, "web")
}

func TestComposeValidate(t *testing.T) {
	is := is.New(t)

	t.Run("OK", func(t *testing.T) {
		is := is.New(t)
		c := jailer.Compose{
			Version: "0.1",
			Services: []jailer.Service{
				{Label: "test1"},
				{Label: "test2"},
			},
		}
		is.NoErr(c.Validate())
	})

	t.Run("FAIL: Missing version", func(t *testing.T) {
		is := is.New(t)
		c := jailer.Compose{}
		is.True(c.Validate() != nil)

	})

	t.Run("FAIL: Label used twice", func(t *testing.T) {
		is := is.New(t)
		c := jailer.Compose{
			Version: "0.1",
			Services: []jailer.Service{
				{Label: "test1"},
				{Label: "test1"},
			},
		}
		is.True(c.Validate() != nil)
	})
}
