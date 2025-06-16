package config

import (
	"fmt"
	"os"
	"server-tools/logger"
	"sync"
	"time"

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
		fmt.Println("3 秒后自动退出...")
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}

	err = viper.Unmarshal(&appConfig)
	if err != nil {
		logger.GetLogger().Fatal("❌ 配置解析失败: ", zap.Error(err))
		fmt.Println("3 秒后自动退出...")
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}
}

func GetConfig() *Config {
	once.Do(func() {
		initConfig()
	})
	return &appConfig
}
