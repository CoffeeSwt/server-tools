package config

import (
	"server-tools/logger"
	"sync"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	Port       int      `mapstructure:"port"`
	ServerName string   `mapstructure:"server_name"`
	Mission    string   `mapstructure:"mission"`
	ClientMods []string `mapstructure:"client_mods"`
	ServerMods []string `mapstructure:"server_mods"`
}

var (
	appConfig Config
	once      sync.Once
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	// viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		logger.GetLogger().Fatal("❌ 读取配置文件失败: ", zap.Error(err))
	}

	err = viper.Unmarshal(&appConfig)
	if err != nil {
		logger.GetLogger().Fatal("❌ 配置解析失败: ", zap.Error(err))
	}
}

func GetConfig() *Config {
	once.Do(func() {
		initConfig()
	})
	return &appConfig
}
