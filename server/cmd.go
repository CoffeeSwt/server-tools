package server

import (
	"fmt"
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
	l.Info("如果遇到小白框自动关闭又自动重启，请检查你的mpmissions目录下是否有任务文件，以及任务文件是否合法，不合法会导致服务器无法启动")
	l.Info("解决方案:去 https://github.com/BohemiaInteractive/DayZ-Central-Economy 下载官方的任务文件放到mpmissions目录下")
	l.Info("请使用小白使用这三个任务文件夹: dayzOffline.chernarusplus / dayzOffline.enoch / dayzOffline.sakhal（需要购买DLC），注意前缀和文件夹名，不要复制错了")

	// 获取服务器路径和启动参数(只需获取一次)
	slp := GetServerLaunchParameters()
	paths := GetDayZPaths()

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
		l.Info("如果遇到BattleEye相关报错，请确保防火墙，杀毒等工具已经关闭，代码完全开源无毒，请放心使用")
		cmd := exec.Command(paths.DayZServerExecutable, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// 启动进程
		if err := cmd.Start(); err != nil {
			l.Error("DayZ 服务器启动失败", zap.Error(err))
			fmt.Println("3 秒后自动退出...")
			os.Exit(1)
			return
		}

		l.Info("✅ "+cfg.ServerName+"服务器已启动", zap.Int("PID", cmd.Process.Pid))

		// 等待进程结束
		if err := cmd.Wait(); err != nil {
			l.Error("服务器进程关闭", zap.Error(err))
		}

		l.Info("⏳ 等待 3 秒后重启服务器...")
		time.Sleep(3 * time.Second)
	}
}
