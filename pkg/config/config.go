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
