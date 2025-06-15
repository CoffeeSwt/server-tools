package server

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"server-tools/logger"
	"strings"
	"sync"

	"go.uber.org/zap"
	"golang.org/x/sys/windows/registry"
)

type DayZPaths struct {
	SteamPath            string `steam:"true"`                                                        // 从注册表获取
	DayZPath             string `path:"steamapps/common/DayZ"`                                        // DayZ 客户端路径
	DayZServerPath       string `path:"steamapps/common/DayZServer"`                                  // DayZ 服务端路径
	DayZServerExecutable string `path:"steamapps/common/DayZServer/DayZServer_x64.exe" isDir:"false"` // 可执行文件
	MissionsPath         string `path:"mpmissions" create:"true"`                                     // 任务目录（当前工作目录）
	ModsPath             string `path:"steamapps/common/DayZ/!Workshop"`                              // 模组目录
	ProfilePath          string `path:"profiles" create:"true"`                                       // 配置文件目录
	CfgPath              string `path:"serverCfgs" create:"true"`                                     // 服务端配置目录
	KeysPath             string `path:"steamapps/common/DayZServer/keys"`                             // 密钥目录
}

var (
	paths         *DayZPaths
	pathsOnce     sync.Once
	hasCreatedDir bool // 全局标志，表示是否创建过目录
)

// 单例访问
func GetDayZPaths() (*DayZPaths, error) {
	var err error
	pathsOnce.Do(func() {
		paths, err = buildPaths()
		if err != nil {
			logger.GetLogger().Error("获取 DayZ 路径失败", zap.Error(err))
			// 这里直接退出程序，因为无法继续进行后续操作
			os.Exit(1)
			return
		}
		// 如果创建过目录，则提示并退出程序
		if hasCreatedDir {
			logger.GetLogger().Info("程序首次启动，缺少部分目录，已自动创建，请重新运行程序。")
			os.Exit(0)
		}
	})
	return paths, err
}

// 读取注册表，获取 Steam 安装路径
func getSteamInstallPath() (string, error) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\WOW6432Node\Valve\Steam`, registry.QUERY_VALUE)
	if err != nil {
		key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Valve\Steam`, registry.QUERY_VALUE)
		if err != nil {
			return "", fmt.Errorf("failed to open registry: %w", err)
		}
	}
	defer key.Close()

	path, _, err := key.GetStringValue("InstallPath")
	if err != nil {
		return "", fmt.Errorf("failed to read InstallPath: %w", err)
	}
	return path, nil
}

// 使用标签自动构造路径
func buildPaths() (*DayZPaths, error) {
	steamPath, err := getSteamInstallPath()
	if err != nil {
		return nil, err
	}
	cwd, _ := os.Getwd()

	result := &DayZPaths{}
	val := reflect.ValueOf(result).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tagPath := field.Tag.Get("path")
		isDir := field.Tag.Get("isDir") != "false" // 默认为目录
		create := field.Tag.Get("create") == "true"
		isSteam := field.Tag.Get("steam") == "true"

		switch {
		case isSteam:
			val.Field(i).SetString(steamPath)

		case tagPath != "":
			var fullPath string
			if filepath.IsAbs(tagPath) || filepath.VolumeName(tagPath) != "" {
				fullPath = tagPath
			} else if strings.HasPrefix(tagPath, "steamapps") {
				fullPath = filepath.Join(steamPath, filepath.FromSlash(tagPath))
			} else {
				fullPath = filepath.Join(cwd, filepath.FromSlash(tagPath))
			}

			if err := ensurePath(fullPath, isDir, create); err != nil {
				return nil, fmt.Errorf("%s init failed: %w", field.Name, err)
			}
			val.Field(i).SetString(fullPath)
		}
	}
	return result, nil
}

// 检查路径是否存在，可选创建
func ensurePath(path string, isDir, create bool) error {
	info, err := os.Stat(path)
	if err == nil {
		if isDir && !info.IsDir() {
			return fmt.Errorf("%s exists but is not a directory", path)
		}
		if !isDir && info.IsDir() {
			return fmt.Errorf("%s exists but is a directory", path)
		}
		return nil
	}
	if os.IsNotExist(err) && create && isDir {
		// 创建目录
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
		// 标记已经创建过目录
		hasCreatedDir = true
		return nil
	}
	return err
}
