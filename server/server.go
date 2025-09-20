package server

import (
	"errors"
	"os"
	"path/filepath"
	"server-tools/config"
	"server-tools/defaultCfg"
	"server-tools/logger"
	"server-tools/utils"
	"strings"
	"sync"

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
		initServerLaunchParameters()
	})
	return serverLaunchParameters
}

func initServerLaunchParameters() {
	dpaths := GetDayZPaths()
	cfg := config.GetConfig()
	if defaultCfg.IsUseDefaultConfig() {
		config.UseDefaultConfig()
	}

	port := cfg.Port

	// 统一处理模组路径和 keys 复制
	clientMods := resolveModsWithKeys(cfg.ClientMods, dpaths.ModsPath)
	serverMods := resolveModsWithKeys(cfg.ServerMods, dpaths.ModsPath)

	// 构造任务路径
	// missionPath := filepath.Join()
	missionPath := buildMissionPath(dpaths, cfg.Mission)

	// 构造配置文件路径
	profilePath := filepath.Join(dpaths.ProfilePath, cfg.ServerName)

	serverLaunchParameters = &ServerLaunchParameters{
		Port:       port,
		ClientMods: strings.Join(clientMods, ";"),
		ServerMods: strings.Join(serverMods, ";"),
		Mission:    missionPath,
		Profiles:   profilePath,
		Config:     filepath.Join(dpaths.CfgPath, cfg.ServerName+".cfg"),
	}
}

// 复制模组中的 keys 到 DayZServer 的 keys 目录
func copyKeys(modOriginPath string) error {
	modName := filepath.Base(modOriginPath)
	paths := GetDayZPaths()

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
		return errors.New("没有找到" + modName + "模组的的 keys 目录，订阅模组后请先打开一次DayZ客户端，进行模组下载与更新")
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
			return errors.New("复制 " + modName + " 模组的 key 文件 " + entry.Name() + " 失败: " + err.Error() + "，请检查DayZServer的keys目录是否存在，以及权限是否正确")
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
