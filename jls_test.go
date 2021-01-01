package jailer_test

import (
	"testing"

	"github.com/NeoBSD/jailer"
)

func TestParseJLSOutputMockEmpty(t *testing.T) {
	jls := jailer.JLS{Path: "testdata/jls_empty"}
	jails, err := jls.GetActiveJails()

	if err != nil {
		t.Error(err)
	}

	if len(jails) != 0 {
		t.Errorf("Expected: %d, Got: %d", 0, len(jails))
	}

}

func TestParseJLSOutputMock(t *testing.T) {
	jls := jailer.JLS{Path: "testdata/jls"}
	jails, err := jls.GetActiveJails()

	if err != nil {
		t.Error(err)
	}

	if len(jails) != 2 {
		t.Errorf("Expected: %d, Got: %d", 2, len(jails))
	}

}

// func TestParseJLSOutputJID(t *testing.T) {

// 	expected := map[string]string{"jid": "2"}

// 	actual, err := jailer.ParseJLSOutput("jid=2")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if actual["jid"] != expected["jid"] {
// 		t.Errorf("Expected: %v, Got: %v", expected, actual)
// 	}

// }

// func TestParseJLSOutputMultiplePairs(t *testing.T) {

// 	expected := map[string]string{"jid": "2", "name": "test_jail", "path": "/jailer/test_jail"}

// 	actual, err := jailer.ParseJLSOutput("jid=2 name=test_jail path=/jailer/test_jail")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if actual["jid"] != expected["jid"] {
// 		t.Errorf("Expected: %v, Got: %v", expected, actual)
// 	}

// 	if actual["name"] != expected["name"] {
// 		t.Errorf("Expected: %v, Got: %v", expected, actual)
// 	}

// 	if actual["path"] != expected["path"] {
// 		t.Errorf("Expected: %v, Got: %v", expected, actual)
// 	}
// }

// func TestParseJLSOutput(t *testing.T) {

// 	var tests = []struct {
// 		name     string
// 		input    string
// 		expected map[string]string
// 	}{
// 		{"empty", "", nil},
// 		{"jid", "jid=2", map[string]string{"jid": "2"}},
// 		{"multi", "jid=2 name=test_jail path=/jailer/test_jail", map[string]string{
// 			"jid":  "2",
// 			"name": "2",
// 			"path": "2",
// 		}},
// 	}

// 	for _, tt := range tests {

// 		t.Run(tt.input, func(t *testing.T) {

// 			actual, err := jailer.ParseJLSOutput(tt.input)

// 			if err != nil {
// 				t.Error(err)
// 			}

// 			if reflect.DeepEqual(actual, tt.expected) {
// 				t.Errorf("Expected: %q, Got: %q", tt.expected, actual)
// 			}

// 		})

// 	}

// }
