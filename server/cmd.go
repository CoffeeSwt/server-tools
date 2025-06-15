package server

import (
	"os"
	"os/exec"
	"server-tools/config"
	"server-tools/logger"
	"strconv"
	"time"

	"go.uber.org/zap"
)

func StartServer() {
	l := logger.GetLogger()
	cfg := config.GetConfig()
	l.Info("🚀 启动 " + cfg.ServerName + "服务器监控程序...")

	// 获取服务器路径和启动参数(只需获取一次)
	slp := GetServerLaunchParameters()
	paths, err := GetDayZPaths()
	if err != nil {
		l.Error("获取 DayZ 路径失败", zap.Error(err))
		return
	}

	args := []string{
		"-port=" + strconv.Itoa(slp.Port),
		"-mission=" + slp.Mission,
		"-profiles=" + slp.Profiles,
		"-mod=" + slp.ClientMods,
		"-serverMod=" + slp.ServerMods,
		"-config=" + slp.Config,
		"-dologs",
		"-adminlog",
		"-netlog",
		"-freezecheck",
	}

	for {
		l.Info("正在启动服务器...")
		cmd := exec.Command(paths.DayZServerExecutable, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// 启动进程
		if err := cmd.Start(); err != nil {
			l.Error("DayZ 服务器启动失败", zap.Error(err))
			return
		}

		l.Info("✅ "+cfg.ServerName+"服务器已启动", zap.Int("PID", cmd.Process.Pid))

		// 等待进程结束
		if err := cmd.Wait(); err != nil {
			l.Error("服务器进程异常退出", zap.Error(err))
		}

		l.Info("⏳ 等待 3 秒后重启服务器...")
		time.Sleep(3 * time.Second)
	}
}
