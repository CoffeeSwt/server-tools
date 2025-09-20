[English](#dayz-server-tools) | [中文](#dayz-服务器工具)

# DayZ 服务器工具

轻松管理 DayZ 服务器

## 前置要求

1. 确保你通过 Steam 安装好了 DayZ 和 DayZ Server。
2. 确保你至少启动过一次 DayZ 客户端启动器。
3. 如果你需要安装模组，请确保你订阅模组后，启动过游戏启动器，以确保模组更新完成。

## 快速启动

1. 建立一个空文件夹（不要在桌面，不要使用中文），将该 exe 程序拖入这个空文件夹。
2. 双击启动一次 exe 程序，等待程序会自动关闭。
3. 再次启动 exe 程序，等待服务器开启。
4. 打开 DayZ 游戏，在 LAN 找到你的服务器，设置 DLC 并加入。
5. 开始玩吧~

## 高级使用（自定义配置）

1. 将 `server-tools.exe` 放置在空文件夹中。
2. 在该文件夹中创建 `config.yaml` 文件。
3. 在 `config.yaml` 中填写以下配置项

```yaml
server_name: "ServerTest" # 服务器名称
port: 2302 # 端口号
mission: "dayzOffline.chernarusplus" # 地图任务
client_mods: ["@CF"] # 客户端模组
server_mods: [] # 服务器模组
```

4. 首次运行 `server-tools.exe`，确保目录中已生成 `mpmissions`、`profiles` 和 `serverCfgs` 文件夹。
5. 根据 `config.yaml` 中的 `server_name` 创建对应的配置文件。例如，如果 `server_name` 为 `ServerTest`，则在 `serverCfgs` 文件夹中创建 `ServerTest.cfg` 文件。
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

---

# Dayz Server Tools

Easily manage your DayZ server

## Prerequisites

1. Make sure you have installed DayZ and DayZ Server via Steam.
2. Make sure you have launched the DayZ client launcher at least once.
3. If you need to install mods, make sure you have subscribed to the mods and launched the game launcher to complete the mod update.

## Quick Start

1. Create an empty folder (do not use Desktop, do not use Chinese characters), and put the exe program into this folder.
2. Double-click to run the exe program once, wait for the program to close automatically.
3. Run the exe program again, wait for the server to start.
4. Open DayZ, find your server in LAN, set DLC and join.
5. Start playing!

## Advanced Usage (Custom Configuration)

1. Put `server-tools.exe` in an empty folder.
2. Create a `config.yaml` file in this folder.
3. Fill in the following configuration in `config.yaml`:

```yaml
server_name: "ServerTest" # Server name
port: 2302 # Port number
mission: "dayzOffline.chernarusplus" # Map mission
client_mods: ["@CF"] # Client mods
server_mods: [] # Server mods
```

4. Run `server-tools.exe` for the first time, make sure the folders `mpmissions`, `profiles`, and `serverCfgs` are generated in the directory.
5. According to the `server_name` in `config.yaml`, create the corresponding configuration file. For example, if `server_name` is `ServerTest`, create a `ServerTest.cfg` file in the `serverCfgs` folder.
6. Fill in the cfg file content, official reference:

```cfg
hostname = "Example Server";     // Server name
password = "";                   // Password to connect to the server
passwordAdmin = "";              // Admin password

description = "";                // Server description, shown in client server browser

enableWhitelist = 0;             // Enable/disable whitelist (0-1)

maxPlayers = 60;                 // Maximum number of players

verifySignatures = 2;            // Verify .pbo file signatures (only 2 is supported)
forceSameBuild = 1;              // Force client version to match server version (0-1)

disableVoN = 0;                  // Enable/disable voice (0-1)
vonCodecQuality = 20;            // Voice codec quality (0-30)

shardId = "123abc";              // Six-character alphanumeric identifier for private server

disable3rdPerson=0;              // Toggle third-person view (0-1)
disableCrosshair=0;              // Toggle crosshair (0-1)

disablePersonalLight = 1;        // Disable personal light for all clients
lightingConfig = 0;              // 0 for brighter night, 1 for darker night

serverTime="SystemTime";         // Initial server time, "SystemTime" uses local machine time
serverTimeAcceleration=12;       // Time acceleration multiplier (0-24)
serverNightTimeAcceleration=1;   // Night time acceleration multiplier (0.1-64)
serverTimePersistent=0;          // Persistent time (0-1)

guaranteedUpdates=1;             // Game server communication protocol (use only 1)

loginQueueConcurrentPlayers=5;   // Number of players processed concurrently during login
loginQueueMaxPlayers=500;        // Maximum number of players in login queue

instanceId = 1;                  // DayZ server instance ID

storageAutoFix = 1;              // Check and fix corrupted persistence files (0-1)

class Missions
{
    class DayZ
    {
        template="dayzOffline.chernarusplus"; // Mission loaded at server startup
    };
};
```

7. Put the map mission file into the mpmissions folder, make sure the file name matches the name in the configuration file.

**Note: The map configuration in the cfg will not take effect 👇 (the following code is invalid)**

```
class Missions
{
    class DayZ
    {
        template="dayzOffline.chernarusplus"; // Mission loaded at server startup
    };
};
```

8. Run the exe file to start using.

---

