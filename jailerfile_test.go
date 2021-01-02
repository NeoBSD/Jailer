package jailer_test

import (
	"testing"

	"github.com/NeoBSD/jailer"
	"github.com/matryer/is"
)

func TestReadFromFile(t *testing.T) {

	var tests = []struct {
		input       string
		expected    *jailer.Jailerfile
		expectError bool
	}{
		{"testdata/jailer/noexist/Jailerfile", &jailer.Jailerfile{}, true},
		{"testdata/jailer/label/Jailerfile", &jailer.Jailerfile{BaseImage: jailer.BaseImage{Name: "freebsd", Version: "latest"}}, false},
	}

	for _, tt := range tests {
		is := is.New(t)
		actual, err := jailer.ReadFromFile(tt.input)
		is.Equal(err != nil, tt.expectError)
		is.Equal(actual.BaseImage, tt.expected.BaseImage)
	}
}

func TestLabelParsing(t *testing.T) {
	is := is.New(t)
	jf, err := jailer.ReadFromFile("testdata/jailer/label/Jailerfile")
	is.NoErr(err)
	is.Equal(jf.Labels["maintainer"], `"example@example.com"`)
	is.Equal(jf.Labels["version"], `"1.0"`)
}

func TestFromWithImplicitLatestParsing(t *testing.T) {
	is := is.New(t)
	jf, err := jailer.ReadFromFile("testdata/jailer/from/Jailerfile")
	is.NoErr(err)
	is.Equal(jf.BaseImage.Name, "freebsd")
	is.Equal(jf.BaseImage.Version, "latest")

}

func TestFromWithExplicitLatestParsing(t *testing.T) {
	is := is.New(t)
	jf, err := jailer.ReadFromFile("testdata/jailer/from_with_latest/Jailerfile")
	is.NoErr(err)
	is.Equal(jf.BaseImage.Name, "freebsd")
	is.Equal(jf.BaseImage.Version, "latest")
}

func TestFromWithExplicitVersionParsing(t *testing.T) {
	is := is.New(t)
	jf, err := jailer.ReadFromFile("testdata/jailer/from_with_version/Jailerfile")
	is.NoErr(err)
	is.Equal(jf.BaseImage.Name, "freebsd")
	is.Equal(jf.BaseImage.Version, "12.1")
}

func TestRunParsing(t *testing.T) {
	is := is.New(t)
	jf, err := jailer.ReadFromFile("testdata/jailer/run/Jailerfile")
	is.NoErr(err)
	is.Equal(len(jf.Instructions), 2)

	t.Run("first", func(t *testing.T) {
		is := is.New(t)
		val, ok := jf.Instructions[0].(*jailer.RunInstruction)
		is.True(ok)
		is.Equal(val.Name(), "RUN")
		is.Equal(val.Command, "echo \"Hello Jailer!\"")
	})

	t.Run("second", func(t *testing.T) {
		is := is.New(t)
		val, ok := jf.Instructions[1].(*jailer.RunInstruction)
		is.True(ok)
		is.Equal(val.Name(), "RUN")
		is.Equal(val.Command, "pkg install -y nano")
	})

}

func TestWorkDirParsing(t *testing.T) {
	is := is.New(t)
	jf, err := jailer.ReadFromFile("testdata/jailer/workdir/Jailerfile")
	is.NoErr(err)
	is.Equal(len(jf.Instructions), 1)
	val, ok := jf.Instructions[0].(*jailer.WorkDirInstruction)
	is.True(ok)
	is.Equal(val.Name(), "WORKDIR")
	is.Equal(val.Command, "/work")
}
