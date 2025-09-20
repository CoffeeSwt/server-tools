package server

import (
	"fmt"
	"os"
	"path/filepath"
	"server-tools/defaultCfg"
	"server-tools/logger"
	"server-tools/utils"

	"go.uber.org/zap/zapcore"
)

func buildMissionPath(dayzPath *DayZPaths, mission string) string {
	missionPath := filepath.Join(dayzPath.MissionsPath, mission)
	if defaultCfg.IsUseDefaultConfig() {
		missionPath = filepath.Join(dayzPath.MissionsPath, "default.chernarusplus")
		logger.GetLogger().Info("正在使用默认任务: default.chernarusplus")

		// 启动动画
		stopChan := make(chan struct{})
		go utils.Spinner(stopChan)

		// 执行拷贝
		err := utils.CopyFolder(
			filepath.Join(dayzPath.MissionsDefaultPath, "dayzOffline.chernarusplus"),
			missionPath,
			[]string{`^storage_.*`},
		)

		// 停止动画
		close(stopChan)
		fmt.Println() // 换行

		if err != nil {
			logger.GetLogger().Error("拷贝失败", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
		} else {
			logger.GetLogger().Info("拷贝完成")
		}

		logger.GetLogger().Info("正在创建默认服务器Cfg文件...")
		CreateDefaultCfg()

	}
	return missionPath
}

func CreateDefaultCfg() {
	dayzPath := GetDayZPaths()
	cfgPath := filepath.Join(dayzPath.CfgPath, "MyDayZServer.cfg")
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		err := os.WriteFile(cfgPath, []byte(getDefaultCfgTemplate()), 0644)
		if err != nil {
			logger.GetLogger().Error("创建默认Cfg文件失败", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
		} else {
			logger.GetLogger().Info("默认Cfg文件创建成功，路径: " + cfgPath)
		}
	} else {
		logger.GetLogger().Info("默认Cfg文件已存在，路径: " + cfgPath)
	}
}

func getDefaultCfgTemplate() string {
	return `
hostname = "默认切尔那鲁斯 by coffeesw";     // 服务器名称
password = "";              // 连接服务器密码
passwordAdmin = "";         // 管理员密码

description = "";          // 服务器描述，将显示在客户端服务器浏览器中

enableWhitelist = 0;       // 启用/禁用白名单 (0-1)

maxPlayers = 60;          // 最大玩家数

verifySignatures = 2;     // 验证 .pbo 文件签名 (仅支持值 2)
forceSameBuild = 1;       // 强制客户端版本与服务器版本一致 (0-1)

disableVoN = 0;          // 启用/禁用语音 (0-1)
vonCodecQuality = 20;    // 语音编码质量，越高越好 (0-30)

shardId = "123abc";      // 私人服务器的六位字母数字标识符

disable3rdPerson=0;      // 切换第三人称视角 (0-1)
disableCrosshair=0;     // 切换准星显示 (0-1)

disablePersonalLight = 1;   // 禁用所有客户端的个人光源
lightingConfig = 0;         // 0 为较亮的夜晚设置，1 为较暗的夜晚设置

serverTime="SystemTime";    // 服务器初始时间，"SystemTime"表示使用本机时间
serverTimeAcceleration=8;  // 时间加速倍率 (0-24)
serverNightTimeAcceleration=12;  // 夜间时间加速倍率 (0.1-64)
serverTimePersistent=1;     // 时间持久化 (0-1)

guaranteedUpdates=1;        // 游戏服务器通信协议 (仅使用 1)

loginQueueConcurrentPlayers=5;  // 同时处理的登录玩家数量
loginQueueMaxPlayers=500;       // 登录队列最大玩家数

instanceId = 1;             // DayZ 服务器实例 ID

storageAutoFix = 1;         // 检查并修复损坏的持久化文件 (0-1)

class Missions
{
    class DayZ
    {
        template="dayzOffline.chernarusplus"; // 服务器启动时加载的任务
    };
};
	
	`
}
