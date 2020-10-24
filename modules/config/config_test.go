package config

import (
	"reflect"
	"testing"
)

func TestGetDefault(t *testing.T) {
	configs := NewFromFile("./testdata/example.toml", "./testdata/example2.toml")

	value := configs.GetDefault("not_exists", "default")
	if value.(string) != "default" {
		t.Error("we expect default as a value, we got: ", value)
	}
}

func TestGet(t *testing.T) {
	configs := NewFromFile("./testdata/example.toml", "./testdata/example2.toml")

	if configs.GetString("app.version") != "0.0.1" {
		t.Error("we expect 0.0.1, we got: ", configs.GetString("app.version"))
	}
}

func TestGetString(t *testing.T) {
	configs := NewFromFile("./testdata/example.toml", "./testdata/example2.toml")

	if reflect.TypeOf(configs.GetString("string")).String() != "string" {
		t.Errorf("we expect string type, we got: %s", reflect.TypeOf(configs.GetString("string")))
	}

	if configs.GetString("string") != "string" {
		t.Error("we expect string, we got: ", configs.GetString("string"))
	}

	if configs.GetString("app.version") != "0.0.1" {
		t.Error("we expect 0.0.1, we got: ", configs.GetString("app.version"))
	}

	if configs.GetString("string_not_exists") != "" {
		t.Errorf("we expect empty string, we got: %v", configs.GetString("string_not_exists"))
	}
}

func TestGetStringWithDefault(t *testing.T) {
	configs := NewFromFile("./testdata/example.toml", "./testdata/example2.toml")

	// Not exists
	value := configs.GetString("app.not_exists", "default")
	if value != "default" {
		t.Errorf("we expect default, we got: %v", value)
	}

	// Exists
	existValue := configs.GetString("app.version", "default")
	if existValue != "0.0.1" {
		t.Errorf("we expect default, we got: %v", existValue)
	}
}

func TestOverride(t *testing.T) {
	configs := NewFromFile("./testdata/example.toml", "./testdata/example2.toml")

	value := configs.GetString("override.value")
	if value != "exmaple2" {
		t.Errorf("we expect example2, we got: %s", value)
	}
}

func TestGetInt(t *testing.T) {
	configs := NewFromFile("./testdata/example.toml", "./testdata/example2.toml")

	if reflect.TypeOf(configs.GetInt("database.port")).String() != "int" {
		t.Errorf("we expect int type, we got: %s", reflect.TypeOf(configs.GetInt("database.port")))
	}

	if configs.GetInt("database.port") != 1653 {
		t.Error("we expect 1653, we got: ", configs.GetInt("database.port"))
	}
}

func TestGetInt64(t *testing.T) {
	configs := NewFromFile("./testdata/example.toml", "./testdata/example2.toml")

	if reflect.TypeOf(configs.GetInt64("database.port")).String() != "int64" {
		t.Errorf("we expect int64 type, we got: %s", reflect.TypeOf(configs.GetInt64("database.port")))
	}

	if configs.GetInt64("database.port") != 1653 {
		t.Error("we expect 1653, we got: ", configs.GetInt64("database.port"))
	}
}

func TestGetFloat(t *testing.T) {
	configs := NewFromFile("./testdata/example.toml", "./testdata/example2.toml")

	if reflect.TypeOf(configs.GetFloat64("server.port")).String() != "float64" {
		t.Errorf("we expect float64 type, we got: %s", reflect.TypeOf(configs.GetFloat64("server.port")))
	}

	if configs.GetFloat64("server.port") != 12.5 {
		t.Error("we expect 12.5, we got: ", configs.GetFloat64("server.port"))
	}
}

func TestGetBool(t *testing.T) {
	configs := NewFromFile("./testdata/example.toml", "./testdata/example2.toml")

	if reflect.TypeOf(configs.GetBool("server.tls")).String() != "bool" {
		t.Errorf("we expect bool type, we got: %s", reflect.TypeOf(configs.GetBool("server.tls")))
	}

	if configs.GetBool("server.tls") != true {
		t.Error("we expect true, we got: ", configs.GetBool("server.tls"))
	}

	if configs.GetBool("server.acme") != false {
		t.Error("we expect false, we got: ", configs.GetBool("server.acme"))
	}
}
