package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

type configuration struct {
	port         int
	databaseFile string
}

type jsonConfiguration struct {
	Port         int    `json:"port"`
	DatabaseFile string `json:"databaseFile"`
}

type Configuration interface {
	Port() int
	DatabaseFile() string
}

func (c configuration) Port() int {
	return c.port
}
func (c configuration) DatabaseFile() string {
	return c.databaseFile
}

type ConfigurationBuilder struct {
	err    error
	config configuration
}

func NewConfigurationBuilder() *ConfigurationBuilder {
	b := ConfigurationBuilder{err: nil, config: configuration{port: 8080, databaseFile: "users.db"}}
	return &b
}

func (b *ConfigurationBuilder) AddJsonFile(filename string) *ConfigurationBuilder {
	if b.err != nil {
		return b
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		b.err = err
		return b
	}

	var fileConfig jsonConfiguration
	err = json.Unmarshal(content, &fileConfig)
	if err != nil {
		b.err = err
		return b
	}

	if fileConfig.Port != 0 {
		b.config.port = fileConfig.Port
	}
	if fileConfig.DatabaseFile != "" {
		b.config.databaseFile = fileConfig.DatabaseFile
	}
	return b
}

//goland:noinspection ALL,SpellCheckingInspection,SpellCheckingInspection
func (b *ConfigurationBuilder) AddEnvironment() *ConfigurationBuilder {
	if b.err != nil {
		return b
	}
	port := 0
	databaseFile := ""
	var err error

	value, present := os.LookupEnv("USERAPI:PORT")
	if present {
		port, err = strconv.Atoi(value)
		if err != nil {
			port = 0
		}
	}

	value, present = os.LookupEnv("USERAPI:DATABASEFILE")
	if present {
		databaseFile = value
	}

	if port != 0 {
		b.config.port = port
	}

	if databaseFile != "" {
		b.config.databaseFile = databaseFile
	}

	return b
}

func (b *ConfigurationBuilder) Build() Configuration {
	var c Configuration = b.config
	b.config = configuration{port: 8080, databaseFile: "users.db"}
	return c
}
