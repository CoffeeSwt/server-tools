package config

import (
	"server-tools/defaultCfg"
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
		logger.GetLogger().Error("❌ 读取配置文件失败: ", zap.Error(err))
		logger.GetLogger().Info("正在使用默认的服务器配置...")
		defaultCfg.SetUseDefaultConfig(true)
		return
	}

	err = viper.Unmarshal(&appConfig)
	if err != nil {
		logger.GetLogger().Error("❌ 配置解析失败: ", zap.Error(err))
		logger.GetLogger().Info("正在使用默认的服务器配置...")
		defaultCfg.SetUseDefaultConfig(true)
	}
}

func GetConfig() *Config {
	once.Do(func() {
		initConfig()
	})
	return &appConfig
}

func UseDefaultConfig() {
	appConfig = Config{
		Port:       2302,
		ServerName: "MyDayZServer",
		Mission:    "default.chernarusplus", //这个是唯一值
		ClientMods: []string{},
		ServerMods: []string{},
	}
}
