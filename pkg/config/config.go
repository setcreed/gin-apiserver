/*
Copyright 2023 The gin-apiserver Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var cfg Config

type flagOpt struct {
	optName         string
	optDefaultValue interface{}
	optUsage        string
}

var flagsOpts = []flagOpt{
	{
		optName:         FLAG_KEY_SERVER_HOST,
		optDefaultValue: "0.0.0.0",
		optUsage:        "server listen host",
	},
	{
		optName:         FLAG_KEY_SERVER_PORT,
		optDefaultValue: 8081,
		optUsage:        "server listen port",
	},
	{
		optName:         FLAG_KEY_GIN_MODE,
		optDefaultValue: "debug",
		optUsage:        "gin mode",
	},
	{
		optName:         FLAG_KEY_LOG_LEVEL,
		optDefaultValue: "info",
		optUsage:        "log level",
	},
}

func init() {
	for _, opt := range flagsOpts {
		viper.SetDefault(opt.optName, opt.optDefaultValue)
	}
	viper.SetConfigType("yaml")
	viper.SetConfigName(SERVICE_NAME)
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", SERVICE_NAME))
	viper.AddConfigPath("./etc/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error config file: %s \n", err)
	}
	viper.SetEnvPrefix(SERVICE_NAME)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	for _, opt := range flagsOpts {
		switch opt.optDefaultValue.(type) {
		case int:
			pflag.Int(opt.optName, opt.optDefaultValue.(int), opt.optUsage)
		case string:
			pflag.String(opt.optName, opt.optDefaultValue.(string), opt.optUsage)
		case bool:
			pflag.Bool(opt.optName, opt.optDefaultValue.(bool), opt.optUsage)
		default:
			continue
		}
	}

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	err = viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to decode into struct, %v\n", err)
	}
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetConfig() *Config {
	return &cfg
}
