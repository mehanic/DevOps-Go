package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// Config will hold the config we
// capture from a json file and env vars
type Config struct {
	Version string `json:"version" required:"true"`
	IsSafe  bool   `json:"is_safe" default:"true"`
	Secret  string `json:"secret"`
}

func main() {
	var err error

	// create a temporary file to hold
	// an example json file
	tf, err := os.CreateTemp("", "tmp")
	if err != nil {
		panic(err)
	}
	defer tf.Close()
	defer os.Remove(tf.Name())

	// create a json file to hold
	// our secrets
	secrets := `{
        "secret": "so so secret"
    }`

	if _, err = tf.Write(bytes.NewBufferString(secrets).Bytes()); err != nil {
		panic(err)
	}

	// We can easily set environment variables
	// as needed
	if err = os.Setenv("EXAMPLE_VERSION", "1.0.0"); err != nil {
		panic(err)
	}
	if err = os.Setenv("EXAMPLE_ISSAFE", "false"); err != nil {
		panic(err)
	}

	c := Config{}
	if err = LoadConfig(tf.Name(), "EXAMPLE", &c); err != nil {
		panic(err)
	}

	fmt.Println("secrets file contains =", secrets)

	// We can also read them
	fmt.Println("EXAMPLE_VERSION =", os.Getenv("EXAMPLE_VERSION"))
	fmt.Println("EXAMPLE_ISSAFE =", os.Getenv("EXAMPLE_ISSAFE"))

	// The final config is a mix of json and environment
	// variables
	fmt.Printf("Final Config: %#v\n", c)

}

func LoadConfig(path, envPrefix string, config interface{}) error {
	if path != "" {
		err := LoadFile(path, config)
		if err != nil {
			return errors.Wrap(err, "error loading config from file")
		}
	}
	err := envconfig.Process(envPrefix, config)
	return errors.Wrap(err, "error loading config from env")
}

// LoadFile unmarshalls a json file into a config struct
func LoadFile(path string, config interface{}) error {
	configFile, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "failed to read config file")
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	if err = decoder.Decode(config); err != nil {
		return errors.Wrap(err, "failed to decode config file")
	}
	return nil
}

// secrets file contains = {
//         "secret": "so so secret"
//     }
// EXAMPLE_VERSION = 1.0.0
// EXAMPLE_ISSAFE = false
// Final Config: main.Config{Version:"1.0.0", IsSafe:false, Secret:"so so secret"}
