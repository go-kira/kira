package config

import (
	"bytes"
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/go-kira/kog"

	yaml "gopkg.in/yaml.v2"
)

var errInvalidEnv = errors.New("invalid environment, use only: development, production, test")

// DefaultPath to the configs folder.
var DefaultPath = "./config/"

// DefaultVariablesPath enviroments path.
var DefaultVariablesPath = "./config/environments/"

// Env type
type Env int

const (
	// Development enviroment
	Development Env = iota
	// Production enviroment
	Production
	// Test enviroment
	Test
)

var envStrings = map[string]Env{
	"development": Development,
	"production":  Production,
	"test":        Test,
}
var envNames = [...]string{
	Production:  "production",
	Development: "development",
	Test:        "test",
}

// to store all config files
var globalData map[string]interface{}

// New return Config instance
func New(env string) map[string]interface{} {
	// first check if the env supported.
	if _, ok := envStrings[env]; !ok {
		kog.Panic(errInvalidEnv)
	}

	// load configs
	configs := load(env)

	// return configs
	if configs == nil {
		// panic here
		return globalData
	}
	// panic if there an error.
	kog.Panic(configs)

	return globalData
}

// Load config file data
func load(env string) error {
	var buf bytes.Buffer

	// to store all config files
	var files []string

	// first get all config files
	configs, err := filepath.Glob(DefaultPath + "*.yaml")
	if err != nil {
		return err
	}
	files = append(files, configs...)

	// add env file to the config files.
	// this will change any config value.
	if env != "" {
		files = append(files, DefaultVariablesPath+env+".yaml")
	}

	// loop throw all
	for _, config := range files {
		// read file
		data, err := ioutil.ReadFile(config)
		if err != nil {
			return err
		}

		buf.Write(data)
		buf.WriteString("\n") // add new line to the end of the file
	}

	// Unmarshal toml data
	yaml.Unmarshal(buf.Bytes(), &globalData)

	return nil
}

// Get ...
func Get(key string) interface{} {
	return globalData[key]
}

// GetDefault return value by key, if the value empty set a default value.
func GetDefault(key string, def interface{}) interface{} {
	val := globalData[key]
	if val == nil {
		return def
	}

	return val
}

// GetString - return config as string type.
func GetString(key string) string {
	if globalData[key] == nil {
		return ""
	}

	return globalData[key].(string)
}

// GetInt - return config as int type.
func GetInt(key string) int {
	if globalData[key] == nil {
		return 0
	}

	return globalData[key].(int)
}

// GetBool - return config as int type.
func GetBool(key string) bool {
	if globalData[key] == nil {
		return false
	}

	return globalData[key].(bool)
}
