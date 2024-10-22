package configs

import (
	"path"
	"path/filepath"
	"runtime"

	"github.com/conflux-fans/cmctool/pkg/cfgutil"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var (
	_config Config
)

func Init() {
	// 获取当前文件的绝对路径
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("Error: Unable to get current file path")
	}
	logrus.Info("Current file path:", file)
	cwd := filepath.Dir(file) // get curent path

	viper.SetConfigName("config")                // name of config file (without extension)
	viper.SetConfigType("yaml")                  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(cwd)                     // optionally look for config in the working directory
	viper.AddConfigPath(path.Join(cwd, ".."))    // optionally look for config in the working directory
	viper.AddConfigPath(path.Join(cwd, "../..")) // optionally look for config in the working directory
	_config = *cfgutil.MustLoadViper[Config]()
}

func InitByFile(configPath string) {
	viper.SetConfigFile(configPath)
	_config = *cfgutil.MustLoadViper[Config]()
}

func Get() *Config {
	return &_config
}
