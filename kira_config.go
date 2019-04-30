package kira

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/go-kira/config"
)

// getEnv for set the framework environment.
func getEnv() string {
	// Get the environment from .kira_env file.
	if _, err := os.Stat("./.kira_env"); !os.IsNotExist(err) {
		kiraEnv, err := ioutil.ReadFile("./.kira_env")
		if err != nil {
			log.Panic(err)
		}
		return strings.TrimSpace(string(kiraEnv))
	}

	// Get the environment from system variable
	osEnv := os.Getenv("KIRA_ENV")
	if osEnv == "" {
		return "development"
	}
	return osEnv
}

func getConfig() *config.Config {
	var files = []string{"./testdata/kira.conf", "./kira.conf", "./config/kira.conf"}
	var env = fmt.Sprintf("./kira.conf.%s", getEnv())

	if _, err := os.Stat(env); !os.IsNotExist(err) {
		files = append(files, env)
	}

	return config.NewFromFile(files...)
}
