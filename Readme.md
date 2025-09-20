# Dayz Server Tools

[English](#dayz-server-tools) | [中文](#dayz-服务器工具)

# Dayz Server Tools

Easy management of Dayz's servers

## Requirement
1. Ensure you have installed DayZ and DayZServer via Steam
2. Ensure you have lanuched DayZ Lanucher at least once
3. After you change your mod, ensure you you have lanuched DayZ Lanucher to make the mod update successfully

## How to use

1. Put `server-tools.exe` in an empty folder
2. Create the `config.yaml` file in this folder
3. Fill in the `config.yaml` file with these items

```yaml
server_name: "ServerTest"
port: 2302
mission: "dayzOffline.chernarusplus"
client_mods: ["@CF"]
server_mods: []
```

4. Start `server-tools.exe` once, then make sure you have the folders `mpmissions`, `profiles`, and `serverCfgs` in your directory.
5. Note the `server_name` in your `config.yaml`, and then create a new .cfg file with the same name as it. For example, in my code above, my `server_name` is `ServerTest`, so I will create a new `ServerTest.cfg` file in the `serverCfgs` folder
6. Fill out this cfg file, official reference:

```cfg
hostname = "EXAMPLE NAME";  // Server name
password = "";              // Password to connect to the server
passwordAdmin = "";         // Password to become a server admin

description = "";			// Description of the server. Gets displayed to users in client server browser.

enableWhitelist = 0;        // Enable/disable whitelist (value 0-1)

maxPlayers = 60;            // Maximum amount of players

verifySignatures = 2;       // Verifies .pbos against .bisign files. (only 2 is supported)
forceSameBuild = 1;         // When enabled, the server will allow the connection only to clients with same the .exe revision as the server (value 0-1)

disableVoN = 0;             // Enable/disable voice over network (value 0-1)
vonCodecQuality = 20;       // Voice over network codec quality, the higher the better (values 0-30)

shardId = "123abc";			// Six alphanumeric characters for Private server

disable3rdPerson=0;         // Toggles the 3rd person view for players (value 0-1)
disableCrosshair=0;         // Toggles the cross-hair (value 0-1)

disablePersonalLight = 1;   // Disables personal light for all clients connected to server
lightingConfig = 0;         // 0 for brighter night setup, 1 for darker night setup

serverTime="SystemTime";    // Initial in-game time of the server. "SystemTime" means the local time of the machine. Another possibility is to set the time to some value in "YYYY/MM/DD/HH/MM" format, f.e. "2015/4/8/17/23" .
serverTimeAcceleration=12;  // Accelerated Time (value 0-24)// This is a time multiplier for in-game time. In this case, the time would move 24 times faster than normal, so an entire day would pass in one hour.
serverNightTimeAcceleration=1;  // Accelerated Nigh Time - The numerical value being a multiplier (0.1-64) and also multiplied by serverTimeAcceleration value. Thus, in case it is set to 4 and serverTimeAcceleration is set to 2, night time would move 8 times faster than normal. An entire night would pass in 3 hours.
serverTimePersistent=0;     // Persistent Time (value 0-1)// The actual server time is saved to storage, so when active, the next server start will use the saved time value.

guaranteedUpdates=1;        // Communication protocol used with game server (use only number 1)

loginQueueConcurrentPlayers=5;  // The number of players concurrently processed during the login process. Should prevent massive performance drop during connection when a lot of people are connecting at the same time.
loginQueueMaxPlayers=500;       // The maximum number of players that can wait in login queue

instanceId = 1;             // DayZ server instance id, to identify the number of instances per box and their storage folders with persistence files

storageAutoFix = 1;         // Checks if the persistence files are corrupted and replaces corrupted ones with empty ones (value 0-1)


class Missions
{
    class DayZ
    {
        template="dayzOffline.chernarusplus"; // Mission to load on server startup. <MissionName>.<TerrainName>
					      // Vanilla mission: dayzOffline.chernarusplus
					      // DLC mission: dayzOffline.enoch
    };
};
```

7. Put your map mission file in the mpmissions folder and make sure its name is the same as the name you have in the configuration file.

**！Please note that the map configured in the cfg will not take effect 👇（This code below will not work）**

```
class Missions
{
    class DayZ
    {
        template="dayzOffline.chernarusplus"; // Mission to load on server startup. <MissionName>.<TerrainName>
					      // Vanilla mission: dayzOffline.chernarusplus
					      // DLC mission: dayzOffline.enoch
    };
};
```
8. Run the exe file and enjoy.

---

# DayZ 服务器工具

轻松管理 DayZ 服务器

## 前置要求
1. 确保你通过Steam安装好了DayZ和DayZ Server
2. 确保你至少启动过一次DayZ客户端启动器
3. 如果你需要安装模组，请确保你订阅模组后，启动过游戏启动器，以确保模组更新完成

## 使用方法

1. 将 `server-tools.exe` 放置在空文件夹中
2. 在该文件夹中创建 `config.yaml` 文件
3. 在 `config.yaml` 中填写以下配置项

```yaml
server_name: "ServerTest"    # 服务器名称
port: 2302                   # 端口号
mission: "dayzOffline.chernarusplus"    # 地图任务
client_mods: ["@CF"]         # 客户端模组
server_mods: []              # 服务器模组
```

4. 首次运行 `server-tools.exe`，确保目录中已生成 `mpmissions`、`profiles` 和 `serverCfgs` 文件夹
5. 根据 `config.yaml` 中的 `server_name` 创建对应的配置文件。例如，如果 `server_name` 为 `ServerTest`，则在 `serverCfgs` 文件夹中创建 `ServerTest.cfg` 文件
6. 填写 cfg 文件内容，官方参考配置：

```cfg
hostname = "示例服务器";     // 服务器名称
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
serverTimeAcceleration=12;  // 时间加速倍率 (0-24)
serverNightTimeAcceleration=1;  // 夜间时间加速倍率 (0.1-64)
serverTimePersistent=0;     // 时间持久化 (0-1)

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
```

7. 将地图任务文件放入 mpmissions 文件夹，确保文件名与配置文件中的名称一致

**！请注意：cfg 中配置的地图设置将不会生效 👇（以下代码无效）**

```
class Missions
{
    class DayZ
    {
        template="dayzOffline.chernarusplus"; // 服务器启动时加载的任务
    };
};
```

8. 运行 exe 文件即可开始使用