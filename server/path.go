package server

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"server-tools/logger"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sys/windows/registry"
)

type DayZPaths struct {
	SteamPath            string `steam:"true"`                                                        // 主 Steam 路径（注册表）
	DayZPath             string `path:"steamapps/common/DayZ"`                                        // DayZ 客户端路径
	DayZServerPath       string `path:"steamapps/common/DayZServer"`                                  // DayZ 服务端路径
	DayZServerExecutable string `path:"steamapps/common/DayZServer/DayZServer_x64.exe" isDir:"false"` // 服务端可执行文件
	MissionsPath         string `path:"mpmissions" create:"true"`                                     // 当前工作目录
	ModsPath             string `path:"steamapps/common/DayZ/!Workshop"`                              // 客户端模组目录
	ProfilePath          string `path:"profiles" create:"true"`                                       // 配置目录
	CfgPath              string `path:"serverCfgs" create:"true"`                                     // 服务端配置目录
	KeysPath             string `path:"steamapps/common/DayZServer/keys"`                             // 密钥目录
}

var (
	paths         *DayZPaths
	pathsOnce     sync.Once
	hasCreatedDir bool
)

func GetDayZPaths() (*DayZPaths, error) {
	var err error
	pathsOnce.Do(func() {
		paths, err = buildPaths()
		if err != nil {
			logger.GetLogger().Error("获取 DayZ 路径失败", zap.Error(err))
			fmt.Println("3 秒后自动退出...")
			time.Sleep(3 * time.Second)
			os.Exit(1)
			return
		}
		if hasCreatedDir {
			logger.GetLogger().Info("首次运行，部分目录已自动创建，请重新运行程序。")
			fmt.Println("3 秒后自动退出...")
			time.Sleep(3 * time.Second)
			os.Exit(0)
		}
	})
	return paths, err
}

// 获取 Steam 所有库路径（包括主库和附加库）
func getSteamLibraries() ([]string, error) {
	var libs []string

	// 注册表读取主路径
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Valve\Steam`, registry.QUERY_VALUE)
	if err != nil {
		return nil, fmt.Errorf("打开注册表失败: %w", err)
	}
	defer key.Close()

	mainPath, _, err := key.GetStringValue("SteamPath")
	if err != nil {
		return nil, fmt.Errorf("读取 SteamPath 失败: %w", err)
	}
	libs = append(libs, mainPath)

	// libraryfolders.vdf 文件路径
	vdfPath := filepath.Join(mainPath, "steamapps", "libraryfolders.vdf")
	content, err := os.ReadFile(vdfPath)
	if err != nil {
		return libs, nil
	}
	data := string(content)

	// 1. 解析新版 JSON 格式
	type Folder struct {
		Path string `json:"path"`
	}
	type LibraryFolders struct {
		LibraryFolders map[string]Folder `json:"libraryfolders"`
	}
	var parsed LibraryFolders
	if err := json.Unmarshal(content, &parsed); err == nil {
		for _, folder := range parsed.LibraryFolders {
			if folder.Path != "" {
				libs = append(libs, folder.Path)
			}
		}
		return uniqueStrings(libs), nil
	}

	// 2. fallback 到旧格式正则
	re := regexp.MustCompile(`"path"\s+"([^"]+)"`)
	matches := re.FindAllStringSubmatch(data, -1)
	for _, match := range matches {
		if len(match) > 1 {
			path := strings.ReplaceAll(match[1], `\\`, `\`)
			libs = append(libs, path)
		}
	}

	return uniqueStrings(libs), nil
}

func uniqueStrings(arr []string) []string {
	seen := make(map[string]struct{})
	var result []string
	for _, v := range arr {
		if _, ok := seen[v]; !ok {
			result = append(result, v)
			seen[v] = struct{}{}
		}
	}
	return result
}

// 构建 DayZPaths 结构体并自动判断路径
func buildPaths() (*DayZPaths, error) {
	steamLibraries, err := getSteamLibraries()
	if err != nil || len(steamLibraries) == 0 {
		return nil, fmt.Errorf("无法获取 Steam 库路径: %w", err)
	}
	steamPath := steamLibraries[0]
	cwd, _ := os.Getwd()

	result := &DayZPaths{}
	val := reflect.ValueOf(result).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tagPath := field.Tag.Get("path")
		isDir := field.Tag.Get("isDir") != "false" // 默认为 true
		create := field.Tag.Get("create") == "true"
		isSteam := field.Tag.Get("steam") == "true"

		switch {
		case isSteam:
			val.Field(i).SetString(steamPath)

		case tagPath != "":
			var fullPath string

			// 绝对路径直接用
			if filepath.IsAbs(tagPath) || filepath.VolumeName(tagPath) != "" {
				fullPath = tagPath

			} else if strings.HasPrefix(tagPath, "steamapps") {
				// 遍历所有 Steam 库路径
				for _, lib := range steamLibraries {
					p := filepath.Join(lib, filepath.FromSlash(tagPath))
					if _, err := os.Stat(p); err == nil || (create && isDir) {
						fullPath = p
						break
					}
				}
				if fullPath == "" {
					return nil, fmt.Errorf("未在任何 Steam 库中找到: %s", tagPath)
				}

			} else {
				fullPath = filepath.Join(cwd, filepath.FromSlash(tagPath))
			}

			if err := ensurePath(fullPath, isDir, create); err != nil {
				return nil, fmt.Errorf("%s 初始化失败: %w", field.Name, err)
			}
			val.Field(i).SetString(fullPath)
		}
	}

	return result, nil
}

// 检查路径是否存在；可选创建
func ensurePath(path string, isDir, create bool) error {
	info, err := os.Stat(path)
	if err == nil {
		if isDir && !info.IsDir() {
			return fmt.Errorf("%s 存在但不是目录", path)
		}
		if !isDir && info.IsDir() {
			return fmt.Errorf("%s 存在但是目录", path)
		}
		return nil
	}
	if os.IsNotExist(err) && create && isDir {
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
		hasCreatedDir = true
		return nil
	}
	return err
}
