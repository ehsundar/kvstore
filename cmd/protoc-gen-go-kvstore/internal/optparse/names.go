package optparse

import (
	"strings"

	"github.com/iancoleman/strcase"
)

var keySuffixes = []string{
	"key",
	"Key",
}

var valueSuffixes = []string{
	"val",
	"Val",
	"value",
	"Value",
}

func extractKeyName(name string) string {
	sfx := ""
	for _, s := range keySuffixes {
		if strings.HasSuffix(name, s) {
			sfx = s
			break
		}
	}

	return strcase.ToKebab(name[:len(name)-len(sfx)])
}

func extractValueName(name string) string {
	sfx := ""
	for _, s := range valueSuffixes {
		if strings.HasSuffix(name, s) {
			sfx = s
			break
		}
	}

	return strcase.ToKebab(name[:len(name)-len(sfx)])
}
