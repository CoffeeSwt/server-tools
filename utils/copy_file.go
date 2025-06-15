package utils

import (
	"io"
	"os"
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
