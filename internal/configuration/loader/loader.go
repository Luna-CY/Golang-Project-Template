package loader

import (
	"errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/runtime"
	"github.com/spf13/viper"
)

func LoadConfig(path string, dst any) error {
	if "" == path {
		path = "config"
	}

	// Load main config
	if err := loadConfig(path, "main", dst); nil != err {
		return err
	}

	// Load config for special environment
	var env = runtime.GetEnvironment()
	if err := loadConfig(path, env, dst); nil != err {
		return err
	}

	// Load custom config
	if err := loadConfig(path, "custom", dst); nil != err {
		return err
	}

	return nil
}

func loadConfig(path string, name string, dst any) error {
	var v = viper.New()
	v.AddConfigPath(path)
	v.SetConfigType("toml")
	v.SetConfigName(name)

	if err := v.ReadInConfig(); nil != err {
		var configFileNotFoundError viper.ConfigFileNotFoundError

		if !errors.As(err, &configFileNotFoundError) {
			return err
		}
	}

	if err := v.Unmarshal(&dst); nil != err {
		return err
	}

	return nil
}
