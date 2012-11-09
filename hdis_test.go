package hdis

import (
	"testing"
)

type KeyTest struct {
	input string
	key   string
	field string
}

var keyTests = []KeyTest{
	{"object", "object", ""},
	{"object:1", "object:", "1"},
	{"object:12", "object:", "12"},
	{"object:123", "object:12", "3"},
	{"object:1234", "object:12", "34"},
}

func TestGetKeyField(t *testing.T) {
	for _, entry := range keyTests {
		key, field := getKeyField(entry.input)

		if key != entry.key || field != entry.field {
			t.Errorf("getKeyField(%q) = %q,%q want %q,%a", entry.input, key, field, entry.key, entry.field)
		}
	}
}
