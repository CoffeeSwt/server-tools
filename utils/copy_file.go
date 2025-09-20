package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

// CopyFile 安全复制文件：如果目标已存在，则跳过（返回特定错误）
func CopyKeyFile(src, dst, name string) error {
	// 如果目标文件已存在，直接返回（不覆盖）
	if _, err := os.Stat(dst); err == nil {
		// return errors.New(name + "模组的" + filepath.Base(dst) + "key文件" + "已存在，跳过复制")
		return nil
	}

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	// 保留原权限
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	return os.Chmod(dst, srcInfo.Mode())
}

// CopyFolder 递归拷贝 src 目录下的所有文件到 dst 目录，相同文件会被覆盖
// blacklistPatterns 是黑名单正则列表，匹配的文件/目录会被跳过
func CopyFolder(src string, dst string, blacklistPatterns []string) error {
	// 编译黑名单正则
	var blacklist []*regexp.Regexp
	for _, pattern := range blacklistPatterns {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return fmt.Errorf("黑名单正则编译失败: %v", err)
		}
		blacklist = append(blacklist, re)
	}

	// 确保目标文件夹存在
	err := os.MkdirAll(dst, os.ModePerm)
	if err != nil {
		return fmt.Errorf("创建目标目录失败: %v", err)
	}

	// 遍历源目录
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 构建目标路径
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)

		// 检查黑名单
		for _, re := range blacklist {
			if re.MatchString(relPath) {
				if info.IsDir() {
					// 如果是目录，跳过整个目录
					return filepath.SkipDir
				}
				// 如果是文件，跳过文件
				return nil
			}
		}

		if info.IsDir() {
			// 创建子目录
			return os.MkdirAll(dstPath, os.ModePerm)
		} else {
			// 拷贝文件
			return copyFile(path, dstPath)
		}
	})
}

// copyFile 将单个文件从 src 拷贝到 dst
func copyFile(srcFile string, dstFile string) error {
	srcF, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer srcF.Close()

	dstF, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer dstF.Close()

	_, err = io.Copy(dstF, srcF)
	if err != nil {
		return err
	}

	// 保持文件权限
	info, err := os.Stat(srcFile)
	if err != nil {
		return err
	}
	return os.Chmod(dstFile, info.Mode())
}

// Spinner 控制台旋转动画
func Spinner(stopChan <-chan struct{}) {
	chars := []rune{'|', '/', '-', '\\'}
	i := 0
	for {
		select {
		case <-stopChan:
			fmt.Print("\r") // 清空动画
			return
		default:
			fmt.Printf("\r%c 正在拷贝官方文件到任务目录，请稍等...", chars[i%len(chars)])
			time.Sleep(100 * time.Millisecond)
			i++
		}
	}
}
