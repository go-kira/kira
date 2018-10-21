package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	DefaultPath = "./testdata"
	DefaultVariablesPath = "./testdata/enviroments"

	New("production")
	t.Fail()
}

func TestGetDefault(t *testing.T) {
	DefaultPath = "./testdata/"
	DefaultVariablesPath = "./testdata/environments/"

	New("development")
	def := GetDefault("SESSION_NAMEe", "default")

	if def != "default" {
		t.Error("the value not 'default'")
	}
}
