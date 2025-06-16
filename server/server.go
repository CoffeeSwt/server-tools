package server

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"server-tools/config"
	"server-tools/logger"
	"server-tools/utils"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
)

var (
	serverLaunchParameters *ServerLaunchParameters
	serverLaunchOnce       sync.Once
)

type ServerLaunchParameters struct {
	Port       int
	ClientMods string
	ServerMods string
	Mission    string
	Profiles   string
	Config     string
}

// 获取启动项
func GetServerLaunchParameters() *ServerLaunchParameters {
	serverLaunchOnce.Do(func() {
		paths, err := GetDayZPaths()
		if err != nil {
			logger.GetLogger().Error("获取 DayZ 路径失败", zap.Error(err))
			fmt.Println("3 秒后自动退出...")
			time.Sleep(3 * time.Second)
			os.Exit(1)
			// 这里直接退出程序，因为无法继续进行后续操作
			return
		}
		cfg := config.GetConfig()

		port := cfg.Port

		// 统一处理模组路径和 keys 复制
		clientMods := resolveModsWithKeys(cfg.ClientMods, paths.ModsPath)
		serverMods := resolveModsWithKeys(cfg.ServerMods, paths.ModsPath)

		// 构造任务路径
		missionPath := filepath.Join(paths.MissionsPath, cfg.Mission)

		// 构造配置文件路径
		profilePath := filepath.Join(paths.ProfilePath, cfg.ServerName)

		serverLaunchParameters = &ServerLaunchParameters{
			Port:       port,
			ClientMods: strings.Join(clientMods, ";"),
			ServerMods: strings.Join(serverMods, ";"),
			Mission:    missionPath,
			Profiles:   profilePath,
			Config:     filepath.Join(paths.CfgPath, cfg.ServerName+".cfg"),
		}
	})

	return serverLaunchParameters
}

// 复制模组中的 keys 到 DayZServer 的 keys 目录
func copyKeys(modOriginPath string) error {
	modName := filepath.Base(modOriginPath)
	paths, _ := GetDayZPaths()

	// 所有可能的 key 文件夹名（不区分大小写）
	keyDirs := []string{"keys", "Keys", "key", "Key"}

	var foundKeyDir string

	// 遍历尝试找到存在的 key 目录
	for _, dirName := range keyDirs {
		tryPath := filepath.Join(modOriginPath, dirName)
		info, err := os.Stat(tryPath)
		if err == nil && info.IsDir() {
			foundKeyDir = tryPath
			break
		}
	}

	if foundKeyDir == "" {
		// 没找到 key 目录，直接返回
		// return errors.New("未找到模组 " + modName + " 的 keys 目录")
		return errors.New("没有找到" + modName + "模组的的 keys 目录")
	}

	// 遍历 key 目录下的所有文件，复制到目标目录
	entries, err := os.ReadDir(foundKeyDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		srcFile := filepath.Join(foundKeyDir, entry.Name())
		dstFile := filepath.Join(paths.KeysPath, entry.Name())

		if err := utils.CopyKeyFile(srcFile, dstFile, modName); err != nil {
			return err
		}
	}

	return nil
}

func resolveModsWithKeys(mods []string, basePath string) []string {
	var result []string
	for _, mod := range mods {
		modPath := filepath.Join(basePath, mod)
		if _, err := os.Stat(modPath); os.IsNotExist(err) {
			logger.GetLogger().Warn("模组路径不存在", zap.String("modPath", modPath))
			continue
		}
		result = append(result, modPath)
		if err := copyKeys(modPath); err != nil {
			logger.GetLogger().Info("复制模组 keys 失败", zap.String("mod", mod), zap.Error(err))
		}
	}
	return result
}
