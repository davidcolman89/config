package config

import (
	"reflect"
	"runtime"
	"path"
	"github.com/spf13/viper"
)

var configPath []string

func getConfigNames (c interface{}) {
	rv := reflect.ValueOf(c).Elem()
	for i := 0; i < rv.NumField(); i++ {
		varName := rv.Type().Field(i).Name
		viper.BindEnv(varName)
	}
}

func GetConfigPath() []string{
	return configPath
}

func AddConfigPath(path string) {
	configPath = append(configPath,path)
}

func SetConfig(prefix string, c interface{}, configFile string) {
	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()
	viper.BindEnv("Message")
	getConfigNames(c)
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	viper.SetConfigName(configFile)
	// name of config file (without extension)
	for _,cPath := range configPath {
		viper.AddConfigPath(cPath)
	}
	viper.AddConfigPath(".")
	// optionally look for config in the working directory
	viper.AddConfigPath(path.Dir(filename))
	err := viper.ReadInConfig()
	// Find and read the config file
	if err != nil { // Handle errors reading the config file

	}
	viper.Unmarshal(c)
}
